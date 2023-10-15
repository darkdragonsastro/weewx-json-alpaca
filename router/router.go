// Package router is used to define the routes to our request handlers.
package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/darkdragonsastro/weewx-json-alpaca/middleware"
)

// Handler exposes the functions for handling web requests.
type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
	ApiVersions(w http.ResponseWriter, r *http.Request)
	Description(w http.ResponseWriter, r *http.Request)
	ConfiguredDevices(w http.ResponseWriter, r *http.Request)
	PutAction(w http.ResponseWriter, r *http.Request)
	PutCommandBlind(w http.ResponseWriter, r *http.Request)
	PutCommandBool(w http.ResponseWriter, r *http.Request)
	PutCommandString(w http.ResponseWriter, r *http.Request)
	GetConnected(w http.ResponseWriter, r *http.Request)
	PutConnected(w http.ResponseWriter, r *http.Request)
	GetDescription(w http.ResponseWriter, r *http.Request)
	GetDriverInfo(w http.ResponseWriter, r *http.Request)
	GetDriverVersion(w http.ResponseWriter, r *http.Request)
	GetInterfaceVersion(w http.ResponseWriter, r *http.Request)
	GetName(w http.ResponseWriter, r *http.Request)
	GetSupportedActions(w http.ResponseWriter, r *http.Request)
	GetAveragePeriod(w http.ResponseWriter, r *http.Request)
	PutAveragePeriod(w http.ResponseWriter, r *http.Request)
	GetCloudCover(w http.ResponseWriter, r *http.Request)
	GetDewPoint(w http.ResponseWriter, r *http.Request)
	GetHumidity(w http.ResponseWriter, r *http.Request)
	GetPressure(w http.ResponseWriter, r *http.Request)
	GetRainRate(w http.ResponseWriter, r *http.Request)
	GetSkyBrightness(w http.ResponseWriter, r *http.Request)
	GetSkyQuality(w http.ResponseWriter, r *http.Request)
	GetSkyTemperature(w http.ResponseWriter, r *http.Request)
	GetStarFWHM(w http.ResponseWriter, r *http.Request)
	GetTemperature(w http.ResponseWriter, r *http.Request)
	GetWindDirection(w http.ResponseWriter, r *http.Request)
	GetWindGust(w http.ResponseWriter, r *http.Request)
	GetWindSpeed(w http.ResponseWriter, r *http.Request)
	PutRefresh(w http.ResponseWriter, r *http.Request)
	GetSensorDescription(w http.ResponseWriter, r *http.Request)
	GetTimeSinceLastUpdate(w http.ResponseWriter, r *http.Request)
}

// NewRouter creates a new CORS enabled router for our API. All requests will be logged and
// instrumented with New Relic.
func NewRouter(h Handler, log *zap.Logger) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.TraceID)
	r.Use(middleware.Logger(log))
	r.Use(middleware.Alpaca)
	r.NotFound(h.NotFound)

	r.Get("/health", h.Health)

	r.Get("/management/apiversions", h.ApiVersions)
	r.Get("/management/v1/description", h.Description)
	r.Get("/management/v1/configureddevices", h.ConfiguredDevices)

	r.Put("/api/v1/observingconditions/0/action", h.PutAction)
	r.Put("/api/v1/observingconditions/0/commandblind", h.PutCommandBlind)
	r.Put("/api/v1/observingconditions/0/commandbool", h.PutCommandBool)
	r.Put("/api/v1/observingconditions/0/commandstring", h.PutCommandString)
	r.Get("/api/v1/observingconditions/0/connected", h.GetConnected)
	r.Put("/api/v1/observingconditions/0/connected", h.PutConnected)
	r.Get("/api/v1/observingconditions/0/description", h.GetDescription)
	r.Get("/api/v1/observingconditions/0/driverinfo", h.GetDriverInfo)
	r.Get("/api/v1/observingconditions/0/driverversion", h.GetDriverVersion)
	r.Get("/api/v1/observingconditions/0/interfaceversion", h.GetInterfaceVersion)
	r.Get("/api/v1/observingconditions/0/name", h.GetName)
	r.Get("/api/v1/observingconditions/0/supportedactions", h.GetSupportedActions)
	r.Get("/api/v1/observingconditions/0/averageperiod", h.GetAveragePeriod)
	r.Put("/api/v1/observingconditions/0/averageperiod", h.PutAveragePeriod)
	r.Get("/api/v1/observingconditions/0/cloudcover", h.GetCloudCover)
	r.Get("/api/v1/observingconditions/0/dewpoint", h.GetDewPoint)
	r.Get("/api/v1/observingconditions/0/humidity", h.GetHumidity)
	r.Get("/api/v1/observingconditions/0/pressure", h.GetPressure)
	r.Get("/api/v1/observingconditions/0/rainrate", h.GetRainRate)
	r.Get("/api/v1/observingconditions/0/skybrightness", h.GetSkyBrightness)
	r.Get("/api/v1/observingconditions/0/skyquality", h.GetSkyQuality)
	r.Get("/api/v1/observingconditions/0/skytemperature", h.GetSkyTemperature)
	r.Get("/api/v1/observingconditions/0/starfwhm", h.GetStarFWHM)
	r.Get("/api/v1/observingconditions/0/temperature", h.GetTemperature)
	r.Get("/api/v1/observingconditions/0/winddirection", h.GetWindDirection)
	r.Get("/api/v1/observingconditions/0/windgust", h.GetWindGust)
	r.Get("/api/v1/observingconditions/0/windspeed", h.GetWindSpeed)
	r.Put("/api/v1/observingconditions/0/refresh", h.PutRefresh)
	r.Get("/api/v1/observingconditions/0/sensordescription", h.GetSensorDescription)
	r.Get("/api/v1/observingconditions/0/timesincelastupdate", h.GetTimeSinceLastUpdate)

	return r
}
