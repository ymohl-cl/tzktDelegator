// Package pgsql provides a wrapper for the pgsql driver
package pgsql

import (
	"database/sql"
	"fmt"
)

// PGSQL is the interface that wraps the pgsql driver
//
//go:generate mockery --name=PGSQL --output=mocks --filename=pgsql.go --outpkg=mocks
type PGSQL interface {
	// Close closes the database and prevents new queries from starting.
	Close() error
	// Driver returns the underlying driver
	Driver() *sql.DB
}

type pgsql struct {
	driver *sql.DB
}

// New returns a new instance of PGSQL
// It will return an error if the config is invalid
// Call New assume that you have initialized the config with viper.
// If you want to initialize the config with a different package, use NewWithConfig
func New() (PGSQL, error) {
	var err error
	var cfg Config

	if cfg, err = config(); err != nil {
		return nil, err
	}

	return NewWithConfig(cfg)
}

// NewWithConfig returns a new instance of PGSQL
// It will return an error if the config is invalid
// Connection should be established and a ping should be performed
// for the connection to be valid.
func NewWithConfig(c Config) (PGSQL, error) {
	var err error
	pg := pgsql{}

	conn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.Host,
		c.Port,
		c.User,
		c.Pass,
		c.DBName)
	if pg.driver, err = sql.Open(databaseDriver, conn); err != nil {
		return nil, err
	}
	if err = pg.driver.Ping(); err != nil {
		return nil, err
	}

	return &pg, nil
}

func (pg pgsql) Driver() *sql.DB {
	return pg.driver
}

func (pg pgsql) Close() error {
	return pg.driver.Close()
}
