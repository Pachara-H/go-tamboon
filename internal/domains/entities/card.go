package entities

import (
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"github.com/shopspring/decimal"
)

// CardDetails represent user card details
type CardDetails struct {
	Name       *utilities.SecureString
	Amount     decimal.Decimal
	CardNumber *utilities.SecureString
	CVV        *utilities.SecureString
	ExpMonth   int
	ExpYear    int
}

// Clear remove secret value from memory
func (c *CardDetails) Clear() {
	c.Name.Clear()
	c.CardNumber.Clear()
	c.CVV.Clear()
}
