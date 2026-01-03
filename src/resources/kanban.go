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

func RegisterKanbanResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register kanban spaces resource
	s.AddResource(mcp.NewResource(
		"zentao://kanban/spaces",
		"ZenTao Kanban Spaces",
		mcp.WithResourceDescription("List of all kanban spaces"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=kanban&f=space&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get kanban spaces: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://kanban/spaces",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register kanban space details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/spaces/{spaceID}", "ZenTao Kanban Space Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "kanban" || parts[1] != "spaces" {
				return nil, fmt.Errorf("invalid kanban space URI format: %s", uri)
			}
			// Get space details (this might need adjustment based on actual API)
			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=space&t=json"))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban space details: %w", err)
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

	// Register kanban board details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}", "ZenTao Kanban Board Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 2 || parts[0] != "kanban" {
				return nil, fmt.Errorf("invalid kanban URI format: %s", uri)
			}
			kanbanID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=view&t=json&kanbanID=%s", kanbanID))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban board details: %w", err)
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

	// Register kanban region details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}/regions/{regionID}", "ZenTao Kanban Region Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "kanban" || parts[2] != "regions" {
				return nil, fmt.Errorf("invalid kanban region URI format: %s", uri)
			}
			kanbanID := parts[1]
			regionID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=view&t=json&kanbanID=%s&regionID=%s", kanbanID, regionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban region details: %w", err)
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

	// Register kanban lane details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}/regions/{regionID}/lanes/{laneID}", "ZenTao Kanban Lane Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 6 || parts[0] != "kanban" || parts[2] != "regions" || parts[4] != "lanes" {
				return nil, fmt.Errorf("invalid kanban lane URI format: %s", uri)
			}
			regionID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=ajaxGetLanes&t=json&regionID=%s", regionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban lane details: %w", err)
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

	// Register kanban column details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}/regions/{regionID}/lanes/{laneID}/columns/{columnID}", "ZenTao Kanban Column Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 8 || parts[0] != "kanban" || parts[2] != "regions" || parts[4] != "lanes" || parts[6] != "columns" {
				return nil, fmt.Errorf("invalid kanban column URI format: %s", uri)
			}
			laneID := parts[5]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=ajaxGetColumns&t=json&laneID=%s", laneID))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban column details: %w", err)
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

	// Register kanban card details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/cards/{cardID}", "ZenTao Kanban Card Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "kanban" || parts[1] != "cards" {
				return nil, fmt.Errorf("invalid kanban card URI format: %s", uri)
			}
			cardID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=viewCard&t=json&cardID=%s", cardID))
			if err != nil {
				return nil, fmt.Errorf("failed to get kanban card details: %w", err)
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

	// Register archived cards resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}/regions/{regionID}/archived-cards", "ZenTao Kanban Archived Cards"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 5 || parts[0] != "kanban" || parts[2] != "regions" || parts[4] != "archived-cards" {
				return nil, fmt.Errorf("invalid archived cards URI format: %s", uri)
			}
			regionID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=viewArchivedCard&t=json&regionID=%s", regionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get archived kanban cards: %w", err)
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

	// Register archived columns resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://kanban/{kanbanID}/regions/{regionID}/archived-columns", "ZenTao Kanban Archived Columns"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 5 || parts[0] != "kanban" || parts[2] != "regions" || parts[4] != "archived-columns" {
				return nil, fmt.Errorf("invalid archived columns URI format: %s", uri)
			}
			regionID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=viewArchivedColumn&t=json&regionID=%s", regionID))
			if err != nil {
				return nil, fmt.Errorf("failed to get archived kanban columns: %w", err)
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
