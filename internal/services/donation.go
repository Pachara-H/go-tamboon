package services

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
	"go.openly.dev/pointy"
)

// ProcessDonations implement donation logic
func (w *donationService) ProcessDonations(ctx context.Context) error { //nolint
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
	records = records[:20] // for test only
	defer func() {
		for _, r := range records {
			r.Clear()
		}
	}()

	// 4. call charge with goroutine
	donations := make([]*entities.Donation, 0)
	summary := entities.NewDonationSummary()

	var wg sync.WaitGroup
	wg.Add(len(records))

	ch := make(chan *entities.Donation, len(records))

	sleepTime := 200 * time.Millisecond
	for _, r := range records {
		go func(c chan *entities.Donation, r *entities.CardDetails) {
			defer wg.Done()

			select {
			case <-ctx.Done():
				return
			default:
				// do nothing
			}

			donation := entities.NewDonation(r.Name, r.Amount)

			result, err := w.omiseClient.Charge(ctx, r.Name, r.CardNumber, r.CVV, r.Amount, r.ExpMonth, r.ExpYear)
			if err != nil {
				if strings.Contains(err.Error(), "too_many_requests") {
					sleepTime += 100 * time.Millisecond // increase sleep time
				}
				donation.MarkAsFailed(err.Error())
			} else if result.Status != entities.ChargeSuccessful {
				donation.MarkAsFailed(pointy.PointerValue(result.FailureMessage, ""), result.Transaction)
			} else {
				donation.MarkAsSuccessful(result.Transaction)
			}

			select {
			case <-ctx.Done():
				return
			default:
				c <- donation
			}
		}(ch, r)

		// for prevent api rate limit
		time.Sleep(sleepTime)
	}

	// 5. wait go routine all tasks
	err = handleWaitGroup(ctx, &wg, ch)
	if err != nil {
		return err
	}

	// 6. print summary report
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		for d := range ch {
			donations = append(donations, d)
			summary.AddDonation(d)
		}
		summary.CalculateAveragePerPerson()
		summary.GenerateTopDonors(donations, 3)
		w.reporterAgent.PrintSummaryReport(ctx, summary)
	}

	return nil
}

func handleWaitGroup(ctx context.Context, wg *sync.WaitGroup, ch chan *entities.Donation) error {
	allDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(allDone)
	}()

	select {
	case <-ctx.Done():
		close(ch)
		return nil
	case <-allDone:
		close(ch)
		return nil
	case <-time.After(30 * time.Second):
		// Cancel the context to signal all goroutines and database operations to stop
		close(ch)
		log.Println("[ERROR] Go routine process timeout")
		return Error.NewInternalServerError(Code.FailGoroutineTimeout)
	}
}
