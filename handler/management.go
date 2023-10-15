// Package handler implements the request handlers for the API.
package handler

import (
	"net/http"

	"github.com/darkdragons/weewx-json-alpaca/alpaca"
)

func (h *Handler) ApiVersions(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaIntArrayResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: []int{1},
	})
}

type DescriptionValue struct {
	ServerName          string `json:"ServerName"`
	Manufacturer        string `json:"Manufacturer"`
	ManufacturerVersion string `json:"ManufacturerVersion"`
	Location            string `json:"Location"`
}

type DescriptionResponse struct {
	AlpacaResponse
	Value DescriptionValue
}

func (h *Handler) Description(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &DescriptionResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: DescriptionValue{
			ServerName:          "weewx-json-alpaca",
			Manufacturer:        "darkdragons",
			ManufacturerVersion: "0.0.1",
			Location:            "Earth",
		},
	})
}

type ConfiguredDevicesValue struct {
	DeviceName   string `json:"DeviceName"`
	DeviceType   string `json:"DeviceType"`
	DeviceNumber int    `json:"DeviceNumber"`
	UniqueID     string `json:"UniqueID"`
}

type ConfiguredDevicesResponse struct {
	AlpacaResponse
	Value []ConfiguredDevicesValue
}

func (h *Handler) ConfiguredDevices(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &ConfiguredDevicesResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: []ConfiguredDevicesValue{
			{
				DeviceName:   "weewx-json-alpaca",
				DeviceType:   "observingconditions",
				DeviceNumber: 0,
				UniqueID:     h.weewx.Url,
			},
		},
	})
}
