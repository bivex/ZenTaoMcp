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

func RegisterProjectTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createProjectTool := mcp.NewTool("create_project",
		mcp.WithDescription("Create a new project in ZenTao"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Project name"),
		),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("Project code"),
		),
		mcp.WithString("begin",
			mcp.Required(),
			mcp.Description("Planned start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Required(),
			mcp.Description("Planned end date (YYYY-MM-DD)"),
		),
		mcp.WithArray("products",
			mcp.Required(),
			mcp.Description("Associated product IDs"),
			mcp.Items(map[string]any{"type": "number"}),
		),
		mcp.WithString("model",
			mcp.Description("Project model (scrum|agileplus|waterfall|kanban)"),
			mcp.Enum("scrum", "agileplus", "waterfall", "kanban"),
		),
		mcp.WithNumber("parent",
			mcp.Description("Parent program, 0 means no parent"),
		),
	)

	s.AddTool(createProjectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		var products []int
		if prods, ok := args["products"].([]interface{}); ok {
			for _, p := range prods {
				products = append(products, int(p.(float64)))
			}
		}

		body := map[string]interface{}{
			"name":     args["name"],
			"code":     args["code"],
			"begin":    args["begin"],
			"end":      args["end"],
			"products": products,
		}

		if v, ok := args["model"]; ok && v != nil {
			body["model"] = v
		}
		if v, ok := args["parent"]; ok && v != nil {
			body["parent"] = int(v.(float64))
		}

		resp, err := client.Post("/projects", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create project: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateProjectTool := mcp.NewTool("update_project",
		mcp.WithDescription("Update an existing project in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("name",
			mcp.Description("Project name"),
		),
		mcp.WithString("code",
			mcp.Description("Project code"),
		),
		mcp.WithNumber("parent",
			mcp.Description("Parent program"),
		),
		mcp.WithNumber("PM",
			mcp.Description("Project Manager ID"),
		),
		mcp.WithNumber("budget",
			mcp.Description("Project budget amount"),
		),
		mcp.WithString("budgetUnit",
			mcp.Description("Budget currency (CNY|USD)"),
			mcp.Enum("CNY", "USD"),
		),
		mcp.WithNumber("days",
			mcp.Description("Available workdays"),
		),
		mcp.WithString("desc",
			mcp.Description("Project description"),
		),
		mcp.WithString("acl",
			mcp.Description("Access control (open|private)"),
			mcp.Enum("open", "private"),
		),
	)

	s.AddTool(updateProjectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}
		if v, ok := args["parent"]; ok && v != nil {
			body["parent"] = int(v.(float64))
		}
		if v, ok := args["PM"]; ok && v != nil {
			body["PM"] = int(v.(float64))
		}
		if v, ok := args["budget"]; ok && v != nil {
			body["budget"] = int(v.(float64))
		}
		if v, ok := args["budgetUnit"]; ok && v != nil {
			body["budgetUnit"] = v
		}
		if v, ok := args["days"]; ok && v != nil {
			body["days"] = int(v.(float64))
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["acl"]; ok && v != nil {
			body["acl"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/projects/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update project: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProjectStoriesTool := mcp.NewTool("get_project_stories",
		mcp.WithDescription("Get stories for a specific project"),
		mcp.WithNumber("project_id",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of stories to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getProjectStoriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/projects/%d/stories%s", projectID, queryString))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get project stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteProjectTool := mcp.NewTool("delete_project",
		mcp.WithDescription("Delete a project from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Project ID to delete"),
		),
	)

	s.AddTool(deleteProjectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/projects/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete project: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createExecutionTool := mcp.NewTool("create_execution",
		mcp.WithDescription("Create a new execution (sprint/iteration) in ZenTao"),
		mcp.WithNumber("project",
			mcp.Required(),
			mcp.Description("Parent project ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Execution name"),
		),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("Execution code"),
		),
		mcp.WithString("begin",
			mcp.Required(),
			mcp.Description("Planned start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Required(),
			mcp.Description("Planned end date (YYYY-MM-DD)"),
		),
		mcp.WithNumber("days",
			mcp.Description("Available workdays"),
		),
		mcp.WithString("lifetime",
			mcp.Description("Type (short|long|ops)"),
			mcp.Enum("short", "long", "ops"),
		),
		mcp.WithString("PO",
			mcp.Description("Product Owner"),
		),
		mcp.WithString("PM",
			mcp.Description("Iteration Manager"),
		),
		mcp.WithString("QD",
			mcp.Description("Quality Director"),
		),
		mcp.WithString("RD",
			mcp.Description("Release Director"),
		),
		mcp.WithArray("teamMembers",
			mcp.Description("Team members"),
		),
		mcp.WithString("desc",
			mcp.Description("Iteration description"),
		),
	)

	s.AddTool(createExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"project": int(args["project"].(float64)),
			"name":    args["name"],
			"code":    args["code"],
			"begin":   args["begin"],
			"end":     args["end"],
		}

		if v, ok := args["days"]; ok && v != nil {
			body["days"] = int(v.(float64))
		}
		if v, ok := args["lifetime"]; ok && v != nil {
			body["lifetime"] = v
		}
		if v, ok := args["PO"]; ok && v != nil {
			body["PO"] = v
		}
		if v, ok := args["PM"]; ok && v != nil {
			body["PM"] = v
		}
		if v, ok := args["QD"]; ok && v != nil {
			body["QD"] = v
		}
		if v, ok := args["RD"]; ok && v != nil {
			body["RD"] = v
		}
		if v, ok := args["teamMembers"]; ok && v != nil {
			body["teamMembers"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/projects/%d/executions", int(args["project"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateExecutionTool := mcp.NewTool("update_execution",
		mcp.WithDescription("Update an execution (sprint/iteration) in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("name",
			mcp.Description("Execution name"),
		),
		mcp.WithString("code",
			mcp.Description("Execution code"),
		),
		mcp.WithString("begin",
			mcp.Description("Planned start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("Planned end date (YYYY-MM-DD)"),
		),
		mcp.WithNumber("days",
			mcp.Description("Available workdays"),
		),
		mcp.WithString("lifetime",
			mcp.Description("Type (short|long|ops)"),
			mcp.Enum("short", "long", "ops"),
		),
		mcp.WithString("PO",
			mcp.Description("Product Owner"),
		),
		mcp.WithString("PM",
			mcp.Description("Iteration Manager"),
		),
		mcp.WithString("QD",
			mcp.Description("Quality Director"),
		),
		mcp.WithString("RD",
			mcp.Description("Release Director"),
		),
		mcp.WithArray("teamMembers",
			mcp.Description("Team members"),
		),
		mcp.WithString("desc",
			mcp.Description("Iteration description"),
		),
	)

	s.AddTool(updateExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["days"]; ok && v != nil {
			body["days"] = int(v.(float64))
		}
		if v, ok := args["lifetime"]; ok && v != nil {
			body["lifetime"] = v
		}
		if v, ok := args["PO"]; ok && v != nil {
			body["PO"] = v
		}
		if v, ok := args["PM"]; ok && v != nil {
			body["PM"] = v
		}
		if v, ok := args["QD"]; ok && v != nil {
			body["QD"] = v
		}
		if v, ok := args["RD"]; ok && v != nil {
			body["RD"] = v
		}
		if v, ok := args["teamMembers"]; ok && v != nil {
			body["teamMembers"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/executions/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionStoriesTool := mcp.NewTool("get_execution_stories",
		mcp.WithDescription("Get stories for a specific execution (sprint)"),
		mcp.WithNumber("execution_id",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of stories to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getExecutionStoriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		executionID := int(args["execution_id"].(float64))

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

		resp, err := client.Get(fmt.Sprintf("/executions/%d/stories%s", executionID, queryString))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteExecutionTool := mcp.NewTool("delete_execution",
		mcp.WithDescription("Delete an execution from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Execution ID to delete"),
		),
	)

	s.AddTool(deleteExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/executions/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get projects list tool
	getProjectsTool := mcp.NewTool("get_projects",
		mcp.WithDescription("Get list of projects in ZenTao"),
		mcp.WithString("status",
			mcp.Description("Filter by project status"),
			mcp.Enum("wait", "doing", "suspended", "closed"),
		),
		mcp.WithNumber("program",
			mcp.Description("Filter by program ID"),
		),
		mcp.WithString("model",
			mcp.Description("Filter by project model"),
			mcp.Enum("scrum", "agileplus", "waterfall", "kanban"),
		),
		mcp.WithNumber("PM",
			mcp.Description("Filter by project manager ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of projects to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getProjectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if program, ok := args["program"].(float64); ok && program > 0 {
			params["program"] = fmt.Sprintf("%.0f", program)
		}
		if model, ok := args["model"].(string); ok && model != "" {
			params["model"] = model
		}
		if pm, ok := args["PM"].(float64); ok && pm > 0 {
			params["PM"] = fmt.Sprintf("%.0f", pm)
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/projects"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get projects: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get project details tool
	getProjectTool := mcp.NewTool("get_project",
		mcp.WithDescription("Get details of a specific project by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getProjectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/project/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get project: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get executions list tool
	getExecutionsTool := mcp.NewTool("get_executions",
		mcp.WithDescription("Get list of executions (sprints/iterations) in ZenTao"),
		mcp.WithNumber("project",
			mcp.Description("Filter by project ID"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by execution status"),
			mcp.Enum("wait", "doing", "suspended", "closed"),
		),
		mcp.WithString("type",
			mcp.Description("Filter by execution type"),
			mcp.Enum("sprint", "stage", "kanban"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of executions to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getExecutionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if project, ok := args["project"].(float64); ok && project > 0 {
			params["project"] = fmt.Sprintf("%.0f", project)
		}
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if execType, ok := args["type"].(string); ok && execType != "" {
			params["type"] = execType
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/executions"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get executions: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get execution details tool
	getExecutionTool := mcp.NewTool("get_execution",
		mcp.WithDescription("Get details of a specific execution by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
	)

	s.AddTool(getExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/execution/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
