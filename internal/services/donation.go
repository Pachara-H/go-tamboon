package services

import (
	"context"
	"log"
	"strings"
	"sync"
	"sync/atomic"
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
			log.Printf("panic err: %+v\n", x)
		}
	}()

	processingCtx, cancelProcessing := context.WithCancel(ctx)
	defer cancelProcessing()

	// 1. validate file
	if err := w.validatorAgent.IsFileExist(w.config.CSVFilePath); err != nil {
		return err
	}
	if err := w.validatorAgent.IsCSVRot128Extension(w.config.CSVFilePath); err != nil {
		return err
	}

	// 2. readfile and decrypt
	content, err := w.cipherAgent.Rot128DecryptFileContent(processingCtx, w.config.CSVFilePath)
	if err != nil {
		return err
	}
	defer content.Clear()

	// 3. parse csv
	records, err := w.csvParser.ParseCSV(processingCtx, content)
	if err != nil {
		return err
	}
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
	defer close(ch)

	sleepTime := 200 * time.Millisecond
	errCount := int32(0) // prevent blocking from 3rd party

	for idx, r := range records {
		go func(c chan *entities.Donation, r *entities.CardDetails) {
			donation := entities.NewDonation(r.Name, r.Amount)

			defer func() {
				if x := recover(); x != nil {
					log.Printf("panic err: %+v\n", x)
					donation.MarkAsFailed(Error.NewInternalServerError(Code.FailPanic).Error())
					c <- donation
				}
				wg.Done()
			}()

			select {
			case <-processingCtx.Done():
				donation.MarkAsFailed(Error.NewInternalServerError(Code.FailContextCancel).Error())
				c <- donation
				return
			default:
				// do nothing
			}

			result, err := w.omiseClient.Charge(processingCtx, r.Name, r.CardNumber, r.CVV, r.Amount, r.ExpMonth, r.ExpYear)
			if err != nil {
				newCount := atomic.AddInt32(&errCount, 1)

				if strings.Contains(err.Error(), "too_many_requests") {
					atomic.StoreInt64((*int64)(&sleepTime), int64(sleepTime*2)) // exponential increase sleep time
				}
				if newCount > 30 {
					log.Println("[ERROR]: too many error was occurred when charge")
					log.Println("[ERROR]: stopping all remaining tasks...")
					log.Println("[ERROR]: please waiting for final summary...")
					cancelProcessing()
				}
				donation.MarkAsFailed(err.Error())
			} else if result.Status != entities.ChargeSuccessful {
				donation.MarkAsFailed(pointy.PointerValue(result.FailureMessage, ""), result.Transaction)
			} else {
				atomic.StoreInt32(&errCount, 0) // reset count
				donation.MarkAsSuccessful(result.Transaction)
				log.Printf("[SUCCESS]: charge transaction Seq.%d was success with amount: %s", idx, result.Amount)
			}

			c <- donation
		}(ch, r)

		// for prevent api rate limit
		time.Sleep(sleepTime)
	}

	// 5. wait go routine all tasks
	err = handleWaitGroup(processingCtx, cancelProcessing, &wg)
	if err != nil {
		return err
	}

	// 6. finally print summary report
	length := len(ch)
	for i := 0; i < length; i++ {
		d, ok := <-ch
		if !ok {
			break
		}
		donations = append(donations, d)
		summary.AddDonation(d)
	}
	summary.CalculateAveragePerPerson()
	summary.GenerateTopDonors(donations, 3)
	w.reporterAgent.PrintSummaryReport(processingCtx, summary)

	return nil
}

func handleWaitGroup(processingCtx context.Context, cancelProcessing context.CancelFunc, wg *sync.WaitGroup) error {
	allDone := make(chan struct{})
	go func() {
		wg.Wait()
		close(allDone)
	}()

	select {
	case <-processingCtx.Done():
		return nil
	case <-allDone:
		return nil
	case <-time.After(30 * time.Second):
		// Cancel the context to signal all goroutines and database operations to stop
		cancelProcessing()
		log.Println("[ERROR] Go routine process timeout")
		return Error.NewInternalServerError(Code.FailGoroutineTimeout)
	}
}
