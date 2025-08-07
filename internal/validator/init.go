// Package validator is a function for data validation
package validator

// Agent is validator agent interface
type Agent interface {
	IsFileExist(filePath string) error
	IsCSVExtension(filePath string) error
	IsCSVRot128Extension(filePath string) error
}

// agent is validator agent struct
type agent struct{}

// NewAgent creates a new validator agent
func NewAgent() Agent {
	return &agent{}
}
