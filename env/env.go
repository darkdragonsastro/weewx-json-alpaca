package env

import (
	"github.com/caarlos0/env"
)

// Config represents the common environment variables needed for all apps.
type Config struct {
}

// Load will bind the environment variables to the given config.
func Load(c interface{}) error {
	return env.Parse(c)
}
