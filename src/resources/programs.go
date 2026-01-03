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

func RegisterProgramResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Programs list resource
	programsResource := mcp.NewResource(
		"zentao://programs",
		"ZenTao Programs List",
		mcp.WithResourceDescription("List of all programs in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(programsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=program&f=browse&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get programs list: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://programs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Program kanban resource
	programKanbanResource := mcp.NewResource(
		"zentao://programs/kanban",
		"ZenTao Programs Kanban",
		mcp.WithResourceDescription("Programs kanban view"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(programKanbanResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=program&f=kanban&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get program kanban: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://programs/kanban",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Program detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://programs/{id}", "ZenTao Program Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "programs")

			if id == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=view&t=json&programID=%s", id))
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
		},
	)

	// Program products resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://programs/{programId}/products", "ZenTao Program Products"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			programID := extractProgramIDFromURI(uri, "products")

			if programID == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=product&t=json&programID=%s", programID))
			if err != nil {
				return nil, fmt.Errorf("failed to get program products: %w", err)
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

	// Program projects resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://programs/{programId}/projects", "ZenTao Program Projects"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			programID := extractProgramIDFromURI(uri, "projects")

			if programID == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=project&t=json&programID=%s", programID))
			if err != nil {
				return nil, fmt.Errorf("failed to get program projects: %w", err)
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

	// Program stakeholders resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://programs/{programId}/stakeholders", "ZenTao Program Stakeholders"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			programID := extractProgramIDFromURI(uri, "stakeholders")

			if programID == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=stakeholder&t=json&programID=%s", programID))
			if err != nil {
				return nil, fmt.Errorf("failed to get program stakeholders: %w", err)
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

// Helper function to extract program ID from URI for nested resources
func extractProgramIDFromURI(uri, resourceType string) string {
	// For URI like zentao://programs/123/products, extract 123
	parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
	if len(parts) >= 3 && parts[0] == "programs" && parts[2] == resourceType {
		return parts[1]
	}
	return ""
}