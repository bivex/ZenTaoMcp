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

func RegisterTestTaskResources(s *server.MCPServer, client *client.ZenTaoClient) {
	testTaskListResource := mcp.NewResource(
		"zentao://testtasks",
		"ZenTao Test Task List",
		mcp.WithResourceDescription("List of all test tasks in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(testTaskListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/testtasks")
		if err != nil {
			return nil, fmt.Errorf("failed to get test tasks: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://testtasks",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register test task detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://testtask/{id}", "ZenTao Test Task Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI template variables (MCP library handles this)
			id := ""
			if idVal, ok := request.Params.Arguments["id"]; ok {
				if idStr, ok := idVal.(string); ok {
					id = idStr
				}
			}

			if id == "" {
				return nil, fmt.Errorf("test task ID not found in URI template")
			}

			resp, err := client.Get(fmt.Sprintf("/testtasks/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get test task details: %w", err)
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

	// Register project test tasks resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{id}/testtasks", "ZenTao Project Test Tasks"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI template variables (MCP library handles this)
			id := ""
			if idVal, ok := request.Params.Arguments["id"]; ok {
				if idStr, ok := idVal.(string); ok {
					id = idStr
				}
			}

			if id == "" {
				return nil, fmt.Errorf("project ID not found in URI template")
			}

			resp, err := client.Get(fmt.Sprintf("/projects/%s/testtasks", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get project test tasks: %w", err)
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
