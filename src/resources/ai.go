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
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterAiResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register AI admin index resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/admin",
		"AI Module Admin Overview",
		mcp.WithResourceDescription("AI module administrative interface and overview"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=adminIndex&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get AI admin index: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/admin",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register mini programs list resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/mini-programs",
		"AI Mini Programs List",
		mcp.WithResourceDescription("List of available AI mini programs"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=miniPrograms&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get mini programs: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/mini-programs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register prompts list resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/prompts",
		"AI Prompts List",
		mcp.WithResourceDescription("List of available AI prompts"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=prompts&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get prompts: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/prompts",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register role templates resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/role-templates",
		"AI Role Templates",
		mcp.WithResourceDescription("Available AI role templates"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=roleTemplates&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get role templates: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/role-templates",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register prompt view resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://ai/prompts/{id}", "AI Prompt Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "ai" || parts[1] != "prompts" {
				return nil, fmt.Errorf("invalid prompt URI format: %s", uri)
			}
			promptID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptView&t=json&id=%s", promptID))
			if err != nil {
				return nil, fmt.Errorf("failed to get prompt %s: %w", promptID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register mini program details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://ai/mini-programs/{appID}", "Mini Program Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "ai" || parts[1] != "mini-programs" {
				return nil, fmt.Errorf("invalid mini program URI format: %s", uri)
			}
			appID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=miniProgramView&t=json&appID=%s", appID))
			if err != nil {
				return nil, fmt.Errorf("failed to get mini program %s: %w", appID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register prompt execution status resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/prompt-execution/status",
		"Prompt Execution Status",
		mcp.WithResourceDescription("Current status of prompt executions"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=promptExecutionStatus&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get prompt execution status: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/prompt-execution/status",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register AI analytics resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/analytics",
		"AI Module Analytics",
		mcp.WithResourceDescription("Analytics and usage statistics for AI features"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=analytics&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get AI analytics: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/analytics",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register published prompts resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/prompts/published",
		"Published AI Prompts",
		mcp.WithResourceDescription("List of published AI prompts"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=prompts&t=json&status=published")
		if err != nil {
			return nil, fmt.Errorf("failed to get published prompts: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/prompts/published",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register testing prompts resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/prompts/testing",
		"Testing AI Prompts",
		mcp.WithResourceDescription("List of prompts in testing phase"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=prompts&t=json&status=testing")
		if err != nil {
			return nil, fmt.Errorf("failed to get testing prompts: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/prompts/testing",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register published mini programs resource
	s.AddResource(mcp.NewResource(
		"zentao://ai/mini-programs/published",
		"Published Mini Programs",
		mcp.WithResourceDescription("List of published AI mini programs"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=ai&f=miniPrograms&t=json&status=published")
		if err != nil {
			return nil, fmt.Errorf("failed to get published mini programs: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://ai/mini-programs/published",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
