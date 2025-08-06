// Package utilities is a common function for general purpose
package utilities

import (
	"os"
	"strconv"
)

// GetEnvCfgStringOrDefault read environment variable and return as string type
func GetEnvCfgStringOrDefault(key string, defaultValue ...string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return ""
}

// GetEnvCfgInt64OrDefault read environment variable and return as int64 type
func GetEnvCfgInt64OrDefault(key string, defaultValue ...int64) int64 {
	var defValue int64
	if len(defaultValue) > 0 {
		defValue = defaultValue[0]
	}

	if value := os.Getenv(key); value != "" {
		valueInt64, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return defValue
		}
		return valueInt64
	}
	return defValue
}
