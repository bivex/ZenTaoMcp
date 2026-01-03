// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
//
// Licensed under MIT License.
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

func RegisterReleaseResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register project releases resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectId}/releases", "ZenTao Project Releases"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract project ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://projects/123/releases, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "releases" {
				return nil, fmt.Errorf("invalid project releases URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/projects/%s/releases", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get project releases: %w", err)
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

	// Register product releases resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{productId}/releases", "ZenTao Product Releases"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract product ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://products/123/releases, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "products" || parts[2] != "releases" {
				return nil, fmt.Errorf("invalid product releases URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/products/%s/releases", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get product releases: %w", err)
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
