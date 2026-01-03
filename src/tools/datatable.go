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

// RegisterDatatableTools registers all datatable and report-related tools
func RegisterDatatableTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Datatable display and management
	registerDatatableDisplayTools(s, client)
	// Datatable save operations
	registerDatatableSaveTools(s, client)
	// Datatable custom operations
	registerDatatableCustomTools(s, client)
	// Datatable reset operations
	registerDatatableResetTools(s, client)
	// Report tools
	registerReportTools(s, client)
}

func registerDatatableDisplayTools(s *server.MCPServer, client *client.ZenTaoClient) {
	ajaxDisplayTool := mcp.NewTool("datatable_ajax_display",
		mcp.WithDescription("Display datatable with specified configuration"),
		mcp.WithString("datatableID",
			mcp.Required(),
			mcp.Description("Datatable ID"),
		),
		mcp.WithString("moduleName", mcp.Description("Module name")),
		mcp.WithString("methodName", mcp.Description("Method name")),
		mcp.WithString("currentModule", mcp.Description("Current module")),
		mcp.WithString("currentMethod", mcp.Description("Current method")),
	)

	s.AddTool(ajaxDisplayTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		datatableID := args["datatableID"].(string)

		params := make(map[string]interface{})
		params["datatableID"] = datatableID

		if v, ok := args["moduleName"].(string); ok {
			params["moduleName"] = v
		}
		if v, ok := args["methodName"].(string); ok {
			params["methodName"] = v
		}
		if v, ok := args["currentModule"].(string); ok {
			params["currentModule"] = v
		}
		if v, ok := args["currentMethod"].(string); ok {
			params["currentMethod"] = v
		}

		resp, err := client.Get("/index.php?m=datatable&f=ajaxDisplay&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to display datatable: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDatatableSaveTools(s *server.MCPServer, client *client.ZenTaoClient) {
	ajaxSaveTool := mcp.NewTool("datatable_ajax_save",
		mcp.WithDescription("Save datatable configuration"),
	)

	s.AddTool(ajaxSaveTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := make(map[string]interface{})

		resp, err := client.Post("/index.php?m=datatable&f=ajaxSave&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save datatable: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxSaveFieldsTool := mcp.NewTool("datatable_ajax_save_fields",
		mcp.WithDescription("Save datatable field configuration"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
	)

	s.AddTool(ajaxSaveFieldsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}

		resp, err := client.Post("/index.php?m=datatable&f=ajaxSaveFields&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save datatable fields: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxOldSaveTool := mcp.NewTool("datatable_ajax_old_save",
		mcp.WithDescription("Save datatable using old format"),
	)

	s.AddTool(ajaxOldSaveTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		params := make(map[string]interface{})

		resp, err := client.Post("/index.php?m=datatable&f=ajaxOldSave&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save datatable (old): %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDatatableCustomTools(s *server.MCPServer, client *client.ZenTaoClient) {
	ajaxCustomTool := mcp.NewTool("datatable_ajax_custom",
		mcp.WithDescription("Perform custom datatable operation"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
	)

	s.AddTool(ajaxCustomTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}

		resp, err := client.Get("/index.php?m=datatable&f=ajaxCustom&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to perform custom operation: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxOldCustomTool := mcp.NewTool("datatable_ajax_old_custom",
		mcp.WithDescription("Perform old custom datatable operation"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
	)

	s.AddTool(ajaxOldCustomTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}

		resp, err := client.Get("/index.php?m=datatable&f=ajaxOldCustom&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to perform old custom operation: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDatatableResetTools(s *server.MCPServer, client *client.ZenTaoClient) {
	ajaxResetTool := mcp.NewTool("datatable_ajax_reset",
		mcp.WithDescription("Reset datatable configuration"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithNumber("system", mcp.Description("System flag")),
		mcp.WithString("confirm", mcp.Description("Confirmation string")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
	)

	s.AddTool(ajaxResetTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["system"]; ok && v != nil {
			params["system"] = int(v.(float64))
		}
		if v, ok := args["confirm"].(string); ok {
			params["confirm"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}

		resp, err := client.Get("/index.php?m=datatable&f=ajaxReset&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reset datatable: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxOldResetTool := mcp.NewTool("datatable_ajax_old_reset",
		mcp.WithDescription("Reset datatable using old format"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithNumber("system", mcp.Description("System flag")),
		mcp.WithString("confirm", mcp.Description("Confirmation string")),
	)

	s.AddTool(ajaxOldResetTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["system"]; ok && v != nil {
			params["system"] = int(v.(float64))
		}
		if v, ok := args["confirm"].(string); ok {
			params["confirm"] = v
		}

		resp, err := client.Get("/index.php?m=datatable&f=ajaxOldReset&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reset datatable (old): %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	ajaxSaveGlobalTool := mcp.NewTool("datatable_ajax_save_global",
		mcp.WithDescription("Save global datatable settings"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithString("extra", mcp.Description("Extra parameters")),
		mcp.WithString("confirm", mcp.Description("Confirmation string")),
	)

	s.AddTool(ajaxSaveGlobalTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["extra"].(string); ok {
			params["extra"] = v
		}
		if v, ok := args["confirm"].(string); ok {
			params["confirm"] = v
		}

		resp, err := client.Post("/index.php?m=datatable&f=ajaxSaveGlobal&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save global settings: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerReportTools(s *server.MCPServer, client *client.ZenTaoClient) {
	reportIndexTool := mcp.NewTool("report_index",
		mcp.WithDescription("Get report index"),
	)

	s.AddTool(reportIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=report&f=index&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get report index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	reportRemindTool := mcp.NewTool("report_remind",
		mcp.WithDescription("Get report reminders"),
	)

	s.AddTool(reportRemindTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=report&f=remind&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get report reminders: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	annualDataTool := mcp.NewTool("report_annual_data",
		mcp.WithDescription("Get annual report data"),
		mcp.WithString("year", mcp.Description("Year (e.g., 2026)")),
		mcp.WithString("dept", mcp.Description("Department")),
		mcp.WithString("account", mcp.Description("Account name")),
	)

	s.AddTool(annualDataTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["year"].(string); ok {
			params["year"] = v
		}
		if v, ok := args["dept"].(string); ok {
			params["dept"] = v
		}
		if v, ok := args["account"].(string); ok {
			params["account"] = v
		}

		resp, err := client.Get("/index.php?m=report&f=annualData&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get annual data: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
