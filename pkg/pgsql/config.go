package pgsql

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
	CONF_PGSQL_HOST     = "PGSQL_HOST"
	CONF_PGSQL_PORT     = "PGSQL_PORT"
	CONF_PGSQL_USER     = "PGSQL_USER"
	CONF_PGSQL_PASSWORD = "PGSQL_PASSWORD"
	CONF_PGSQL_DATABASE = "PGSQL_DATABASE"

	databaseDriver = "postgres"
)

// Config is the configuration for the instanciate a pgsql driver.
type Config struct {
	Host   string
	Port   string
	User   string
	Pass   string
	DBName string
}

func config() (Config, error) {
	var c Config
	var prefix string

	if ConfigurationPrefix != "" {
		prefix = ConfigurationPrefix + "_"
		ConfigurationPrefix = ""
	}

	if !viper.IsSet(prefix + CONF_PGSQL_HOST) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_PGSQL_HOST)
	}
	if !viper.IsSet(prefix + CONF_PGSQL_PORT) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_PGSQL_PORT)
	}
	if !viper.IsSet(prefix + CONF_PGSQL_USER) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_PGSQL_USER)
	}
	if !viper.IsSet(prefix + CONF_PGSQL_PASSWORD) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_PGSQL_PASSWORD)
	}
	if !viper.IsSet(prefix + CONF_PGSQL_DATABASE) {
		return Config{}, errorx.Wrap(ErrMissingConfig, prefix+CONF_PGSQL_DATABASE)
	}

	c.Host = viper.GetString(prefix + CONF_PGSQL_HOST)
	c.Port = viper.GetString(prefix + CONF_PGSQL_PORT)
	c.User = viper.GetString(prefix + CONF_PGSQL_USER)
	c.Pass = viper.GetString(prefix + CONF_PGSQL_PASSWORD)
	c.DBName = viper.GetString(prefix + CONF_PGSQL_DATABASE)

	return c, nil
}
