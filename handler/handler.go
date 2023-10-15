// Package handler implements the request handlers for the API.
package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"

	"github.com/darkdragons/weewx-json-alpaca/tracing"
	"github.com/darkdragons/weewx-json-alpaca/weewx"
)

var decoder = schema.NewDecoder()

// Handler has all the functions needed to serve our api.
type Handler struct {
	weewx *weewx.Client
}

// New creates a new handler.
func New(weewxClient *weewx.Client) *Handler {
	return &Handler{
		weewx: weewxClient,
	}
}

// Health always returns a 200 response.
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	writeResponse(r, w, http.StatusOK, &SimpleResponse{
		TraceID: tracing.FromContext(r.Context()),
		Message: "OK",
	})
}

// Unauthorized is called when the request is not authorized.
func (h *Handler) Unauthorized(w http.ResponseWriter, r *http.Request) {
	writeResponse(r, w, http.StatusUnauthorized, &SimpleResponse{
		TraceID: tracing.FromContext(r.Context()),
		Message: "unauthorized",
	})
}

// NotFound is called when the request is for an unknown resource.
func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	writeResponse(r, w, http.StatusNotFound, &SimpleResponse{
		TraceID: tracing.FromContext(r.Context()),
		Message: "not found",
	})
}

func routeParamInt(ctx context.Context, name string) int {
	// this func should only be called for params that are guaranteed to be ints.
	val, _ := strconv.Atoi(chi.RouteContext(ctx).URLParam("id")) // nolint
	return val
}
