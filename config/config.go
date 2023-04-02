// YOU CAN EDIT YOUR CUSTOM CONFIG HERE

package config

import (
	"fmt"
	"strconv"
)

// Config ...
type Config struct {
	Base         `mapstructure:",squash"`
	SentryConfig SentryConfig `json:"sentry_config" mapstructure:"sentry_config"`
	// Custom here

}

// GetHTTPAddress ...
func (c *Config) GetHTTPAddress() string {
	if _, err := strconv.Atoi(c.HTTPAddress); err == nil {
		return fmt.Sprintf(":%v", c.HTTPAddress)
	}
	return c.HTTPAddress
}

// SentryConfig ...
type SentryConfig struct {
	Enabled bool   `json:"enabled" mapstructure:"enabled"`
	DNS     string `json:"dns" mapstructure:"dns"`
	Trace   bool   `json:"trace" mapstructure:"trace"`
}
