package configs

import (
	"time"

	"github.com/Pachara-H/go-tamboon/pkg/constants"
	"github.com/Pachara-H/go-tamboon/pkg/utilities"
	"go.openly.dev/pointy"
)

// loader is struct for loading configuration
type loader struct{}

// Config holds all configuration for the application
type Config struct {
	Environment constants.Environment
	CSVFilePath string
	Timezone    string
	Omise       OmiseConfig
}

// OmiseConfig holds Omise API configuration
type OmiseConfig struct {
	PublicKey string
	SecretKey string
	BaseURL   string
	Timeout   time.Duration
}

// LoadConfig loads configuration from environment variables
func (e *loader) LoadConfig() (*Config, error) {
	config := &Config{
		Environment: constants.Environment(utilities.GetEnvCfgStringOrDefault("ENV", constants.Local.String())),
		Timezone:    utilities.GetEnvCfgStringOrDefault("TZ", "Asia/Bangkok"),
	}

	// Load Omise configuration
	config.Omise = pointy.PointerValue(e.loadOmiseConfig(), OmiseConfig{})

	return config, nil
}

func (e *loader) loadOmiseConfig() *OmiseConfig {
	return &OmiseConfig{
		PublicKey: utilities.GetEnvCfgStringOrDefault("OMISE_PUBLIC_KEY"),
		SecretKey: utilities.GetEnvCfgStringOrDefault("OMISE_SECRET_KEY"),
		BaseURL:   utilities.GetEnvCfgStringOrDefault("OMISE_BASE_URL"),
		Timeout:   time.Duration(utilities.GetEnvCfgInt64OrDefault("OMISE_TIMEOUT") * int64(time.Second)),
	}
}
