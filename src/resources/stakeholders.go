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

func RegisterStakeholderResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Stakeholders list resource
	stakeholdersResource := mcp.NewResource(
		"zentao://stakeholders",
		"ZenTao Stakeholders List",
		mcp.WithResourceDescription("List of all stakeholders in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(stakeholdersResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=stakeholder&f=browse&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get stakeholders list: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://stakeholders",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Stakeholder detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://stakeholders/{id}", "ZenTao Stakeholder Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract stakeholder ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "stakeholders")

			if id == "" {
				return nil, fmt.Errorf("stakeholder ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=view&t=json&stakeholderID=%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get stakeholder details: %w", err)
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

	// Project stakeholders resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectId}/stakeholders", "ZenTao Project Stakeholders"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract project ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://projects/123/stakeholders, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "stakeholders" {
				return nil, fmt.Errorf("invalid project stakeholders URI format: %s", uri)
			}
			projectID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=browse&t=json&projectID=%s", projectID))
			if err != nil {
				return nil, fmt.Errorf("failed to get project stakeholders: %w", err)
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

	// Stakeholder issues resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://stakeholders/{id}/issues", "ZenTao Stakeholder Issues"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract stakeholder ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://stakeholders/123/issues, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "stakeholders" || parts[2] != "issues" {
				return nil, fmt.Errorf("invalid stakeholder issues URI format: %s", uri)
			}
			stakeholderID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=userIssue&t=json&stakeholderID=%s", stakeholderID))
			if err != nil {
				return nil, fmt.Errorf("failed to get stakeholder issues: %w", err)
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
