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

func RegisterTestTaskTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTestTaskTool := mcp.NewTool("create_testtask",
		mcp.WithDescription("Create a new test task in ZenTao"),
		mcp.WithNumber("project",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Test task name"),
		),
		mcp.WithString("begin",
			mcp.Required(),
			mcp.Description("Start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Required(),
			mcp.Description("End date (YYYY-MM-DD)"),
		),
		mcp.WithString("owner",
			mcp.Description("Owner user account"),
		),
		mcp.WithString("desc",
			mcp.Description("Test task description"),
		),
	)

	s.AddTool(createTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"project": int(args["project"].(float64)),
			"name":    args["name"],
			"begin":   args["begin"],
			"end":     args["end"],
		}

		if v, ok := args["owner"]; ok && v != nil {
			body["owner"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post("/testtasks", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create test task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestTasksTool := mcp.NewTool("get_testtasks",
		mcp.WithDescription("Get list of test tasks in ZenTao"),
		mcp.WithNumber("project",
			mcp.Description("Filter by project ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of test tasks to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getTestTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)
		if v, ok := args["project"]; ok && v != nil {
			params["project"] = fmt.Sprintf("%d", int(v.(float64)))
		}
		if v, ok := args["limit"]; ok && v != nil {
			params["limit"] = fmt.Sprintf("%d", int(v.(float64)))
		}
		if v, ok := args["offset"]; ok && v != nil {
			params["offset"] = fmt.Sprintf("%d", int(v.(float64)))
		}

		queryString := ""
		if len(params) > 0 {
			queryString = "?"
			for k, v := range params {
				queryString += fmt.Sprintf("%s=%s&", k, v)
			}
			queryString = queryString[:len(queryString)-1] // Remove trailing &
		}

		resp, err := client.Get("/testtasks" + queryString)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get test tasks: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestTaskTool := mcp.NewTool("get_testtask",
		mcp.WithDescription("Get details of a specific test task by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
	)

	s.AddTool(getTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/testtasks/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get test task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProjectTestTasksTool := mcp.NewTool("get_project_testtasks",
		mcp.WithDescription("Get test tasks for a specific project"),
		mcp.WithNumber("project_id",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of test tasks to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getProjectTestTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectID := int(args["project_id"].(float64))

		params := make(map[string]string)
		if v, ok := args["limit"]; ok && v != nil {
			params["limit"] = fmt.Sprintf("%d", int(v.(float64)))
		}
		if v, ok := args["offset"]; ok && v != nil {
			params["offset"] = fmt.Sprintf("%d", int(v.(float64)))
		}

		queryString := ""
		if len(params) > 0 {
			queryString = "?"
			for k, v := range params {
				queryString += fmt.Sprintf("%s=%s&", k, v)
			}
			queryString = queryString[:len(queryString)-1] // Remove trailing &
		}

		resp, err := client.Get(fmt.Sprintf("/projects/%d/testtasks%s", projectID, queryString))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get project test tasks: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTestTaskTool := mcp.NewTool("delete_testtask",
		mcp.WithDescription("Delete a test task from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test Task ID to delete"),
		),
	)

	s.AddTool(deleteTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=delete&t=json&taskID=%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	browseTestTasksTool := mcp.NewTool("browse_testtasks",
		mcp.WithDescription("Browse test tasks with filtering and pagination"),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("type",
			mcp.Description("Task type filter"),
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
		mcp.WithString("beginTime",
			mcp.Description("Begin time filter"),
		),
		mcp.WithString("endTime",
			mcp.Description("End time filter"),
		),
	)

	s.AddTool(browseTestTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
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
		if v, ok := args["beginTime"]; ok && v != nil {
			queryParams += fmt.Sprintf("&beginTime=%s", v)
		}
		if v, ok := args["endTime"]; ok && v != nil {
			queryParams += fmt.Sprintf("&endTime=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse test tasks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editTestTaskTool := mcp.NewTool("edit_testtask",
		mcp.WithDescription("Edit an existing test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("task_data",
			mcp.Required(),
			mcp.Description("Updated test task data as JSON string"),
		),
	)

	s.AddTool(editTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"task_data": args["task_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=edit&t=json&taskID=%d", int(args["taskID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	startTestTaskTool := mcp.NewTool("start_testtask",
		mcp.WithDescription("Start a test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
	)

	s.AddTool(startTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=start&t=json&taskID=%d", int(args["taskID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	closeTestTaskTool := mcp.NewTool("close_testtask",
		mcp.WithDescription("Close a test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
	)

	s.AddTool(closeTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=close&t=json&taskID=%d", int(args["taskID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	blockTestTaskTool := mcp.NewTool("block_testtask",
		mcp.WithDescription("Block a test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
	)

	s.AddTool(blockTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=block&t=json&taskID=%d", int(args["taskID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to block test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	activateTestTaskTool := mcp.NewTool("activate_testtask",
		mcp.WithDescription("Activate a blocked test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
	)

	s.AddTool(activateTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=activate&t=json&taskID=%d", int(args["taskID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestTaskCasesTool := mcp.NewTool("get_testtask_cases",
		mcp.WithDescription("Get test cases for a test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
			mcp.Enum("all", "assignedtome", "bysuite", "byModule"),
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
	)

	s.AddTool(getTestTaskCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&taskID=%d", int(args["taskID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=cases&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get test task cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestTaskUnitCasesTool := mcp.NewTool("get_testtask_unit_cases",
		mcp.WithDescription("Get unit test cases for a test task"),
		mcp.WithNumber("testtaskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
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

	s.AddTool(getTestTaskUnitCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&testtaskID=%d", int(args["testtaskID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=unitCases&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get unit cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	linkCaseToTestTaskTool := mcp.NewTool("link_case_to_testtask",
		mcp.WithDescription("Link test case to test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("type",
			mcp.Description("Link type"),
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

	s.AddTool(linkCaseToTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&taskID=%d", int(args["taskID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=linkCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link case to test task: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkCaseFromTestTaskTool := mcp.NewTool("unlink_case_from_testtask",
		mcp.WithDescription("Unlink test case from test task"),
		mcp.WithNumber("runID",
			mcp.Required(),
			mcp.Description("Test run ID"),
		),
	)

	s.AddTool(unlinkCaseFromTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=unlinkCase&t=json&runID=%d", int(args["runID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkCasesFromTestTaskTool := mcp.NewTool("batch_unlink_cases_from_testtask",
		mcp.WithDescription("Unlink multiple test cases from test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("runs_data",
			mcp.Required(),
			mcp.Description("Test runs data as JSON array"),
		),
	)

	s.AddTool(batchUnlinkCasesFromTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"runs_data": args["runs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=batchUnlinkCases&t=json&taskID=%d", int(args["taskID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	runCaseTool := mcp.NewTool("run_testcase",
		mcp.WithDescription("Run a test case in test task"),
		mcp.WithNumber("runID",
			mcp.Required(),
			mcp.Description("Test run ID"),
		),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Case version"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
		),
		mcp.WithString("result_data",
			mcp.Required(),
			mcp.Description("Test result data as JSON string"),
		),
	)

	s.AddTool(runCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&runID=%d&caseID=%d", int(args["runID"].(float64)), int(args["caseID"].(float64)))
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}

		body := map[string]interface{}{
			"result_data": args["result_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=runCase&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to run test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchRunCasesTool := mcp.NewTool("batch_run_testcases",
		mcp.WithDescription("Run multiple test cases at once"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
		),
		mcp.WithString("runs_data",
			mcp.Required(),
			mcp.Description("Test runs data as JSON array"),
		),
	)

	s.AddTool(batchRunCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&taskID=%d", int(args["productID"].(float64)), int(args["taskID"].(float64)))
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}

		body := map[string]interface{}{
			"runs_data": args["runs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=batchRun&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch run test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestTaskResultsTool := mcp.NewTool("get_testtask_results",
		mcp.WithDescription("Get test results for a test task"),
		mcp.WithNumber("runID",
			mcp.Required(),
			mcp.Description("Test run ID"),
		),
		mcp.WithNumber("caseID",
			mcp.Description("Test case ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Case version"),
		),
		mcp.WithString("status",
			mcp.Description("Result status filter"),
			mcp.Enum("all", "done"),
		),
		mcp.WithString("type",
			mcp.Description("Result type filter"),
			mcp.Enum("all", "fail"),
		),
		mcp.WithNumber("deployID",
			mcp.Description("Deploy ID"),
		),
	)

	s.AddTool(getTestTaskResultsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&runID=%d", int(args["runID"].(float64)))
		if v, ok := args["caseID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&caseID=%d", int(v.(float64)))
		}
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["deployID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&deployID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=results&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get test results: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	assignCaseTool := mcp.NewTool("assign_testcase",
		mcp.WithDescription("Assign test case to user"),
		mcp.WithNumber("runID",
			mcp.Required(),
			mcp.Description("Test run ID"),
		),
		mcp.WithString("assignedTo",
			mcp.Description("User account to assign to"),
		),
	)

	s.AddTool(assignCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=assignCase&t=json&runID=%d", int(args["runID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchAssignCasesTool := mcp.NewTool("batch_assign_testcases",
		mcp.WithDescription("Assign multiple test cases to user"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("account",
			mcp.Required(),
			mcp.Description("User account to assign to"),
		),
		mcp.WithString("runs_data",
			mcp.Required(),
			mcp.Description("Test runs data as JSON array"),
		),
	)

	s.AddTool(batchAssignCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"runs_data": args["runs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=batchAssign&t=json&taskID=%d&account=%s", int(args["taskID"].(float64)), args["account"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch assign test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	reportTestTaskTool := mcp.NewTool("report_testtask",
		mcp.WithDescription("Generate test task report"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithNumber("branchID",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("chartType",
			mcp.Description("Chart type"),
		),
	)

	s.AddTool(reportTestTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&taskID=%d", int(args["productID"].(float64)), int(args["taskID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["branchID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branchID=%d", int(v.(float64)))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["chartType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&chartType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=report&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to generate test task report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupTestTaskCasesTool := mcp.NewTool("group_testtask_cases",
		mcp.WithDescription("Group test cases in test task"),
		mcp.WithNumber("taskID",
			mcp.Required(),
			mcp.Description("Test task ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
	)

	s.AddTool(groupTestTaskCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&taskID=%d", int(args["taskID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=groupCase&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to group test task cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseUnitsTool := mcp.NewTool("browse_testtask_units",
		mcp.WithDescription("Browse unit test tasks"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
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

	s.AddTool(browseUnitsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testtask&f=browseUnits&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse unit test tasks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importUnitResultTool := mcp.NewTool("import_unit_test_result",
		mcp.WithDescription("Import unit test result"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("result_data",
			mcp.Required(),
			mcp.Description("Unit test result data as JSON string"),
		),
	)

	s.AddTool(importUnitResultTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"result_data": args["result_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testtask&f=importUnitResult&t=json&productID=%d", int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import unit test result: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
