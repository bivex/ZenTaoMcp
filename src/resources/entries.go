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

func RegisterEntryResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register entries list resource
	entriesListResource := mcp.NewResource(
		"zentao://entries",
		"ZenTao Entries List",
		mcp.WithResourceDescription("List of all entries in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(entriesListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=entry&f=browse&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get entries list: %w", err)
		}

		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://entries",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register entry detail resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://entries/{id}", "ZenTao Entry Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract entry ID from URI manually
			uri := request.Params.URI
			id := extractIDFromURI(uri, "entries")

			if id == "" {
				return nil, fmt.Errorf("entry ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=entry&f=log&t=json&id=%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get entry details: %w", err)
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

	// Register entry log resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://entries/{id}/log", "ZenTao Entry Log"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract entry ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://entries/123/log, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "entries" || parts[2] != "log" {
				return nil, fmt.Errorf("invalid entry log URI format: %s", uri)
			}
			id := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=entry&f=log&t=json&id=%s", id))
			if err != nil {
				return nil, fmt.Errorf("failed to get entry log: %w", err)
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
