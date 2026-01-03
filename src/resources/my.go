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

func RegisterMyResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Dashboard resource
	myDashboardResource := mcp.NewResource(
		"zentao://my/dashboard",
		"ZenTao My Dashboard",
		mcp.WithResourceDescription("User's personal dashboard in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myDashboardResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=index&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get dashboard: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/dashboard",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Profile resource
	myProfileResource := mcp.NewResource(
		"zentao://my/profile",
		"ZenTao My Profile",
		mcp.WithResourceDescription("User's profile information"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myProfileResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=profile&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get profile: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/profile",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Calendar resource
	myCalendarResource := mcp.NewResource(
		"zentao://my/calendar",
		"ZenTao My Calendar",
		mcp.WithResourceDescription("User's calendar view"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myCalendarResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=calendar&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get calendar: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/calendar",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal work resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://my/work/{mode}", "ZenTao My Work"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract mode from URI manually
			uri := request.Params.URI
			mode := extractModeFromURI(uri, "work")

			if mode == "" {
				return nil, fmt.Errorf("work mode not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=work&t=json&mode=%s", mode))
			if err != nil {
				return nil, fmt.Errorf("failed to get work: %w", err)
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

	// Personal todos resource
	myTodosResource := mcp.NewResource(
		"zentao://my/todos",
		"ZenTao My Todos",
		mcp.WithResourceDescription("User's personal todos"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myTodosResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=todo&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get todos: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/todos",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal stories resource
	myStoriesResource := mcp.NewResource(
		"zentao://my/stories",
		"ZenTao My Stories",
		mcp.WithResourceDescription("User's personal stories"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myStoriesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=story&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get stories: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/stories",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal tasks resource
	myTasksResource := mcp.NewResource(
		"zentao://my/tasks",
		"ZenTao My Tasks",
		mcp.WithResourceDescription("User's personal tasks"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myTasksResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=task&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get tasks: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/tasks",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal bugs resource
	myBugsResource := mcp.NewResource(
		"zentao://my/bugs",
		"ZenTao My Bugs",
		mcp.WithResourceDescription("User's personal bugs"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myBugsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=bug&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get bugs: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/bugs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal projects resource
	myProjectsResource := mcp.NewResource(
		"zentao://my/projects",
		"ZenTao My Projects",
		mcp.WithResourceDescription("User's personal projects"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myProjectsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=project&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get projects: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/projects",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Personal executions resource
	myExecutionsResource := mcp.NewResource(
		"zentao://my/executions",
		"ZenTao My Executions",
		mcp.WithResourceDescription("User's personal executions"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myExecutionsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=execution&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get executions: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/executions",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Team resource
	myTeamResource := mcp.NewResource(
		"zentao://my/team",
		"ZenTao My Team",
		mcp.WithResourceDescription("User's team information"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myTeamResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=team&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get team: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/team",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Dynamic/Activity feed resource
	myDynamicResource := mcp.NewResource(
		"zentao://my/dynamic",
		"ZenTao My Dynamic",
		mcp.WithResourceDescription("User's activity feed"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(myDynamicResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=my&f=dynamic&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get dynamic: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://my/dynamic",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

// Helper function to extract mode from URI
func extractModeFromURI(uri, resourceType string) string {
	// For URI like zentao://my/work/contribution, extract "contribution"
	parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
	if len(parts) >= 3 && parts[0] == "my" && parts[1] == resourceType {
		return parts[2]
	}
	return ""
}
