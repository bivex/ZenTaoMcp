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

func RegisterSpaceTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseSpacesTool := mcp.NewTool("browse_spaces",
		mcp.WithDescription("Browse spaces with filtering and pagination"),
		mcp.WithNumber("spaceID",
			mcp.Description("Space ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(browseSpacesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["spaceID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&spaceID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=space&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse spaces: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createApplicationTool := mcp.NewTool("create_application",
		mcp.WithDescription("Create an application within a space"),
		mcp.WithNumber("appID",
			mcp.Required(),
			mcp.Description("Application ID"),
		),
		mcp.WithNumber("spaceID",
			mcp.Description("Space ID"),
		),
		mcp.WithString("name",
			mcp.Description("Application name"),
		),
		mcp.WithString("code",
			mcp.Description("Application code"),
		),
	)

	s.AddTool(createApplicationTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("appID=%d", int(args["appID"].(float64)))
		if v, ok := args["spaceID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&spaceID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{}
		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=space&f=createApplication&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create application: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getStoreAppInfoTool := mcp.NewTool("get_store_app_info",
		mcp.WithDescription("Get application store information"),
		mcp.WithNumber("appID",
			mcp.Required(),
			mcp.Description("Application ID"),
		),
	)

	s.AddTool(getStoreAppInfoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=space&f=getStoreAppInfo&t=json&appID=%d", int(args["appID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get store app info: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
