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
	"strings"

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

	testTaskDetailResource := mcp.NewResource(
		"zentao://testtask/*",
		"ZenTao Test Task Details",
		mcp.WithResourceDescription("Details of a specific test task (use zentao://testtask/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(testTaskDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid test task URI format. Use: zentao://testtask/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("test task ID not specified in URI. Use format: zentao://testtask/123")
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
	})

	projectTestTasksResource := mcp.NewResource(
		"zentao://projects/{id}/testtasks",
		"ZenTao Project Test Tasks",
		mcp.WithResourceDescription("Test tasks for a specific project (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(projectTestTasksResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "projects")

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
	})
}
