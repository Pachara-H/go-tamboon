package services

import (
	"context"
	"log"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	"go.openly.dev/pointy"
)

// ProcessDonations implement donation logic
func (w *donationService) ProcessDonations(ctx context.Context) error {
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
	if err := w.validatorAgent.IsCSVRot128Extension(w.config.CSVFilePath); err != nil {
		return err
	}

	// 2. readfile and decrypt
	content, err := w.cipherAgent.Rot128DecryptFileContent(ctx, w.config.CSVFilePath)
	if err != nil {
		return err
	}
	defer content.Clear()

	// 3. parse csv
	records, err := w.csvParser.ParseCSV(ctx, content)
	if err != nil {
		return err
	}
	records = records[:10] // for test only
	defer func() {
		for _, r := range records {
			r.Clear()
		}
	}()

	// 4. call charge
	donations := []*entities.Donation{}
	summary := entities.NewDonationSummary()

	for _, r := range records {
		donation := entities.NewDonation(r.Name, r.Amount)

		result, err := w.omiseClient.Charge(ctx, r.Name, r.CardNumber, r.CVV, r.Amount, r.ExpMonth, r.ExpYear)
		if err != nil {
			donation.MarkAsFailed(err.Error())
		} else if result.Status != entities.ChargeSuccessful {
			donation.MarkAsFailed(pointy.PointerValue(result.FailureMessage, ""), result.Transaction)
		} else {
			donation.MarkAsSuccessful(result.Transaction)
		}

		log.Println(donation.Details())

		donations = append(donations, donation)
		summary.AddDonation(donation)
	}

	// 5. print summary report
	summary.CalculateAveragePerPerson()
	summary.GenerateTopDonors(donations, 3)
	w.reporterAgent.PrintSummaryReport(ctx, summary)
	return nil
}
