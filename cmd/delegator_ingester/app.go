package main

import (
	"context"
	"fmt"

	"github.com/ymohl-cl/tzktDelegator/internal/dto"
	"github.com/ymohl-cl/tzktDelegator/pkg/logger"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
	"github.com/ymohl-cl/tzktDelegator/pkg/tzkt"
)

// App connector
// Run, run the main process
// The main process is a loop that listen the tzkt new delegation
// and save them in the database
// To create a new app connector, use NewApp
type App interface {
	Run(ctx context.Context) error
}

type app struct {
	bdd        dto.Querier
	tzktDriver tzkt.TzKT
	logger     logger.Logger
}

func NewApp(pgDriver pgsql.PGSQL, tzktDriver tzkt.TzKT, l logger.Logger) App {
	return &app{
		bdd:        dto.New(pgDriver.Driver()),
		tzktDriver: tzktDriver,
		logger:     l,
	}
}

// Run, read message and check if type is a data or reorg message
func (a *app) Run(ctx context.Context) error {
	for {
		msg, err := a.tzktDriver.Listen()
		if err != nil {
			return err
		}

		switch msg.Type {
		case tzkt.MessageTypeData:
			if err = a.processDataMessage(ctx, msg); err != nil {
				return err
			}
		case tzkt.MessageTypeReorg:
			if err = a.processReorgMessage(ctx, msg); err != nil {
				return err
			}
		}
	}
}

func (a *app) processDataMessage(ctx context.Context, msg tzkt.Message) error {
	for _, data := range msg.Delegations {
		delegator := ""
		if data.Sender != nil {
			delegator = data.Sender.Address
		} else if data.Initiator != nil {
			delegator = data.Initiator.Address
		}

		_, err := a.bdd.InsertDelegator(ctx, dto.InsertDelegatorParams{
			DelegationDate:   data.Timestamp,
			DelegatorAddress: delegator,
			BlockHash:        data.Block,
			Amount:           data.Amount.String(),
			BlockState:       int64(msg.Blocks),
			ExternalID:       int64(data.ID),
		})
		if err != nil {
			a.logger.Error(err)
		}

		a.logger.WithField("delegation", fmt.Sprintf("%v", data)).Info("delegation message")
	}

	return nil
}

// processReorgMessage, delete all delegator from the last block
// to update the database
func (a *app) processReorgMessage(ctx context.Context, msg tzkt.Message) error {
	var err error
	var rowsID []int64

	if rowsID, err = a.bdd.DeleteDelegator(ctx, int32(msg.Blocks)); err != nil {
		return err
	}

	a.logger.WithField("nb rows deleted", fmt.Sprintf("%d", len(rowsID))).Info("reorg message")

	return nil
}
