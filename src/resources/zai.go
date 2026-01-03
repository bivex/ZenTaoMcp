// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
//
// Licensed under the MIT License.
// Commercial licensing available upon request.

package resources

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterZaiResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register ZAI settings resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/settings",
		"ZenTao AI Settings",
		mcp.WithResourceDescription("Current ZenTao AI module configuration and settings"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=setting&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get ZAI settings: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/settings",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register ZAI token resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/token",
		"ZenTao AI Authentication Token",
		mcp.WithResourceDescription("Current authentication token for AI services"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=ajaxGetToken&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get ZAI token: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/token",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register vectorization status resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/vectorization/status",
		"Vectorization Status",
		mcp.WithResourceDescription("Current vectorization status and progress"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=vectorized&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get vectorization status: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/vectorization/status",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

}
