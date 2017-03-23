package control

import (
	"fmt"

	"github.com/joyent/containerpilot/utils"
)

const (
	DEFAULT_SOCKET="/var/run/containerpilot.socket"
)

// Config represents the location on the file system which serves the Unix
// control socket file.
type Config struct {
	Socket    string      `mapstructure:"socket"`
}

// NewConfig parses a json config into a validated Config used by control
// Server.
func NewConfig(raw interface{}) (*Config, error) {
	if raw == nil {
		return nil, nil
	}

	cfg := &Config{Socket: DEFAULT_SOCKET} // defaults
	if err := utils.DecodeRaw(raw, cfg); err != nil {
		return nil, fmt.Errorf("control configuration error: %v", err)
	}

	// TODO: Validate file system location and handle accordingly...

	return cfg, nil
}
