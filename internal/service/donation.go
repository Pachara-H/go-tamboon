package service

import (
	"context"
	"log"

	"github.com/Pachara-H/go-tamboon/internal/reporter"
	"github.com/omise/omise-go"
	"github.com/shopspring/decimal"
)

// ProcessDonations implement donation logic
func (w *worker) ProcessDonations(ctx context.Context) error {
	// 0. defer recover
	defer func() {
		if x := recover(); x != nil {
			log.Fatalf("panic err: %+v", x)
		}
	}()

	// 1. validate file
	if err := w.validatorAgent.IsFileExist(w.config.CSVFilePath); err != nil {
		return err
	}
	if err := w.validatorAgent.IsCSVExtension(w.config.CSVFilePath); err != nil {
		return err
	}

	// 2. readfile and decrypt
	content, err := w.cipherAgent.Rot128DecryptFileContent(ctx, w.config.CSVFilePath)
	if err != nil {
		return err
	}
	defer content.Clear()

	// 3. parse csv
	records, err := w.parserAgent.ConvertCSV(ctx, content)
	if err != nil {
		return err
	}
	defer w.parserAgent.ClearCSVData(ctx, records)

	// 4. call charge
	var summary reporter.SummaryData
	for _, r := range records {
		summary.TotalReceived.Add(r.Amount)

		result, err := w.omiseClient.Charge(ctx, r.Name, r.CardNumber, r.CVV, r.Amount, r.ExpMonth, r.ExpYear)
		if err != nil || result.Status != omise.ChargeSuccessful {
			summary.FaultyDonation.Add(r.Amount)
			continue
		}

		summary.SuccessfullyDonated.Add(decimal.NewFromInt(result.Amount))
	}
	summary.AveragePerPerson = summary.SuccessfullyDonated.Div(decimal.NewFromInt(int64(len(records))))

	// 5. print summary report
	w.reporterAgent.PrintSummaryReport(ctx, summary)
	return nil
}
