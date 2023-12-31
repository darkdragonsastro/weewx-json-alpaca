package logging

import (
	"os"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/darkdragonsastro/weewx-json-alpaca/env"
)

// Initialize creates a new JSON zap logger.
func Initialize(c env.Config) *zap.Logger {
	logEnc := zap.NewProductionEncoderConfig()
	logEnc.EncodeTime = zapcore.ISO8601TimeEncoder
	logEnc.TimeKey = "timestamp"

	log := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(logEnc),
		zapcore.Lock(os.Stdout),
		zap.NewAtomicLevel(),
	))

	log = log.With(
		zap.String("run_id", uuid.New().String()),
		zap.Time("start_time", time.Now()),
	)

	return log
}
