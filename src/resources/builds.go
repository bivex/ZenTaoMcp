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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterBuildResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /builds endpoint in ZenTao API
	// Builds are accessed via projects/executions

	// Register build detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://build/{id}", "ZenTao Build Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "build")

			if id == "" {
				return nil, fmt.Errorf("build ID not found in URI: %s", uri)
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
		},
	)
}
