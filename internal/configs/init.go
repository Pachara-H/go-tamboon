// Package configs is a environment config loader
package configs

// Loader interface for loading configuration
type Loader interface {
	LoadConfig() (*Config, error)
}

// NewLoader creates a new config loader
func NewLoader() Loader {
	return &loader{}
}
