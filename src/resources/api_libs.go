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

func RegisterApiLibResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register API library index resource
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs", "ZenTao API Libraries Index"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			resp, err := client.Get("/index.php?m=api&f=index&t=json")
			if err != nil {
				return nil, fmt.Errorf("failed to get API libraries index: %w", err)
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

	// Register API library releases resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs/{libId}/releases", "ZenTao API Library Releases"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract library ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://api-libs/123/releases, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "api-libs" || parts[2] != "releases" {
				return nil, fmt.Errorf("invalid API library releases URI format: %s", uri)
			}
			libID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=releases&t=json&libID=%s", libID))
			if err != nil {
				return nil, fmt.Errorf("failed to get API library releases: %w", err)
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

	// Register API library structures resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs/{libId}/structs", "ZenTao API Library Structures"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract library ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://api-libs/123/structs, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "api-libs" || parts[2] != "structs" {
				return nil, fmt.Errorf("invalid API library structs URI format: %s", uri)
			}
			libID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=struct&t=json&libID=%s", libID))
			if err != nil {
				return nil, fmt.Errorf("failed to get API library structures: %w", err)
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

	// Register API library structures by release resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs/{libId}/releases/{releaseId}/structs", "ZenTao API Library Structures by Release"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract library ID and release ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://api-libs/123/releases/456/structs, extract 123 and 456
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 5 || parts[0] != "api-libs" || parts[2] != "releases" || parts[4] != "structs" {
				return nil, fmt.Errorf("invalid API library release structs URI format: %s", uri)
			}
			libID := parts[1]
			releaseID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=struct&t=json&libID=%s&releaseID=%s", libID, releaseID))
			if err != nil {
				return nil, fmt.Errorf("failed to get API library structures by release: %w", err)
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

	// Register API resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs/{libId}/apis/{apiId}", "ZenTao API Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract library ID and API ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://api-libs/123/apis/456, extract 123 and 456
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "api-libs" || parts[2] != "apis" {
				return nil, fmt.Errorf("invalid API details URI format: %s", uri)
			}
			libID := parts[1]
			apiID := parts[3]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=view&t=json&libID=%s&apiID=%s", libID, apiID))
			if err != nil {
				return nil, fmt.Errorf("failed to get API details: %w", err)
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

	// Register API library API list resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://api-libs/{libId}/apis", "ZenTao API Library APIs"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract library ID from URI manually
			uri := request.Params.URI
			// For URI like zentao://api-libs/123/apis, extract 123
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "api-libs" || parts[2] != "apis" {
				return nil, fmt.Errorf("invalid API library APIs URI format: %s", uri)
			}
			libID := parts[1]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=index&t=json&libID=%s", libID))
			if err != nil {
				return nil, fmt.Errorf("failed to get API library APIs: %w", err)
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
