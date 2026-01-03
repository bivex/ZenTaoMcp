// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
//
// Licensed under the MIT License.
// Commercial licensing available upon request.

package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterReleaseTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getProjectReleasesTool := mcp.NewTool("get_project_releases",
		mcp.WithDescription("Get releases for a specific project"),
		mcp.WithNumber("project_id",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of releases to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getProjectReleasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectID := int(args["project_id"].(float64))

		params := make(map[string]string)
		if v, ok := args["limit"]; ok && v != nil {
			params["limit"] = fmt.Sprintf("%d", int(v.(float64)))
		}
		if v, ok := args["offset"]; ok && v != nil {
			params["offset"] = fmt.Sprintf("%d", int(v.(float64)))
		}

		queryString := ""
		if len(params) > 0 {
			queryString = "?"
			for k, v := range params {
				queryString += fmt.Sprintf("%s=%s&", k, v)
			}
			queryString = queryString[:len(queryString)-1] // Remove trailing &
		}

		resp, err := client.Get(fmt.Sprintf("/projects/%d/releases%s", projectID, queryString))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get project releases: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProductReleasesTool := mcp.NewTool("get_product_releases",
		mcp.WithDescription("Get releases for a specific product"),
		mcp.WithNumber("product_id",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of releases to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getProductReleasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product_id"].(float64))

		params := make(map[string]string)
		if v, ok := args["limit"]; ok && v != nil {
			params["limit"] = fmt.Sprintf("%d", int(v.(float64)))
		}
		if v, ok := args["offset"]; ok && v != nil {
			params["offset"] = fmt.Sprintf("%d", int(v.(float64)))
		}

		queryString := ""
		if len(params) > 0 {
			queryString = "?"
			for k, v := range params {
				queryString += fmt.Sprintf("%s=%s&", k, v)
			}
			queryString = queryString[:len(queryString)-1] // Remove trailing &
		}

		resp, err := client.Get(fmt.Sprintf("/products/%d/releases%s", productID, queryString))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get product releases: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
