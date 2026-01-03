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

// RegisterTreeTools registers all tree/module management tools
func RegisterTreeTools(s *server.MCPServer, client *client.ZenTaoClient) {
	registerTreeBrowseTools(s, client)
	registerTreeEditTools(s, client)
	registerTreeManagementTools(s, client)
	registerTreeOptionTools(s, client)
}

func registerTreeBrowseTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseTool := mcp.NewTool("tree_browse",
		mcp.WithDescription("Browse tree structure"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("viewType",
			mcp.Description("View type: story|bug|case|doc"),
			mcp.Enum("story", "bug", "case", "doc"),
		),
		mcp.WithNumber("currentModuleID", mcp.Description("Current module ID")),
		mcp.WithString("branch", mcp.Description("Branch")),
		mcp.WithString("from", mcp.Description("From filter")),
	)

	s.AddTool(browseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["currentModuleID"]; ok && v != nil {
			params["currentModuleID"] = int(v.(float64))
		}
		if v, ok := args["branch"].(string); ok {
			params["branch"] = v
		}
		if v, ok := args["from"].(string); ok {
			params["from"] = v
		}

		resp, err := client.Get("/index.php?m=tree&f=browse&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse tree: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseTaskTool := mcp.NewTool("tree_browse_task",
		mcp.WithDescription("Browse tree tasks"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithNumber("productID", mcp.Description("Product ID")),
		mcp.WithNumber("currentModuleID", mcp.Description("Current module ID")),
	)

	s.AddTool(browseTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["productID"]; ok && v != nil {
			params["productID"] = int(v.(float64))
		}
		if v, ok := args["currentModuleID"]; ok && v != nil {
			params["currentModuleID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=tree&f=browseTask&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse tree tasks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerTreeEditTools(s *server.MCPServer, client *client.ZenTaoClient) {
	editTool := mcp.NewTool("tree_edit",
		mcp.WithDescription("Edit tree module"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
		mcp.WithString("type", mcp.Description("Module type")),
		mcp.WithString("branch", mcp.Description("Branch")),
	)

	s.AddTool(editTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["moduleID"] = int(args["moduleID"].(float64))

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["branch"].(string); ok {
			params["branch"] = v
		}

		resp, err := client.Post("/index.php?m=tree&f=edit&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit tree module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	fixTool := mcp.NewTool("tree_fix",
		mcp.WithDescription("Fix tree module"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("type", mcp.Description("Module type")),
	)

	s.AddTool(fixTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=tree&f=fix&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to fix tree module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerTreeManagementTools(s *server.MCPServer, client *client.ZenTaoClient) {
	updateOrderTool := mcp.NewTool("tree_update_order",
		mcp.WithDescription("Update tree module order"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("viewType",
			mcp.Description("View type: story|bug|case|doc"),
			mcp.Enum("story", "bug", "case", "doc"),
		),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
	)

	s.AddTool(updateOrderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=tree&f=updateOrder&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update order: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	manageChildTool := mcp.NewTool("tree_manage_child",
		mcp.WithDescription("Manage tree child modules"),
		mcp.WithNumber("rootID",
			mcp.Required(),
			mcp.Description("Root ID"),
		),
		mcp.WithString("viewType",
			mcp.Description("View type: story|bug|case|doc"),
			mcp.Enum("story", "bug", "case", "doc"),
		),
		mcp.WithString("oldPage",
			mcp.Description("Old page: yes|no"),
			mcp.Enum("yes", "no"),
		),
	)

	s.AddTool(manageChildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["rootID"] = int(args["rootID"].(float64))

		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["oldPage"].(string); ok {
			params["oldPage"] = v
		}

		resp, err := client.Post("/index.php?m=tree&f=manageChild&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage child: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewHistoryTool := mcp.NewTool("tree_view_history",
		mcp.WithDescription("View tree history"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
	)

	s.AddTool(viewHistoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["productID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=tree&f=viewHistory&t=json&productID=%d", productID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view history: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTool := mcp.NewTool("tree_delete",
		mcp.WithDescription("Delete tree module"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation: yes|no"),
			mcp.Enum("yes", "no"),
		),
	)

	s.AddTool(deleteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		moduleID := int(args["moduleID"].(float64))
		params := make(map[string]interface{})

		if v, ok := args["confirm"].(string); ok {
			params["confirm"] = v
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=tree&f=delete&t=json&moduleID=%d", moduleID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete tree module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxCreateModuleTool := mcp.NewTool("tree_ajax_create_module",
		mcp.WithDescription("Create tree module"),
	)

	s.AddTool(ajaxCreateModuleTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=tree&f=ajaxCreateModule&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerTreeOptionTools(s *server.MCPServer, client *client.ZenTaoClient) {
	ajaxGetOptionMenuTool := mcp.NewTool("tree_ajax_get_option_menu",
		mcp.WithDescription("Get tree option menu"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("viewType",
			mcp.Description("View type: story|bug|case|doc"),
			mcp.Enum("story", "bug", "case", "doc"),
		),
		mcp.WithString("branch", mcp.Description("Branch")),
		mcp.WithNumber("rootModuleID", mcp.Description("Root module ID")),
		mcp.WithString("returnType", mcp.Description("Return type")),
		mcp.WithString("fieldID", mcp.Description("Field ID")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
		mcp.WithNumber("currentModuleID", mcp.Description("Current module ID")),
		mcp.WithString("grade", mcp.Description("Grade")),
	)

	s.AddTool(ajaxGetOptionMenuTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["branch"].(string); ok {
			params["branch"] = v
		}
		if v, ok := args["rootModuleID"]; ok && v != nil {
			params["rootModuleID"] = int(v.(float64))
		}
		if v, ok := args["returnType"].(string); ok {
			params["returnType"] = v
		}
		if v, ok := args["fieldID"].(string); ok {
			params["fieldID"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}
		if v, ok := args["currentModuleID"]; ok && v != nil {
			params["currentModuleID"] = int(v.(float64))
		}
		if v, ok := args["grade"].(string); ok {
			params["grade"] = v
		}

		resp, err := client.Get("/index.php?m=tree&f=ajaxGetOptionMenu&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get option menu: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxGetDropMenuTool := mcp.NewTool("tree_ajax_get_drop_menu",
		mcp.WithDescription("Get tree drop menu"),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("module", mcp.Description("Module")),
		mcp.WithString("method", mcp.Description("Method")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
	)

	s.AddTool(ajaxGetDropMenuTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}

		resp, err := client.Get("/index.php?m=tree&f=ajaxGetDropMenu&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get drop menu: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxGetModulesTool := mcp.NewTool("tree_ajax_get_modules",
		mcp.WithDescription("Get tree modules"),
		mcp.WithNumber("productID", mcp.Description("Product ID")),
		mcp.WithString("viewType",
			mcp.Description("View type: story|bug|case|doc"),
			mcp.Enum("story", "bug", "case", "doc"),
		),
		mcp.WithString("branch", mcp.Description("Branch")),
		mcp.WithNumber("number", mcp.Description("Number")),
		mcp.WithNumber("currentModuleID", mcp.Description("Current module ID")),
		mcp.WithString("from",
			mcp.Description("From: showImport"),
		),
	)

	s.AddTool(ajaxGetModulesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["productID"]; ok && v != nil {
			params["productID"] = int(v.(float64))
		}
		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["branch"].(string); ok {
			params["branch"] = v
		}
		if v, ok := args["number"]; ok && v != nil {
			params["number"] = int(v.(float64))
		}
		if v, ok := args["currentModuleID"]; ok && v != nil {
			params["currentModuleID"] = int(v.(float64))
		}
		if v, ok := args["from"].(string); ok {
			params["from"] = v
		}

		resp, err := client.Get("/index.php?m=tree&f=ajaxGetModules&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get modules: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxGetSonModulesTool := mcp.NewTool("tree_ajax_get_son_modules",
		mcp.WithDescription("Get son modules"),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithNumber("rootID", mcp.Description("Root ID")),
		mcp.WithString("type", mcp.Description("Module type")),
	)

	s.AddTool(ajaxGetSonModulesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["rootID"]; ok && v != nil {
			params["rootID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Get("/index.php?m=tree&f=ajaxGetSonModules&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get son modules: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
