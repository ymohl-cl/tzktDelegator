package delegator

import (
	"time"
)

// DelegationFilter to get delegations with criterias
type DelegationFilter struct {
	Year int32
}

// Delegation details of a delegation
type Delegation struct {
	Identifier int64
	Date       time.Time
	Address    string
	Block      []byte
	Amount     float64
}
