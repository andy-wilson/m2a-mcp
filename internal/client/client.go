package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/andy-wilson/m2a-mcp/internal/config"
)

// M2AClient is the HTTP client for M2A Media API
type M2AClient struct {
	config     *config.Config
	httpClient *http.Client
}

// NewM2AClient creates a new M2A API client
func NewM2AClient(cfg *config.Config) *M2AClient {
	return &M2AClient{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// Get performs a GET request
func (c *M2AClient) Get(endpoint string) ([]byte, error) {
	url := c.config.BaseURL + endpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.doRequest(req)
}

// Post performs a POST request with JSON body
func (c *M2AClient) Post(endpoint string, body interface{}) ([]byte, error) {
	url := c.config.BaseURL + endpoint

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// Put performs a PUT request with JSON body
func (c *M2AClient) Put(endpoint string, body interface{}) ([]byte, error) {
	url := c.config.BaseURL + endpoint

	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	return c.doRequest(req)
}

// Delete performs a DELETE request
func (c *M2AClient) Delete(endpoint string) ([]byte, error) {
	url := c.config.BaseURL + endpoint
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	return c.doRequest(req)
}

// doRequest executes the HTTP request with authentication
func (c *M2AClient) doRequest(req *http.Request) ([]byte, error) {
	// Add authentication header
	req.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	return body, nil
}

// GetConfig returns the client configuration
func (c *M2AClient) GetConfig() *config.Config {
	return c.config
}
