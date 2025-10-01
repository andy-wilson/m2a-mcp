package tools

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/andy-wilson/m2a-mcp/internal/client"
)

// VODTools handles M2A VOD API operations
type VODTools struct {
	client *client.M2AClient
}

// NewVODTools creates a new VODTools instance
func NewVODTools(client *client.M2AClient) *VODTools {
	return &VODTools{client: client}
}

// ListVODAssets lists all VOD assets
func (t *VODTools) ListVODAssets(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v1/vod/assets"

	queryParams := ""
	if limit, _ := arguments["limit"].(float64); limit > 0 {
		queryParams += fmt.Sprintf("?limit=%d", int(limit))
	}
	if offset, _ := arguments["offset"].(float64); offset > 0 {
		if queryParams == "" {
			queryParams += "?"
		} else {
			queryParams += "&"
		}
		queryParams += fmt.Sprintf("offset=%d", int(offset))
	}
	endpoint += queryParams

	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list VOD assets: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetVODAsset gets details of a specific VOD asset
func (t *VODTools) GetVODAsset(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	assetID, ok := arguments["asset_id"].(string)
	if !ok || assetID == "" {
		return mcp.NewToolResultError("asset_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/vod/assets/%s", assetID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get VOD asset: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// UpdateVODMetadata updates metadata for a VOD asset
func (t *VODTools) UpdateVODMetadata(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	assetID, ok := arguments["asset_id"].(string)
	if !ok || assetID == "" {
		return mcp.NewToolResultError("asset_id is required"), nil
	}

	body := make(map[string]interface{})
	if title, _ := arguments["title"].(string); title != "" {
		body["title"] = title
	}
	if desc, _ := arguments["description"].(string); desc != "" {
		body["description"] = desc
	}
	if tags, _ := arguments["tags"].(string); tags != "" {
		body["tags"] = tags
	}

	if len(body) == 0 {
		return mcp.NewToolResultError("at least one metadata field is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/vod/assets/%s", assetID)
	data, err := t.client.Put(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to update VOD metadata: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// DeleteVODAsset deletes a VOD asset
func (t *VODTools) DeleteVODAsset(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	assetID, ok := arguments["asset_id"].(string)
	if !ok || assetID == "" {
		return mcp.NewToolResultError("asset_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v1/vod/assets/%s", assetID)
	_, err := t.client.Delete(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to delete VOD asset: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("VOD asset %s deleted successfully", assetID),
	}
	jsonData, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonData)), nil
}

// GetPlaybackURL gets streaming playback URL for a VOD asset
func (t *VODTools) GetPlaybackURL(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	assetID, ok := arguments["asset_id"].(string)
	if !ok || assetID == "" {
		return mcp.NewToolResultError("asset_id is required"), nil
	}

	format, _ := arguments["format"].(string)
	if format == "" {
		format = "hls" // default format
	}

	endpoint := fmt.Sprintf("/api/v1/vod/assets/%s/playback?format=%s", assetID, format)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get playback URL: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
