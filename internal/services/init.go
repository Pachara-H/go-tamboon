// Package services implement main business logic
package services

import (
	"github.com/Pachara-H/go-tamboon/internal/adapter/csv"
	"github.com/Pachara-H/go-tamboon/internal/adapter/omise"
	"github.com/Pachara-H/go-tamboon/internal/cipher"
	"github.com/Pachara-H/go-tamboon/internal/configs"
	"github.com/Pachara-H/go-tamboon/internal/domains"
	"github.com/Pachara-H/go-tamboon/internal/reporter"
	"github.com/Pachara-H/go-tamboon/internal/validator"
)

// donationService is service struct
type donationService struct {
	config *configs.Config

	// Infrastructure layer
	cipherAgent    cipher.Agent
	validatorAgent validator.Agent
	reporterAgent  reporter.Agent

	// Adapter layer
	omiseClient omise.Client
	csvParser   csv.Parser
}

// New creates a new service
func New(config *configs.Config, cipherAgent cipher.Agent, validatorAgent validator.Agent, reporterAgent reporter.Agent, omiseClient omise.Client, csvParser csv.Parser) domains.DonationService {
	return &donationService{
		config:         config,
		cipherAgent:    cipherAgent,
		validatorAgent: validatorAgent,
		reporterAgent:  reporterAgent,
		omiseClient:    omiseClient,
		csvParser:      csvParser,
	}
}
