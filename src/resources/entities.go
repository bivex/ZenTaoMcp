package resources

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterStoryResources(s *server.MCPServer, client *client.ZenTaoClient) {
	productStoriesResource := mcp.NewResource(
		"zentao://products/{id}/stories",
		"ZenTao Product Stories",
		mcp.WithResourceDescription("Stories for a specific product (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productStoriesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "products")

		resp, err := client.Get(fmt.Sprintf("/products/%s/stories", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get product stories: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	storyDetailResource := mcp.NewResource(
		"zentao://stories/{id}",
		"ZenTao Story Details",
		mcp.WithResourceDescription("Details of a specific story by ID (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(storyDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "stories")

		resp, err := client.Get(fmt.Sprintf("/stories/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get story details: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

func RegisterTaskResources(s *server.MCPServer, client *client.ZenTaoClient) {
	executionTasksResource := mcp.NewResource(
		"zentao://executions/{id}/tasks",
		"ZenTao Execution Tasks",
		mcp.WithResourceDescription("Tasks for a specific execution (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(executionTasksResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "executions")

		resp, err := client.Get(fmt.Sprintf("/executions/%s/tasks", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get execution tasks: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	taskDetailResource := mcp.NewResource(
		"zentao://tasks/{id}",
		"ZenTao Task Details",
		mcp.WithResourceDescription("Details of a specific task by ID (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(taskDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "tasks")

		resp, err := client.Get(fmt.Sprintf("/tasks/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get task details: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

func RegisterBugResources(s *server.MCPServer, client *client.ZenTaoClient) {
	productBugsResource := mcp.NewResource(
		"zentao://products/{id}/bugs",
		"ZenTao Product Bugs",
		mcp.WithResourceDescription("Bugs for a specific product (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(productBugsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "products")

		resp, err := client.Get(fmt.Sprintf("/products/%s/bugs", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get product bugs: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	bugDetailResource := mcp.NewResource(
		"zentao://bugs/{id}",
		"ZenTao Bug Details",
		mcp.WithResourceDescription("Details of a specific bug by ID (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(bugDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "bugs")

		resp, err := client.Get(fmt.Sprintf("/bugs/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get bug details: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

func RegisterUserResources(s *server.MCPServer, client *client.ZenTaoClient) {
	userDetailResource := mcp.NewResource(
		"zentao://users/{id}",
		"ZenTao User Details",
		mcp.WithResourceDescription("Details of a specific user by ID (use ID in URI)"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(userDetailResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		id := extractIDFromURI(request.Params.Uri, "users")

		resp, err := client.Get(fmt.Sprintf("/users/%s", id))
		if err != nil {
			return nil, fmt.Errorf("failed to get user details: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				Uri:      request.Params.Uri,
				MimeType: "application/json",
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
				Uri:      "zentao://user",
				MimeType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}
