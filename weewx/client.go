// https://weather.ghro.club/weewx.json
package weewx

import (
	"encoding/json"
	"net/http"
	"sync/atomic"
	"time"

	"go.uber.org/zap"

	"github.com/darkdragonsastro/weewx-json-alpaca/httputil"
)

type ObservingConditions struct {
	Connected     bool
	AveragePeriod float64
	DewPoint      *float64
	Humidity      *float64
	Pressure      *float64
	RainRate      *float64
	Temperature   *float64
	WindDirection *float64
	WindGust      *float64
	WindSpeed     *float64
	LastUpdated   time.Time
}

type Station struct {
	Location  string  `json:"location"`
	Altitude  float64 `json:"altitude (meters)"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Link      string  `json:"link"`
}

type WeewxTime struct {
	time.Time
}

// Unmarshal time with the format of Fri, 13 Oct 2023 22:55:00 EDT
func (t *WeewxTime) UnmarshalJSON(b []byte) error {
	tt, err := time.Parse(`"Mon, 2 Jan 2006 15:04:05 MST"`, string(b))
	if err != nil {
		return err
	}

	*t = WeewxTime{
		Time: tt,
	}

	return nil
}

type Generation struct {
	Time      WeewxTime `json:"time"`
	Generator string    `json:"generator"`
}

type Value struct {
	Value float64 `json:"value"`
	Units string  `json:"units"`
}

func (v *Value) AsPercent() *float64 {
	if v == nil {
		return nil
	}

	if v.Units == "%" {
		return &v.Value
	}

	return nil
}

func (v *Value) AsPressure() *float64 {
	if v == nil {
		return nil
	}

	if v.Units == "hPa" {
		return &v.Value
	}

	if v.Units == "inHg" {
		vv := v.Value
		vv = vv * 33.863886666667
		return &vv
	}

	return nil
}

func (v *Value) AsRate() *float64 {
	if v == nil {
		return nil
	}

	if v.Units == "mm/h" {
		return &v.Value
	}

	if v.Units == "in/h" {
		vv := v.Value
		vv = vv * 25.4
		return &vv
	}

	return nil
}

func (v *Value) AsDirection() *float64 {
	if v == nil {
		vv := 0.0
		return &vv
	}

	if v.Units == "°" {
		return &v.Value
	}

	return &v.Value
}

func (v *Value) AsSpeed() *float64 {
	if v == nil {
		return nil
	}

	if v.Units == "km/h" {
		return &v.Value
	}

	if v.Units == "mph" {
		vv := v.Value
		vv = vv * 1.609344
		return &vv
	}

	return nil
}

func (v *Value) AsTemperature() *float64 {
	if v == nil {
		return nil
	}

	if v.Units == "°C" {
		return &v.Value
	}

	if v.Units == "°F" {
		vv := v.Value
		vv = (vv - 32) * 5 / 9
		return &vv
	}

	return nil
}

type Current struct {
	Temperature       *Value `json:"temperature"`
	DewPoint          *Value `json:"dewpoint"`
	Humidity          *Value `json:"humidity"`
	HeatIndex         *Value `json:"heat index"`
	Barometer         *Value `json:"barometer"`
	WindSpeed         *Value `json:"wind speed"`
	WindGust          *Value `json:"wind gust"`
	WindChill         *Value `json:"wind chill"`
	WindDirection     *Value `json:"wind direction"`
	RainRate          *Value `json:"rain rate"`
	InsideTemperature *Value `json:"inside temperature"`
	InsideHumidity    *Value `json:"inside humidity"`
}

type WeeWx struct {
	Station    Station    `json:"station"`
	Generation Generation `json:"generation"`
	Current    Current    `json:"current"`
}

type Client struct {
	Url string
	c   *http.Client
	log *zap.Logger

	val  *atomic.Value
	done chan bool
}

func NewClient(url string, log *zap.Logger) *Client {
	c := &http.Client{
		Timeout:   10 * time.Second,
		Transport: httputil.DefaultLogTransport(log, httputil.LogTransport(http.DefaultTransport)),
	}

	val := &atomic.Value{}
	val.Store(&ObservingConditions{
		Connected: false,
	})

	return &Client{
		Url:  url,
		c:    c,
		log:  log,
		val:  val,
		done: make(chan bool),
	}
}

func (c *Client) refresh() (*WeeWx, error) {
	resp, err := c.c.Get(c.Url)
	if err != nil {
		return nil, err
	}

	var weewx WeeWx
	err = json.NewDecoder(resp.Body).Decode(&weewx)
	if err != nil {
		return nil, err
	}

	conditions := ObservingConditions{
		Connected:     true,
		AveragePeriod: 0,
		DewPoint:      weewx.Current.DewPoint.AsTemperature(),
		Humidity:      weewx.Current.Humidity.AsPercent(),
		Pressure:      weewx.Current.Barometer.AsPressure(),
		RainRate:      weewx.Current.RainRate.AsRate(),
		Temperature:   weewx.Current.Temperature.AsTemperature(),
		WindDirection: weewx.Current.WindDirection.AsDirection(),
		WindGust:      weewx.Current.WindGust.AsSpeed(),
		WindSpeed:     weewx.Current.WindSpeed.AsSpeed(),
		LastUpdated:   weewx.Generation.Time.Time,
	}

	c.val.Store(&conditions)

	return &weewx, nil
}

func (c *Client) GetCurrent() *ObservingConditions {
	conditions := c.val.Load().(*ObservingConditions)
	return conditions
}

func (c *Client) Start() {
	c.refresh()

	timer := time.NewTicker(5 * time.Second)

	c.log.Info("starting weewx client")

	go func() {
		for range timer.C {
			select {
			case <-c.done:
				c.val.Store(&ObservingConditions{
					Connected: false,
				})

				timer.Stop()

				c.log.Info("stopping weewx client")

				return
			case <-timer.C:
				_, err := c.refresh()
				if err != nil {
					c.log.Error("error getting weewx", zap.Error(err))
				}
			}
		}
	}()
}

func (c *Client) Stop() {
	c.done <- true
}
