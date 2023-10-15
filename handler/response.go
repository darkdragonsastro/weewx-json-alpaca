package handler

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"github.com/darkdragons/weewx-json-alpaca/logging"
)

// SimpleResponse is used to send a meaningful message back to the caller, with
// trace id to debug later.
type SimpleResponse struct {
	XMLName xml.Name `json:"-" xml:"simpleResponse"`
	TraceID string   `json:"trace_id" xml:"traceId,attr"`
	Message string   `json:"message" xml:",innerxml"`
}

type AlpacaResponse struct {
	ClientTransactionID uint64  `json:"ClientTransactionID"`
	ServerTransactionID uint64  `json:"ServerTransactionID"`
	ErrorNumber         *int    `json:"ErrorNumber,omitempty"`
	ErrorMessage        *string `json:"ErrorMessage,omitempty"`
}

type AlpacaIntArrayResponse struct {
	AlpacaResponse
	Value []int `json:"Value"`
}

type AlpacaStringArrayResponse struct {
	AlpacaResponse
	Value []string `json:"Value"`
}

type AlpacaStringResponse struct {
	AlpacaResponse
	Value string `json:"Value"`
}

type AlpacaFloatResponse struct {
	AlpacaResponse
	Value float64 `json:"Value"`
}

type AlpacaIntResponse struct {
	AlpacaResponse
	Value int `json:"Value"`
}

type AlpacaBooleanResponse struct {
	AlpacaResponse
	Value bool `json:"Value"`
}

type AlpacaRequest struct {
	ClientID            uint64 `json:"ClientID"`
	ClientTransactionID uint64 `json:"ClientTransactionID"`
}

type AlpacaConnectedRequest struct {
	AlpacaRequest
	Connected bool `json:"Connected"`
}

func writeResponse(r *http.Request, w http.ResponseWriter, status int, resp interface{}) {
	accept := r.Header.Get("Accept")

	if strings.HasPrefix(accept, "text/xml") {
		writeXMLResponse(r.Context(), w, status, resp)
	} else {
		writeJSONResponse(r.Context(), w, status, resp)
	}
}

func writeXMLResponse(ctx context.Context, w http.ResponseWriter, status int, resp interface{}) {
	if resp != nil {
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}

	w.WriteHeader(status)

	if resp != nil {
		err := xml.NewEncoder(w).Encode(resp)
		if err != nil {
			logging.FromContext(ctx).Error("error writing response", zap.Error(err))
		}
	}
}

func writeJSONResponse(ctx context.Context, w http.ResponseWriter, status int, resp interface{}) {
	if resp != nil {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
	}

	w.WriteHeader(status)

	if resp != nil {
		err := json.NewEncoder(w).Encode(resp)
		if err != nil {
			logging.FromContext(ctx).Error("error writing response", zap.Error(err))
		}
	}
}
