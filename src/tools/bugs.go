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

func RegisterBugTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createBugTool := mcp.NewTool("create_bug",
		mcp.WithDescription("Create a new bug in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Required(),
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Required(),
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(createBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title":    args["title"],
			"severity": int(args["severity"].(float64)),
			"pri":      int(args["pri"].(float64)),
			"type":     args["type"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/bugs", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateBugTool := mcp.NewTool("update_bug",
		mcp.WithDescription("Update an existing bug in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("title",
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(updateBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["severity"]; ok && v != nil {
			body["severity"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/bugs/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteBugTool := mcp.NewTool("delete_bug",
		mcp.WithDescription("Delete a bug from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID to delete"),
		),
	)

	s.AddTool(deleteBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/bugs/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get bugs list tool
	getBugsTool := mcp.NewTool("get_bugs",
		mcp.WithDescription("Get list of bugs in ZenTao"),
		mcp.WithNumber("product",
			mcp.Description("Filter by product ID"),
		),
		mcp.WithNumber("project",
			mcp.Description("Filter by project ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Filter by execution ID"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by bug status"),
			mcp.Enum("active", "resolved", "closed"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Filter by assigned user ID"),
		),
		mcp.WithNumber("openedBy",
			mcp.Description("Filter by opened by user ID"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of bugs to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if product, ok := args["product"].(float64); ok && product > 0 {
			params["product"] = fmt.Sprintf("%.0f", product)
		}
		if project, ok := args["project"].(float64); ok && project > 0 {
			params["project"] = fmt.Sprintf("%.0f", project)
		}
		if execution, ok := args["execution"].(float64); ok && execution > 0 {
			params["execution"] = fmt.Sprintf("%.0f", execution)
		}
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if assignedTo, ok := args["assignedTo"].(float64); ok && assignedTo > 0 {
			params["assignedTo"] = fmt.Sprintf("%.0f", assignedTo)
		}
		if openedBy, ok := args["openedBy"].(float64); ok && openedBy > 0 {
			params["openedBy"] = fmt.Sprintf("%.0f", openedBy)
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/bugs"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get bug details tool
	getBugTool := mcp.NewTool("get_bug",
		mcp.WithDescription("Get details of a specific bug by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
	)

	s.AddTool(getBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/bug/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
