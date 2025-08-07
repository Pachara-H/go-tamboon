package reporter

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// SummaryData holds summary data for reporting
type SummaryData struct {
	TotalReceived       decimal.Decimal
	SuccessfullyDonated decimal.Decimal
	FaultyDonation      decimal.Decimal
	AveragePerPerson    decimal.Decimal
	TopDonors           []string
}

// PrintSummaryReport print result
func (a *agent) PrintSummaryReport(data SummaryData) {
	// Print summary report
	fmt.Println("done.")
	fmt.Printf("Total received: %d\n", data.TotalReceived.IntPart())
	fmt.Printf("successfully donated: %d\n", data.SuccessfullyDonated.IntPart())
	fmt.Printf("faulty donation: %d\n", data.FaultyDonation.IntPart())
	fmt.Printf("average per person: %.2f\n", data.AveragePerPerson.InexactFloat64())

	if len(data.TopDonors) <= 0 {
		return
	}
	fmt.Printf("top donors: %s\n", data.TopDonors[0])
	for i := 1; i < len(data.TopDonors); i++ {
		fmt.Printf("				%s\n", data.TopDonors[i])
	}
}
