package logger

import (
	"errors"

	errorx "github.com/pkg/errors"
	"github.com/spf13/viper"
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

const (
	CONF_LOGGER_LEVEL    = "LOGGER_LEVEL"
	CONF_LOGGER_ENCODING = "LOGGER_ENCODING"
)

// Config logger.
type Config struct {
	Level    Level
	Encoding Encoding
}

// NewConfig creates a logger config
func NewConfig(opts ...Option) (Config, error) {
	var c Config
	var prefix string
	var err error

	if ConfigurationPrefix != "" {
		prefix = ConfigurationPrefix + "_"
		ConfigurationPrefix = ""
	}

	if !viper.IsSet(prefix + CONF_LOGGER_LEVEL) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_LOGGER_LEVEL)
	}
	if !viper.IsSet(prefix + CONF_LOGGER_ENCODING) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_LOGGER_ENCODING)
	}

	level := viper.GetString(prefix + CONF_LOGGER_LEVEL)
	encoding := viper.GetString(prefix + CONF_LOGGER_ENCODING)

	if c.Level, err = ParseLevel(level); err != nil {
		return Config{}, errorx.Wrap(err, "unable to parse logger level")
	}
	if c.Encoding, err = ParseEncoding(encoding); err != nil {
		return Config{}, errorx.Wrap(err, "unable to parse logger encoding")
	}

	for _, opt := range opts {
		c = opt(c)
	}

	return c, nil
}
