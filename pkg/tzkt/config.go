package tzkt

import (
	"errors"

	errorx "github.com/pkg/errors"
	"github.com/spf13/viper"
)

const (
	CONF_TZKT_HOST = "TZKT_HOST"
)

var (
	// ErrMissingConfig is the error returned when a config key is missing.
	ErrMissingConfig = errors.New("missing config key")
)

var (
	// ConfigurationPrefix is the prefix for the configuration keys.
	// Default is empty.
	// Example:
	//   - if ConfigurationPrefix is "APP", the configuration key for host will be "APP_PGSQL_HOST".
	// After the call to config(), the ConfigurationPrefix is reset to empty.
	ConfigurationPrefix = ""
)

type Config struct {
	Host string
}

// NewConfig creates a tzkt config
func NewConfig() (Config, error) {
	var c Config
	var prefix string

	if ConfigurationPrefix != "" {
		prefix = ConfigurationPrefix + "_"
		ConfigurationPrefix = ""
	}

	if !viper.IsSet(prefix + CONF_TZKT_HOST) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_TZKT_HOST)
	}

	return c, nil
}
