// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
//
//
// Licensed under MIT License.
// Commercial licensing available upon request.

package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

// RegisterSearchTools registers all search/query tools
func RegisterSearchTools(s *server.MCPServer, client *client.ZenTaoClient) {
	registerSearchFormTools(s, client)
	registerSearchQueryTools(s, client)
	registerSearchIndexTools(s, client)
}

func registerSearchFormTools(s *server.MCPServer, client *client.ZenTaoClient) {
	buildFormTool := mcp.NewTool("search_build_form",
		mcp.WithDescription("Build search form"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
		mcp.WithString("mode",
			mcp.Description("Mode: new20 (new page) | old20 (old page)"),
			mcp.Enum("new20", "old20"),
		),
	)

	s.AddTool(buildFormTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["module"] = args["module"].(string)
		params["mode"] = args["mode"].(string)

		resp, err := client.Get("/index.php?m=search&f=buildForm&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to build form: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	buildOldFormTool := mcp.NewTool("search_build_old_form",
		mcp.WithDescription("Build old search form"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
	)

	s.AddTool(buildOldFormTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["module"] = args["module"].(string)

		resp, err := client.Get("/index.php?m=search&f=buildOldForm&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to build old form: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerSearchQueryTools(s *server.MCPServer, client *client.ZenTaoClient) {
	buildQueryTool := mcp.NewTool("search_build_query",
		mcp.WithDescription("Build search query"),
		mcp.WithString("mode",
			mcp.Description("Mode: new20 (new page) | old20 (old page)"),
			mcp.Enum("new20", "old20"),
		),
	)

	s.AddTool(buildQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["mode"] = args["mode"].(string)

		resp, err := client.Post("/index.php?m=search&f=buildQuery&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to build query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	buildOldQueryTool := mcp.NewTool("search_build_old_query",
		mcp.WithDescription("Build old search query"),
	)

	s.AddTool(buildOldQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=search&f=buildOldQuery&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to build old query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	saveQueryTool := mcp.NewTool("search_save_query",
		mcp.WithDescription("Save search query"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
		mcp.WithString("onMenuBar",
			mcp.Description("On menu bar flag"),
		),
	)

	s.AddTool(saveQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["module"] = args["module"].(string)

		if v, ok := args["onMenuBar"].(string); ok {
			params["onMenuBar"] = v
		}

		resp, err := client.Post("/index.php?m=search&f=saveQuery&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	saveOldQueryTool := mcp.NewTool("search_save_old_query",
		mcp.WithDescription("Save old search query"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
		mcp.WithString("onMenuBar",
			mcp.Description("On menu bar flag"),
		),
	)

	s.AddTool(saveOldQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["module"] = args["module"].(string)

		if v, ok := args["onMenuBar"].(string); ok {
			params["onMenuBar"] = v
		}

		resp, err := client.Post("/index.php?m=search&f=saveOldQuery&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save old query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteQueryTool := mcp.NewTool("search_delete_query",
		mcp.WithDescription("Delete search query"),
		mcp.WithNumber("queryID",
			mcp.Required(),
			mcp.Description("Query ID"),
		),
	)

	s.AddTool(deleteQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		queryID := int(args["queryID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=search&f=deleteQuery&t=json&queryID=%d", queryID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxGetQueryTool := mcp.NewTool("search_ajax_get_query",
		mcp.WithDescription("Get search query details"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module name"),
		),
		mcp.WithNumber("queryID",
			mcp.Required(),
			mcp.Description("Query ID"),
		),
	)

	s.AddTool(ajaxGetQueryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=search&f=ajaxGetQuery&t=json"

		url += fmt.Sprintf("&module=%s", args["module"].(string))
		url += fmt.Sprintf("&queryID=%d", int(args["queryID"].(float64)))

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get query: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxRemoveMenuTool := mcp.NewTool("search_ajax_remove_menu",
		mcp.WithDescription("Remove query from menu"),
		mcp.WithNumber("queryID",
			mcp.Required(),
			mcp.Description("Query ID"),
		),
	)

	s.AddTool(ajaxRemoveMenuTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		queryID := int(args["queryID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=search&f=ajaxRemoveMenu&t=json&queryID=%d", queryID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to remove menu: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerSearchIndexTools(s *server.MCPServer, client *client.ZenTaoClient) {
	buildIndexTool := mcp.NewTool("search_build_index",
		mcp.WithDescription("Build search index"),
		mcp.WithString("mode",
			mcp.Description("Mode: show|build"),
			mcp.Enum("show", "build"),
		),
		mcp.WithString("type", mcp.Description("Type")),
		mcp.WithNumber("lastID", mcp.Description("Last ID")),
	)

	s.AddTool(buildIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["mode"].(string); ok {
			params["mode"] = v
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["lastID"]; ok && v != nil {
			params["lastID"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=search&f=buildIndex&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to build index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	indexTool := mcp.NewTool("search_index",
		mcp.WithDescription("Search index"),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(indexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=search&f=index&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to index search: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
