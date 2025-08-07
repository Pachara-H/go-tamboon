package omise

import (
	"context"
	"log"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	omiseLib "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/shopspring/decimal"
)

func (c *client) Charge(ctx context.Context, name, cardNumber, cvv *utilities.SecureString, amount decimal.Decimal, expMonth, expYear int) (*entities.ChargeResult, error) {
	tokenID, err := c.Token(ctx, name, cardNumber, cvv, expMonth, expYear)
	if err != nil {
		return nil, err
	}

	result := &omiseLib.Charge{}
	if err := c.omiseClient.Do(result, &operations.CreateCharge{
		Amount:   amount.IntPart(),
		Currency: "THB",
		Card:     tokenID,
	}); err != nil {
		log.Printf("[ERROR]: charge error: %v\n", err)
		return nil, Error.NewInternalServerError(Code.FailChargeError)
	}

	return &entities.ChargeResult{
		Transaction:    result.Transaction,
		Status:         entities.ChargeStatus(result.Status),
		Amount:         decimal.NewFromInt(result.Amount),
		FailureMessage: result.FailureMessage,
	}, nil
}
