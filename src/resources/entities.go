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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterStoryResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /stories endpoint in ZenTao API
	// Stories are accessed via projects/products/executions

	// Register product stories resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{id}/stories", "ZenTao Product Stories"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "products")

			if id == "" {
				return nil, fmt.Errorf("product ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/products/%s/stories", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get product stories: %w", err)
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

	// Register story detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://story/{id}", "ZenTao Story Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "story")

			if id == "" {
				return nil, fmt.Errorf("story ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/stories/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get story details: %w", err)
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

func RegisterTaskResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /tasks endpoint in ZenTao API
	// Tasks are accessed via executions

	// Register execution tasks resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://executions/{id}/tasks", "ZenTao Execution Tasks"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "executions")

			if id == "" {
				return nil, fmt.Errorf("execution ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/executions/%s/tasks", id))
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

	// Register task detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://task/{id}", "ZenTao Task Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "task")

			if id == "" {
				return nil, fmt.Errorf("task ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/tasks/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get task details: %w", err)
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

func RegisterBugResources(s *server.MCPServer, client *client.ZenTaoClient) {
	bugsListResource := mcp.NewResource(
		"zentao://bugs",
		"ZenTao Bugs List",
		mcp.WithResourceDescription("List of all bugs in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(bugsListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/bugs")
		if err != nil {
			return nil, fmt.Errorf("failed to get bugs: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://bugs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register product bugs resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://products/{id}/bugs", "ZenTao Product Bugs"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "products")

			if id == "" {
				return nil, fmt.Errorf("product ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/products/%s/bugs", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get product bugs: %w", err)
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

	// Register bug detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://bug/{id}", "ZenTao Bug Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "bug")

			if id == "" {
				return nil, fmt.Errorf("bug ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/bugs/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get bug details: %w", err)
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

func RegisterUserResources(s *server.MCPServer, client *client.ZenTaoClient) {
	usersListResource := mcp.NewResource(
		"zentao://users",
		"ZenTao Users List",
		mcp.WithResourceDescription("List of all users in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(usersListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/users")
		if err != nil {
			return nil, fmt.Errorf("failed to get users: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://users",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register user detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://user/{id}", "ZenTao User Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract ID from URI manually since MCP library may not populate Arguments
			uri := request.Params.URI
			id := extractIDFromURI(uri, "user")

			if id == "" {
				return nil, fmt.Errorf("user ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/users/%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get user details: %w", err)
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

	myProfileResource := mcp.NewResource(
		"zentao://user",
		"ZenTao My Profile",
		mcp.WithResourceDescription("Current user's profile"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myProfileResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/user")
		if err != nil {
			return nil, fmt.Errorf("failed to get user profile: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://user",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
