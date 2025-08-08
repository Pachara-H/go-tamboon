package configs

import (
	"encoding/base64"
	"log"

	Code "github.com/Pachara-H/go-tamboon/internal/errorcode"
	"github.com/Pachara-H/go-tamboon/pkg/constants"
	Error "github.com/Pachara-H/go-tamboon/pkg/errors"
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
	PublicKey *utilities.SecureString
	SecretKey *utilities.SecureString
}

// LoadConfig loads configuration from environment variables
func (e *loader) LoadConfig() (*Config, error) {
	config := &Config{
		Environment: constants.Environment(utilities.GetEnvCfgStringOrDefault("ENV", constants.Local.String())),
		Timezone:    utilities.GetEnvCfgStringOrDefault("TZ", "Asia/Bangkok"),
	}

	// Load Omise configuration
	omiseCfg, err := e.loadOmiseConfig()
	if err != nil {
		return nil, err
	}
	config.Omise = pointy.PointerValue(omiseCfg, OmiseConfig{})

	return config, nil
}

func (e *loader) loadOmiseConfig() (*OmiseConfig, error) {
	pKeyByte, err := base64.StdEncoding.DecodeString(utilities.GetEnvCfgStringOrDefault("OMISE_PUBLIC_KEY"))
	if err != nil {
		log.Println("[ERROR]: load Omise public key failed")
		return nil, Error.NewInternalServerError(Code.FailToLoadOmiseConfigPublicKey)
	}
	sKeyByte, err := base64.StdEncoding.DecodeString(utilities.GetEnvCfgStringOrDefault("OMISE_SECRET_KEY"))
	if err != nil {
		log.Println("[ERROR]: load Omise secret key failed")
		return nil, Error.NewInternalServerError(Code.FailToLoadOmiseConfigSecretKey)
	}

	return &OmiseConfig{
		PublicKey: utilities.NewSecureStringFromByte(pKeyByte),
		SecretKey: utilities.NewSecureStringFromByte(sKeyByte),
	}, nil
}

// ClearConfig set pointer variable to null
func (e *loader) ClearConfig(cfg *Config) {
	cfg.Omise.PublicKey.Clear()
	cfg.Omise.SecretKey.Clear()
}
