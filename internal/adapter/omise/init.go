// Package omise is a function for integration with omise external layer
package omise

import (
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	omiseLib "github.com/omise/omise-go"
	"github.com/shopspring/decimal"
)

// Client is omise client interface
type Client interface {
	Charge(name, cardNumber, cvv *utilities.SecureString, amount decimal.Decimal, expMonth, expYear int) (*omiseLib.Charge, error)
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
