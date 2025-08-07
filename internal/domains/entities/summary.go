package entities

import (
	"sort"

	"github.com/shopspring/decimal"
)

// DonationSummary represents the summary of donation processing
type DonationSummary struct {
	TotalReceived       decimal.Decimal
	SuccessfullyDonated decimal.Decimal
	FaultyDonation      decimal.Decimal
	AveragePerPerson    decimal.Decimal
	TopDonors           []string
	TotalCount          int
	SuccessfulCount     int
	FailedCount         int
}

// NewDonationSummary creates a new donation summary
func NewDonationSummary() *DonationSummary {
	return &DonationSummary{
		TotalReceived:       decimal.Zero,
		SuccessfullyDonated: decimal.Zero,
		FaultyDonation:      decimal.Zero,
		AveragePerPerson:    decimal.Zero,
		TopDonors:           make([]string, 0),
	}
}

// AddDonation adds a donation to the summary
func (s *DonationSummary) AddDonation(donation *Donation) {
	s.TotalCount++
	s.TotalReceived = s.TotalReceived.Add(donation.Amount)

	if donation.IsSuccessful() {
		s.SuccessfulCount++
		s.SuccessfullyDonated = s.SuccessfullyDonated.Add(donation.Amount)
	} else if donation.IsFailed() {
		s.FailedCount++
		s.FaultyDonation = s.FaultyDonation.Add(donation.Amount)
	}
}

// CalculateAveragePerPerson calculates the average donation per person
func (s *DonationSummary) CalculateAveragePerPerson() {
	if s.SuccessfulCount > 0 {
		s.AveragePerPerson = s.SuccessfullyDonated.Div(decimal.NewFromInt(int64(s.SuccessfulCount)))
	}
}

// GenerateTopDonors generates top donors list from donations
func (s *DonationSummary) GenerateTopDonors(donations []*Donation, limit int) {
	// Create a map to aggregate donations by donor name
	donorAmounts := make(map[string]decimal.Decimal)

	for _, donation := range donations {
		if donation.IsSuccessful() {
			if existing, exists := donorAmounts[donation.DonorName.String()]; exists {
				donorAmounts[donation.DonorName.String()] = existing.Add(donation.Amount)
			} else {
				donorAmounts[donation.DonorName.String()] = donation.Amount
			}
		}
	}

	// Convert to slice for sorting
	type donorAmount struct {
		name   string
		amount decimal.Decimal
	}

	donors := make([]donorAmount, 0, len(donorAmounts))
	for name, amount := range donorAmounts {
		donors = append(donors, donorAmount{name: name, amount: amount})
	}

	// Sort by amount (descending)
	sort.Slice(donors, func(i, j int) bool {
		return donors[i].amount.GreaterThan(donors[j].amount)
	})

	// Extract names and limit results
	topDonors := make([]string, 0, limit)
	for i, donor := range donors {
		if i >= limit {
			break
		}
		topDonors = append(topDonors, donor.name)
	}

	s.TopDonors = topDonors
}
