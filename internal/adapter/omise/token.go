package omise

import (
	"context"
	"log"
	"time"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	omiseLib "github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

func (c *client) Token(_ context.Context, name, cardNumber, cvv *utilities.SecureString, expMonth, expYear int) (string, error) {
	result := &omiseLib.Card{}
	if err := c.omiseClient.Do(result, &operations.CreateToken{
		Name:            name.String(),
		Number:          cardNumber.String(),
		ExpirationMonth: time.Month(expMonth),
		ExpirationYear:  expYear,
		SecurityCode:    cvv.String(),
	}); err != nil {
		log.Printf("[ERROR]: get token error: %v\n", err)
		return "", Error.NewInternalServerError(Code.FailGetTokenError)
	}

	return result.ID, nil
}
