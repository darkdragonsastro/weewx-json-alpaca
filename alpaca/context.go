package alpaca

import "context"

type AlpacaContext struct {
	ClientID            uint64
	ClientTransactionID uint64
	ServerTransactionID uint64
	DeviceType          *string
	DeviceNumber        *int
}

type contextKey string

func (k contextKey) String() string {
	return "context key: " + string(k)
}

var (
	traceIDKey = contextKey("alpaca")
)

func WithAlpacaContext(ctx context.Context, alpacaContext AlpacaContext) context.Context {
	return context.WithValue(ctx, traceIDKey, alpacaContext)
}

// FromContext retrieves the AlpacaContext from the request context.
func FromContext(ctx context.Context) AlpacaContext {
	if val, ok := ctx.Value(traceIDKey).(AlpacaContext); ok {
		return val
	}

	return AlpacaContext{}
}
