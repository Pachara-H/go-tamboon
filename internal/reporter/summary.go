package reporter

import "fmt"

// SummaryData holds summary data for reporting
type SummaryData struct {
	TotalReceived       int64
	SuccessfullyDonated int64
	FaultyDonation      int64
	AveragePerPerson    float64
	TopDonors           []string
}

// PrintSummaryReport print result
func (a *agent) PrintSummaryReport(data SummaryData) {
	// Print summary report
	fmt.Println("done.")
	fmt.Printf("Total received: %d\n", data.TotalReceived)
	fmt.Printf("successfully donated: %d\n", data.SuccessfullyDonated)
	fmt.Printf("faulty donation: %d\n", data.FaultyDonation)
	fmt.Printf("average per person: %f\n", data.AveragePerPerson)

	if len(data.TopDonors) <= 0 {
		return
	}
	fmt.Printf("top donors: %s\n", data.TopDonors[0])
	for i := 1; i < len(data.TopDonors); i++ {
		fmt.Printf("				%s\n", data.TopDonors[i])
	}
}
