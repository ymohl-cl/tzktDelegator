package pgsql

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	hostVal := "localhost"
	portVal := "5432"
	userVal := "postgres"
	passVal := "postgres"
	dbNameVal := "postgres"

	t.Run("should return the config with prefix", func(t *testing.T) {
		prefix := "APP"
		ConfigurationPrefix = prefix

		prefix += "_"
		viper.Set(prefix+CONF_PGSQL_HOST, hostVal)
		viper.Set(prefix+CONF_PGSQL_PORT, portVal)
		viper.Set(prefix+CONF_PGSQL_USER, userVal)
		viper.Set(prefix+CONF_PGSQL_PASSWORD, passVal)
		viper.Set(prefix+CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		config, err := config()
		if assert.NoError(t, err) {
			assert.NotEmpty(t, config)
			assert.Equal(t, config.Host, hostVal)
			assert.Equal(t, config.Port, portVal)
			assert.Equal(t, config.User, userVal)
			assert.Equal(t, config.Pass, passVal)
			assert.Equal(t, config.DBName, dbNameVal)
		}
	})
	t.Run("should return the config without prefix", func(t *testing.T) {
		viper.Set(CONF_PGSQL_HOST, hostVal)
		viper.Set(CONF_PGSQL_PORT, portVal)
		viper.Set(CONF_PGSQL_USER, userVal)
		viper.Set(CONF_PGSQL_PASSWORD, passVal)
		viper.Set(CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		config, err := config()
		if assert.NoError(t, err) {
			assert.NotEmpty(t, config)
			assert.Equal(t, config.Host, hostVal)
			assert.Equal(t, config.Port, portVal)
			assert.Equal(t, config.User, userVal)
			assert.Equal(t, config.Pass, passVal)
			assert.Equal(t, config.DBName, dbNameVal)
		}
	})
	t.Run("should return error when missing host", func(t *testing.T) {
		viper.Set(CONF_PGSQL_PORT, portVal)
		viper.Set(CONF_PGSQL_USER, userVal)
		viper.Set(CONF_PGSQL_PASSWORD, passVal)
		viper.Set(CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		cfg, err := config()
		if assert.Error(t, err); err != nil {
			assert.Empty(t, cfg)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
	t.Run("should return error when missing port", func(t *testing.T) {
		viper.Set(CONF_PGSQL_HOST, hostVal)
		viper.Set(CONF_PGSQL_USER, userVal)
		viper.Set(CONF_PGSQL_PASSWORD, passVal)
		viper.Set(CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		cfg, err := config()
		if assert.Error(t, err); err != nil {
			assert.Empty(t, cfg)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
	t.Run("should return error when missing user", func(t *testing.T) {
		viper.Set(CONF_PGSQL_HOST, hostVal)
		viper.Set(CONF_PGSQL_PORT, portVal)
		viper.Set(CONF_PGSQL_PASSWORD, passVal)
		viper.Set(CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		cfg, err := config()
		if assert.Error(t, err); err != nil {
			assert.Empty(t, cfg)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
	t.Run("should return error when missing password", func(t *testing.T) {
		viper.Set(CONF_PGSQL_HOST, hostVal)
		viper.Set(CONF_PGSQL_PORT, portVal)
		viper.Set(CONF_PGSQL_USER, userVal)
		viper.Set(CONF_PGSQL_DATABASE, dbNameVal)
		defer viper.Reset()

		cfg, err := config()
		if assert.Error(t, err); err != nil {
			assert.Empty(t, cfg)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
	t.Run("should return error when missing database", func(t *testing.T) {
		viper.Set(CONF_PGSQL_HOST, hostVal)
		viper.Set(CONF_PGSQL_PORT, portVal)
		viper.Set(CONF_PGSQL_USER, userVal)
		viper.Set(CONF_PGSQL_PASSWORD, passVal)
		defer viper.Reset()

		cfg, err := config()
		if assert.Error(t, err); err != nil {
			assert.Empty(t, cfg)
			assert.ErrorIs(t, err, ErrMissingConfig)
		}
	})
}
