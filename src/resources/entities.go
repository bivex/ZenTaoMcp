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
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterStoryResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /stories endpoint in ZenTao API
	// Stories are accessed via projects/products/executions

	productStoriesResource := mcp.NewResource(
		"zentao://products/{id}/stories",
		"ZenTao Product Stories",
		mcp.WithResourceDescription("Stories for a specific product (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productStoriesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "products")

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
	})

	storyDetailResource := mcp.NewResource(
		"zentao://story/*",
		"ZenTao Story Details",
		mcp.WithResourceDescription("Details of a specific story (use zentao://story/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(storyDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid story URI format. Use: zentao://story/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("story ID not specified in URI. Use format: zentao://story/123")
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
	})
}

func RegisterTaskResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Note: No general /tasks endpoint in ZenTao API
	// Tasks are accessed via executions

	executionTasksResource := mcp.NewResource(
		"zentao://executions/{id}/tasks",
		"ZenTao Execution Tasks",
		mcp.WithResourceDescription("Tasks for a specific execution (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(executionTasksResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "executions")

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
	})

	taskDetailResource := mcp.NewResource(
		"zentao://task/*",
		"ZenTao Task Details",
		mcp.WithResourceDescription("Details of a specific task (use zentao://task/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(taskDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid task URI format. Use: zentao://task/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("task ID not specified in URI. Use format: zentao://task/123")
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
	})
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

	productBugsResource := mcp.NewResource(
		"zentao://products/{id}/bugs",
		"ZenTao Product Bugs",
		mcp.WithResourceDescription("Bugs for a specific product (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productBugsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.URI, "products")

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
	})

	bugDetailResource := mcp.NewResource(
		"zentao://bug/*",
		"ZenTao Bug Details",
		mcp.WithResourceDescription("Details of a specific bug (use zentao://bug/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(bugDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid bug URI format. Use: zentao://bug/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("bug ID not specified in URI. Use format: zentao://bug/123")
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
	})
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

	userDetailResource := mcp.NewResource(
		"zentao://user/*",
		"ZenTao User Details",
		mcp.WithResourceDescription("Details of a specific user (use zentao://user/123)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(userDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Extract ID from URI path
		uriParts := strings.Split(request.Params.URI, "/")
		if len(uriParts) < 3 {
			return nil, fmt.Errorf("invalid user URI format. Use: zentao://user/123")
		}
		id := uriParts[len(uriParts)-1]

		if id == "" || id == "*" {
			return nil, fmt.Errorf("user ID not specified in URI. Use format: zentao://user/123")
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
	})

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
