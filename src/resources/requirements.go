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

func RegisterRequirementResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register requirement details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://requirements/{requirementID}", "ZenTao Requirement Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "requirements" {
				return nil, fmt.Errorf("invalid requirement URI format: %s", uri)
			}
			requirementID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=view&t=json&storyID=%s", requirementID))
			if err != nil {
				return nil, fmt.Errorf("failed to get requirement details: %w", err)
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

	// Register requirement stories resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://requirements/{requirementID}/stories", "ZenTao Requirement Stories"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "requirements" || parts[2] != "stories" {
				return nil, fmt.Errorf("invalid requirement stories URI format: %s", uri)
			}
			requirementID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=linkStory&t=json&storyID=%s", requirementID))
			if err != nil {
				return nil, fmt.Errorf("failed to get requirement stories: %w", err)
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

	// Register requirement linked requirements resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://requirements/{requirementID}/linked-requirements", "ZenTao Requirement Linked Requirements"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "requirements" || parts[2] != "linked-requirements" {
				return nil, fmt.Errorf("invalid requirement linked requirements URI format: %s", uri)
			}
			requirementID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=linkRequirements&t=json&storyID=%s", requirementID))
			if err != nil {
				return nil, fmt.Errorf("failed to get requirement linked requirements: %w", err)
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

	// Register product requirements resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{productID}/requirements", "ZenTao Product Requirements"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "products" || parts[2] != "requirements" {
				return nil, fmt.Errorf("invalid product requirements URI format: %s", uri)
			}
			productID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=import&t=json&productID=%s", productID))
			if err != nil {
				return nil, fmt.Errorf("failed to get product requirements: %w", err)
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

	// Register project requirements resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectID}/requirements", "ZenTao Project Requirements"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "requirements" {
				return nil, fmt.Errorf("invalid project requirements URI format: %s", uri)
			}
			projectID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=import&t=json&projectID=%s", projectID))
			if err != nil {
				return nil, fmt.Errorf("failed to get project requirements: %w", err)
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

	// Register execution requirements resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{executionID}/requirements", "ZenTao Execution Requirements"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "executions" || parts[2] != "requirements" {
				return nil, fmt.Errorf("invalid execution requirements URI format: %s", uri)
			}
			executionID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=import&t=json&executionID=%s", executionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get execution requirements: %w", err)
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
