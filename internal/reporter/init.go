// Package reporter is a function for parse content data to struct
package reporter

import (
	"context"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
)

// Agent is reporter agent interface
type Agent interface {
	PrintSummaryReport(ctx context.Context, data *entities.DonationSummary)
}

// agent is parser agent struct
type agent struct{}

// NewAgent creates a new parser agent
func NewAgent() Agent {
	return &agent{}
}
