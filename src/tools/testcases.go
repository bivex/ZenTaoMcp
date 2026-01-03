// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
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

func RegisterTestCaseTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTestCaseTool := mcp.NewTool("create_testcase",
		mcp.WithDescription("Create a new test case in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Test case title"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Test case type"),
			mcp.Enum("feature", "performance", "config", "install", "security", "interface", "unit", "other"),
		),
		mcp.WithArray("steps",
			mcp.Required(),
			mcp.Description("Test case steps"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Story ID"),
		),
		mcp.WithString("stage",
			mcp.Description("Stage"),
			mcp.Enum("unittest", "feature", "intergrate", "system", "smoke", "bvt"),
		),
		mcp.WithString("precondition",
			mcp.Description("Precondition"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(createTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title": args["title"],
			"type":  args["type"],
			"steps": args["steps"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["stage"]; ok && v != nil {
			body["stage"] = v
		}
		if v, ok := args["precondition"]; ok && v != nil {
			body["precondition"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/testcases", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateTestCaseTool := mcp.NewTool("update_testcase",
		mcp.WithDescription("Update an existing test case in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithString("title",
			mcp.Description("Test case title"),
		),
		mcp.WithString("type",
			mcp.Description("Test case type"),
			mcp.Enum("feature", "performance", "config", "install", "security", "interface", "unit", "other"),
		),
		mcp.WithArray("steps",
			mcp.Description("Test case steps"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Story ID"),
		),
		mcp.WithString("stage",
			mcp.Description("Stage"),
			mcp.Enum("unittest", "feature", "intergrate", "system", "smoke", "bvt"),
		),
		mcp.WithString("precondition",
			mcp.Description("Precondition"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(updateTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["stage"]; ok && v != nil {
			body["stage"] = v
		}
		if v, ok := args["precondition"]; ok && v != nil {
			body["precondition"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/testcases/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTestCaseTool := mcp.NewTool("delete_testcase",
		mcp.WithDescription("Delete a test case from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test case ID to delete"),
		),
	)

	s.AddTool(deleteTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=delete&t=json&caseID=%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	browseTestCasesTool := mcp.NewTool("browse_testcases",
		mcp.WithDescription("Browse test cases with filtering and pagination"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithNumber("param",
			mcp.Description("Additional filter parameter"),
		),
		mcp.WithString("caseType",
			mcp.Description("Case type filter"),
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
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
		mcp.WithNumber("blockID",
			mcp.Description("Block ID"),
		),
	)

	s.AddTool(browseTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}
		if v, ok := args["caseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&caseType=%s", v)
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
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["blockID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&blockID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewTestCaseTool := mcp.NewTool("view_testcase",
		mcp.WithDescription("View test case details"),
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

	s.AddTool(viewTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=view&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateTestCasesTool := mcp.NewTool("batch_create_testcases",
		mcp.WithDescription("Create multiple test cases at once"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("storyID",
			mcp.Description("Story ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchCreateTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["storyID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=batchCreate&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditTestCasesTool := mcp.NewTool("batch_edit_testcases",
		mcp.WithDescription("Edit multiple test cases at once"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("type",
			mcp.Description("Case type filter"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Updated test cases data as JSON array"),
		),
	)

	s.AddTool(batchEditTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=batchEdit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchDeleteTestCasesTool := mcp.NewTool("batch_delete_testcases",
		mcp.WithDescription("Delete multiple test cases at once"),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchDeleteTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post("/index.php?m=testcase&f=batchDelete&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch delete test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	reviewTestCaseTool := mcp.NewTool("review_testcase",
		mcp.WithDescription("Review a test case"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithString("review_data",
			mcp.Description("Review data as JSON string"),
		),
	)

	s.AddTool(reviewTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["review_data"]; ok && v != nil {
			body["review_data"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=review&t=json&caseID=%d", int(args["caseID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to review test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchReviewTestCasesTool := mcp.NewTool("batch_review_testcases",
		mcp.WithDescription("Review multiple test cases at once"),
		mcp.WithString("result",
			mcp.Required(),
			mcp.Description("Review result"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchReviewTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"result":     args["result"],
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post("/index.php?m=testcase&f=batchReview&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch review test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeBranchTool := mcp.NewTool("batch_change_testcase_branch",
		mcp.WithDescription("Change branch for multiple test cases"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("New branch ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchChangeBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=batchChangeBranch&t=json&branchID=%d", int(args["branchID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change branch: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeModuleTool := mcp.NewTool("batch_change_testcase_module",
		mcp.WithDescription("Change module for multiple test cases"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("New module ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchChangeModuleTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=batchChangeModule&t=json&moduleID=%d", int(args["moduleID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeTypeTool := mcp.NewTool("batch_change_testcase_type",
		mcp.WithDescription("Change type for multiple test cases"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("New case type"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(batchChangeTypeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=batchChangeType&t=json&type=%s", args["type"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change type: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	linkCasesTool := mcp.NewTool("link_testcases",
		mcp.WithDescription("Link related test cases"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
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
	)

	s.AddTool(linkCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&caseID=%d", int(args["caseID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=linkCases&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	linkBugsToCaseTool := mcp.NewTool("link_bugs_to_testcase",
		mcp.WithDescription("Link bugs to a test case"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
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
	)

	s.AddTool(linkBugsToCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&caseID=%d", int(args["caseID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=linkBugs&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link bugs to test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createBugFromCaseTool := mcp.NewTool("create_bug_from_testcase",
		mcp.WithDescription("Create a bug from a test case"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Case version"),
		),
		mcp.WithNumber("runID",
			mcp.Description("Test run ID"),
		),
	)

	s.AddTool(createBugFromCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&caseID=%d", int(args["productID"].(float64)), int(args["caseID"].(float64)))
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["runID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&runID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=createBug&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create bug from test case: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	exportTestCasesTool := mcp.NewTool("export_testcases",
		mcp.WithDescription("Export test cases to file"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithNumber("taskID",
			mcp.Description("Task ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
	)

	s.AddTool(exportTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["taskID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&taskID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=export&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	exportTestCaseTemplateTool := mcp.NewTool("export_testcase_template",
		mcp.WithDescription("Export import template for test cases"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
	)

	s.AddTool(exportTestCaseTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=exportTemplate&t=json&productID=%d", int(args["productID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importTestCasesTool := mcp.NewTool("import_testcases",
		mcp.WithDescription("Import test cases from file"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("import_data",
			mcp.Required(),
			mcp.Description("Import data as JSON string"),
		),
	)

	s.AddTool(importTestCasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}

		body := map[string]interface{}{
			"import_data": args["import_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=import&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importFromLibTool := mcp.NewTool("import_testcases_from_lib",
		mcp.WithDescription("Import test cases from case library"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
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
		mcp.WithNumber("queryID",
			mcp.Description("Query ID"),
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
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithString("cases_data",
			mcp.Required(),
			mcp.Description("Test cases data as JSON array"),
		),
	)

	s.AddTool(importFromLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&libID=%d", int(args["productID"].(float64)), int(args["libID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["queryID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&queryID=%d", int(v.(float64)))
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
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{
			"cases_data": args["cases_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=importFromLib&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import from library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importToLibTool := mcp.NewTool("import_testcase_to_lib",
		mcp.WithDescription("Import test case to case library"),
		mcp.WithNumber("caseID",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
	)

	s.AddTool(importToLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=importToLib&t=json&caseID=%d", int(args["caseID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import to library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseSceneTool := mcp.NewTool("browse_testcase_scenes",
		mcp.WithDescription("Browse test case scenes"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
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

	s.AddTool(browseSceneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=browseScene&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse scenes: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createSceneTool := mcp.NewTool("create_testcase_scene",
		mcp.WithDescription("Create a new test case scene"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branch",
			mcp.Required(),
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("scene_data",
			mcp.Required(),
			mcp.Description("Scene data as JSON string"),
		),
	)

	s.AddTool(createSceneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&branch=%d", int(args["productID"].(float64)), int(args["branch"].(float64)))
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{
			"scene_data": args["scene_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=createScene&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create scene: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editSceneTool := mcp.NewTool("edit_testcase_scene",
		mcp.WithDescription("Edit a test case scene"),
		mcp.WithNumber("sceneID",
			mcp.Required(),
			mcp.Description("Scene ID"),
		),
		mcp.WithString("scene_data",
			mcp.Required(),
			mcp.Description("Updated scene data as JSON string"),
		),
	)

	s.AddTool(editSceneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"scene_data": args["scene_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testcase&f=editScene&t=json&sceneID=%d", int(args["sceneID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit scene: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteSceneTool := mcp.NewTool("delete_testcase_scene",
		mcp.WithDescription("Delete a test case scene"),
		mcp.WithNumber("sceneID",
			mcp.Required(),
			mcp.Description("Scene ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
		),
	)

	s.AddTool(deleteSceneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&sceneID=%d", int(args["sceneID"].(float64)))
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=deleteScene&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete scene: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupCaseTool := mcp.NewTool("group_testcases",
		mcp.WithDescription("Group test cases by criteria"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("groupBy",
			mcp.Required(),
			mcp.Description("Group by field"),
		),
		mcp.WithNumber("objectID",
			mcp.Description("Object ID"),
		),
		mcp.WithString("caseType",
			mcp.Description("Case type filter"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
	)

	s.AddTool(groupCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d&groupBy=%s", int(args["productID"].(float64)), args["groupBy"])
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["objectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectID=%d", int(v.(float64)))
		}
		if v, ok := args["caseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&caseType=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=groupCase&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to group test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	zeroCaseTool := mcp.NewTool("get_zero_testcases",
		mcp.WithDescription("Get test cases with zero execution"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branchID",
			mcp.Description("Branch ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithNumber("objectID",
			mcp.Description("Object ID"),
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

	s.AddTool(zeroCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branchID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branchID=%d", int(v.(float64)))
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["objectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectID=%d", int(v.(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testcase&f=zeroCase&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get zero test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
