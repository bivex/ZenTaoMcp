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

func RegisterBuildTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createBuildTool := mcp.NewTool("create_build",
		mcp.WithDescription("Create a new build in ZenTao"),
		mcp.WithNumber("project",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("execution",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Build name"),
		),
		mcp.WithString("builder",
			mcp.Required(),
			mcp.Description("Builder user account"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithString("date",
			mcp.Description("Build date (YYYY-MM-DD)"),
		),
		mcp.WithString("scmPath",
			mcp.Description("Source code path"),
		),
		mcp.WithString("filePath",
			mcp.Description("Download URL"),
		),
		mcp.WithString("desc",
			mcp.Description("Build description"),
		),
	)

	s.AddTool(createBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		projectID := int(args["project"].(float64))

		body := map[string]interface{}{
			"execution": int(args["execution"].(float64)),
			"product":   int(args["product"].(float64)),
			"name":      args["name"],
			"builder":   args["builder"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["date"]; ok && v != nil {
			body["date"] = v
		}
		if v, ok := args["scmPath"]; ok && v != nil {
			body["scmPath"] = v
		}
		if v, ok := args["filePath"]; ok && v != nil {
			body["filePath"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/projects/%d/builds", projectID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateBuildTool := mcp.NewTool("update_build",
		mcp.WithDescription("Update an existing build in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("name",
			mcp.Description("Build name"),
		),
		mcp.WithString("builder",
			mcp.Description("Builder user account"),
		),
		mcp.WithString("date",
			mcp.Description("Build date (YYYY-MM-DD)"),
		),
		mcp.WithString("scmPath",
			mcp.Description("Source code path"),
		),
		mcp.WithString("filePath",
			mcp.Description("Download URL"),
		),
		mcp.WithString("desc",
			mcp.Description("Build description"),
		),
	)

	s.AddTool(updateBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["builder"]; ok && v != nil {
			body["builder"] = v
		}
		if v, ok := args["date"]; ok && v != nil {
			body["date"] = v
		}
		if v, ok := args["scmPath"]; ok && v != nil {
			body["scmPath"] = v
		}
		if v, ok := args["filePath"]; ok && v != nil {
			body["filePath"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/builds/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteBuildTool := mcp.NewTool("delete_build",
		mcp.WithDescription("Delete a build from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Build ID to delete"),
		),
	)

	s.AddTool(deleteBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/builds/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
