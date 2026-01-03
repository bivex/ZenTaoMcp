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

func RegisterRequirementTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Requirement CRUD operations
	createRequirementTool := mcp.NewTool("create_requirement",
		mcp.WithDescription("Create a new requirement"),
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
			mcp.Description("Requirement title"),
		),
		mcp.WithString("spec",
			mcp.Description("Requirement specification"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimate"),
		),
	)

	s.AddTool(createRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=create&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateRequirementsTool := mcp.NewTool("batch_create_requirements",
		mcp.WithDescription("Create multiple requirements"),
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
			mcp.Description("Requirement titles"),
		),
		mcp.WithArray("specs",
			mcp.Description("Requirement specifications"),
		),
	)

	s.AddTool(batchCreateRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchCreate&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewRequirementTool := mcp.NewTool("view_requirement",
		mcp.WithDescription("View a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Version"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter (executionID|projectID)"),
		),
	)

	s.AddTool(viewRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editRequirementTool := mcp.NewTool("edit_requirement",
		mcp.WithDescription("Edit a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithString("kanbanGroup",
			mcp.Description("Kanban group"),
		),
		mcp.WithString("title",
			mcp.Description("Requirement title"),
		),
		mcp.WithString("spec",
			mcp.Description("Requirement specification"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimate"),
		),
	)

	s.AddTool(editRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=edit&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditRequirementsTool := mcp.NewTool("batch_edit_requirements",
		mcp.WithDescription("Batch edit requirements"),
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
		mcp.WithArray("requirementIDs",
			mcp.Required(),
			mcp.Description("Requirement IDs to edit"),
		),
		mcp.WithArray("titles",
			mcp.Description("New titles"),
		),
		mcp.WithArray("specs",
			mcp.Description("New specifications"),
		),
	)

	s.AddTool(batchEditRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
			"requirementIDs": args["requirementIDs"],
		}

		if v, ok := args["titles"]; ok && v != nil {
			body["titles"] = v
		}
		if v, ok := args["specs"]; ok && v != nil {
			body["specs"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchEdit&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteRequirementTool := mcp.NewTool("delete_requirement",
		mcp.WithDescription("Delete a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
	)

	s.AddTool(deleteRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=delete&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Requirement linking operations
	linkStoryToRequirementTool := mcp.NewTool("link_story_to_requirement",
		mcp.WithDescription("Link a story to a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithString("type",
			mcp.Description("Link type"),
			mcp.Enum("linkStories", "linkRelateUR", "linkRelateSR"),
		),
		mcp.WithNumber("linkedStoryID",
			mcp.Description("Story ID to link"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("queryID",
			mcp.Description("Query ID"),
		),
	)

	s.AddTool(linkStoryToRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["linkedStoryID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&linkedStoryID=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["queryID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&queryID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=linkStory&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link story to requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkRequirementsTool := mcp.NewTool("link_requirements",
		mcp.WithDescription("Link requirements together"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
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

	s.AddTool(linkRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=linkRequirements&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Import/Export operations
	importRequirementsTool := mcp.NewTool("import_requirements",
		mcp.WithDescription("Import requirements"),
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

	s.AddTool(importRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=import&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportRequirementsTool := mcp.NewTool("export_requirements",
		mcp.WithDescription("Export requirements"),
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

	s.AddTool(exportRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=export&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportRequirementTemplateTool := mcp.NewTool("export_requirement_template",
		mcp.WithDescription("Export requirement template"),
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

	s.AddTool(exportRequirementTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%d", int(v.(float64)))
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=requirement&f=exportTemplate&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export requirement template: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Workflow operations
	assignRequirementTool := mcp.NewTool("assign_requirement",
		mcp.WithDescription("Assign a requirement to a user"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
	)

	s.AddTool(assignRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": int(args["assignedTo"].(float64)),
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=assignTo&t=json&storyID=%d", int(args["storyID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchAssignRequirementsTool := mcp.NewTool("batch_assign_requirements",
		mcp.WithDescription("Batch assign requirements"),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
			mcp.Enum("story", "requirement"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
		mcp.WithArray("requirementIDs",
			mcp.Required(),
			mcp.Description("Requirement IDs to assign"),
		),
	)

	s.AddTool(batchAssignRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": args["assignedTo"],
			"requirementIDs": args["requirementIDs"],
		}

		if v, ok := args["storyType"]; ok && v != nil {
			body["storyType"] = v
		}

		resp, err := client.Post("/index.php?m=requirement&f=batchAssignTo&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch assign requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeRequirementTool := mcp.NewTool("close_requirement",
		mcp.WithDescription("Close a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
	)

	s.AddTool(closeRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=close&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCloseRequirementsTool := mcp.NewTool("batch_close_requirements",
		mcp.WithDescription("Batch close requirements"),
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

	s.AddTool(batchCloseRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchClose&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch close requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateRequirementTool := mcp.NewTool("activate_requirement",
		mcp.WithDescription("Activate a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
	)

	s.AddTool(activateRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=activate&t=json&storyID=%d", int(args["storyID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Review workflow
	reviewRequirementTool := mcp.NewTool("review_requirement",
		mcp.WithDescription("Review a requirement"),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Requirement ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("product", "project"),
		),
	)

	s.AddTool(reviewRequirementTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("storyID=%d", int(args["storyID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=review&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to review requirement: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchReviewRequirementsTool := mcp.NewTool("batch_review_requirements",
		mcp.WithDescription("Batch review requirements"),
		mcp.WithString("result",
			mcp.Required(),
			mcp.Description("Review result"),
		),
		mcp.WithString("reason",
			mcp.Description("Review reason"),
		),
	)

	s.AddTool(batchReviewRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"result": args["result"],
		}

		if v, ok := args["reason"]; ok && v != nil {
			body["reason"] = v
		}

		resp, err := client.Post("/index.php?m=requirement&f=batchReview&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch review requirements: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Reporting
	reportRequirementsTool := mcp.NewTool("report_requirements",
		mcp.WithDescription("Generate requirement reports"),
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

	s.AddTool(reportRequirementsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=report&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to generate requirement report: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Batch operations
	batchChangeRequirementBranchTool := mcp.NewTool("batch_change_requirement_branch",
		mcp.WithDescription("Batch change requirement branch"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("New branch ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
		mcp.WithString("storyIdList",
			mcp.Description("Story ID list"),
		),
	)

	s.AddTool(batchChangeRequirementBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("branchID=%d", int(args["branchID"].(float64)))
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}
		if v, ok := args["storyIdList"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyIdList=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchChangeBranch&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change requirement branch: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeRequirementModuleTool := mcp.NewTool("batch_change_requirement_module",
		mcp.WithDescription("Batch change requirement module"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("New module ID"),
		),
	)

	s.AddTool(batchChangeRequirementModuleTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchChangeModule&t=json&moduleID=%d", int(args["moduleID"].(float64))), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change requirement module: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeRequirementParentTool := mcp.NewTool("batch_change_requirement_parent",
		mcp.WithDescription("Batch change requirement parent"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
	)

	s.AddTool(batchChangeRequirementParentTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchChangeParent&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change requirement parent: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeRequirementGradeTool := mcp.NewTool("batch_change_requirement_grade",
		mcp.WithDescription("Batch change requirement grade"),
		mcp.WithNumber("grade",
			mcp.Required(),
			mcp.Description("New grade"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
		),
	)

	s.AddTool(batchChangeRequirementGradeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("grade=%d", int(args["grade"].(float64)))
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchChangeGrade&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change requirement grade: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchChangeRequirementPlanTool := mcp.NewTool("batch_change_requirement_plan",
		mcp.WithDescription("Batch change requirement plan"),
		mcp.WithNumber("planID",
			mcp.Required(),
			mcp.Description("New plan ID"),
		),
		mcp.WithNumber("oldPlanID",
			mcp.Description("Old plan ID"),
		),
	)

	s.AddTool(batchChangeRequirementPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("planID=%d", int(args["planID"].(float64)))
		if v, ok := args["oldPlanID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&oldPlanID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=requirement&f=batchChangePlan&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch change requirement plan: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
