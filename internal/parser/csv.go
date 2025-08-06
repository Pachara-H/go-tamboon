package parser

import (
	"bytes"
	"encoding/csv"
	"errors"
	"io"
	"strconv"
	"strings"
	"time"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"github.com/shopspring/decimal"
)

// CSVRowData store csv row data
type CSVRowData struct {
	Name       *utilities.SecureString
	Amount     decimal.Decimal
	CardNumber *utilities.SecureString
	CVV        *utilities.SecureString
	ExpMonth   int
	ExpYear    int
}

// CSVRowsData store csv rows data
type CSVRowsData []CSVRowData

// ConvertCSV convert .csv content to struct type
func (a *agent) ConvertCSV(content *utilities.SecureByte) (CSVRowsData, error) { //nolint
	if content.IsEmpty() {
		return nil, Error.NewNotFoundError(Code.FailEmptyCSVContent)
	}

	reader := csv.NewReader(bytes.NewReader(content.Bytes()))

	var rows CSVRowsData
	var isHeaderExits bool
	var nameIdx, amountIdx, cardNumberIdx, cvvIdx, expMonthIdx, expYearIdx = -1, -1, -1, -1, -1, -1

	start := time.Now()
	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, Error.NewInternalServerError(Code.FailReadingCSVRecord)
		}
		if time.Since(start) > 30*time.Second { // Timeout after 30 seconds to prevent infinite loop
			return nil, Error.NewInternalServerError(Code.FailReadingCSVTimeout)
		}

		if len(record) <= 0 {
			continue
		}

		if !isHeaderExits {
			// Find the index of each columns
			// assume first not empty row is header
			isHeaderExits = true
			for i, h := range record {
				switch {
				case strings.EqualFold(h, "Name"):
					nameIdx = i
				case strings.EqualFold(h, "AmountSubunits"):
					amountIdx = i
				case strings.EqualFold(h, "CCNumber"):
					cardNumberIdx = i
				case strings.EqualFold(h, "CVV"):
					cvvIdx = i
				case strings.EqualFold(h, "ExpMonth"):
					expMonthIdx = i
				case strings.EqualFold(h, "ExpYear"):
					expYearIdx = i
				}
			}
			if nameIdx == -1 || amountIdx == -1 || cardNumberIdx == -1 || cvvIdx == -1 || expMonthIdx == -1 || expYearIdx == -1 {
				return nil, Error.NewNotFoundError(Code.FailMissingCSVColumnName)
			}

			continue
		}

		name := utilities.NewSecureString(record[nameIdx])
		a, err := decimal.NewFromString(record[amountIdx])
		if err != nil {
			return nil, Error.NewInternalServerError(Code.FailConvertingCSVAmount)
		}
		amount, err := decimal.NewFromString(a.String()) // for remove zero tailing in digits
		if err != nil || amount.IsNegative() || -amount.Exponent() > 0 {
			return nil, Error.NewInternalServerError(Code.FailConvertingCSVAmount)
		}
		cardNumber := utilities.NewSecureString(record[cardNumberIdx])
		cvv := utilities.NewSecureString(record[cvvIdx])
		if cvv.Len() != 3 {
			return nil, Error.NewInternalServerError(Code.FailConvertingCSVCVV)
		}
		expMonth, err := strconv.Atoi(record[expMonthIdx])
		if err != nil || expMonth < 1 || expMonth > 12 {
			return nil, Error.NewInternalServerError(Code.FailConvertingCSVExpMonth)
		}
		expYear, err := strconv.Atoi(record[expYearIdx])
		if err != nil {
			return nil, Error.NewInternalServerError(Code.FailConvertingCSVExpYear)
		}

		rows = append(rows, CSVRowData{
			Name:       name,
			Amount:     amount,
			CardNumber: cardNumber,
			CVV:        cvv,
			ExpMonth:   expMonth,
			ExpYear:    expYear,
		})
	}

	return rows, nil
}

// ClearCSVData set pointer variable to null
func (a *agent) ClearCSVData(data CSVRowsData) {
	for _, d := range data {
		d.Name.Clear()
		d.CardNumber.Clear()
		d.CVV.Clear()
	}
}
