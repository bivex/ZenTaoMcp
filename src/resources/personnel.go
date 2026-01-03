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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterPersonnelResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Personnel accessible list resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://personnel/accessible/{programID}", "ZenTao Personnel Accessible"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			programID := extractIDFromURI(uri, "accessible")

			if programID == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=accessible&t=json&programID=%s", programID))
			if err != nil {
				return nil, fmt.Errorf("failed to get accessible personnel: %w", err)
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

	// Personnel invest resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://personnel/invest/{programID}", "ZenTao Personnel Invest"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract program ID from URI manually
			uri := request.Params.URI
			programID := extractIDFromURI(uri, "invest")

			if programID == "" {
				return nil, fmt.Errorf("program ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=invest&t=json&programID=%s", programID))
			if err != nil {
				return nil, fmt.Errorf("failed to get personnel invest: %w", err)
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

	// Personnel whitelist resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://personnel/whitelist/{objectID}", "ZenTao Personnel Whitelist"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			// Extract object ID from URI manually
			uri := request.Params.URI
			objectID := extractIDFromURI(uri, "whitelist")

			if objectID == "" {
				return nil, fmt.Errorf("object ID not found in URI: %s", uri)
			}

			resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=whitelist&t=json&objectID=%s", objectID))
			if err != nil {
				return nil, fmt.Errorf("failed to get personnel whitelist: %w", err)
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
