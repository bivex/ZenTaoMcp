// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
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

func RegisterBuildResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /builds endpoint in ZenTao API
	// Builds are accessed via projects/executions

	buildDetailResource := mcp.NewResource(
		"zentao://build/*",
		"ZenTao Build Details",
		mcp.WithResourceDescription("Details of a specific build (use zentao://build/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(buildDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid build URI format. Use: zentao://build/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("build ID not specified in URI. Use format: zentao://build/123")
		}

		resp, err := client.Get(fmt.Sprintf("/builds/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get build details: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
