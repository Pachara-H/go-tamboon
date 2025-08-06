// Package cipher is a function for data decryption/encryption
package cipher

// Agent is cipher agent interface
type Agent interface {
	Rot128Decrypt(cipherText []byte) ([]byte, error)
}

// agent is cipher agent struct
type agent struct{}

// NewAgent creates a new cipher agent
func NewAgent() Agent {
	return &agent{}
}
