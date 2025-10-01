package main

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/andy-wilson/m2a-mcp/internal/client"
	"github.com/andy-wilson/m2a-mcp/internal/config"
	"github.com/andy-wilson/m2a-mcp/internal/tools"
)

const (
	serverName    = "m2a-media-mcp"
	serverVersion = "0.1.0"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Create M2A API client
	m2aClient := client.NewM2AClient(cfg)

	// Create MCP server
	mcpServer := server.NewMCPServer(
		serverName,
		serverVersion,
	)

	// Register all tools
	if err := registerTools(mcpServer, m2aClient); err != nil {
		log.Fatalf("Failed to register tools: %v", err)
	}

	// Start server with stdio transport
	if err := server.ServeStdio(mcpServer); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func registerTools(s *server.MCPServer, client *client.M2AClient) error {
	// M2A Connect tools
	connectTools := tools.NewConnectTools(client)
	s.AddTool(mcp.NewTool("list_sources",
		mcp.WithDescription("List all video sources in M2A Connect"),
		mcp.WithString("status", mcp.Description("Filter by status (active, inactive, all)"), mcp.Enum("active", "inactive", "all")),
	), connectTools.ListSources)

	s.AddTool(mcp.NewTool("get_source",
		mcp.WithDescription("Get details of a specific video source"),
		mcp.WithString("source_id", mcp.Required(), mcp.Description("The ID of the source")),
	), connectTools.GetSource)

	s.AddTool(mcp.NewTool("create_source",
		mcp.WithDescription("Create a new video source in M2A Connect"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Name of the source")),
		mcp.WithString("type", mcp.Required(), mcp.Description("Source type (rtmp, srt, udp, etc.)"), mcp.Enum("rtmp", "srt", "udp", "rtp")),
		mcp.WithString("url", mcp.Required(), mcp.Description("Source URL or endpoint")),
		mcp.WithString("description", mcp.Description("Optional description")),
	), connectTools.CreateSource)

	s.AddTool(mcp.NewTool("update_source",
		mcp.WithDescription("Update an existing video source"),
		mcp.WithString("source_id", mcp.Required(), mcp.Description("The ID of the source")),
		mcp.WithString("name", mcp.Description("New name for the source")),
		mcp.WithString("url", mcp.Description("New source URL")),
		mcp.WithString("description", mcp.Description("New description")),
	), connectTools.UpdateSource)

	s.AddTool(mcp.NewTool("delete_source",
		mcp.WithDescription("Delete a video source"),
		mcp.WithString("source_id", mcp.Required(), mcp.Description("The ID of the source to delete")),
	), connectTools.DeleteSource)

	s.AddTool(mcp.NewTool("list_subscribers",
		mcp.WithDescription("List all subscribers in M2A Connect"),
		mcp.WithNumber("limit", mcp.Description("Maximum number of results to return")),
		mcp.WithNumber("offset", mcp.Description("Offset for pagination")),
	), connectTools.ListSubscribers)

	s.AddTool(mcp.NewTool("get_subscriber",
		mcp.WithDescription("Get details of a specific subscriber"),
		mcp.WithString("subscriber_id", mcp.Required(), mcp.Description("The ID of the subscriber")),
	), connectTools.GetSubscriber)

	s.AddTool(mcp.NewTool("create_subscriber",
		mcp.WithDescription("Create a new subscriber"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Subscriber name")),
		mcp.WithString("email", mcp.Required(), mcp.Description("Subscriber email")),
		mcp.WithString("organization", mcp.Description("Organization name")),
	), connectTools.CreateSubscriber)

	s.AddTool(mcp.NewTool("list_subscriptions",
		mcp.WithDescription("List all subscription packages"),
	), connectTools.ListSubscriptions)

	s.AddTool(mcp.NewTool("get_subscription",
		mcp.WithDescription("Get details of a specific subscription package"),
		mcp.WithString("subscription_id", mcp.Required(), mcp.Description("The ID of the subscription")),
	), connectTools.GetSubscription)

	s.AddTool(mcp.NewTool("create_subscription",
		mcp.WithDescription("Create a new subscription package"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Subscription name")),
		mcp.WithString("subscriber_id", mcp.Required(), mcp.Description("Subscriber ID")),
		mcp.WithString("source_ids", mcp.Required(), mcp.Description("Comma-separated list of source IDs")),
	), connectTools.CreateSubscription)

	s.AddTool(mcp.NewTool("list_schedules",
		mcp.WithDescription("List all scheduled events"),
		mcp.WithString("start_date", mcp.Description("Filter by start date (ISO 8601 format)")),
		mcp.WithString("end_date", mcp.Description("Filter by end date (ISO 8601 format)")),
	), connectTools.ListSchedules)

	s.AddTool(mcp.NewTool("get_schedule",
		mcp.WithDescription("Get details of a specific schedule"),
		mcp.WithString("schedule_id", mcp.Required(), mcp.Description("The ID of the schedule")),
	), connectTools.GetSchedule)

	s.AddTool(mcp.NewTool("create_schedule",
		mcp.WithDescription("Create a new scheduled event"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Schedule name")),
		mcp.WithString("source_id", mcp.Required(), mcp.Description("Source ID")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Start time (ISO 8601 format)")),
		mcp.WithString("end_time", mcp.Required(), mcp.Description("End time (ISO 8601 format)")),
	), connectTools.CreateSchedule)

	// M2A Live tools
	liveTools := tools.NewLiveTools(client)
	s.AddTool(mcp.NewTool("list_channels",
		mcp.WithDescription("List all MediaLive channels"),
		mcp.WithString("state", mcp.Description("Filter by channel state"), mcp.Enum("IDLE", "CREATING", "STARTING", "RUNNING", "STOPPING", "DELETING")),
	), liveTools.ListChannels)

	s.AddTool(mcp.NewTool("get_channel",
		mcp.WithDescription("Get details of a specific MediaLive channel"),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("The ID of the channel")),
	), liveTools.GetChannel)

	s.AddTool(mcp.NewTool("create_channel",
		mcp.WithDescription("Create a new MediaLive channel"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Channel name")),
		mcp.WithString("input_type", mcp.Required(), mcp.Description("Input type"), mcp.Enum("RTMP_PUSH", "RTP_PUSH", "UDP_PUSH", "MEDIACONNECT")),
		mcp.WithString("encoder_config_id", mcp.Description("Encoder configuration ID to use")),
	), liveTools.CreateChannel)

	s.AddTool(mcp.NewTool("start_channel",
		mcp.WithDescription("Start a MediaLive channel"),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("The ID of the channel to start")),
	), liveTools.StartChannel)

	s.AddTool(mcp.NewTool("stop_channel",
		mcp.WithDescription("Stop a MediaLive channel"),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("The ID of the channel to stop")),
	), liveTools.StopChannel)

	s.AddTool(mcp.NewTool("delete_channel",
		mcp.WithDescription("Delete a MediaLive channel"),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("The ID of the channel to delete")),
	), liveTools.DeleteChannel)

	s.AddTool(mcp.NewTool("list_encoder_configs",
		mcp.WithDescription("List encoder configuration fragments"),
	), liveTools.ListEncoderConfigs)

	s.AddTool(mcp.NewTool("get_encoder_config",
		mcp.WithDescription("Get details of a specific encoder configuration"),
		mcp.WithString("config_id", mcp.Required(), mcp.Description("The ID of the encoder configuration")),
	), liveTools.GetEncoderConfig)

	s.AddTool(mcp.NewTool("list_workflows",
		mcp.WithDescription("List all live streaming workflows"),
	), liveTools.ListWorkflows)

	s.AddTool(mcp.NewTool("get_workflow",
		mcp.WithDescription("Get details of a specific workflow"),
		mcp.WithString("workflow_id", mcp.Required(), mcp.Description("The ID of the workflow")),
	), liveTools.GetWorkflow)

	s.AddTool(mcp.NewTool("create_workflow",
		mcp.WithDescription("Create a new live streaming workflow"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Workflow name")),
		mcp.WithString("description", mcp.Description("Workflow description")),
	), liveTools.CreateWorkflow)

	// M2A Capture tools
	captureTools := tools.NewCaptureTools(client)
	s.AddTool(mcp.NewTool("list_captures",
		mcp.WithDescription("List all capture jobs (live-to-VOD)"),
		mcp.WithString("status", mcp.Description("Filter by status"), mcp.Enum("PENDING", "IN_PROGRESS", "COMPLETED", "FAILED", "CANCELLED")),
	), captureTools.ListCaptures)

	s.AddTool(mcp.NewTool("get_capture",
		mcp.WithDescription("Get details of a specific capture job"),
		mcp.WithString("capture_id", mcp.Required(), mcp.Description("The ID of the capture job")),
	), captureTools.GetCapture)

	s.AddTool(mcp.NewTool("create_capture",
		mcp.WithDescription("Create a new live-to-VOD capture job"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Capture job name")),
		mcp.WithString("channel_id", mcp.Required(), mcp.Description("Source channel ID")),
		mcp.WithString("start_time", mcp.Required(), mcp.Description("Capture start time (ISO 8601)")),
		mcp.WithString("end_time", mcp.Required(), mcp.Description("Capture end time (ISO 8601)")),
	), captureTools.CreateCapture)

	s.AddTool(mcp.NewTool("cancel_capture",
		mcp.WithDescription("Cancel an in-progress capture job"),
		mcp.WithString("capture_id", mcp.Required(), mcp.Description("The ID of the capture job to cancel")),
	), captureTools.CancelCapture)

	s.AddTool(mcp.NewTool("list_capture_exports",
		mcp.WithDescription("List all completed VOD exports from captures"),
	), captureTools.ListCaptureExports)

	s.AddTool(mcp.NewTool("get_capture_export",
		mcp.WithDescription("Get details of a specific capture export"),
		mcp.WithString("export_id", mcp.Required(), mcp.Description("The ID of the export")),
	), captureTools.GetCaptureExport)

	s.AddTool(mcp.NewTool("create_clip",
		mcp.WithDescription("Create a frame-accurate clip from a capture"),
		mcp.WithString("capture_id", mcp.Required(), mcp.Description("Source capture ID")),
		mcp.WithString("start_timecode", mcp.Required(), mcp.Description("Start timecode (HH:MM:SS:FF)")),
		mcp.WithString("end_timecode", mcp.Required(), mcp.Description("End timecode (HH:MM:SS:FF)")),
		mcp.WithString("name", mcp.Required(), mcp.Description("Clip name")),
	), captureTools.CreateClip)

	// VOD tools
	vodTools := tools.NewVODTools(client)
	s.AddTool(mcp.NewTool("list_vod_assets",
		mcp.WithDescription("List all VOD assets"),
		mcp.WithNumber("limit", mcp.Description("Maximum number of results")),
		mcp.WithNumber("offset", mcp.Description("Offset for pagination")),
	), vodTools.ListVODAssets)

	s.AddTool(mcp.NewTool("get_vod_asset",
		mcp.WithDescription("Get details of a specific VOD asset"),
		mcp.WithString("asset_id", mcp.Required(), mcp.Description("The ID of the VOD asset")),
	), vodTools.GetVODAsset)

	s.AddTool(mcp.NewTool("update_vod_metadata",
		mcp.WithDescription("Update metadata for a VOD asset"),
		mcp.WithString("asset_id", mcp.Required(), mcp.Description("The ID of the VOD asset")),
		mcp.WithString("title", mcp.Description("Asset title")),
		mcp.WithString("description", mcp.Description("Asset description")),
		mcp.WithString("tags", mcp.Description("Comma-separated tags")),
	), vodTools.UpdateVODMetadata)

	s.AddTool(mcp.NewTool("delete_vod_asset",
		mcp.WithDescription("Delete a VOD asset"),
		mcp.WithString("asset_id", mcp.Required(), mcp.Description("The ID of the asset to delete")),
	), vodTools.DeleteVODAsset)

	s.AddTool(mcp.NewTool("get_playback_url",
		mcp.WithDescription("Get streaming playback URL for a VOD asset"),
		mcp.WithString("asset_id", mcp.Required(), mcp.Description("The ID of the VOD asset")),
		mcp.WithString("format", mcp.Description("Playback format"), mcp.Enum("hls", "dash", "mp4")),
	), vodTools.GetPlaybackURL)

	return nil
}
