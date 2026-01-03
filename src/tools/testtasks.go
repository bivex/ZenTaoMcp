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
