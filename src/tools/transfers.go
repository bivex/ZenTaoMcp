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

func RegisterTransferTools(s *server.MCPServer, client *client.ZenTaoClient) {
	exportDataTool := mcp.NewTool("export_data",
		mcp.WithDescription("Export data from a ZenTao module"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module to export from (e.g., story, task, bug, product, project)"),
		),
		mcp.WithString("filters",
			mcp.Description("Export filters as JSON string"),
		),
		mcp.WithString("fields",
			mcp.Description("Fields to export (comma-separated)"),
		),
	)

	s.AddTool(exportDataTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("module=%s", args["module"])
		if v, ok := args["filters"]; ok && v != nil {
			queryParams += fmt.Sprintf("&filters=%s", v)
		}
		if v, ok := args["fields"]; ok && v != nil {
			queryParams += fmt.Sprintf("&fields=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=transfer&f=export&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export data: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportTemplateTool := mcp.NewTool("export_template",
		mcp.WithDescription("Export template for data import"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module to export template for (e.g., story, task, bug, product, project)"),
		),
		mcp.WithString("params",
			mcp.Description("Template parameters as JSON string"),
		),
		mcp.WithString("format",
			mcp.Description("Export format (csv, xlsx)"),
			mcp.Enum("csv", "xlsx"),
		),
	)

	s.AddTool(exportTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("module=%s", args["module"])
		if v, ok := args["params"]; ok && v != nil {
			queryParams += fmt.Sprintf("&params=%s", v)
		}
		if v, ok := args["format"]; ok && v != nil {
			queryParams += fmt.Sprintf("&format=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=transfer&f=exportTemplate&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export template: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	importDataTool := mcp.NewTool("import_data",
		mcp.WithDescription("Import data to a ZenTao module"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module to import to (e.g., story, task, bug, product, project)"),
		),
		mcp.WithString("locate",
			mcp.Description("Import location/target"),
		),
		mcp.WithString("data",
			mcp.Required(),
			mcp.Description("Import data as JSON string"),
		),
		mcp.WithString("encoding",
			mcp.Description("Data encoding"),
			mcp.Enum("utf-8", "gbk", "big5"),
		),
	)

	s.AddTool(importDataTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("module=%s", args["module"])
		if v, ok := args["locate"]; ok && v != nil {
			queryParams += fmt.Sprintf("&locate=%s", v)
		}
		if v, ok := args["encoding"]; ok && v != nil {
			queryParams += fmt.Sprintf("&encoding=%s", v)
		}

		body := map[string]interface{}{
			"data": args["data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=transfer&f=import&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import data: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getImportTableBodyTool := mcp.NewTool("get_import_table_body",
		mcp.WithDescription("Get table body data for import preview"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module for import"),
		),
		mcp.WithNumber("lastID",
			mcp.Description("Last processed ID"),
		),
		mcp.WithNumber("pagerID",
			mcp.Description("Pager ID for pagination"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Number of records to fetch"),
		),
	)

	s.AddTool(getImportTableBodyTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("module=%s", args["module"])
		if v, ok := args["lastID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&lastID=%d", int(v.(float64)))
		}
		if v, ok := args["pagerID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pagerID=%d", int(v.(float64)))
		}
		if v, ok := args["limit"]; ok && v != nil {
			queryParams += fmt.Sprintf("&limit=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=transfer&f=ajaxGetTbody&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get import table body: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getImportOptionsTool := mcp.NewTool("get_import_options",
		mcp.WithDescription("Get options for import fields"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module for import"),
		),
		mcp.WithString("field",
			mcp.Required(),
			mcp.Description("Field name"),
		),
		mcp.WithString("value",
			mcp.Description("Current field value"),
		),
		mcp.WithString("index",
			mcp.Description("Field index"),
		),
		mcp.WithString("search",
			mcp.Description("Search term for filtering options"),
		),
	)

	s.AddTool(getImportOptionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("module=%s&field=%s", args["module"], args["field"])
		if v, ok := args["value"]; ok && v != nil {
			queryParams += fmt.Sprintf("&value=%s", v)
		}
		if v, ok := args["index"]; ok && v != nil {
			queryParams += fmt.Sprintf("&index=%s", v)
		}
		if v, ok := args["search"]; ok && v != nil {
			queryParams += fmt.Sprintf("&search=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=transfer&f=ajaxGetOptions&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get import options: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	validateImportDataTool := mcp.NewTool("validate_import_data",
		mcp.WithDescription("Validate import data before actual import"),
		mcp.WithString("module",
			mcp.Required(),
			mcp.Description("Module to validate for"),
		),
		mcp.WithString("data",
			mcp.Required(),
			mcp.Description("Data to validate as JSON string"),
		),
		mcp.WithString("rules",
			mcp.Description("Validation rules as JSON string"),
		),
	)

	s.AddTool(validateImportDataTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"module": args["module"],
			"data":   args["data"],
		}

		if v, ok := args["rules"]; ok && v != nil {
			body["rules"] = v
		}

		resp, err := client.Post("/index.php?m=transfer&f=validate&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to validate import data: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
