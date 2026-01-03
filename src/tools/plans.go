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

func RegisterPlanTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createPlanTool := mcp.NewTool("create_plan",
		mcp.WithDescription("Create a new product plan in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Plan name"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithString("begin",
			mcp.Description("Plan start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("Plan end date (YYYY-MM-DD)"),
		),
		mcp.WithString("desc",
			mcp.Description("Plan description"),
		),
		mcp.WithNumber("parent",
			mcp.Description("Parent plan ID"),
		),
	)

	s.AddTool(createPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title": args["title"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["parent"]; ok && v != nil {
			body["parent"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/plans", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create plan: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updatePlanTool := mcp.NewTool("update_plan",
		mcp.WithDescription("Update an existing product plan in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Plan ID"),
		),
		mcp.WithString("title",
			mcp.Description("Plan name"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithString("begin",
			mcp.Description("Plan start date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("Plan end date (YYYY-MM-DD)"),
		),
		mcp.WithString("desc",
			mcp.Description("Plan description"),
		),
	)

	s.AddTool(updatePlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/productplans/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update plan: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deletePlanTool := mcp.NewTool("delete_plan",
		mcp.WithDescription("Delete a product plan from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Plan ID to delete"),
		),
	)

	s.AddTool(deletePlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/productsplan/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete plan: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkStoriesToPlanTool := mcp.NewTool("link_stories_to_plan",
		mcp.WithDescription("Link stories to a product plan in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Plan ID"),
		),
		mcp.WithArray("stories",
			mcp.Required(),
			mcp.Description("Story IDs to link"),
			mcp.Items(map[string]any{"type": "number"}),
		),
	)

	s.AddTool(linkStoriesToPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := map[string]interface{}{
			"stories": args["stories"],
		}

		resp, err := client.Post(fmt.Sprintf("/productplans/%d/linkstories", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkStoriesFromPlanTool := mcp.NewTool("unlink_stories_from_plan",
		mcp.WithDescription("Unlink stories from a product plan in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Plan ID"),
		),
		mcp.WithArray("stories",
			mcp.Required(),
			mcp.Description("Story IDs to unlink"),
			mcp.Items(map[string]any{"type": "number"}),
		),
	)

	s.AddTool(unlinkStoriesFromPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := map[string]interface{}{
			"stories": args["stories"],
		}

		resp, err := client.Post(fmt.Sprintf("/productplans/%d/unlinkstories", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkBugsToPlanTool := mcp.NewTool("link_bugs_to_plan",
		mcp.WithDescription("Link bugs to a product plan in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithArray("bugs",
			mcp.Required(),
			mcp.Description("Bug IDs to link"),
			mcp.Items(map[string]any{"type": "number"}),
		),
	)

	s.AddTool(linkBugsToPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"bugs": args["bugs"],
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/linkBugs", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkBugsFromPlanTool := mcp.NewTool("unlink_bugs_from_plan",
		mcp.WithDescription("Unlink bugs from a product plan in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Plan ID"),
		),
		mcp.WithArray("bugs",
			mcp.Required(),
			mcp.Description("Bug IDs to unlink"),
			mcp.Items(map[string]any{"type": "number"}),
		),
	)

	s.AddTool(unlinkBugsFromPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := map[string]interface{}{
			"bugs": args["bugs"],
		}

		resp, err := client.Post(fmt.Sprintf("/productplans/%d/unlinkbugs", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
