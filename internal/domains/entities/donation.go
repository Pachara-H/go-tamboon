// Package entities contains core business entities
package entities

import (
	"fmt"
	"time"

	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"github.com/shopspring/decimal"
	"go.openly.dev/pointy"
)

// DonationStatus represents the status of a donation
type DonationStatus string

const (
	// StatusPending represents a pending donation
	StatusPending DonationStatus = "pending"
	// StatusSuccess represents a success donation
	StatusSuccess DonationStatus = "success"
	// StatusFailed represents a failed donation
	StatusFailed DonationStatus = "failed"
)

// Donation represents a donation entity
type Donation struct {
	DonorName   *utilities.SecureString
	Amount      decimal.Decimal
	Currency    string
	Status      DonationStatus
	ChargeID    string
	ProcessedAt time.Time
	Error       *string
}

// NewDonation creates a new donation with pending status
func NewDonation(donorName *utilities.SecureString, amount decimal.Decimal) *Donation {
	return &Donation{
		DonorName: donorName,
		Amount:    amount,
		Currency:  "THB", // fixed currency
		Status:    StatusPending,
	}
}

// MarkAsSuccessful marks the donation as successful
func (d *Donation) MarkAsSuccessful(chargeID string) {
	d.Status = StatusSuccess
	d.ChargeID = chargeID
	d.ProcessedAt = time.Now()
	d.Error = nil
}

// MarkAsFailed marks the donation as failed with error message
func (d *Donation) MarkAsFailed(errorMsg string, chargeID ...string) {
	if len(chargeID) > 0 && chargeID[0] != "" {
		d.ChargeID = chargeID[0]
	}
	d.Status = StatusFailed
	d.Error = pointy.Pointer(errorMsg)
	d.ProcessedAt = time.Now()
}

// IsSuccessful returns true if donation was successful
func (d *Donation) IsSuccessful() bool {
	return d.Status == StatusSuccess
}

// IsFailed returns true if donation failed
func (d *Donation) IsFailed() bool {
	return d.Status == StatusFailed
}

// Details for logging
func (d *Donation) Details() string {
	return fmt.Sprintf("Name: %s, Amount: %s, Status: %s, ChargeID: %s, Error: %v", d.DonorName, d.Amount, d.Status, d.ChargeID, pointy.PointerValue(d.Error, ""))
}
