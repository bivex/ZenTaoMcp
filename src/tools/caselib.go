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

func RegisterCaseLibTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getCaseLibIndexTool := mcp.NewTool("get_caselib_index",
		mcp.WithDescription("Get case library index"),
	)

	s.AddTool(getCaseLibIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=caselib&f=index&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get case library index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createCaseLibTool := mcp.NewTool("create_caselib",
		mcp.WithDescription("Create a new case library"),
		mcp.WithString("lib_data",
			mcp.Required(),
			mcp.Description("Case library data as JSON string"),
		),
	)

	s.AddTool(createCaseLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"lib_data": args["lib_data"],
		}

		resp, err := client.Post("/index.php?m=caselib&f=create&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create case library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editCaseLibTool := mcp.NewTool("edit_caselib",
		mcp.WithDescription("Edit an existing case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithString("lib_data",
			mcp.Required(),
			mcp.Description("Updated case library data as JSON string"),
		),
	)

	s.AddTool(editCaseLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"lib_data": args["lib_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=edit&t=json&libID=%d", int(args["libID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit case library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteCaseLibTool := mcp.NewTool("delete_caselib",
		mcp.WithDescription("Delete a case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
	)

	s.AddTool(deleteCaseLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=delete&t=json&libID=%d", int(args["libID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete case library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseCaseLibTool := mcp.NewTool("browse_caselib",
		mcp.WithDescription("Browse case library with filtering and pagination"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithNumber("param",
			mcp.Description("Additional filter parameter"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID for pagination"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
		mcp.WithNumber("blockID",
			mcp.Description("Block ID"),
		),
	)

	s.AddTool(browseCaseLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
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
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["blockID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&blockID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse case library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewCaseLibTool := mcp.NewTool("view_caselib",
		mcp.WithDescription("View case library details"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
	)

	s.AddTool(viewCaseLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=view&t=json&libID=%d", int(args["libID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view case library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createCaseTool := mcp.NewTool("create_caselib_case",
		mcp.WithDescription("Create a new test case in case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("param",
			mcp.Description("Additional parameter"),
		),
		mcp.WithString("case_data",
			mcp.Required(),
			mcp.Description("Test case data as JSON string"),
		),
	)

	s.AddTool(createCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}

		body := map[string]interface{}{
			"case_data": args["case_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=createCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateCaseTool := mcp.NewTool("batch_create_caselib_cases",
		mcp.WithDescription("Create multiple test cases in case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchCreateCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=batchCreateCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editCaseTool := mcp.NewTool("edit_caselib_case",
		mcp.WithDescription("Edit a test case in case library"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithString("case_data",
			mcp.Required(),
			mcp.Description("Updated test case data as JSON string"),
		),
	)

	s.AddTool(editCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"case_data": args["case_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=editCase&t=json&caseID=%d", int(args["caseID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditCaseTool := mcp.NewTool("batch_edit_caselib_cases",
		mcp.WithDescription("Edit multiple test cases in case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("type",
			mcp.Description("Case type filter"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Updated test cases data as JSON array"),
		),
	)

	s.AddTool(batchEditCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=batchEditCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewCaseTool := mcp.NewTool("view_caselib_case",
		mcp.WithDescription("View test case details in case library"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Case version"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
		mcp.WithNumber("taskID",
			mcp.Description("Task ID"),
		),
		mcp.WithString("stepsType",
			mcp.Description("Steps type"),
		),
	)

	s.AddTool(viewCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&caseID=%d", int(args["caseID"].(float64)))
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["taskID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&taskID=%d", int(v.(float64)))
		}
		if v, ok := args["stepsType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&stepsType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=viewCase&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	exportTemplateTool := mcp.NewTool("export_caselib_template",
		mcp.WithDescription("Export import template for case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
	)

	s.AddTool(exportTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=exportTemplate&t=json&libID=%d", int(args["libID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importCasesTool := mcp.NewTool("import_caselib_cases",
		mcp.WithDescription("Import test cases to case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithString("import_data",
			mcp.Required(),
			mcp.Description("Import data as JSON string"),
		),
	)

	s.AddTool(importCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"import_data": args["import_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=import&t=json&libID=%d", int(args["libID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	showImportTool := mcp.NewTool("show_caselib_import",
		mcp.WithDescription("Show import preview for case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
		mcp.WithNumber("maxImport",
			mcp.Description("Maximum import count"),
		),
		mcp.WithString("insert",
			mcp.Description("Insert mode (0=cover, 1=insert)"),
		),
	)

	s.AddTool(showImportTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}
		if v, ok := args["maxImport"]; ok && v != nil {
			queryParams += fmt.Sprintf("&maxImport=%d", int(v.(float64)))
		}
		if v, ok := args["insert"]; ok && v != nil {
			queryParams += fmt.Sprintf("&insert=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=showImport&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to show import: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	exportCasesTool := mcp.NewTool("export_caselib_cases",
		mcp.WithDescription("Export test cases from case library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Case library ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
	)

	s.AddTool(exportCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&libID=%d", int(args["libID"].(float64)))
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=caselib&f=exportCase&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
