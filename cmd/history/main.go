package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/dipdup-net/go-lib/tzkt/data"
	"github.com/ymohl-cl/tzktDelegator/internal/dto"
	"github.com/ymohl-cl/tzktDelegator/pkg/config"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
)

func main() {
	var err error

	if err = config.Load(); err != nil {
		if err == config.ErrNoError {
			os.Exit(0)
		}

		panic(err)
	}
	ctx := context.Background()
	var pg pgsql.PGSQL
	if pg, err = pgsql.New(); err != nil {
		panic(err)
	}
	defer pg.Close()
	bdd := dto.New(pg.Driver())

	// count the number of delegations
	var nbDelegations int64
	if nbDelegations, err = countDelegation(); err != nil {
		panic(err)
	}

	// browse the delegations
	if err = browseDlegations(ctx, bdd, nbDelegations); err != nil {
		panic(err)
	}
}

func browseDlegations(ctx context.Context, bdd *dto.Queries, count int64) error {
	var err error

	limit := 10000
	offset := int64(0)
	url := fmt.Sprintf("https://api.tzkt.io/v1/operations/delegations?limit=%d", limit)

	for offset < count {
		var response *http.Response
		if response, err = http.Get(fmt.Sprintf("%s&offset=%d", url, offset)); err != nil {
			return err
		}

		if err = reponseReader(ctx, response, bdd); err != nil {
			return err
		}

		offset += int64(limit)
	}

	return nil
}

func reponseReader(ctx context.Context, response *http.Response, bdd *dto.Queries) error {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// parse the delegate data
	var delegates []data.Delegation
	err = json.Unmarshal(body, &delegates)
	if err != nil {
		return err
	}

	for _, delegate := range delegates {
		delegator := ""
		if delegate.Sender != nil {
			delegator = delegate.Sender.Address
		} else if delegate.Initiator != nil {
			delegator = delegate.Initiator.Address
		}

		_, err = bdd.InsertDelegator(ctx, dto.InsertDelegatorParams{
			DelegationDate:   delegate.Timestamp,
			DelegatorAddress: delegator,
			BlockHash:        delegate.Block,
			Amount:           delegate.Amount.String(),
			BlockState:       int64(delegate.Level),
			ExternalID:       int64(delegate.ID),
		})
		if err != nil {
			fmt.Printf("error to insert delegation: %s\n", err.Error())
		}
	}

	return nil
}

func countDelegation() (int64, error) {
	resp, err := http.Get("https://api.tzkt.io/v1/operations/delegations/count")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var countResponse int64
	err = json.Unmarshal(body, &countResponse)
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of delegations: ", countResponse)
	return countResponse, nil
}
