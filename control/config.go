package control

import (
	"fmt"

	"github.com/joyent/containerpilot/utils"
)

const (
	DEFAULT_SOCKET = "/var/run/containerpilot.socket"
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

	// cfg := &Config{Control: &NestedConfig{ Socket: DEFAULT_SOCKET }} // defaults
	cfg := &Config{Socket: DEFAULT_SOCKET} // defaults
	if err := utils.DecodeRaw(raw, cfg); err != nil {
		return nil, fmt.Errorf("control configuration error: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("control configuration error: %v", err)
	}

	return cfg, nil
}

// Validate parsed control configuration and the values contained within.
func (cfg *Config) Validate() error {
	// TODO: Validate NestedConfig and socket's file system location ...
	return nil
}
