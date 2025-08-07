// Package cipher is a function for data decryption/encryption
package cipher

import "github.com/Pachara-H/go-tamboon/pkg/utilities"

// Agent is cipher agent interface
type Agent interface {
	Rot128DecryptFileContent(path string) (*utilities.SecureByte, error)
}

// agent is cipher agent struct
type agent struct{}

// NewAgent creates a new cipher agent
func NewAgent() Agent {
	return &agent{}
}
