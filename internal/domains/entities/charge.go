package entities

import (
	omiseLib "github.com/omise/omise-go"
	"github.com/shopspring/decimal"
)

// ChargeResult represent omise charge result data
type ChargeResult struct {
	Transaction    string
	Status         ChargeStatus
	Amount         decimal.Decimal
	FailureMessage *string
}

// ChargeStatus represents an enumeration of possible status of a Charge object.
type ChargeStatus string

// ChargeStatus can be one of the following list of constants:
const (
	ChargeFailed     ChargeStatus = ChargeStatus(omiseLib.ChargeFailed)
	ChargePending    ChargeStatus = ChargeStatus(omiseLib.ChargePending)
	ChargeSuccessful ChargeStatus = ChargeStatus(omiseLib.ChargeSuccessful)
	ChargeReversed   ChargeStatus = ChargeStatus(omiseLib.ChargeReversed)
)

// IsSuccessful returns true if charge was successful
func (c *ChargeResult) IsSuccessful() bool {
	return c.Status == ChargeSuccessful
}

// IsFailed returns true if charge was failed
func (c *ChargeResult) IsFailed() bool {
	return c.Status == ChargeFailed
}
