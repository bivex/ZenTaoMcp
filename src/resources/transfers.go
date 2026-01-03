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

func RegisterTransferResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register supported modules for transfer resource
	s.AddResource(mcp.NewResource(
		"zentao://transfer/modules",
		"ZenTao Transfer Supported Modules",
		mcp.WithResourceDescription("List of modules that support data transfer operations"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// This would typically fetch from an API, but for now we'll provide known modules
		modules := `{
			"modules": [
				"story", "task", "bug", "product", "project", "execution",
				"requirement", "epic", "user", "build", "release", "testcase",
				"testtask", "feedback", "ticket", "entry"
			],
			"description": "Modules that support import/export operations in ZenTao"
		}`

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://transfer/modules",
				MIMEType: "application/json",
				Text:     modules,
			},
		}, nil
	})

	// Register module export templates resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://transfer/{module}/export-template", "ZenTao Export Template"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "transfer" || parts[2] != "export-template" {
				return nil, fmt.Errorf("invalid export template URI format: %s", uri)
			}
			module := parts[1]

			resp, err := client.Post(fmt.Sprintf("/index.php?m=transfer&f=exportTemplate&t=json&module=%s", module), nil)
			if err != nil {
				return nil, fmt.Errorf("failed to get export template for module %s: %w", module, err)
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

	// Register module import status resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://transfer/{module}/import-status", "ZenTao Import Status"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "transfer" || parts[2] != "import-status" {
				return nil, fmt.Errorf("invalid import status URI format: %s", uri)
			}
			module := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=transfer&f=import&t=json&module=%s", module))
			if err != nil {
				return nil, fmt.Errorf("failed to get import status for module %s: %w", module, err)
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

	// Register export history resource
	s.AddResource(mcp.NewResource(
		"zentao://transfer/exports",
		"ZenTao Export History",
		mcp.WithResourceDescription("History of data exports"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=transfer&f=history&t=json&type=export")
		if err != nil {
			return nil, fmt.Errorf("failed to get export history: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://transfer/exports",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register import history resource
	s.AddResource(mcp.NewResource(
		"zentao://transfer/imports",
		"ZenTao Import History",
		mcp.WithResourceDescription("History of data imports"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=transfer&f=history&t=json&type=import")
		if err != nil {
			return nil, fmt.Errorf("failed to get import history: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://transfer/imports",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
