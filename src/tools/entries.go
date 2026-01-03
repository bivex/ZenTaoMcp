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

func RegisterEntryTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createEntryTool := mcp.NewTool("create_entry",
		mcp.WithDescription("Create a new entry in ZenTao"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Entry name"),
		),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("Entry code"),
		),
		mcp.WithString("key",
			mcp.Description("Entry key"),
		),
		mcp.WithString("ip",
			mcp.Description("IP address"),
		),
		mcp.WithString("desc",
			mcp.Description("Entry description"),
		),
		mcp.WithString("version",
			mcp.Description("Version"),
		),
	)

	s.AddTool(createEntryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
			"code": args["code"],
		}

		if v, ok := args["key"]; ok && v != nil {
			body["key"] = v
		}
		if v, ok := args["ip"]; ok && v != nil {
			body["ip"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["version"]; ok && v != nil {
			body["version"] = v
		}

		resp, err := client.Post("/index.php?m=entry&f=create&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create entry: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editEntryTool := mcp.NewTool("edit_entry",
		mcp.WithDescription("Edit an existing entry in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Entry ID"),
		),
		mcp.WithString("name",
			mcp.Description("Entry name"),
		),
		mcp.WithString("code",
			mcp.Description("Entry code"),
		),
		mcp.WithString("key",
			mcp.Description("Entry key"),
		),
		mcp.WithString("ip",
			mcp.Description("IP address"),
		),
		mcp.WithString("desc",
			mcp.Description("Entry description"),
		),
		mcp.WithString("version",
			mcp.Description("Version"),
		),
	)

	s.AddTool(editEntryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}
		if v, ok := args["key"]; ok && v != nil {
			body["key"] = v
		}
		if v, ok := args["ip"]; ok && v != nil {
			body["ip"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["version"]; ok && v != nil {
			body["version"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=entry&f=edit&t=json&id=%d", int(args["id"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit entry: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteEntryTool := mcp.NewTool("delete_entry",
		mcp.WithDescription("Delete an entry from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Entry ID to delete"),
		),
	)

	s.AddTool(deleteEntryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=entry&f=delete&t=json&id=%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete entry: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	browseEntriesTool := mcp.NewTool("browse_entries",
		mcp.WithDescription("Browse entries in ZenTao"),
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

	s.AddTool(browseEntriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=entry&f=browse&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse entries: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getEntryLogTool := mcp.NewTool("get_entry_log",
		mcp.WithDescription("Get log for an entry in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Entry ID"),
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

	s.AddTool(getEntryLogTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		queryParams := fmt.Sprintf("id=%d", id)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=entry&f=log&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get entry log: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
