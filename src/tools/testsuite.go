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

func RegisterTestSuiteTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getTestSuiteIndexTool := mcp.NewTool("get_testsuite_index",
		mcp.WithDescription("Get test suite index"),
	)

	s.AddTool(getTestSuiteIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=testsuite&f=index&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get test suite index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseTestSuitesTool := mcp.NewTool("browse_testsuites",
		mcp.WithDescription("Browse test suites with filtering and pagination"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("type",
			mcp.Description("Suite type filter"),
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
	)

	s.AddTool(browseTestSuitesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse test suites: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createTestSuiteTool := mcp.NewTool("create_testsuite",
		mcp.WithDescription("Create a new test suite"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("suite_data",
			mcp.Required(),
			mcp.Description("Test suite data as JSON string"),
		),
	)

	s.AddTool(createTestSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"suite_data": args["suite_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testsuite&f=create&t=json&productID=%d", int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create test suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewTestSuiteTool := mcp.NewTool("view_testsuite",
		mcp.WithDescription("View test suite details"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
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
	)

	s.AddTool(viewTestSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&suiteID=%d", int(args["suiteID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=view&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view test suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editTestSuiteTool := mcp.NewTool("edit_testsuite",
		mcp.WithDescription("Edit an existing test suite"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
		),
		mcp.WithString("suite_data",
			mcp.Required(),
			mcp.Description("Updated test suite data as JSON string"),
		),
	)

	s.AddTool(editTestSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"suite_data": args["suite_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testsuite&f=edit&t=json&suiteID=%d", int(args["suiteID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit test suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTestSuiteTool := mcp.NewTool("delete_testsuite",
		mcp.WithDescription("Delete a test suite"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
		),
	)

	s.AddTool(deleteTestSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=delete&t=json&suiteID=%d", int(args["suiteID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	linkCaseToSuiteTool := mcp.NewTool("link_case_to_testsuite",
		mcp.WithDescription("Link test case to test suite"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithNumber("param",
			mcp.Description("Additional filter parameter"),
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
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(linkCaseToSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&suiteID=%d", int(args["suiteID"].(float64)))
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

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testsuite&f=linkCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link case to suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkCaseFromSuiteTool := mcp.NewTool("unlink_case_from_testsuite",
		mcp.WithDescription("Unlink test case from test suite"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
		),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
	)

	s.AddTool(unlinkCaseFromSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=unlinkCase&t=json&suiteID=%d&caseID=%d", int(args["suiteID"].(float64)), int(args["caseID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink case from suite: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkCasesFromSuiteTool := mcp.NewTool("batch_unlink_cases_from_testsuite",
		mcp.WithDescription("Unlink multiple test cases from test suite"),
		mcp.WithNumber("suiteID",
			mcp.Required(),
			mcp.Description("Test suite ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchUnlinkCasesFromSuiteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testsuite&f=batchUnlinkCases&t=json&suiteID=%d", int(args["suiteID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
