package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/darkdragonsastro/weewx-json-alpaca/alpaca"
)

var serverTransactionID atomic.Int64

// TraceID tries to read the X-Trace-Id request header, and if empty, creates
// a new traceID to use. This traceID will be added to the response X-Trace-Id
// header and to the request context.
func Alpaca(next http.Handler) http.Handler {
	badRequest := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusBadRequest)
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx alpaca.AlpacaContext

		if r.Method == http.MethodGet {
			query := r.URL.Query()

			var err error
			var clientID uint64 = 0
			var clientTransactionID uint64 = 0

			for k, vs := range query {
				if strings.ToLower(k) == "clientid" {
					clientID, err = strconv.ParseUint(vs[0], 10, 64)
					if err != nil {
						badRequest.ServeHTTP(w, r)
						return
					}
				} else if strings.ToLower(k) == "clienttransactionid" {
					clientTransactionID, err = strconv.ParseUint(vs[0], 10, 64)
					if err != nil {
						badRequest.ServeHTTP(w, r)
						return
					}
				}
			}

			ctx = alpaca.AlpacaContext{
				ClientID:            clientID,
				ClientTransactionID: clientTransactionID,
				ServerTransactionID: uint64(serverTransactionID.Add(1)),
			}
		} else {
			err := r.ParseForm()
			if err != nil {
				badRequest.ServeHTTP(w, r)
				return
			}

			var clientID uint64 = 0
			var clientTransactionID uint64 = 0

			for k, vs := range r.PostForm {
				if k == "ClientID" {
					clientID, err = strconv.ParseUint(vs[0], 10, 64)
					if err != nil {
						badRequest.ServeHTTP(w, r)
						return
					}
				} else if k == "ClientTransactionID" {
					clientTransactionID, err = strconv.ParseUint(vs[0], 10, 64)
					if err != nil {
						badRequest.ServeHTTP(w, r)
						return
					}
				}
			}

			ctx = alpaca.AlpacaContext{
				ClientID:            clientID,
				ClientTransactionID: clientTransactionID,
				ServerTransactionID: uint64(serverTransactionID.Add(1)),
			}
		}

		r = r.WithContext(alpaca.WithAlpacaContext(r.Context(), ctx))
		next.ServeHTTP(w, r)
	})
}
