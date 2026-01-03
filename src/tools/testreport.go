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

func RegisterTestReportTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseTestReportsTool := mcp.NewTool("browse_testreports",
		mcp.WithDescription("Browse test reports with filtering and pagination"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID (project, execution, or product)"),
		),
		mcp.WithString("objectType",
			mcp.Required(),
			mcp.Description("Object type"),
			mcp.Enum("project", "execution", "product"),
		),
		mcp.WithNumber("extra",
			mcp.Description("Extra filter parameter"),
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

	s.AddTool(browseTestReportsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&objectID=%d&objectType=%s", int(args["objectID"].(float64)), args["objectType"])
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%d", int(v.(float64)))
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testreport&f=browse&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse test reports: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createTestReportTool := mcp.NewTool("create_testreport",
		mcp.WithDescription("Create a new test report"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithString("objectType",
			mcp.Required(),
			mcp.Description("Object type"),
			mcp.Enum("project", "execution", "product"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra parameters"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("End date (YYYY-MM-DD)"),
		),
		mcp.WithString("report_data",
			mcp.Required(),
			mcp.Description("Test report data as JSON string"),
		),
	)

	s.AddTool(createTestReportTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&objectID=%d&objectType=%s", int(args["objectID"].(float64)), args["objectType"])
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}
		if v, ok := args["begin"]; ok && v != nil {
			queryParams += fmt.Sprintf("&begin=%s", v)
		}
		if v, ok := args["end"]; ok && v != nil {
			queryParams += fmt.Sprintf("&end=%s", v)
		}

		body := map[string]interface{}{
			"report_data": args["report_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testreport&f=create&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create test report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editTestReportTool := mcp.NewTool("edit_testreport",
		mcp.WithDescription("Edit an existing test report"),
		mcp.WithNumber("reportID",
			mcp.Required(),
			mcp.Description("Test report ID"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("End date (YYYY-MM-DD)"),
		),
		mcp.WithString("report_data",
			mcp.Required(),
			mcp.Description("Updated test report data as JSON string"),
		),
	)

	s.AddTool(editTestReportTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&reportID=%d", int(args["reportID"].(float64)))
		if v, ok := args["begin"]; ok && v != nil {
			queryParams += fmt.Sprintf("&begin=%s", v)
		}
		if v, ok := args["end"]; ok && v != nil {
			queryParams += fmt.Sprintf("&end=%s", v)
		}

		body := map[string]interface{}{
			"report_data": args["report_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testreport&f=edit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit test report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewTestReportTool := mcp.NewTool("view_testreport",
		mcp.WithDescription("View test report details"),
		mcp.WithNumber("reportID",
			mcp.Required(),
			mcp.Description("Test report ID"),
		),
		mcp.WithString("tab",
			mcp.Description("Tab to display"),
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

	s.AddTool(viewTestReportTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&reportID=%d", int(args["reportID"].(float64)))
		if v, ok := args["tab"]; ok && v != nil {
			queryParams += fmt.Sprintf("&tab=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testreport&f=view&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view test report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTestReportTool := mcp.NewTool("delete_testreport",
		mcp.WithDescription("Delete a test report"),
		mcp.WithNumber("reportID",
			mcp.Required(),
			mcp.Description("Test report ID"),
		),
	)

	s.AddTool(deleteTestReportTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testreport&f=delete&t=json&reportID=%d", int(args["reportID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
