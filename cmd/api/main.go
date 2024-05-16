package main

import (
	"os"

	_ "github.com/lib/pq"

	"github.com/gofast-pkg/api"
	"github.com/ymohl-cl/tzktDelegator/pkg/config"
	"github.com/ymohl-cl/tzktDelegator/pkg/logger"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
)

func main() {
	var err error
	var l logger.Logger
	var pg pgsql.PGSQL
	var app api.API
	var h Handler

	if err = config.Load(); err != nil {
		if err == config.ErrNoError {
			os.Exit(0)
		}

		panic(err)
	}

	if l, err = logger.New(); err != nil {
		panic(err)
	}
	defer l.Close()

	if pg, err = pgsql.New(); err != nil {
		panic(err)
	}
	defer pg.Close()

	h = NewHandler(l, pg)

	if app, err = api.New(); err != nil {
		panic(err)
	}

	group, err := app.SubRouter("/xtz", false)
	if err != nil {
		panic(err)
	}
	group.GET("/delegations", h.GetDelegations)

	if err = app.Start(); err != nil {
		panic(err)
	}
}
