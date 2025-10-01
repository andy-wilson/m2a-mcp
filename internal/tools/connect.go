package tools

import (
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/andy-wilson/m2a-mcp/internal/client"
)

// ConnectTools handles M2A Connect API operations
type ConnectTools struct {
	client *client.M2AClient
}

// NewConnectTools creates a new ConnectTools instance
func NewConnectTools(client *client.M2AClient) *ConnectTools {
	return &ConnectTools{client: client}
}

// ListSources lists all video sources
func (t *ConnectTools) ListSources(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	status, _ := arguments["status"].(string)

	endpoint := "/api/v2/connect/sources"
	if status != "" && status != "all" {
		endpoint += "?status=" + status
	}

	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list sources: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetSource gets details of a specific source
func (t *ConnectTools) GetSource(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	sourceID, ok := arguments["source_id"].(string)
	if !ok || sourceID == "" {
		return mcp.NewToolResultError("source_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/sources/%s", sourceID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get source: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateSource creates a new video source
func (t *ConnectTools) CreateSource(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	sourceType, ok := arguments["type"].(string)
	if !ok || sourceType == "" {
		return mcp.NewToolResultError("type is required"), nil
	}

	url, ok := arguments["url"].(string)
	if !ok || url == "" {
		return mcp.NewToolResultError("url is required"), nil
	}

	body := map[string]interface{}{
		"name": name,
		"type": sourceType,
		"url":  url,
	}

	desc, _ := arguments["description"].(string)
	if desc != "" {
		body["description"] = desc
	}

	endpoint := "/api/v2/connect/sources"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create source: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// UpdateSource updates an existing source
func (t *ConnectTools) UpdateSource(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	sourceID, ok := arguments["source_id"].(string)
	if !ok || sourceID == "" {
		return mcp.NewToolResultError("source_id is required"), nil
	}

	body := make(map[string]interface{})
	if name, _ := arguments["name"].(string); name != "" {
		body["name"] = name
	}
	if url, _ := arguments["url"].(string); url != "" {
		body["url"] = url
	}
	if desc, _ := arguments["description"].(string); desc != "" {
		body["description"] = desc
	}

	if len(body) == 0 {
		return mcp.NewToolResultError("at least one field to update is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/sources/%s", sourceID)
	data, err := t.client.Put(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to update source: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// DeleteSource deletes a source
func (t *ConnectTools) DeleteSource(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	sourceID, ok := arguments["source_id"].(string)
	if !ok || sourceID == "" {
		return mcp.NewToolResultError("source_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/sources/%s", sourceID)
	_, err := t.client.Delete(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to delete source: %v", err)), nil
	}

	result := map[string]interface{}{
		"success": true,
		"message": fmt.Sprintf("Source %s deleted successfully", sourceID),
	}
	jsonData, _ := json.Marshal(result)
	return mcp.NewToolResultText(string(jsonData)), nil
}

// ListSubscribers lists all subscribers
func (t *ConnectTools) ListSubscribers(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v2/connect/subscribers"

	// Add pagination parameters if provided
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
		return mcp.NewToolResultError(fmt.Sprintf("failed to list subscribers: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetSubscriber gets details of a specific subscriber
func (t *ConnectTools) GetSubscriber(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	subscriberID, ok := arguments["subscriber_id"].(string)
	if !ok || subscriberID == "" {
		return mcp.NewToolResultError("subscriber_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/subscribers/%s", subscriberID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get subscriber: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateSubscriber creates a new subscriber
func (t *ConnectTools) CreateSubscriber(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	email, ok := arguments["email"].(string)
	if !ok || email == "" {
		return mcp.NewToolResultError("email is required"), nil
	}

	body := map[string]interface{}{
		"name":  name,
		"email": email,
	}

	org, _ := arguments["organization"].(string)
	if org != "" {
		body["organization"] = org
	}

	endpoint := "/api/v2/connect/subscribers"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create subscriber: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// ListSubscriptions lists all subscriptions
func (t *ConnectTools) ListSubscriptions(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v2/connect/subscriptions"
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list subscriptions: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetSubscription gets details of a specific subscription
func (t *ConnectTools) GetSubscription(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	subscriptionID, ok := arguments["subscription_id"].(string)
	if !ok || subscriptionID == "" {
		return mcp.NewToolResultError("subscription_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/subscriptions/%s", subscriptionID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get subscription: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateSubscription creates a new subscription package
func (t *ConnectTools) CreateSubscription(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	subscriberID, ok := arguments["subscriber_id"].(string)
	if !ok || subscriberID == "" {
		return mcp.NewToolResultError("subscriber_id is required"), nil
	}

	sourceIDs, ok := arguments["source_ids"].(string)
	if !ok || sourceIDs == "" {
		return mcp.NewToolResultError("source_ids is required"), nil
	}

	body := map[string]interface{}{
		"name":          name,
		"subscriber_id": subscriberID,
		"source_ids":    sourceIDs,
	}

	endpoint := "/api/v2/connect/subscriptions"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create subscription: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// ListSchedules lists all schedules
func (t *ConnectTools) ListSchedules(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	endpoint := "/api/v2/connect/schedules"

	queryParams := ""
	startDate, _ := arguments["start_date"].(string)
	if startDate != "" {
		queryParams += "?start_date=" + startDate
	}
	endDate, _ := arguments["end_date"].(string)
	if endDate != "" {
		if queryParams == "" {
			queryParams += "?"
		} else {
			queryParams += "&"
		}
		queryParams += "end_date=" + endDate
	}
	endpoint += queryParams

	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to list schedules: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// GetSchedule gets details of a specific schedule
func (t *ConnectTools) GetSchedule(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	scheduleID, ok := arguments["schedule_id"].(string)
	if !ok || scheduleID == "" {
		return mcp.NewToolResultError("schedule_id is required"), nil
	}

	endpoint := fmt.Sprintf("/api/v2/connect/schedules/%s", scheduleID)
	data, err := t.client.Get(endpoint)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to get schedule: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}

// CreateSchedule creates a new schedule
func (t *ConnectTools) CreateSchedule(arguments map[string]interface{}) (*mcp.CallToolResult, error) {
	name, ok := arguments["name"].(string)
	if !ok || name == "" {
		return mcp.NewToolResultError("name is required"), nil
	}

	sourceID, ok := arguments["source_id"].(string)
	if !ok || sourceID == "" {
		return mcp.NewToolResultError("source_id is required"), nil
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
		"source_id":  sourceID,
		"start_time": startTime,
		"end_time":   endTime,
	}

	endpoint := "/api/v2/connect/schedules"
	data, err := t.client.Post(endpoint, body)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("failed to create schedule: %v", err)), nil
	}

	return mcp.NewToolResultText(string(data)), nil
}
