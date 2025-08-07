// Package omise is a function for integration with omise external layer
package omise

import (
	"context"
	"log"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	omiseLib "github.com/omise/omise-go"
	"github.com/shopspring/decimal"
)

// Client is omise client interface
type Client interface {
	Token(ctx context.Context, name, cardNumber, cvv *utilities.SecureString, expMonth, expYear int) (string, error)
	Charge(ctx context.Context, name, cardNumber, cvv *utilities.SecureString, amount decimal.Decimal, expMonth, expYear int) (*entities.ChargeResult, error)
}

// client is omise client struct
type client struct {
	omiseClient *omiseLib.Client
}

// NewClient creates a new omise client
func NewClient(publicKey, secretKey *utilities.SecureString) (Client, error) {
	c, err := omiseLib.NewClient(publicKey.String(), secretKey.String())
	if err != nil {
		log.Printf("[ERROR]: initial omise client failed: %v\n", err)
		return nil, Error.NewInternalServerError(Code.FailInitOmiseClient)
	}

	return &client{
		omiseClient: c,
	}, nil
}
