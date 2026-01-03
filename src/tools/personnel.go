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

func RegisterPersonnelTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getAccessiblePersonnelTool := mcp.NewTool("get_accessible_personnel",
		mcp.WithDescription("Get accessible personnel list"),
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
		mcp.WithNumber("deptID",
			mcp.Description("Department ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
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

	s.AddTool(getAccessiblePersonnelTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["programID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&programID=%d", int(v.(float64)))
		}
		if v, ok := args["deptID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=accessible&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get accessible personnel: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getPersonnelInvestTool := mcp.NewTool("get_personnel_invest",
		mcp.WithDescription("Get personnel investment information"),
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
	)

	s.AddTool(getPersonnelInvestTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["programID"]; ok && v != nil {
			queryParams = fmt.Sprintf("programID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=invest&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get personnel invest: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getPersonnelWhitelistTool := mcp.NewTool("get_personnel_whitelist",
		mcp.WithDescription("Get personnel whitelist"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithString("module",
			mcp.Description("Module type"),
			mcp.Enum("personnel", "program", "project", "product"),
		),
		mcp.WithString("objectType",
			mcp.Description("Object type"),
			mcp.Enum("program", "project", "product", "sprint"),
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
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("project", "program", "programproject"),
		),
	)

	s.AddTool(getPersonnelWhitelistTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("objectID=%d", int(args["objectID"].(float64)))
		if v, ok := args["module"]; ok && v != nil {
			queryParams += fmt.Sprintf("&module=%s", v)
		}
		if v, ok := args["objectType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectType=%s", v)
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
		if v, ok := args["programID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&programID=%d", int(v.(float64)))
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=whitelist&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get personnel whitelist: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	addPersonnelWhitelistTool := mcp.NewTool("add_personnel_whitelist",
		mcp.WithDescription("Add personnel to whitelist"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithNumber("deptID",
			mcp.Description("Department ID"),
		),
		mcp.WithNumber("copyID",
			mcp.Description("Copy from ID"),
		),
		mcp.WithString("objectType",
			mcp.Description("Object type"),
			mcp.Enum("program", "project", "product", "sprint"),
		),
		mcp.WithString("module",
			mcp.Description("Module"),
		),
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("project", "program", "programproject"),
		),
		mcp.WithArray("users",
			mcp.Description("User IDs to add"),
		),
	)

	s.AddTool(addPersonnelWhitelistTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["users"]; ok && v != nil {
			body["users"] = v
		}

		queryParams := fmt.Sprintf("objectID=%d", int(args["objectID"].(float64)))
		if v, ok := args["deptID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}
		if v, ok := args["copyID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&copyID=%d", int(v.(float64)))
		}
		if v, ok := args["objectType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectType=%s", v)
		}
		if v, ok := args["module"]; ok && v != nil {
			queryParams += fmt.Sprintf("&module=%s", v)
		}
		if v, ok := args["programID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&programID=%d", int(v.(float64)))
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=personnel&f=addWhitelist&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to add personnel whitelist: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unbindPersonnelWhitelistTool := mcp.NewTool("unbind_personnel_whitelist",
		mcp.WithDescription("Remove personnel from whitelist"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Whitelist entry ID"),
		),
	)

	s.AddTool(unbindPersonnelWhitelistTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=personnel&f=unbindWhitelist&t=json&id=%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unbind personnel whitelist: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
