// Package handler implements the request handlers for the API.
package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/darkdragons/weewx-json-alpaca/alpaca"
	"github.com/darkdragons/weewx-json-alpaca/tracing"
)

func (h *Handler) GetAveragePeriod(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: oc.AveragePeriod,
	})
}

func (h *Handler) PutAveragePeriod(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	averagePeriod, err := strconv.ParseFloat(r.Form.Get("AveragePeriod"), 64)
	if err != nil {
		errNumber := 0x401
		errMessage := "Invalid Value"

		writeResponse(r, w, http.StatusBadRequest, &AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
			ErrorNumber:         &errNumber,
			ErrorMessage:        &errMessage,
		})
		return
	}

	if averagePeriod != 0.0 {
		errNumber := 0x401
		errMessage := "Invalid Value"

		writeResponse(r, w, http.StatusOK, &AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
			ErrorNumber:         &errNumber,
			ErrorMessage:        &errMessage,
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
	})
}

func (h *Handler) GetCloudCover(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	errNumber := 0x400
	errMessage := "Not implemented"

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
		ErrorNumber:         &errNumber,
		ErrorMessage:        &errMessage,
	})
}

func (h *Handler) GetDewPoint(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.DewPoint == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.DewPoint,
	})
}

func (h *Handler) GetHumidity(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.Humidity == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.Humidity,
	})
}

func (h *Handler) GetPressure(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.Pressure == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.Pressure,
	})
}

func (h *Handler) GetRainRate(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.RainRate == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.RainRate,
	})
}

func (h *Handler) GetSkyBrightness(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	errNumber := 0x400
	errMessage := "Not implemented"

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
		ErrorNumber:         &errNumber,
		ErrorMessage:        &errMessage,
	})
}

func (h *Handler) GetSkyQuality(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	errNumber := 0x400
	errMessage := "Not implemented"

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
		ErrorNumber:         &errNumber,
		ErrorMessage:        &errMessage,
	})
}

func (h *Handler) GetSkyTemperature(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	errNumber := 0x400
	errMessage := "Not implemented"

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
		ErrorNumber:         &errNumber,
		ErrorMessage:        &errMessage,
	})
}

func (h *Handler) GetStarFWHM(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	errNumber := 0x400
	errMessage := "Not implemented"

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
		ErrorNumber:         &errNumber,
		ErrorMessage:        &errMessage,
	})
}

func (h *Handler) GetTemperature(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.Temperature == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.Temperature,
	})
}

func (h *Handler) GetWindDirection(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.WindDirection == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.WindDirection,
	})
}

func (h *Handler) GetWindGust(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.WindGust == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.WindGust,
	})
}

func (h *Handler) GetWindSpeed(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()
	if oc == nil || oc.WindSpeed == nil {
		writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
			TraceID: tracing.FromContext(r.Context()),
			Message: "observing conditions is nil",
		})
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: *oc.WindSpeed,
	})
}

func (h *Handler) PutRefresh(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())
	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: ctx.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
	})
}

func (h *Handler) GetSensorDescription(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	var sensorName string

	for k, vs := range query {
		if strings.ToLower(k) == "sensorname" {
			sensorName = vs[0]
		}
	}

	ctx := alpaca.FromContext(r.Context())

	switch strings.ToLower(sensorName) {
	case "cloudcover":
		fallthrough
	case "skybrightness":
		fallthrough
	case "skyquality":
		fallthrough
	case "skytemperature":
		fallthrough
	case "starfwhm":
		errNumber := 0x400
		errMessage := "Not implemented"
		writeResponse(r, w, http.StatusOK, &AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
			ErrorNumber:         &errNumber,
			ErrorMessage:        &errMessage,
		})
		return
	}

	switch strings.ToLower(sensorName) {
	case "dewpoint":
		fallthrough
	case "humidity":
		fallthrough
	case "pressure":
		fallthrough
	case "rainrate":
		fallthrough
	case "temperature":
		fallthrough
	case "winddirection":
		fallthrough
	case "windgust":
		fallthrough
	case "windspeed":
		writeResponse(r, w, http.StatusOK, &AlpacaStringResponse{
			AlpacaResponse: AlpacaResponse{
				ClientTransactionID: ctx.ClientTransactionID,
				ServerTransactionID: ctx.ServerTransactionID,
			},
			Value: sensorName,
		})
		return
	}

	writeResponse(r, w, http.StatusBadRequest, nil)
}

func (h *Handler) GetTimeSinceLastUpdate(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	query := r.URL.Query()
	var sensorName string

	for k, vs := range query {
		if strings.ToLower(k) == "sensorname" {
			sensorName = vs[0]
		}
	}

	switch strings.ToLower(sensorName) {
	case "cloudcover":
		fallthrough
	case "skybrightness":
		fallthrough
	case "skyquality":
		fallthrough
	case "skytemperature":
		fallthrough
	case "starfwhm":
		errNumber := 0x400
		errMessage := "Not implemented"
		writeResponse(r, w, http.StatusOK, &AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
			ErrorNumber:         &errNumber,
			ErrorMessage:        &errMessage,
		})
		return
	}

	switch strings.ToLower(sensorName) {
	case "":
		fallthrough
	case "dewpoint":
		fallthrough
	case "humidity":
		fallthrough
	case "pressure":
		fallthrough
	case "rainrate":
		fallthrough
	case "temperature":
		fallthrough
	case "winddirection":
		fallthrough
	case "windgust":
		fallthrough
	case "windspeed":
		oc := h.weewx.GetCurrent()
		if oc == nil {
			writeResponse(r, w, http.StatusInternalServerError, &SimpleResponse{
				TraceID: tracing.FromContext(r.Context()),
				Message: "observing conditions is nil",
			})
			return
		}

		writeResponse(r, w, http.StatusOK, &AlpacaFloatResponse{
			AlpacaResponse: AlpacaResponse{
				ClientTransactionID: ctx.ClientTransactionID,
				ServerTransactionID: ctx.ServerTransactionID,
			},
			Value: time.Since(oc.LastUpdated).Seconds(),
		})
		return
	}

	writeResponse(r, w, http.StatusBadRequest, nil)
}
