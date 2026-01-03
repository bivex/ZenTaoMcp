// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
//
//
// Licensed under MIT License.
// Commercial licensing available upon request.

package resources

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterProgramResources(s *server.MCPServer, client *client.ZenTaoClient) {
	programListResource := mcp.NewResource(
		"zentao://programs",
		"ZenTao Program List",
		mcp.WithResourceDescription("List of all programs in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(programListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/programs")
		if err != nil {
			return nil, fmt.Errorf("failed to get programs: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://programs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	programDetailResource := mcp.NewResource(
		"zentao://programs/{id}",
		"ZenTao Program Details",
		mcp.WithResourceDescription("Details of a specific program by ID (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(programDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "programs")

		resp, err := client.Get(fmt.Sprintf("/programs/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get program details: %w", err)
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
