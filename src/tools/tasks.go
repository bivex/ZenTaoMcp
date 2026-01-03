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
}
