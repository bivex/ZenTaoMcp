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

func RegisterProgramTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createProgramTool := mcp.NewTool("create_program",
		mcp.WithDescription("Create a new program in ZenTao"),
		mcp.WithString("name",
			mcp.Description("Program name"),
		),
		mcp.WithNumber("parent",
			mcp.Description("Parent program, 0 means no parent"),
		),
		mcp.WithString("PM",
			mcp.Description("Project Manager"),
		),
		mcp.WithNumber("budget",
			mcp.Description("Budget amount"),
		),
		mcp.WithString("budgetUnit",
			mcp.Description("Budget currency (CNY|USD)"),
			mcp.Enum("CNY", "USD"),
		),
		mcp.WithString("desc",
			mcp.Description("Program description"),
		),
		mcp.WithString("begin",
			mcp.Description("Estimated start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("Estimated end date (YYYY-MM-DD)"),
		),
		mcp.WithString("acl",
			mcp.Description("Access control (open|private)"),
			mcp.Enum("open", "private"),
		),
		mcp.WithArray("whitelist",
			mcp.Description("Whitelist, only effective when acl = private"),
		),
	)

	s.AddTool(createProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["parent"]; ok && v != nil {
			body["parent"] = int(v.(float64))
		}
		if v, ok := args["PM"]; ok && v != nil {
			body["PM"] = v
		}
		if v, ok := args["budget"]; ok && v != nil {
			body["budget"] = int(v.(float64))
		}
		if v, ok := args["budgetUnit"]; ok && v != nil {
			body["budgetUnit"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["acl"]; ok && v != nil {
			body["acl"] = v
		}
		if v, ok := args["whitelist"]; ok && v != nil {
			body["whitelist"] = v
		}

		resp, err := client.Post("/programs", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateProgramTool := mcp.NewTool("update_program",
		mcp.WithDescription("Update an existing program in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("name",
			mcp.Description("Program name"),
		),
		mcp.WithNumber("parent",
			mcp.Description("Parent program, 0 means no parent"),
		),
		mcp.WithString("PM",
			mcp.Description("Project Manager"),
		),
		mcp.WithNumber("budget",
			mcp.Description("Budget amount"),
		),
		mcp.WithString("budgetUnit",
			mcp.Description("Budget currency (CNY|USD)"),
			mcp.Enum("CNY", "USD"),
		),
		mcp.WithString("desc",
			mcp.Description("Program description"),
		),
		mcp.WithString("begin",
			mcp.Description("Estimated start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("Estimated end date (YYYY-MM-DD)"),
		),
		mcp.WithString("acl",
			mcp.Description("Access control (open|private)"),
			mcp.Enum("open", "private"),
		),
		mcp.WithArray("whitelist",
			mcp.Description("Whitelist, only effective when acl = private"),
		),
	)

	s.AddTool(updateProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["parent"]; ok && v != nil {
			body["parent"] = int(v.(float64))
		}
		if v, ok := args["PM"]; ok && v != nil {
			body["PM"] = v
		}
		if v, ok := args["budget"]; ok && v != nil {
			body["budget"] = int(v.(float64))
		}
		if v, ok := args["budgetUnit"]; ok && v != nil {
			body["budgetUnit"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["acl"]; ok && v != nil {
			body["acl"] = v
		}
		if v, ok := args["whitelist"]; ok && v != nil {
			body["whitelist"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/programs/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteProgramTool := mcp.NewTool("delete_program",
		mcp.WithDescription("Delete a program from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Program ID to delete"),
		),
	)

	s.AddTool(deleteProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/programs/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
