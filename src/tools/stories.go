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

func RegisterStoryTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createStoryTool := mcp.NewTool("create_story",
		mcp.WithDescription("Create a new user story in ZenTao"),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Story title"),
		),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("pri",
			mcp.Required(),
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("category",
			mcp.Required(),
			mcp.Description("Story category"),
		),
		mcp.WithString("spec",
			mcp.Description("Story description"),
		),
		mcp.WithString("verify",
			mcp.Description("Acceptance criteria"),
		),
		mcp.WithString("source",
			mcp.Description("Source"),
		),
		mcp.WithString("sourceNote",
			mcp.Description("Source note"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(createStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"title":    args["title"],
			"product":  int(args["product"].(float64)),
			"pri":      int(args["pri"].(float64)),
			"category": args["category"],
		}

		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["verify"]; ok && v != nil {
			body["verify"] = v
		}
		if v, ok := args["source"]; ok && v != nil {
			body["source"] = v
		}
		if v, ok := args["sourceNote"]; ok && v != nil {
			body["sourceNote"] = v
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Post("/stories", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateStoryTool := mcp.NewTool("update_story",
		mcp.WithDescription("Update an existing story in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithString("source",
			mcp.Description("Source"),
		),
		mcp.WithString("sourceNote",
			mcp.Description("Source note"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("category",
			mcp.Description("Type (feature|interface|performance|safe|experience|improve|other)"),
			mcp.Enum("feature", "interface", "performance", "safe", "experience", "improve", "other"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(updateStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["source"]; ok && v != nil {
			body["source"] = v
		}
		if v, ok := args["sourceNote"]; ok && v != nil {
			body["sourceNote"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["category"]; ok && v != nil {
			body["category"] = v
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/stories/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	changeStoryTool := mcp.NewTool("change_story",
		mcp.WithDescription("Change story content"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
		mcp.WithString("title",
			mcp.Description("Story title"),
		),
		mcp.WithString("spec",
			mcp.Description("Story description"),
		),
		mcp.WithString("verify",
			mcp.Description("Acceptance criteria"),
		),
	)

	s.AddTool(changeStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["verify"]; ok && v != nil {
			body["verify"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/stories/%d/change", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to change story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteStoryTool := mcp.NewTool("delete_story",
		mcp.WithDescription("Delete a story from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID to delete"),
		),
	)

	s.AddTool(deleteStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/stories/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get stories list tool
	getStoriesTool := mcp.NewTool("get_stories",
		mcp.WithDescription("Get list of user stories in ZenTao"),
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
			mcp.Description("Filter by story status"),
			mcp.Enum("draft", "active", "changed", "closed"),
		),
		mcp.WithString("stage",
			mcp.Description("Filter by story stage"),
			mcp.Enum("wait", "planned", "projected", "developing", "developed", "testing", "tested", "verified", "released", "closed"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Filter by priority (1-9)"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of stories to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getStoriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if stage, ok := args["stage"].(string); ok && stage != "" {
			params["stage"] = stage
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

		path := "/stories"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get story details tool
	getStoryTool := mcp.NewTool("get_story",
		mcp.WithDescription("Get details of a specific story by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
	)

	s.AddTool(getStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/story/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
