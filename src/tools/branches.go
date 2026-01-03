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

func RegisterBranchTools(s *server.MCPServer, client *client.ZenTaoClient) {
	manageBranchesTool := mcp.NewTool("manage_branches",
		mcp.WithDescription("Manage branches for a product"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(manageBranchesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=branch&f=manage&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage branches: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createBranchTool := mcp.NewTool("create_branch",
		mcp.WithDescription("Create a new branch for a product"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Branch name"),
		),
		mcp.WithString("desc",
			mcp.Description("Branch description"),
		),
	)

	s.AddTool(createBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=branch&f=create&t=json&productID=%d", int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create branch: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editBranchTool := mcp.NewTool("edit_branch",
		mcp.WithDescription("Edit an existing branch"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("name",
			mcp.Description("Branch name"),
		),
		mcp.WithString("desc",
			mcp.Description("Branch description"),
		),
	)

	s.AddTool(editBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=branch&f=edit&t=json&branchID=%d&productID=%d", int(args["branchID"].(float64)), int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit branch: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchEditBranchesTool := mcp.NewTool("batch_edit_branches",
		mcp.WithDescription("Batch edit branches for a product"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithArray("branchIDs",
			mcp.Required(),
			mcp.Description("Branch IDs to edit"),
		),
		mcp.WithArray("names",
			mcp.Description("New branch names"),
		),
		mcp.WithArray("descs",
			mcp.Description("New branch descriptions"),
		),
	)

	s.AddTool(batchEditBranchesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"branchIDs": args["branchIDs"],
		}

		if v, ok := args["names"]; ok && v != nil {
			body["names"] = v
		}
		if v, ok := args["descs"]; ok && v != nil {
			body["descs"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=branch&f=batchEdit&t=json&productID=%d", int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit branches: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeBranchTool := mcp.NewTool("close_branch",
		mcp.WithDescription("Close a branch"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("Branch ID"),
		),
	)

	s.AddTool(closeBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		branchID := int(args["branchID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=branch&f=close&t=json&branchID=%d", branchID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close branch: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateBranchTool := mcp.NewTool("activate_branch",
		mcp.WithDescription("Activate a branch"),
		mcp.WithNumber("branchID",
			mcp.Required(),
			mcp.Description("Branch ID"),
		),
	)

	s.AddTool(activateBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		branchID := int(args["branchID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=branch&f=activate&t=json&branchID=%d", branchID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate branch: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	sortBranchesTool := mcp.NewTool("sort_branches",
		mcp.WithDescription("Sort branches"),
		mcp.WithArray("branchOrders",
			mcp.Required(),
			mcp.Description("Branch order mapping"),
		),
	)

	s.AddTool(sortBranchesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"orders": args["branchOrders"],
		}

		resp, err := client.Post("/index.php?m=branch&f=sort&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to sort branches: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getBranchesTool := mcp.NewTool("get_branches",
		mcp.WithDescription("Get branches for a product"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("oldBranch",
			mcp.Description("Old branch"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithString("withMainBranch",
			mcp.Description("Include main branch"),
		),
		mcp.WithString("isTwins",
			mcp.Description("Is twins"),
		),
		mcp.WithString("fieldID",
			mcp.Description("Field ID"),
		),
		mcp.WithString("multiple",
			mcp.Description("Multiple selection"),
		),
		mcp.WithNumber("charterID",
			mcp.Description("Charter ID"),
		),
	)

	s.AddTool(getBranchesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["oldBranch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&oldBranch=%s", v)
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}
		if v, ok := args["withMainBranch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&withMainBranch=%s", v)
		}
		if v, ok := args["isTwins"]; ok && v != nil {
			queryParams += fmt.Sprintf("&isTwins=%s", v)
		}
		if v, ok := args["fieldID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&fieldID=%s", v)
		}
		if v, ok := args["multiple"]; ok && v != nil {
			queryParams += fmt.Sprintf("&multiple=%s", v)
		}
		if v, ok := args["charterID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&charterID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=branch&f=ajaxGetBranches&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get branches: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	mergeBranchTool := mcp.NewTool("merge_branch",
		mcp.WithDescription("Merge branches for a product"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("sourceBranch",
			mcp.Required(),
			mcp.Description("Source branch ID"),
		),
		mcp.WithNumber("targetBranch",
			mcp.Required(),
			mcp.Description("Target branch ID"),
		),
	)

	s.AddTool(mergeBranchTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"sourceBranch": int(args["sourceBranch"].(float64)),
			"targetBranch": int(args["targetBranch"].(float64)),
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=branch&f=mergeBranch&t=json&productID=%d", int(args["productID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to merge branches: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
