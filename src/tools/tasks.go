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

func RegisterTaskTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTaskTool := mcp.NewTool("create_task",
		mcp.WithDescription("Create a new task in ZenTao"),
		mcp.WithNumber("execution",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Task name"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Task type (design|devel|request|test|study|discuss|ui|affair|misc)"),
			mcp.Enum("design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"),
		),
		mcp.WithArray("assignedTo",
			mcp.Required(),
			mcp.Description("Assigned to user accounts"),
		),
		mcp.WithString("estStarted",
			mcp.Required(),
			mcp.Description("Estimated start date (YYYY-MM-DD)"),
		),
		mcp.WithString("deadline",
			mcp.Required(),
			mcp.Description("Estimated end date (YYYY-MM-DD)"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Associated story ID"),
		),
		mcp.WithNumber("fromBug",
			mcp.Description("From bug ID"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
	)

	s.AddTool(createTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		executionID := int(args["execution"].(float64))

		body := map[string]interface{}{
			"name":       args["name"],
			"type":       args["type"],
			"assignedTo": args["assignedTo"],
			"estStarted": args["estStarted"],
			"deadline":   args["deadline"],
			"openedBy":   1,
			"openedDate": fmt.Sprintf("%s", args["estStarted"]),
		}

		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["fromBug"]; ok && v != nil {
			body["fromBug"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}

		resp, err := client.Post(fmt.Sprintf("/executions/%d/tasks", executionID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateTaskTool := mcp.NewTool("update_task",
		mcp.WithDescription("Update an existing task in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Task ID"),
		),
		mcp.WithString("name",
			mcp.Description("Task name"),
		),
		mcp.WithString("type",
			mcp.Description("Task type"),
			mcp.Enum("design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"),
		),
		mcp.WithArray("assignedTo",
			mcp.Description("Assigned to user accounts"),
		),
		mcp.WithString("estStarted",
			mcp.Description("Estimated start date (YYYY-MM-DD)"),
		),
		mcp.WithString("deadline",
			mcp.Description("Estimated end date (YYYY-MM-DD)"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Associated story ID"),
		),
		mcp.WithNumber("fromBug",
			mcp.Description("From bug ID"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
	)

	s.AddTool(updateTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = v
		}
		if v, ok := args["estStarted"]; ok && v != nil {
			body["estStarted"] = v
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["fromBug"]; ok && v != nil {
			body["fromBug"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}

		resp, err := client.Put(fmt.Sprintf("/tasks/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTaskTool := mcp.NewTool("delete_task",
		mcp.WithDescription("Delete a task from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Task ID to delete"),
		),
	)

	s.AddTool(deleteTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/tasks/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get tasks list tool
	getTasksTool := mcp.NewTool("get_tasks",
		mcp.WithDescription("Get list of tasks in ZenTao"),
		mcp.WithNumber("execution",
			mcp.Description("Filter by execution ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Filter by story ID"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by task status"),
			mcp.Enum("wait", "doing", "done", "pause", "cancel", "closed"),
		),
		mcp.WithString("type",
			mcp.Description("Filter by task type"),
			mcp.Enum("design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Filter by assigned user ID"),
		),
		mcp.WithNumber("openedBy",
			mcp.Description("Filter by opened by user ID"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Filter by priority (1-9)"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of tasks to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if execution, ok := args["execution"].(float64); ok && execution > 0 {
			params["execution"] = fmt.Sprintf("%.0f", execution)
		}
		if story, ok := args["story"].(float64); ok && story > 0 {
			params["story"] = fmt.Sprintf("%.0f", story)
		}
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if taskType, ok := args["type"].(string); ok && taskType != "" {
			params["type"] = taskType
		}
		if assignedTo, ok := args["assignedTo"].(float64); ok && assignedTo > 0 {
			params["assignedTo"] = fmt.Sprintf("%.0f", assignedTo)
		}
		if openedBy, ok := args["openedBy"].(float64); ok && openedBy > 0 {
			params["openedBy"] = fmt.Sprintf("%.0f", openedBy)
		}
		if pri, ok := args["pri"].(float64); ok && pri > 0 && pri <= 9 {
			params["pri"] = fmt.Sprintf("%.0f", pri)
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/tasks"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get tasks: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get task details tool
	getTaskTool := mcp.NewTool("get_task",
		mcp.WithDescription("Get details of a specific task by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Task ID"),
		),
	)

	s.AddTool(getTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/task/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
