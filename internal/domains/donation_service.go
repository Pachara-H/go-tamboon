// Package domains contains interface for all service
package domains

import "context"

// DonationService is service interface
type DonationService interface {
	ProcessDonations(ctx context.Context) error
}
