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
}
