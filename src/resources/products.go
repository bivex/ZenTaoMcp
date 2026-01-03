// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
//
//
// Licensed under the MIT License.
// Commercial licensing available upon request.

package resources

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func extractIDFromURI(uri, resourceType string) string {
	re := regexp.MustCompile(fmt.Sprintf("%s/([^/]+)", resourceType))
	matches := re.FindStringSubmatch(uri)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func RegisterProductResources(s *server.MCPServer, client *client.ZenTaoClient) {
	productListResource := mcp.NewResource(
		"zentao://products",
		"ZenTao Product List",
		mcp.WithResourceDescription("List of all products in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/products")
		if err != nil {
			return nil, fmt.Errorf("failed to get products: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://products",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Note: MCP library may not support URI templates like {id}
	// For now, we'll implement individual resources using query parameters
	// Register product detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://product/{id}", "ZenTao Product Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "product")

			if id == "" {
				return nil, fmt.Errorf("product ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/products/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get product details: %w", err)
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

func RegisterProjectResources(s *server.MCPServer, client *client.ZenTaoClient) {
	projectListResource := mcp.NewResource(
		"zentao://projects",
		"ZenTao Project List",
		mcp.WithResourceDescription("List of all projects in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(projectListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/projects")
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://projects",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register project detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://project/{id}", "ZenTao Project Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "project")

			if id == "" {
				return nil, fmt.Errorf("project ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/projects/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get project details: %w", err)
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

	// Register project executions resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectId}/executions", "ZenTao Project Executions"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract project ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://projects/123/executions, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "executions" {
				return nil, fmt.Errorf("invalid project executions URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/projects/%s/executions", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get project executions: %w", err)
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

	// Register project stories resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://projects/{projectId}/stories", "ZenTao Project Stories"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract project ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://projects/123/stories, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "projects" || parts[2] != "stories" {
				return nil, fmt.Errorf("invalid project stories URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/projects/%s/stories", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get project stories: %w", err)
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

	// Register execution detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://execution/{id}", "ZenTao Execution Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "execution")

			if id == "" {
				return nil, fmt.Errorf("execution ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/executions/%s", id))
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
}
