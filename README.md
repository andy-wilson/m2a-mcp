# M2A Media MCP Service

** BIG FAT WARNING! ** 

This was put together as an experiment with Claude Code and Sonnet 4.5. ** Here be large, possibly drunken, and definitely unweildy dragons! **

** USE AT YOUR OWN RISK, AND PREFERABLY NOT ON YOU PRODUCTION ACCOUNT! **

---

A Model Context Protocol (MCP) service for interacting with M2A Media's cloud-based video orchestration and automation platform. This service provides comprehensive access to M2A Connect, Live, Capture, and VOD APIs.

## Features

- **M2A Connect**: Manage video sources, subscribers, subscriptions, and schedules
- **M2A Live**: Control MediaLive channels, encoder configurations, and workflows
- **M2A Capture**: Create live-to-VOD captures and frame-accurate clips
- **M2A VOD**: Manage video on demand assets and playback URLs

## Prerequisites

- Go 1.24 or later
- M2A Media account with API credentials
- Associated AWS account

## Installation

### From Source

```bash
git clone https://github.com/andy-wilson/m2a-mcp.git
cd m2a-mcp
go build -o m2a-mcp
```

## Configuration

The service requires the following environment variables:

- `M2A_API_KEY` (required): Your M2A Media API key
- `M2A_BASE_URL` (optional): M2A API base URL (default: `https://cloud.m2amedia.tv`)
- `M2A_AWS_ACCOUNT_ID` (required): Your AWS account ID associated with M2A

### Getting API Credentials

To obtain API credentials:

1. Contact M2A Media at ops@m2amedia.tv
2. Provide your company details and AWS Entitlement information
3. Receive initial login credentials
4. Associate your AWS account with M2A

### Claude Desktop Configuration

Add to your `claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "m2a-media": {
      "command": "/path/to/m2a-mcp",
      "env": {
        "M2A_API_KEY": "your-api-key-here",
        "M2A_AWS_ACCOUNT_ID": "your-aws-account-id"
      }
    }
  }
}
```

## Available Tools

### M2A Connect Tools

#### Source Management
- `list_sources` - List all video sources
- `get_source` - Get source details
- `create_source` - Create a new video source
- `update_source` - Update source configuration
- `delete_source` - Delete a source

#### Subscriber Management
- `list_subscribers` - List all subscribers
- `get_subscriber` - Get subscriber details
- `create_subscriber` - Create a new subscriber

#### Subscription Management
- `list_subscriptions` - List subscription packages
- `get_subscription` - Get subscription details
- `create_subscription` - Create subscription package

#### Schedule Management
- `list_schedules` - List scheduled events
- `get_schedule` - Get schedule details
- `create_schedule` - Create a new schedule

### M2A Live Tools

#### Channel Management
- `list_channels` - List MediaLive channels
- `get_channel` - Get channel details
- `create_channel` - Create a new channel
- `start_channel` - Start a channel
- `stop_channel` - Stop a channel
- `delete_channel` - Delete a channel

#### Encoder Configuration
- `list_encoder_configs` - List encoder configurations
- `get_encoder_config` - Get encoder config details

#### Workflow Management
- `list_workflows` - List live streaming workflows
- `get_workflow` - Get workflow details
- `create_workflow` - Create a new workflow

### M2A Capture Tools

- `list_captures` - List capture jobs
- `get_capture` - Get capture job details
- `create_capture` - Create live-to-VOD capture
- `cancel_capture` - Cancel capture job
- `list_capture_exports` - List completed exports
- `get_capture_export` - Get export details
- `create_clip` - Create frame-accurate clip

### VOD Tools

- `list_vod_assets` - List VOD assets
- `get_vod_asset` - Get VOD asset details
- `update_vod_metadata` - Update asset metadata
- `delete_vod_asset` - Delete VOD asset
- `get_playback_url` - Get streaming URL

## Usage Examples

### List All Sources

```
Can you list all my video sources in M2A Connect?
```

### Create a New Live Channel

```
Create a new MediaLive channel named "Sports Stream" with RTMP_PUSH input type
```

### Create a Capture Job

```
Create a capture job from channel "ch-123" starting at 2025-10-01T14:00:00Z
and ending at 2025-10-01T16:00:00Z, name it "Game Highlights"
```

### Get Playback URL

```
Get the HLS playback URL for VOD asset "asset-456"
```

## API Endpoints

The service interacts with the following M2A Media API endpoints:

- **Connect API**: `/api/v2/connect/*` and `/api/v1/connect/capture/*`
- **Live API**: `/api/v3/live/*`, `/api/v1/live/*`
- **VOD API**: `/api/v1/vod/*`

## Error Handling

The service provides detailed error messages for:

- Authentication failures
- Missing required parameters
- API errors with HTTP status codes
- Network connectivity issues

## Development

### Project Structure

```
m2a-mcp/
├── main.go                    # MCP server entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── client/
│   │   └── client.go         # M2A API HTTP client
│   └── tools/
│       ├── connect.go        # Connect API tools
│       ├── live.go           # Live API tools
│       ├── capture.go        # Capture API tools
│       └── vod.go            # VOD API tools
├── go.mod
└── README.md
```

### Building

```bash
go build -o m2a-mcp
```

### Testing

```bash
go test ./...
```

## Resources

- [M2A Media Official Site](https://m2amedia.tv)
- [M2A Connect API Documentation](https://cloud.m2amedia.tv/docs/html/connect/api.html)
- [M2A Live API Documentation](https://cloud.m2amedia.tv/docs/html/live/api.html)
- [Model Context Protocol](https://modelcontextprotocol.io)

## License

MIT License - see LICENSE file for details

## Support

For M2A Media API support, contact ops@m2amedia.tv

For issues with this MCP service, please open an issue on GitHub.
