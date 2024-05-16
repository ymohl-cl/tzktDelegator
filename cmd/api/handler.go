package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ymohl-cl/tzktDelegator/internal/dto"
	"github.com/ymohl-cl/tzktDelegator/internal/service/delegator"
	"github.com/ymohl-cl/tzktDelegator/pkg/logger"
	"github.com/ymohl-cl/tzktDelegator/pkg/pgsql"
)

const (
	YearQueryParamKey = "year"
)

type Handler interface {
	GetDelegations(ctx echo.Context) error
}

type handler struct {
	logger           logger.Logger
	delegatorService delegator.DelegatorService
}

func NewHandler(l logger.Logger, pg pgsql.PGSQL) Handler {
	return &handler{
		logger:           l,
		delegatorService: delegator.New(pg),
	}
}

type GetDelegationsQueryParams struct {
	Year int32 `json:"year"`
}

func parseGetDelegationsQueryParams(ctx echo.Context) (GetDelegationsQueryParams, error) {
	yearSTR := ctx.QueryParam(YearQueryParamKey)
	if yearSTR == "" {
		return GetDelegationsQueryParams{
			Year: 0,
		}, nil
	}

	year, err := strconv.Atoi(yearSTR)
	if err != nil {
		return GetDelegationsQueryParams{}, err
	}

	return GetDelegationsQueryParams{
		Year: int32(year),
	}, nil
}

func (h handler) GetDelegations(ctx echo.Context) error {
	var delegations []dto.TzktDelegation
	var err error

	params, err := parseGetDelegationsQueryParams(ctx)
	if err != nil {
		h.logger.Warn(err.Error())

		return ctx.JSON(http.StatusBadRequest, &ErrorResponseJSON{
			Message: "failed to parse query params, year should be valid",
		})
	}

	if delegations, err = h.delegatorService.Delegations(
		ctx.Request().Context(),
		delegator.DelegationFilter{
			Year: params.Year,
		},
	); err != nil {
		h.logger.
			WithField("year", fmt.Sprintf("%d", params.Year)).
			WithField("function", "delegatorService.Delegations").
			Error(err)
		return err
	}

	return ctx.JSON(http.StatusOK, &ResponseJSON{
		Data: DelegationsDTOToModel(delegations),
	})
}
