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

func RegisterEpicResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register epic details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://epics/{epicID}", "ZenTao Epic Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "epics" {
				return nil, fmt.Errorf("invalid epic URI format: %s", uri)
			}
			epicID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=view&t=json&storyID=%s", epicID))
			if err != nil {
				return nil, fmt.Errorf("failed to get epic details: %w", err)
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

	// Register epic stories resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://epics/{epicID}/stories", "ZenTao Epic Stories"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "epics" || parts[2] != "stories" {
				return nil, fmt.Errorf("invalid epic stories URI format: %s", uri)
			}
			epicID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=linkStories&t=json&storyID=%s", epicID))
			if err != nil {
				return nil, fmt.Errorf("failed to get epic stories: %w", err)
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

	// Register epic requirements resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://epics/{epicID}/requirements", "ZenTao Epic Requirements"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "epics" || parts[2] != "requirements" {
				return nil, fmt.Errorf("invalid epic requirements URI format: %s", uri)
			}
			epicID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=linkRequirements&t=json&storyID=%s", epicID))
			if err != nil {
				return nil, fmt.Errorf("failed to get epic requirements: %w", err)
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

	// Register product epics resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{productID}/epics", "ZenTao Product Epics"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "products" || parts[2] != "epics" {
				return nil, fmt.Errorf("invalid product epics URI format: %s", uri)
			}
			productID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=import&t=json&productID=%s", productID))
			if err != nil {
				return nil, fmt.Errorf("failed to get product epics: %w", err)
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

	// Register project epics resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectID}/epics", "ZenTao Project Epics"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "epics" {
				return nil, fmt.Errorf("invalid project epics URI format: %s", uri)
			}
			projectID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=import&t=json&projectID=%s", projectID))
			if err != nil {
				return nil, fmt.Errorf("failed to get project epics: %w", err)
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

	// Register execution epics resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/epics", "ZenTao Execution Epics"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "epics" {
				return nil, fmt.Errorf("invalid execution epics URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=import&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution epics: %w", err)
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
