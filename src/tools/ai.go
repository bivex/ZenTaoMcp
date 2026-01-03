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

func RegisterAiTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Mini Programs Tools
	getAiAdminIndexTool := mcp.NewTool("get_ai_admin_index",
		mcp.WithDescription("Get AI module admin interface overview"),
	)

	s.AddTool(getAiAdminIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=ai&f=adminIndex&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get AI admin index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getMiniProgramsTool := mcp.NewTool("get_mini_programs",
		mcp.WithDescription("Get list of AI mini programs with filtering"),
		mcp.WithString("category",
			mcp.Description("Filter by category"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by status"),
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

	s.AddTool(getMiniProgramsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["category"]; ok && v != nil {
			queryParams += fmt.Sprintf("&category=%s", v)
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=miniPrograms&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get mini programs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editMiniProgramCategoryTool := mcp.NewTool("edit_mini_program_category",
		mcp.WithDescription("Edit mini program category"),
		mcp.WithString("category_data",
			mcp.Required(),
			mcp.Description("Category data as JSON string"),
		),
	)

	s.AddTool(editMiniProgramCategoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"category_data": args["category_data"],
		}

		resp, err := client.Post("/index.php?m=ai&f=editMiniProgramCategory&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit mini program category: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	publishMiniProgramTool := mcp.NewTool("publish_mini_program",
		mcp.WithDescription("Publish an AI mini program"),
		mcp.WithString("appID",
			mcp.Required(),
			mcp.Description("Mini program application ID"),
		),
	)

	s.AddTool(publishMiniProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=publishMiniProgram&t=json&appID=%s", args["appID"]))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to publish mini program: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	unpublishMiniProgramTool := mcp.NewTool("unpublish_mini_program",
		mcp.WithDescription("Unpublish an AI mini program"),
		mcp.WithString("appID",
			mcp.Required(),
			mcp.Description("Mini program application ID"),
		),
	)

	s.AddTool(unpublishMiniProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=unpublishMiniProgram&t=json&appID=%s", args["appID"]))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unpublish mini program: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	importMiniProgramTool := mcp.NewTool("import_mini_program",
		mcp.WithDescription("Import AI mini program"),
		mcp.WithString("import_data",
			mcp.Required(),
			mcp.Description("Import data as JSON string"),
		),
	)

	s.AddTool(importMiniProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"import_data": args["import_data"],
		}

		resp, err := client.Post("/index.php?m=ai&f=importMiniProgram&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import mini program: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	// Prompts Tools
	getPromptsTool := mcp.NewTool("get_prompts",
		mcp.WithDescription("Get list of AI prompts with filtering"),
		mcp.WithString("module",
			mcp.Description("Filter by module"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by status"),
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

	s.AddTool(getPromptsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["module"]; ok && v != nil {
			queryParams += fmt.Sprintf("&module=%s", v)
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=prompts&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get prompts: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getPromptViewTool := mcp.NewTool("get_prompt_view",
		mcp.WithDescription("Get detailed view of a specific AI prompt"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
	)

	s.AddTool(getPromptViewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptView&t=json&id=%d", int(args["id"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get prompt view: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createPromptTool := mcp.NewTool("create_prompt",
		mcp.WithDescription("Create a new AI prompt"),
		mcp.WithString("prompt_data",
			mcp.Required(),
			mcp.Description("Prompt data as JSON string"),
		),
	)

	s.AddTool(createPromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"prompt_data": args["prompt_data"],
		}

		resp, err := client.Post("/index.php?m=ai&f=createPrompt&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editPromptTool := mcp.NewTool("edit_prompt",
		mcp.WithDescription("Edit an existing AI prompt"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("prompt_data",
			mcp.Required(),
			mcp.Description("Updated prompt data as JSON string"),
		),
	)

	s.AddTool(editPromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"id":          int(args["id"].(float64)),
			"prompt_data": args["prompt_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptEdit&t=json&id=%d", int(args["id"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deletePromptTool := mcp.NewTool("delete_prompt",
		mcp.WithDescription("Delete an AI prompt"),
		mcp.WithNumber("prompt",
			mcp.Required(),
			mcp.Description("Prompt ID to delete"),
		),
	)

	s.AddTool(deletePromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptDelete&t=json&prompt=%d", int(args["prompt"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	assignPromptRoleTool := mcp.NewTool("assign_prompt_role",
		mcp.WithDescription("Assign role to an AI prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("role_data",
			mcp.Required(),
			mcp.Description("Role assignment data as JSON string"),
		),
	)

	s.AddTool(assignPromptRoleTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"role_data": args["role_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptAssignRole&t=json&promptID=%d", int(args["promptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign prompt role: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	selectPromptDataSourceTool := mcp.NewTool("select_prompt_data_source",
		mcp.WithDescription("Select data source for an AI prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("data_source",
			mcp.Required(),
			mcp.Description("Data source configuration as JSON string"),
		),
	)

	s.AddTool(selectPromptDataSourceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"data_source": args["data_source"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptSelectDataSource&t=json&promptID=%d", int(args["promptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to select prompt data source: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	setPromptPurposeTool := mcp.NewTool("set_prompt_purpose",
		mcp.WithDescription("Set purpose for an AI prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("purpose",
			mcp.Required(),
			mcp.Description("Prompt purpose description"),
		),
	)

	s.AddTool(setPromptPurposeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"purpose": args["purpose"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptSetPurpose&t=json&promptID=%d", int(args["promptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to set prompt purpose: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	setPromptTargetFormTool := mcp.NewTool("set_prompt_target_form",
		mcp.WithDescription("Set target form for an AI prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("target_form",
			mcp.Required(),
			mcp.Description("Target form configuration as JSON string"),
		),
	)

	s.AddTool(setPromptTargetFormTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"target_form": args["target_form"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptSetTargetForm&t=json&promptID=%d", int(args["promptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to set prompt target form: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	finalizePromptTool := mcp.NewTool("finalize_prompt",
		mcp.WithDescription("Finalize an AI prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("final_config",
			mcp.Description("Final configuration as JSON string"),
		),
	)

	s.AddTool(finalizePromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["final_config"]; ok && v != nil {
			body["final_config"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptFinalize&t=json&promptID=%d", int(args["promptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to finalize prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	executePromptTool := mcp.NewTool("execute_prompt",
		mcp.WithDescription("Execute an AI prompt"),
		mcp.WithNumber("promptId",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithNumber("objectId",
			mcp.Required(),
			mcp.Description("Object ID to execute prompt on"),
		),
		mcp.WithBoolean("auto",
			mcp.Description("Auto open target form and apply changes"),
		),
	)

	s.AddTool(executePromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&promptId=%d&objectId=%d", int(args["promptId"].(float64)), int(args["objectId"].(float64)))
		if v, ok := args["auto"]; ok && v != nil {
			if v.(bool) {
				queryParams += "&auto=1"
			} else {
				queryParams += "&auto=0"
			}
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptExecute&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to execute prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	resetPromptExecutionTool := mcp.NewTool("reset_prompt_execution",
		mcp.WithDescription("Reset prompt execution state"),
		mcp.WithBoolean("failed",
			mcp.Description("Whether the execution failed"),
		),
	)

	s.AddTool(resetPromptExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["failed"]; ok && v != nil {
			if v.(bool) {
				queryParams = "&failed=1"
			} else {
				queryParams = "&failed=0"
			}
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptExecutionReset&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reset prompt execution: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	auditPromptTool := mcp.NewTool("audit_prompt",
		mcp.WithDescription("Audit an AI prompt execution"),
		mcp.WithNumber("promptId",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithNumber("objectId",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithBoolean("exit",
			mcp.Description("Exit flag"),
		),
		mcp.WithString("audit_data",
			mcp.Description("Audit data as JSON string"),
		),
	)

	s.AddTool(auditPromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["audit_data"]; ok && v != nil {
			body["audit_data"] = v
		}

		queryParams := fmt.Sprintf("&promptId=%d&objectId=%d", int(args["promptId"].(float64)), int(args["objectId"].(float64)))
		if v, ok := args["exit"]; ok && v != nil {
			if v.(bool) {
				queryParams += "&exit=1"
			} else {
				queryParams += "&exit=0"
			}
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=ai&f=promptAudit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to audit prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	publishPromptTool := mcp.NewTool("publish_prompt",
		mcp.WithDescription("Publish an AI prompt"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithBoolean("backToTestingLocation",
			mcp.Description("Back to testing location flag"),
		),
	)

	s.AddTool(publishPromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&id=%d", int(args["id"].(float64)))
		if v, ok := args["backToTestingLocation"]; ok && v != nil {
			if v.(bool) {
				queryParams += "&backToTestingLocation=1"
			} else {
				queryParams += "&backToTestingLocation=0"
			}
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptPublish&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to publish prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	unpublishPromptTool := mcp.NewTool("unpublish_prompt",
		mcp.WithDescription("Unpublish an AI prompt"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
	)

	s.AddTool(unpublishPromptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=promptUnpublish&t=json&id=%d", int(args["id"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unpublish prompt: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTestingLocationTool := mcp.NewTool("get_testing_location",
		mcp.WithDescription("Get testing location for a prompt"),
		mcp.WithNumber("promptID",
			mcp.Required(),
			mcp.Description("Prompt ID"),
		),
		mcp.WithString("module",
			mcp.Description("Module name"),
		),
		mcp.WithString("targetForm",
			mcp.Description("Target form name"),
		),
	)

	s.AddTool(getTestingLocationTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&promptID=%d", int(args["promptID"].(float64)))
		if v, ok := args["module"]; ok && v != nil {
			queryParams += fmt.Sprintf("&module=%s", v)
		}
		if v, ok := args["targetForm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&targetForm=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=ai&f=ajaxGetTestingLocation&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get testing location: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getRoleTemplatesTool := mcp.NewTool("get_role_templates",
		mcp.WithDescription("Get AI role templates"),
		mcp.WithString("template_data",
			mcp.Description("Template filter data as JSON string"),
		),
	)

	s.AddTool(getRoleTemplatesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["template_data"]; ok && v != nil {
			body["template_data"] = v
		}

		resp, err := client.Post("/index.php?m=ai&f=roleTemplates&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get role templates: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
