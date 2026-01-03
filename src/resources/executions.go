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

func RegisterExecutionResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register execution details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}", "ZenTao Execution Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "executions" {
				return nil, fmt.Errorf("invalid execution URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=view&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution details: %w", err)
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

	// Register execution tasks resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/tasks", "ZenTao Execution Tasks"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "tasks" {
				return nil, fmt.Errorf("invalid execution tasks URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=task&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution tasks: %w", err)
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

	// Register execution stories resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/stories", "ZenTao Execution Stories"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "stories" {
				return nil, fmt.Errorf("invalid execution stories URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=story&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution stories: %w", err)
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

	// Register execution bugs resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/bugs", "ZenTao Execution Bugs"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "bugs" {
				return nil, fmt.Errorf("invalid execution bugs URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=bug&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution bugs: %w", err)
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

	// Register execution team resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/team", "ZenTao Execution Team"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "team" {
				return nil, fmt.Errorf("invalid execution team URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=team&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution team: %w", err)
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

	// Register execution burn chart resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/burn", "ZenTao Execution Burn Chart"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "burn" {
				return nil, fmt.Errorf("invalid execution burn URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=ajaxGetBurn&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution burn chart: %w", err)
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

	// Register execution CFD resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/cfd", "ZenTao Execution CFD"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "cfd" {
				return nil, fmt.Errorf("invalid execution CFD URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=ajaxGetCFD&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution CFD: %w", err)
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

	// Register execution kanban resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/kanban", "ZenTao Execution Kanban"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "kanban" {
				return nil, fmt.Errorf("invalid execution kanban URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=ajaxGetExecutionKanban&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution kanban: %w", err)
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

	// Register all executions resource
	s.AddResource(mcp.NewResource(
		"zentao://executions",
		"ZenTao All Executions",
		mcp.WithResourceDescription("List of all executions in ZenTao"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=execution&f=all&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get all executions: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://executions",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
