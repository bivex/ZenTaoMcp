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

func RegisterTodoResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Todos list resource
	todosResource := mcp.NewResource(
		"zentao://todos",
		"ZenTao Todos List",
		mcp.WithResourceDescription("List of all todos"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(todosResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=todo&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get todos: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://todos",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Todo detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://todos/{id}", "ZenTao Todo Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract todo ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "todos")

			if id == "" {
				return nil, fmt.Errorf("todo ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=view&t=json&todoID=%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get todo details: %w", err)
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

	// Todo detail via AJAX resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://todos/{id}/detail", "ZenTao Todo AJAX Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract todo ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://todos/123/detail, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "todos" || parts[2] != "detail" {
				return nil, fmt.Errorf("invalid todo detail URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=ajaxGetDetail&t=json&todoID=%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get todo AJAX details: %w", err)
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

	// Todos by status resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://todos/status/{status}", "ZenTao Todos by Status"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract status from URI manually
			uri := request.Params.URI
			// For URI like zentao://todos/status/wait, extract "wait"
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "todos" || parts[1] != "status" {
				return nil, fmt.Errorf("invalid todos by status URI format: %s", uri)
			}
			status := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=todo&t=json&status=%s", status))
			if err != nil {
				return nil, fmt.Errorf("failed to get todos by status: %w", err)
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

	// Todos by type resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://todos/type/{type}", "ZenTao Todos by Type"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract type from URI manually
			uri := request.Params.URI
			// For URI like zentao://todos/type/custom, extract "custom"
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "todos" || parts[1] != "type" {
				return nil, fmt.Errorf("invalid todos by type URI format: %s", uri)
			}
			todoType := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=todo&t=json&type=%s", todoType))
			if err != nil {
				return nil, fmt.Errorf("failed to get todos by type: %w", err)
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
