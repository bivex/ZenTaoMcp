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

func RegisterPlanResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /productplans endpoint in ZenTao API
	// Plans are accessed individually by ID

	// Register plan detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://plan/{id}", "ZenTao Product Plan Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "plan")

			if id == "" {
				return nil, fmt.Errorf("plan ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/productplans/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get plan details: %w", err)
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
