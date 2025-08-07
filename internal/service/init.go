// Package service implement main business logic
package service

import (
	"context"

	"github.com/Pachara-H/go-tamboon/internal/adapter/omise"
	"github.com/Pachara-H/go-tamboon/internal/cipher"
	"github.com/Pachara-H/go-tamboon/internal/configs"
	"github.com/Pachara-H/go-tamboon/internal/parser"
	"github.com/Pachara-H/go-tamboon/internal/reporter"
	"github.com/Pachara-H/go-tamboon/internal/validator"
)

// Worker is service interface
type Worker interface {
	ProcessDonations(ctx context.Context) error
}

// worker is service struct
type worker struct {
	config         *configs.Config
	cipherAgent    cipher.Agent
	parserAgent    parser.Agent
	validatorAgent validator.Agent
	reporterAgent  reporter.Agent
	omiseClient    omise.Client
}

// NewWorker creates a new service worker
func NewWorker(config *configs.Config, cipherAgent cipher.Agent, parserAgent parser.Agent, validatorAgent validator.Agent, reporterAgent reporter.Agent, omiseClient omise.Client) Worker {
	return &worker{
		config:         config,
		cipherAgent:    cipherAgent,
		parserAgent:    parserAgent,
		validatorAgent: validatorAgent,
		reporterAgent:  reporterAgent,
		omiseClient:    omiseClient,
	}
}
