// Package parser is a function for parse content data to struct
package parser

import (
	"context"

	"github.com/Pachara-H/go-tamboon/pkg/utilities"
)

// Agent is parser agent interface
type Agent interface {
	ConvertCSV(ctx context.Context, content *utilities.SecureByte) (CSVRowsData, error)
	ClearCSVData(ctx context.Context, data CSVRowsData)
}

// agent is parser agent struct
type agent struct{}

// NewAgent creates a new parser agent
func NewAgent() Agent {
	return &agent{}
}
