// Package csv provides CSV parsing adapter that converts CSV data to domain entities
package csv

import (
	"bytes"
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"github.com/shopspring/decimal"
)

// ParseCSV converts CSV content to domain entities
func (p *parser) ParseCSV(ctx context.Context, content *utilities.SecureByte) ([]*entities.CardDetails, error) { //nolint
	if content.IsEmpty() {
		log.Print("[ERROR]: content of decrypted file is empty")
		return nil, Error.NewNotFoundError(Code.FailEmptyCSVContent)
	}

	reader := csv.NewReader(bytes.NewReader(content.Bytes()))

	var cardDetails []*entities.CardDetails
	var isHeaderExists bool
	var nameIdx, amountIdx, cardNumberIdx, cvvIdx, expMonthIdx, expYearIdx = -1, -1, -1, -1, -1, -1

	start := time.Now()
	for {
		// Check context cancellation
		select {
		case <-ctx.Done():
			log.Println("[ERROR]: context was cancelled")
			return nil, Error.NewInternalServerError(Code.FailContextCancel)
		default:
		}

		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			log.Print("[ERROR]: read data record failed")
			return nil, Error.NewInternalServerError(Code.FailReadingCSVRecord)
		}

		// Timeout protection
		if time.Since(start) > 300*time.Second {
			log.Print("[ERROR]: timeout of reading file data reach")
			return nil, Error.NewInternalServerError(Code.FailReadingCSVTimeout)
		}

		if len(record) <= 0 {
			continue
		}

		// Parse header
		if !isHeaderExists {
			isHeaderExists = true

			if err = p.parseHeader(record, &nameIdx, &amountIdx, &cardNumberIdx, &cvvIdx, &expMonthIdx, &expYearIdx); err != nil {
				return nil, err
			}
			continue
		}

		// Parse data row
		cardDetail, err := p.parseDataRow(record, nameIdx, amountIdx, cardNumberIdx, cvvIdx, expMonthIdx, expYearIdx)
		if err != nil {
			return nil, err // not allow if some data is invalid
		}

		cardDetails = append(cardDetails, cardDetail)
	}

	return cardDetails, nil
}

// parseHeader parses CSV header and finds column indices
func (p *parser) parseHeader(record []string, nameIdx, amountIdx, cardNumberIdx, cvvIdx, expMonthIdx, expYearIdx *int) error { //nolint
	for i, h := range record {
		switch {
		case strings.EqualFold(h, "Name"):
			*nameIdx = i
		case strings.EqualFold(h, "AmountSubunits"):
			*amountIdx = i
		case strings.EqualFold(h, "CCNumber"):
			*cardNumberIdx = i
		case strings.EqualFold(h, "CVV"):
			*cvvIdx = i
		case strings.EqualFold(h, "ExpMonth"):
			*expMonthIdx = i
		case strings.EqualFold(h, "ExpYear"):
			*expYearIdx = i
		}
	}

	// Validate all required columns are present
	if *nameIdx == -1 || *amountIdx == -1 || *cardNumberIdx == -1 || *cvvIdx == -1 || *expMonthIdx == -1 || *expYearIdx == -1 {
		log.Print("[ERROR]: missing some data set")
		return Error.NewNotFoundError(Code.FailMissingCSVColumnName)
	}

	return nil
}

// parseDataRow parses a single CSV data row into domain entities
func (p *parser) parseDataRow(record []string, nameIdx, amountIdx, cardNumberIdx, cvvIdx, expMonthIdx, expYearIdx int) (*entities.CardDetails, error) {
	// Parse donor name
	donorName := utilities.NewSecureString(strings.TrimSpace(record[nameIdx]))
	if donorName.Len() <= 0 {
		log.Print("[ERROR]: some name data is empty")
		return nil, Error.NewInternalServerError(Code.FailConvertingCSVName)
	}

	// Parse amount
	amountStr := strings.TrimSpace(record[amountIdx])
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Print("[ERROR]: some amount data is not numeric")
		return nil, Error.NewInternalServerError(Code.FailConvertingCSVAmount)
	}

	// Validate amount is not negative and is integer (no decimal places for subunits)
	if amount.IsNegative() || amount.Exponent() < 0 {
		log.Print("[ERROR]: some amount data is negative or not subunit")
		return nil, Error.NewInternalServerError(Code.FailConvertingCSVAmount)
	}

	// Parse card details
	cardNumber := utilities.NewSecureString(strings.TrimSpace(record[cardNumberIdx]))

	cvv := utilities.NewSecureString(strings.TrimSpace(record[cvvIdx]))

	expMonth, err := strconv.Atoi(strings.TrimSpace(record[expMonthIdx]))
	if err != nil || expMonth < 1 || expMonth > 12 {
		log.Print("[ERROR]: some card expiry month data is invalid")
		return nil, Error.NewInternalServerError(Code.FailConvertingCSVExpMonth)
	}

	expYear, err := strconv.Atoi(strings.TrimSpace(record[expYearIdx]))
	if err != nil {
		log.Print("[ERROR]: some card expiry year data is invalid")
		return nil, Error.NewInternalServerError(Code.FailConvertingCSVExpYear)
	}

	return &entities.CardDetails{
		Name:       donorName,
		Amount:     amount,
		CardNumber: cardNumber,
		CVV:        cvv,
		ExpMonth:   expMonth,
		ExpYear:    expYear,
	}, nil
}
