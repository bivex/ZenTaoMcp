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

		resp, err := client.Delete(fmt.Sprintf("/testtasks/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test task: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
