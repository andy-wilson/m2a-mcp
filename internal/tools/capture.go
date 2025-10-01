package tools

import (
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/andy-wilson/m2a-mcp/internal/client"
)

// CaptureTools handles M2A Capture API operations
type CaptureTools struct {
	client *client.M2AClient
}

// NewCaptureTools creates a new CaptureTools instance
func NewCaptureTools(client *client.M2AClient) *CaptureTools {
	return &CaptureTools{client: client}
}

// ListCaptures lists all capture jobs
func (t *CaptureTools) ListCaptures(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v1/connect/capture"

	status, _ := arguments["status"].(string)
	if status != "" {
		endpoint += "?status=" + status
	}

	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list captures: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetCapture gets details of a specific capture job
func (t *CaptureTools) GetCapture(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	captureID, ok := arguments["capture_id"].(string)
	if !ok || captureID == "" {
		return mcp.NewToolResultError("capture_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/connect/capture/%s", captureID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get capture: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateCapture creates a new live-to-VOD capture job
func (t *CaptureTools) CreateCapture(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	channelID, ok := arguments["channel_id"].(string)
	if !ok || channelID == "" {
		return mcp.NewToolResultError("channel_id is required"), nil
	}

	startTime, ok := arguments["start_time"].(string)
	if !ok || startTime == "" {
		return mcp.NewToolResultError("start_time is required"), nil
	}

	endTime, ok := arguments["end_time"].(string)
	if !ok || endTime == "" {
		return mcp.NewToolResultError("end_time is required"), nil
	}

	body := map[string]interface{}{
		"name":       name,
		"channel_id": channelID,
		"start_time": startTime,
		"end_time":   endTime,
	}

	endpoint := "/api/v1/connect/capture"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create capture: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CancelCapture cancels an in-progress capture job
func (t *CaptureTools) CancelCapture(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	captureID, ok := arguments["capture_id"].(string)
	if !ok || captureID == "" {
		return mcp.NewToolResultError("capture_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/connect/capture/%s/cancel", captureID)
	data, err := t.client.Post(endpoint, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to cancel capture: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// ListCaptureExports lists all completed VOD exports from captures
func (t *CaptureTools) ListCaptureExports(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v1/connect/capture/exports"
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list capture exports: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetCaptureExport gets details of a specific capture export
func (t *CaptureTools) GetCaptureExport(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	exportID, ok := arguments["export_id"].(string)
	if !ok || exportID == "" {
		return mcp.NewToolResultError("export_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/connect/capture/exports/%s", exportID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get capture export: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateClip creates a frame-accurate clip from a capture
func (t *CaptureTools) CreateClip(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	captureID, ok := arguments["capture_id"].(string)
	if !ok || captureID == "" {
		return mcp.NewToolResultError("capture_id is required"), nil
	}

	startTimecode, ok := arguments["start_timecode"].(string)
	if !ok || startTimecode == "" {
		return mcp.NewToolResultError("start_timecode is required"), nil
	}

	endTimecode, ok := arguments["end_timecode"].(string)
	if !ok || endTimecode == "" {
		return mcp.NewToolResultError("end_timecode is required"), nil
	}

	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	body := map[string]interface{}{
		"capture_id":     captureID,
		"start_timecode": startTimecode,
		"end_timecode":   endTimecode,
		"name":           name,
	}

	endpoint := "/api/v1/connect/capture/clips"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create clip: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
