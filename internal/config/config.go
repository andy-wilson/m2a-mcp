package config

import (
	"fmt"
	"os"
)

// Config holds the configuration for the M2A MCP service
type Config struct {
	APIKey       string
	BaseURL      string
	AWSAccountID string
}

// Load reads configuration from environment variables
func Load() (*Config, error) {
	apiKey := os.Getenv("M2A_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("M2A_API_KEY environment variable is required")
	}

	baseURL := os.Getenv("M2A_BASE_URL")
	if baseURL == "" {
		baseURL = "https://cloud.m2amedia.tv"
	}

	awsAccountID := os.Getenv("M2A_AWS_ACCOUNT_ID")
	if awsAccountID == "" {
		return nil, fmt.Errorf("M2A_AWS_ACCOUNT_ID environment variable is required")
	}

	return &Config{
		APIKey:       apiKey,
		BaseURL:      baseURL,
		AWSAccountID: awsAccountID,
	}, nil
}
