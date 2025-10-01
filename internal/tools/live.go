package tools

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/andy-wilson/m2a-mcp/internal/client"
)

// LiveTools handles M2A Live API operations
type LiveTools struct {
	client *client.M2AClient
}

// NewLiveTools creates a new LiveTools instance
func NewLiveTools(client *client.M2AClient) *LiveTools {
	return &LiveTools{client: client}
}

// ListChannels lists all MediaLive channels
func (t *LiveTools) ListChannels(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v3/live/channels"

	state, _ := arguments["state"].(string)
	if state != "" {
		endpoint += "?state=" + state
	}

	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list channels: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetChannel gets details of a specific channel
func (t *LiveTools) GetChannel(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	channelID, ok := arguments["channel_id"].(string)
	if !ok || channelID == "" {
		return mcp.NewToolResultError("channel_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v3/live/channels/%s", channelID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get channel: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateChannel creates a new MediaLive channel
func (t *LiveTools) CreateChannel(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	inputType, ok := arguments["input_type"].(string)
	if !ok || inputType == "" {
		return mcp.NewToolResultError("input_type is required"), nil
	}

	body := map[string]interface{}{
		"name":       name,
		"input_type": inputType,
	}

	encoderConfigID, _ := arguments["encoder_config_id"].(string)
	if encoderConfigID != "" {
		body["encoder_config_id"] = encoderConfigID
	}

	endpoint := "/api/v3/live/channels"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create channel: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// StartChannel starts a MediaLive channel
func (t *LiveTools) StartChannel(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	channelID, ok := arguments["channel_id"].(string)
	if !ok || channelID == "" {
		return mcp.NewToolResultError("channel_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v3/live/channels/%s/start", channelID)
	data, err := t.client.Post(endpoint, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to start channel: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// StopChannel stops a MediaLive channel
func (t *LiveTools) StopChannel(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	channelID, ok := arguments["channel_id"].(string)
	if !ok || channelID == "" {
		return mcp.NewToolResultError("channel_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v3/live/channels/%s/stop", channelID)
	data, err := t.client.Post(endpoint, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to stop channel: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// DeleteChannel deletes a MediaLive channel
func (t *LiveTools) DeleteChannel(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	channelID, ok := arguments["channel_id"].(string)
	if !ok || channelID == "" {
		return mcp.NewToolResultError("channel_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v3/live/channels/%s", channelID)
	_, err := t.client.Delete(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to delete channel: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Channel %s deleted successfully", channelID),
	}
	jsonData, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonData)), nil
}

// ListEncoderConfigs lists encoder configuration fragments
func (t *LiveTools) ListEncoderConfigs(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v1/live/encoder-configs"
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list encoder configs: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetEncoderConfig gets details of a specific encoder configuration
func (t *LiveTools) GetEncoderConfig(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	configID, ok := arguments["config_id"].(string)
	if !ok || configID == "" {
		return mcp.NewToolResultError("config_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/live/encoder-configs/%s", configID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get encoder config: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// ListWorkflows lists all live streaming workflows
func (t *LiveTools) ListWorkflows(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v1/live/workflows"
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list workflows: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetWorkflow gets details of a specific workflow
func (t *LiveTools) GetWorkflow(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	workflowID, ok := arguments["workflow_id"].(string)
	if !ok || workflowID == "" {
		return mcp.NewToolResultError("workflow_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/live/workflows/%s", workflowID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get workflow: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateWorkflow creates a new live streaming workflow
func (t *LiveTools) CreateWorkflow(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	body := map[string]interface{}{
		"name": name,
	}

	desc, _ := arguments["description"].(string)
	if desc != "" {
		body["description"] = desc
	}

	endpoint := "/api/v1/live/workflows"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create workflow: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
