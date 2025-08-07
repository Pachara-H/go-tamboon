package omise

import (
	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	omiseLib "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
	"github.com/shopspring/decimal"
)

func (c *client) Charge(_, cardNumber, _ *utilities.SecureString, amount decimal.Decimal, _, _ int) (*omiseLib.Charge, error) {
	omiseClient, err := omiseLib.NewClient(c.publicKey.String(), c.secretKey.String())
	if err != nil {
		return nil, Error.NewInternalServerError(Code.FailInitOmiseClient)
	}

	result := &omiseLib.Charge{}
	if err := omiseClient.Do(result, &operations.CreateCharge{
		Amount:   amount.IntPart(),
		Currency: "thb",
		Card:     cardNumber.String(),
	}); err != nil {
		return nil, Error.NewInternalServerError(Code.FailChargeError)
	}

	return result, nil
}
