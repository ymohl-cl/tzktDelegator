package delegator

import (
	"context"

	"github.com/ymohl-cl/tzktDelegator/internal/dto"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
)

const (
	defaultLimitItem = 100
)

type DelegatorService interface {
	Delegations(ctx context.Context, filter DelegationFilter) ([]dto.TzktDelegation, error)
	Write(ctx context.Context, delegation dto.TzktDelegation) error
}

type delegator struct {
	bdd dto.Querier
}

func New(driver pgsql.PGSQL) DelegatorService {
	return &delegator{
		bdd: dto.New(driver.Driver()),
	}
}

func (d delegator) Delegations(ctx context.Context, filter DelegationFilter) ([]dto.TzktDelegation, error) {
	var data []dto.TzktDelegation
	var err error

	dtoParams := dto.SearchDelegatorParams{
		LimitItem: defaultLimitItem,
	}
	if filter.Year != 0 {
		dtoParams.DelegationYear = filter.Year
	}

	if data, err = d.bdd.SearchDelegator(ctx, dtoParams); err != nil {
		return nil, err
	}

	return data, nil
}

func (d delegator) Write(ctx context.Context, delegation dto.TzktDelegation) error {
	var err error

	if _, err = d.bdd.InsertDelegator(ctx, dto.InsertDelegatorParams{
		DelegationDate:   delegation.DelegationDate,
		DelegatorAddress: delegation.DelegatorAddress,
		BlockHash:        delegation.BlockHash,
		Amount:           delegation.Amount,
	}); err != nil {
		return err
	}

	return nil
}
