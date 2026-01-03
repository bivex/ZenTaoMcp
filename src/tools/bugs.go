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

func RegisterBugTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createBugTool := mcp.NewTool("create_bug",
		mcp.WithDescription("Create a new bug in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Required(),
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Required(),
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(createBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title":    args["title"],
			"severity": int(args["severity"].(float64)),
			"pri":      int(args["pri"].(float64)),
			"type":     args["type"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/bugs", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateBugTool := mcp.NewTool("update_bug",
		mcp.WithDescription("Update an existing bug in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("title",
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(updateBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["severity"]; ok && v != nil {
			body["severity"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/bugs/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteBugTool := mcp.NewTool("delete_bug",
		mcp.WithDescription("Delete a bug from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID to delete"),
		),
	)

	s.AddTool(deleteBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/bugs/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get bugs list tool
	getBugsTool := mcp.NewTool("get_bugs",
		mcp.WithDescription("Get list of bugs in ZenTao"),
		mcp.WithNumber("product",
			mcp.Description("Filter by product ID"),
		),
		mcp.WithNumber("project",
			mcp.Description("Filter by project ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Filter by execution ID"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by bug status"),
			mcp.Enum("active", "resolved", "closed"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Filter by assigned user ID"),
		),
		mcp.WithNumber("openedBy",
			mcp.Description("Filter by opened by user ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of bugs to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if product, ok := args["product"].(float64); ok && product > 0 {
			params["product"] = fmt.Sprintf("%.0f", product)
		}
		if project, ok := args["project"].(float64); ok && project > 0 {
			params["project"] = fmt.Sprintf("%.0f", project)
		}
		if execution, ok := args["execution"].(float64); ok && execution > 0 {
			params["execution"] = fmt.Sprintf("%.0f", execution)
		}
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if assignedTo, ok := args["assignedTo"].(float64); ok && assignedTo > 0 {
			params["assignedTo"] = fmt.Sprintf("%.0f", assignedTo)
		}
		if openedBy, ok := args["openedBy"].(float64); ok && openedBy > 0 {
			params["openedBy"] = fmt.Sprintf("%.0f", openedBy)
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/bugs"
		if len(params) > 0 {
			query := ""
			for k, v := range params {
				if query != "" {
					query += "&"
				}
				query += fmt.Sprintf("%s=%s", k, v)
			}
			path += "?" + query
		}

		resp, err := client.Get(path)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get bug details tool
	getBugTool := mcp.NewTool("get_bug",
		mcp.WithDescription("Get details of a specific bug by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
	)

	s.AddTool(getBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=bug&f=view&t=json&bugID=%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	browseBugsTool := mcp.NewTool("browse_bugs",
		mcp.WithDescription("Browse bugs with filtering and pagination"),
		mcp.WithNumber("productID",
			mcp.Description("Filter by product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Filter by branch"),
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

	s.AddTool(browseBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=bug&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	assignBugTool := mcp.NewTool("assign_bug",
		mcp.WithDescription("Assign a bug to a user"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("assignedTo",
			mcp.Description("User account to assign to"),
		),
	)

	s.AddTool(assignBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=assignTo&t=json&bugID=%d", int(args["bugID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign bug: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	confirmBugTool := mcp.NewTool("confirm_bug",
		mcp.WithDescription("Confirm a bug"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("kanbanParams",
			mcp.Description("Kanban parameters"),
		),
		mcp.WithString("from",
			mcp.Description("Source context"),
		),
	)

	s.AddTool(confirmBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["kanbanParams"]; ok && v != nil {
			body["kanbanParams"] = v
		}

		queryParams := fmt.Sprintf("&bugID=%d", int(args["bugID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=confirm&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to confirm bug: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	resolveBugTool := mcp.NewTool("resolve_bug",
		mcp.WithDescription("Resolve a bug"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Additional parameters"),
		),
	)

	s.AddTool(resolveBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["extra"]; ok && v != nil {
			body["extra"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=resolve&t=json&bugID=%d", int(args["bugID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to resolve bug: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	activateBugTool := mcp.NewTool("activate_bug",
		mcp.WithDescription("Activate a closed bug"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("kanbanInfo",
			mcp.Description("Kanban information"),
		),
	)

	s.AddTool(activateBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["kanbanInfo"]; ok && v != nil {
			body["kanbanInfo"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=activate&t=json&bugID=%d", int(args["bugID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate bug: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	closeBugTool := mcp.NewTool("close_bug",
		mcp.WithDescription("Close a bug"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Additional parameters"),
		),
	)

	s.AddTool(closeBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["extra"]; ok && v != nil {
			body["extra"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=close&t=json&bugID=%d", int(args["bugID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close bug: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	exportBugsTool := mcp.NewTool("export_bugs",
		mcp.WithDescription("Export bugs to file"),
		mcp.WithNumber("productID",
			mcp.Description("Filter by product ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Filter by execution ID"),
		),
	)

	s.AddTool(exportBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=export&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	reportBugsTool := mcp.NewTool("report_bugs",
		mcp.WithDescription("Generate bug report"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
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

	s.AddTool(reportBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=report&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to generate bug report: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateBugsTool := mcp.NewTool("batch_create_bugs",
		mcp.WithDescription("Create multiple bugs at once"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Additional parameters"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchCreateBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchCreate&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditBugsTool := mcp.NewTool("batch_edit_bugs",
		mcp.WithDescription("Edit multiple bugs at once"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Updated bugs data as JSON array"),
		),
	)

	s.AddTool(batchEditBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchEdit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeBranchTool := mcp.NewTool("batch_change_bug_branch",
		mcp.WithDescription("Change branch for multiple bugs"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("New branch ID"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchChangeBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchChangeBranch&t=json&branchID=%d", int(args["branchID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change branch: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeModuleTool := mcp.NewTool("batch_change_bug_module",
		mcp.WithDescription("Change module for multiple bugs"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("New module ID"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchChangeModuleTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchChangeModule&t=json&moduleID=%d", int(args["moduleID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change module: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangePlanTool := mcp.NewTool("batch_change_bug_plan",
		mcp.WithDescription("Change plan for multiple bugs"),
		mcp.WithNumber("planID",
			mcp.Required(),
			mcp.Description("New plan ID"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchChangePlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchChangePlan&t=json&planID=%d", int(args["planID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change plan: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchAssignBugsTool := mcp.NewTool("batch_assign_bugs",
		mcp.WithDescription("Assign multiple bugs to a user"),
		mcp.WithString("assignedTo",
			mcp.Required(),
			mcp.Description("User account to assign to"),
		),
		mcp.WithNumber("objectID",
			mcp.Description("Object ID (projectID or executionID)"),
		),
		mcp.WithString("type",
			mcp.Description("Object type"),
			mcp.Enum("execution", "project", "product", "my"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchAssignBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&assignedTo=%s", args["assignedTo"])
		if v, ok := args["objectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectID=%d", int(v.(float64)))
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchAssignTo&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch assign bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchConfirmBugsTool := mcp.NewTool("batch_confirm_bugs",
		mcp.WithDescription("Confirm multiple bugs"),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchConfirmBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post("/index.php?m=bug&f=batchConfirm&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch confirm bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchResolveBugsTool := mcp.NewTool("batch_resolve_bugs",
		mcp.WithDescription("Resolve multiple bugs"),
		mcp.WithString("resolution",
			mcp.Description("Resolution type"),
		),
		mcp.WithString("resolvedBuild",
			mcp.Description("Resolved build"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchResolveBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["resolution"]; ok && v != nil {
			queryParams += fmt.Sprintf("&resolution=%s", v)
		}
		if v, ok := args["resolvedBuild"]; ok && v != nil {
			queryParams += fmt.Sprintf("&resolvedBuild=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchResolve&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch resolve bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCloseBugsTool := mcp.NewTool("batch_close_bugs",
		mcp.WithDescription("Close multiple bugs"),
		mcp.WithNumber("releaseID",
			mcp.Description("Release ID"),
		),
		mcp.WithString("viewType",
			mcp.Description("View type"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchCloseBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["releaseID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&releaseID=%d", int(v.(float64)))
		}
		if v, ok := args["viewType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&viewType=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchClose&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch close bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	batchActivateBugsTool := mcp.NewTool("batch_activate_bugs",
		mcp.WithDescription("Activate multiple bugs"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("bugs_data",
			mcp.Required(),
			mcp.Description("Bugs data as JSON array"),
		),
	)

	s.AddTool(batchActivateBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}

		body := map[string]interface{}{
			"bugs_data": args["bugs_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=bug&f=batchActivate&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch activate bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	linkBugsTool := mcp.NewTool("link_bugs",
		mcp.WithDescription("Link related bugs"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("bySearch",
			mcp.Description("Search filter"),
		),
		mcp.WithString("excludeBugs",
			mcp.Description("Bugs to exclude"),
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
	)

	s.AddTool(linkBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&bugID=%d", int(args["bugID"].(float64)))
		if v, ok := args["bySearch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&bySearch=%s", v)
		}
		if v, ok := args["excludeBugs"]; ok && v != nil {
			queryParams += fmt.Sprintf("&excludeBugs=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=bug&f=linkBugs&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	confirmStoryChangeTool := mcp.NewTool("confirm_bug_story_change",
		mcp.WithDescription("Confirm story change for a bug"),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
	)

	s.AddTool(confirmStoryChangeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=bug&f=confirmStoryChange&t=json&bugID=%d", int(args["bugID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to confirm story change: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
