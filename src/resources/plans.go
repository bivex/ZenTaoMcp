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

func RegisterPlanResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /productplans endpoint in ZenTao API
	// Plans are accessed individually by ID

	planDetailResource := mcp.NewResource(
		"zentao://plan/*",
		"ZenTao Product Plan Details",
		mcp.WithResourceDescription("Details of a specific product plan (use zentao://plan/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(planDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid plan URI format. Use: zentao://plan/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("plan ID not specified in URI. Use format: zentao://plan/123")
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
	})
}
