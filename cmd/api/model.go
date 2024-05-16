package main

import (
	"time"

	"github.com/ymohl-cl/tzktDelegator/internal/dto"
)

const limitItemAPI = 100

type ResponseJSON struct {
	Data []TzKTDelegationJSON `json:"data"`
}

type ErrorResponseJSON struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type TzKTDelegationJSON struct {
	Date    time.Time `json:"timestamp"`
	Amount  string    `json:"amount"`
	Address string    `json:"delegator"`
	Block   string    `json:"block"`
}

func DelegationsDTOToModel(delegations []dto.TzktDelegation) []TzKTDelegationJSON {
	var data []TzKTDelegationJSON

	nDelegations := len(delegations)
	if nDelegations > limitItemAPI {
		nDelegations = limitItemAPI
	}
	data = make([]TzKTDelegationJSON, nDelegations)

	for i := 0; i < nDelegations; i++ {
		data[i].Address = delegations[i].DelegatorAddress
		data[i].Amount = delegations[i].Amount
		data[i].Block = delegations[i].BlockHash
		data[i].Date = delegations[i].DelegationDate
	}

	return data
}
