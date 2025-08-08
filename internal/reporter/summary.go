package reporter

import (
	"context"
	"fmt"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
)

// PrintSummaryReport print result
func (a *agent) PrintSummaryReport(_ context.Context, data *entities.DonationSummary) {
	// Print summary report
	fmt.Println("done.")
	fmt.Printf("   Total transaction: THB %d\n", data.TotalCount)
	fmt.Printf("      Total received: THB %d\n", data.TotalReceived.IntPart())
	fmt.Printf("successfully donated: THB %d\n", data.SuccessfullyDonated.IntPart())
	fmt.Printf("     faulty donation: THB %d\n", data.FaultyDonation.IntPart())
	fmt.Println()
	fmt.Printf("  average per person: THB %.2f\n", data.AveragePerPerson.InexactFloat64())

	if len(data.TopDonors) <= 0 {
		return
	}
	fmt.Printf("          top donors: %s\n", data.TopDonors[0])
	for i := 1; i < len(data.TopDonors); i++ {
		fmt.Printf("		      %s\n", data.TopDonors[i])
	}
}
