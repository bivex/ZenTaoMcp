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

func RegisterEpicTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Epic CRUD operations
	createEpicTool := mcp.NewTool("create_epic",
		mcp.WithDescription("Create a new epic"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("storyID",
			mcp.Description("Story ID"),
		),
		mcp.WithNumber("objectID",
			mcp.Description("Object ID (projectID|executionID)"),
		),
		mcp.WithNumber("bugID",
			mcp.Description("Bug ID"),
		),
		mcp.WithNumber("planID",
			mcp.Description("Plan ID"),
		),
		mcp.WithNumber("todoID",
			mcp.Description("Todo ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra parameters"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Epic title"),
		),
		mcp.WithString("spec",
			mcp.Description("Epic specification"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimate"),
		),
	)

	s.AddTool(createEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%d", int(v.(float64)))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["storyID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyID=%d", int(v.(float64)))
		}
		if v, ok := args["objectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&objectID=%d", int(v.(float64)))
		}
		if v, ok := args["bugID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&bugID=%d", int(v.(float64)))
		}
		if v, ok := args["planID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&planID=%d", int(v.(float64)))
		}
		if v, ok := args["todoID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&todoID=%d", int(v.(float64)))
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}

		body := map[string]interface{}{
			"title": args["title"],
		}

		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=create&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateEpicsTool := mcp.NewTool("batch_create_epics",
		mcp.WithDescription("Create multiple epics"),
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
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("plan",
			mcp.Description("Plan ID"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra parameters"),
		),
		mcp.WithArray("titles",
			mcp.Required(),
			mcp.Description("Epic titles"),
		),
		mcp.WithArray("specs",
			mcp.Description("Epic specifications"),
		),
	)

	s.AddTool(batchCreateEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["storyID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyID=%d", int(v.(float64)))
		}
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}
		if v, ok := args["plan"]; ok && v != nil {
			queryParams += fmt.Sprintf("&plan=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}

		body := map[string]interface{}{
			"titles": args["titles"],
		}

		if v, ok := args["specs"]; ok && v != nil {
			body["specs"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=batchCreate&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewEpicTool := mcp.NewTool("view_epic",
		mcp.WithDescription("View an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Version"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter (executionID|projectID)"),
		),
	)

	s.AddTool(viewEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editEpicTool := mcp.NewTool("edit_epic",
		mcp.WithDescription("Edit an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("kanbanGroup",
			mcp.Description("Kanban group"),
		),
		mcp.WithString("title",
			mcp.Description("Epic title"),
		),
		mcp.WithString("spec",
			mcp.Description("Epic specification"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimate"),
		),
	)

	s.AddTool(editEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["kanbanGroup"]; ok && v != nil {
			queryParams += fmt.Sprintf("&kanbanGroup=%s", v)
		}

		body := map[string]interface{}{}

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=edit&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditEpicsTool := mcp.NewTool("batch_edit_epics",
		mcp.WithDescription("Batch edit epics"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
		mcp.WithArray("epicIDs",
			mcp.Required(),
			mcp.Description("Epic IDs to edit"),
		),
		mcp.WithArray("titles",
			mcp.Description("New titles"),
		),
		mcp.WithArray("specs",
			mcp.Description("New specifications"),
		),
	)

	s.AddTool(batchEditEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{
			"epicIDs": args["epicIDs"],
		}

		if v, ok := args["titles"]; ok && v != nil {
			body["titles"] = v
		}
		if v, ok := args["specs"]; ok && v != nil {
			body["specs"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=batchEdit&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteEpicTool := mcp.NewTool("delete_epic",
		mcp.WithDescription("Delete an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
	)

	s.AddTool(deleteEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=delete&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Epic linking operations
	linkStoriesToEpicTool := mcp.NewTool("link_stories_to_epic",
		mcp.WithDescription("Link stories to an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithString("excludeStories",
			mcp.Description("Exclude stories"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(linkStoriesToEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["excludeStories"]; ok && v != nil {
			queryParams += fmt.Sprintf("&excludeStories=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=linkStories&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link stories to epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkRequirementsToEpicTool := mcp.NewTool("link_requirements_to_epic",
		mcp.WithDescription("Link requirements to an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithString("excludeStories",
			mcp.Description("Exclude stories"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(linkRequirementsToEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["excludeStories"]; ok && v != nil {
			queryParams += fmt.Sprintf("&excludeStories=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=linkRequirements&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link requirements to epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Import/Export operations
	importEpicsTool := mcp.NewTool("import_epics",
		mcp.WithDescription("Import epics"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(importEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=import&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportEpicsTool := mcp.NewTool("export_epics",
		mcp.WithDescription("Export epics"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
	)

	s.AddTool(exportEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=export&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportEpicTemplateTool := mcp.NewTool("export_epic_template",
		mcp.WithDescription("Export epic template"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
	)

	s.AddTool(exportEpicTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=epic&f=exportTemplate&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export epic template: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Workflow operations
	assignEpicTool := mcp.NewTool("assign_epic",
		mcp.WithDescription("Assign an epic to a user"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
	)

	s.AddTool(assignEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": int(args["assignedTo"].(float64)),
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=assignTo&t=json&storyID=%d", int(args["storyID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchAssignEpicsTool := mcp.NewTool("batch_assign_epics",
		mcp.WithDescription("Batch assign epics"),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
			mcp.Enum("story", "epic"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
		mcp.WithArray("epicIDs",
			mcp.Required(),
			mcp.Description("Epic IDs to assign"),
		),
	)

	s.AddTool(batchAssignEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": args["assignedTo"],
			"epicIDs":    args["epicIDs"],
		}

		if v, ok := args["storyType"]; ok && v != nil {
			body["storyType"] = v
		}

		resp, err := client.Post("/index.php?m=epic&f=batchAssignTo&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch assign epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeEpicTool := mcp.NewTool("close_epic",
		mcp.WithDescription("Close an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
	)

	s.AddTool(closeEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=close&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCloseEpicsTool := mcp.NewTool("batch_close_epics",
		mcp.WithDescription("Batch close epics"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("contribute", "work"),
		),
	)

	s.AddTool(batchCloseEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=batchClose&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch close epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateEpicTool := mcp.NewTool("activate_epic",
		mcp.WithDescription("Activate an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
	)

	s.AddTool(activateEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=activate&t=json&storyID=%d", int(args["storyID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Review workflow
	reviewEpicTool := mcp.NewTool("review_epic",
		mcp.WithDescription("Review an epic"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Epic ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("product", "project"),
		),
	)

	s.AddTool(reviewEpicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=review&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to review epic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchReviewEpicsTool := mcp.NewTool("batch_review_epics",
		mcp.WithDescription("Batch review epics"),
		mcp.WithString("result",
			mcp.Required(),
			mcp.Description("Review result"),
		),
		mcp.WithString("reason",
			mcp.Description("Review reason"),
		),
	)

	s.AddTool(batchReviewEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"result": args["result"],
		}

		if v, ok := args["reason"]; ok && v != nil {
			body["reason"] = v
		}

		resp, err := client.Post("/index.php?m=epic&f=batchReview&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch review epics: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Reporting
	reportEpicsTool := mcp.NewTool("report_epics",
		mcp.WithDescription("Generate epic reports"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("branchID",
			mcp.Description("Branch ID"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("chartType",
			mcp.Description("Chart type"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(reportEpicsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branchID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branchID=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["chartType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&chartType=%s", v)
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=epic&f=report&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to generate epic report: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
