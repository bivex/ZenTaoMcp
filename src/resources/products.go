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
	productDetailResource := mcp.NewResource(
		"zentao://product/*",
		"ZenTao Product Details",
		mcp.WithResourceDescription("Details of a specific product (use zentao://product/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid product URI format. Use: zentao://product/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("product ID not specified in URI. Use format: zentao://product/123")
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
	})
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

	projectDetailResource := mcp.NewResource(
		"zentao://project/*",
		"ZenTao Project Details",
		mcp.WithResourceDescription("Details of a specific project (use zentao://project/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(projectDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid project URI format. Use: zentao://project/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("project ID not specified in URI. Use format: zentao://project/123")
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
	})

	projectExecutionsResource := mcp.NewResource(
		"zentao://projects/{id}/executions",
		"ZenTao Project Executions",
		mcp.WithResourceDescription("Executions for a specific project (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(projectExecutionsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "projects")

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
	})

	projectStoriesResource := mcp.NewResource(
		"zentao://projects/{id}/stories",
		"ZenTao Project Stories",
		mcp.WithResourceDescription("Stories for a specific project (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(projectStoriesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "projects")

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
	})

	executionDetailResource := mcp.NewResource(
		"zentao://execution/*",
		"ZenTao Execution Details",
		mcp.WithResourceDescription("Details of a specific execution (use zentao://execution/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(executionDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid execution URI format. Use: zentao://execution/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("execution ID not specified in URI. Use format: zentao://execution/123")
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
	})
}
