// Package handler implements the request handlers for the API.
package handler

import (
	"net/http"
	"strings"

	"github.com/darkdragonsastro/weewx-json-alpaca/alpaca"
	"github.com/darkdragonsastro/weewx-json-alpaca/logging"
	"go.uber.org/zap"
)

func (h *Handler) PutAction(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) PutCommandBlind(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) PutCommandBool(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) PutCommandString(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) GetConnected(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	oc := h.weewx.GetCurrent()

	writeResponse(r, w, http.StatusOK, &AlpacaBooleanResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: oc.Connected,
	})
}

func (h *Handler) PutConnected(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	err := r.ParseForm()
	if err != nil {
		logger := logging.FromContext(r.Context())
		logger.Error("failed to parse form", zap.Error(err))

		writeResponse(r, w, http.StatusBadRequest, nil)
		return
	}

	var req AlpacaConnectedRequest

	if strings.ToLower(r.PostForm.Get("Connected")) != "true" && strings.ToLower(r.PostForm.Get("Connected")) != "false" {
		writeResponse(r, w, http.StatusBadRequest, nil)
		return
	}

	err = decoder.Decode(&req, r.PostForm)
	if err != nil {
		logger := logging.FromContext(r.Context())
		logger.Error("failed to parse form", zap.Error(err))

		writeResponse(r, w, http.StatusBadRequest, nil)
		return
	}

	writeResponse(r, w, http.StatusOK, &AlpacaResponse{
		ClientTransactionID: req.ClientTransactionID,
		ServerTransactionID: ctx.ServerTransactionID,
	})
}

func (h *Handler) GetDescription(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaStringResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: "weewx-json-alpaca",
	})
}

func (h *Handler) GetDriverInfo(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaStringResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: "weewx-json-alpaca",
	})
}

func (h *Handler) GetDriverVersion(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaStringResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: "1.0",
	})
}

func (h *Handler) GetInterfaceVersion(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaIntResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: 1,
	})
}

func (h *Handler) GetName(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaStringResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: "weewx-json-alpaca",
	})
}

func (h *Handler) GetSupportedActions(w http.ResponseWriter, r *http.Request) {
	ctx := alpaca.FromContext(r.Context())

	writeResponse(r, w, http.StatusOK, &AlpacaStringArrayResponse{
		AlpacaResponse: AlpacaResponse{
			ClientTransactionID: ctx.ClientTransactionID,
			ServerTransactionID: ctx.ServerTransactionID,
		},
		Value: []string{},
	})
}
