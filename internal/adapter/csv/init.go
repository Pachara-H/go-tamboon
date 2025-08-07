// Package csv is a function for parse content data to struct
package csv

import (
	"context"

	"github.com/Pachara-H/go-tamboon/internal/domains/entities"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
)

// Parser is csv parser interface
type Parser interface {
	ParseCSV(ctx context.Context, content *utilities.SecureByte) ([]*entities.CardDetails, error)
}

// Parser is csv parser struct
type parser struct{}

// NewParser creates a new csv parser
func NewParser() Parser {
	return &parser{}
}
