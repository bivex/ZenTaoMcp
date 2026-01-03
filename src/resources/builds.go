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

func RegisterBuildResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register build details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://builds/{buildID}", "ZenTao Build Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "builds" {
				return nil, fmt.Errorf("invalid build URI format: %s", uri)
			}
			buildID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=view&t=json&buildID=%s", buildID))
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

	// Register product builds resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{productID}/builds", "ZenTao Product Builds"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "products" || parts[2] != "builds" {
				return nil, fmt.Errorf("invalid product builds URI format: %s", uri)
			}
			productID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetProductBuilds&t=json&productID=%s", productID))
			if err != nil {
				return nil, fmt.Errorf("failed to get product builds: %w", err)
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

	// Register project builds resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectID}/builds", "ZenTao Project Builds"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "builds" {
				return nil, fmt.Errorf("invalid project builds URI format: %s", uri)
			}
			projectID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetProjectBuilds&t=json&projectID=%s", projectID))
			if err != nil {
				return nil, fmt.Errorf("failed to get project builds: %w", err)
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

	// Register execution builds resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/builds", "ZenTao Execution Builds"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "builds" {
				return nil, fmt.Errorf("invalid execution builds URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetExecutionBuilds&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution builds: %w", err)
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