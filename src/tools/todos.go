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

func RegisterTodoTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTodoTool := mcp.NewTool("create_todo",
		mcp.WithDescription("Create a new todo item"),
		mcp.WithString("date",
			mcp.Description("Todo date (YYYY-MM-DD)"),
		),
		mcp.WithString("from",
			mcp.Description("Source of todo"),
			mcp.Enum("todo", "feedback", "block"),
		),
		mcp.WithString("name",
			mcp.Description("Todo name"),
		),
		mcp.WithString("desc",
			mcp.Description("Todo description"),
		),
		mcp.WithString("status",
			mcp.Description("Todo status"),
			mcp.Enum("wait", "doing", "done"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin time (HH:mm)"),
		),
		mcp.WithString("end",
			mcp.Description("End time (HH:mm)"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
	)

	s.AddTool(createTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["status"]; ok && v != nil {
			body["status"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}

		queryParams := ""
		if v, ok := args["date"]; ok && v != nil {
			queryParams += fmt.Sprintf("&date=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=create&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateTodosTool := mcp.NewTool("batch_create_todos",
		mcp.WithDescription("Create multiple todo items at once"),
		mcp.WithString("date",
			mcp.Required(),
			mcp.Description("Todo date (YYYY-MM-DD)"),
		),
		mcp.WithArray("names",
			mcp.Required(),
			mcp.Description("Array of todo names"),
		),
		mcp.WithArray("descs",
			mcp.Description("Array of todo descriptions"),
		),
	)

	s.AddTool(batchCreateTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"names": args["names"],
		}

		if v, ok := args["descs"]; ok && v != nil {
			body["descs"] = v
		}

		queryParams := fmt.Sprintf("&date=%s", args["date"])

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=batchCreate&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editTodoTool := mcp.NewTool("edit_todo",
		mcp.WithDescription("Edit an existing todo item"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
		mcp.WithString("name",
			mcp.Description("Todo name"),
		),
		mcp.WithString("desc",
			mcp.Description("Todo description"),
		),
		mcp.WithString("status",
			mcp.Description("Todo status"),
			mcp.Enum("wait", "doing", "done"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin time (HH:mm)"),
		),
		mcp.WithString("end",
			mcp.Description("End time (HH:mm)"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
	)

	s.AddTool(editTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["status"]; ok && v != nil {
			body["status"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=edit&t=json&todoID=%d", int(args["todoID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditTodosTool := mcp.NewTool("batch_edit_todos",
		mcp.WithDescription("Edit multiple todo items at once"),
		mcp.WithString("from",
			mcp.Description("Source filter"),
		),
		mcp.WithString("type",
			mcp.Description("Type filter"),
		),
		mcp.WithNumber("userID",
			mcp.Description("User ID filter"),
		),
		mcp.WithString("status",
			mcp.Description("Status filter"),
		),
		mcp.WithString("newStatus",
			mcp.Description("New status to set"),
		),
		mcp.WithNumber("pri",
			mcp.Description("New priority"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("New assigned user"),
		),
	)

	s.AddTool(batchEditTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["newStatus"]; ok && v != nil {
			body["status"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}

		queryParams := ""
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["userID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&userID=%d", int(v.(float64)))
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=batchEdit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	startTodoTool := mcp.NewTool("start_todo",
		mcp.WithDescription("Start a todo item"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
	)

	s.AddTool(startTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=start&t=json&todoID=%d", todoID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateTodoTool := mcp.NewTool("activate_todo",
		mcp.WithDescription("Activate a todo item"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
	)

	s.AddTool(activateTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=activate&t=json&todoID=%d", todoID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeTodoTool := mcp.NewTool("close_todo",
		mcp.WithDescription("Close a todo item"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
	)

	s.AddTool(closeTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=close&t=json&todoID=%d", todoID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	assignTodoTool := mcp.NewTool("assign_todo",
		mcp.WithDescription("Assign a todo item to a user"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
		mcp.WithString("comment",
			mcp.Description("Assignment comment"),
		),
	)

	s.AddTool(assignTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": int(args["assignedTo"].(float64)),
		}

		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=assignTo&t=json&todoID=%d", int(args["todoID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewTodoTool := mcp.NewTool("view_todo",
		mcp.WithDescription("View todo details"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
		mcp.WithString("from",
			mcp.Description("View source"),
			mcp.Enum("my", "company"),
		),
	)

	s.AddTool(viewTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		queryParams := ""
		if v, ok := args["from"]; ok && v != nil {
			queryParams = fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=view&t=json&todoID=%d%s", todoID, queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTodoTool := mcp.NewTool("delete_todo",
		mcp.WithDescription("Delete a todo item"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
	)

	s.AddTool(deleteTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		queryParams := ""
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams = fmt.Sprintf("&confirm=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=delete&t=json&todoID=%d%s", todoID, queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	finishTodoTool := mcp.NewTool("finish_todo",
		mcp.WithDescription("Mark a todo as finished"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
	)

	s.AddTool(finishTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=finish&t=json&todoID=%d", todoID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to finish todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchFinishTodosTool := mcp.NewTool("batch_finish_todos",
		mcp.WithDescription("Mark multiple todos as finished"),
	)

	s.AddTool(batchFinishTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=todo&f=batchFinish&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch finish todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCloseTodosTool := mcp.NewTool("batch_close_todos",
		mcp.WithDescription("Close multiple todos"),
	)

	s.AddTool(batchCloseTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=todo&f=batchClose&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch close todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	importTodoToTodayTool := mcp.NewTool("import_todo_to_today",
		mcp.WithDescription("Import a todo to today's list"),
		mcp.WithString("todoID",
			mcp.Required(),
			mcp.Description("Todo ID(s) to import"),
		),
	)

	s.AddTool(importTodoToTodayTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=import2Today&t=json&todoID=%s", args["todoID"]), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import todo to today: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportTodosTool := mcp.NewTool("export_todos",
		mcp.WithDescription("Export todos to file"),
		mcp.WithNumber("userID",
			mcp.Description("User ID to export todos for"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
	)

	s.AddTool(exportTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["userID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&userID=%d", int(v.(float64)))
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=todo&f=export&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getTodoDetailTool := mcp.NewTool("get_todo_detail",
		mcp.WithDescription("Get detailed information about a todo"),
		mcp.WithNumber("todoID",
			mcp.Required(),
			mcp.Description("Todo ID"),
		),
	)

	s.AddTool(getTodoDetailTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		todoID := int(args["todoID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=ajaxGetDetail&t=json&todoID=%d", todoID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get todo detail: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProgramIDTool := mcp.NewTool("get_program_id",
		mcp.WithDescription("Get program ID for a todo object"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithString("objectType",
			mcp.Required(),
			mcp.Description("Object type"),
		),
	)

	s.AddTool(getProgramIDTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=ajaxGetProgramID&t=json&objectID=%d&objectType=%s",
			int(args["objectID"].(float64)), args["objectType"]))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get program ID: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionPairsTool := mcp.NewTool("get_execution_pairs",
		mcp.WithDescription("Get execution pairs for a project"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getExecutionPairsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectID := int(args["projectID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=ajaxGetExecutionPairs&t=json&projectID=%d", projectID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution pairs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProductPairsTool := mcp.NewTool("get_product_pairs",
		mcp.WithDescription("Get product pairs for a project"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getProductPairsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectID := int(args["projectID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=todo&f=ajaxGetProductPairs&t=json&projectID=%d", projectID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get product pairs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createCycleTodoTool := mcp.NewTool("create_cycle_todo",
		mcp.WithDescription("Create a recurring/cycle todo"),
	)

	s.AddTool(createCycleTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=todo&f=createCycle&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create cycle todo: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
