// Package constants is a common constant variables
package constants

// Environment is name of running environment
type Environment string

const (
	// Local is local environment name
	Local Environment = "local"
	// Dev is development environment name
	Dev Environment = "development"
	// Test is testing environment name
	Test Environment = "testing"
	// Prod is production environment name
	Prod Environment = "production"
)

// String convert Environment type to string
func (e Environment) String() string {
	return string(e)
}
