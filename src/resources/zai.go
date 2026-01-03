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

	// Register AI models resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/models",
		"Available AI Models",
		mcp.WithResourceDescription("List of available AI models and their capabilities"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=models&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get AI models: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/models",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register vectorization statistics resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/vectorization/stats",
		"Vectorization Statistics",
		mcp.WithResourceDescription("Statistics and metrics for vectorization operations"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=vectorizationStats&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get vectorization statistics: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/vectorization/stats",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register AI service health resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/health",
		"AI Service Health Status",
		mcp.WithResourceDescription("Health status and connectivity information for AI services"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=health&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get AI service health: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/health",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register AI usage metrics resource
	s.AddResource(mcp.NewResource(
		"zentao://zai/usage",
		"AI Usage Metrics",
		mcp.WithResourceDescription("Usage statistics and metrics for AI features"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zai&f=usage&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get AI usage metrics: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zai/usage",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
