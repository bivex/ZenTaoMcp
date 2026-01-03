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

func RegisterSpaceResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register spaces list resource
	s.AddResource(mcp.NewResource(
		"zentao://spaces",
		"ZenTao Spaces",
		mcp.WithResourceDescription("List of all spaces in ZenTao"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=space&f=browse&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get spaces: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://spaces",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register space details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://spaces/{spaceID}", "ZenTao Space Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "spaces" {
				return nil, fmt.Errorf("invalid space URI format: %s", uri)
			}
			spaceID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=space&f=browse&t=json&spaceID=%s", spaceID))
			if err != nil {
				return nil, fmt.Errorf("failed to get space details: %w", err)
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

	// Register space applications resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://spaces/{spaceID}/applications", "ZenTao Space Applications"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "spaces" || parts[2] != "applications" {
				return nil, fmt.Errorf("invalid space applications URI format: %s", uri)
			}
			spaceID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=space&f=browse&t=json&spaceID=%s", spaceID))
			if err != nil {
				return nil, fmt.Errorf("failed to get space applications: %w", err)
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
}
