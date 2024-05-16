package main

import (
	"context"
	"os"

	_ "github.com/lib/pq"

	"github.com/ymohl-cl/tzktDelegator/pkg/config"
	"github.com/ymohl-cl/tzktDelegator/pkg/logger"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
	"github.com/ymohl-cl/tzktDelegator/pkg/tzkt"
)

func main() {
	var err error
	var l logger.Logger
	var pg pgsql.PGSQL
	var tezoDriver tzkt.TzKT

	if err = config.Load(); err != nil {
		if err == config.ErrNoError {
			os.Exit(0)
		}

		panic(err)
	}
	ctx := context.Background()

	if l, err = logger.New(); err != nil {
		panic(err)
	}
	defer l.Close()
	if pg, err = pgsql.New(); err != nil {
		panic(err)
	}
	defer pg.Close()
	if tezoDriver, err = tzkt.New(ctx); err != nil {
		panic(err)
	}
	defer tezoDriver.Close()

	app := NewApp(pg, tezoDriver, l)
	if err = app.Run(ctx); err != nil {
		panic(err)
	}
}
