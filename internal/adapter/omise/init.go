// Package omise is a function for integration with omise external layer
package omise

import (
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"github.com/shopspring/decimal"
)

// Client is omise client interface
type Client interface {
	Charge(_, cardNumber, _ *utilities.SecureString, amount decimal.Decimal, _, _ int) error
}

// client is omise client struct
type client struct {
	publicKey *utilities.SecureString
	secretKey *utilities.SecureString
}

// NewClient creates a new omise client
func NewClient(publicKey, secretKey *utilities.SecureString) Client {
	return &client{
		publicKey: publicKey,
		secretKey: secretKey,
	}
}
