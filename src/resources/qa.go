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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterQaResources(s *server.MCPServer, client *client.ZenTaoClient) {
	qaIndexResource := mcp.NewResource(
		"zentao://qa",
		"ZenTao QA Index",
		mcp.WithResourceDescription("QA module index"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(qaIndexResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=qa&f=index&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get QA index: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://qa",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

