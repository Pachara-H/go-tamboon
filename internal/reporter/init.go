// Package reporter is a function for parse content data to struct
package reporter

// Agent is reporter agent interface
type Agent interface {
	PrintSummaryReport(data SummaryData)
}

// agent is parser agent struct
type agent struct{}

// NewAgent creates a new parser agent
func NewAgent() Agent {
	return &agent{}
}
