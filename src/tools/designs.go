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

func RegisterDesignTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseDesignsTool := mcp.NewTool("browse_designs",
		mcp.WithDescription("Browse designs for a project/product"),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithString("type",
			mcp.Description("Design type"),
			mcp.Enum("all", "bySearch", "HLDS", "DDS", "DBDS", "ADS"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
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

	s.AddTool(browseDesignsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=browse&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse designs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createDesignTool := mcp.NewTool("create_design",
		mcp.WithDescription("Create a new design document"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Design type"),
			mcp.Enum("all", "bySearch", "HLDS", "DDS", "DBDS", "ADS"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Design name"),
		),
		mcp.WithString("desc",
			mcp.Description("Design description"),
		),
		mcp.WithString("content",
			mcp.Description("Design content"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
	)

	s.AddTool(createDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
			"type": args["type"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["content"]; ok && v != nil {
			body["content"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=design&f=create&t=json&projectID=%d&productID=%d&type=%s",
			int(args["projectID"].(float64)), int(args["productID"].(float64)), args["type"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateDesignsTool := mcp.NewTool("batch_create_designs",
		mcp.WithDescription("Create multiple design documents"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Design type"),
			mcp.Enum("all", "bySearch", "HLDS", "DDS", "DBDS", "ADS"),
		),
		mcp.WithArray("names",
			mcp.Required(),
			mcp.Description("Design names"),
		),
		mcp.WithArray("descs",
			mcp.Description("Design descriptions"),
		),
	)

	s.AddTool(batchCreateDesignsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"names": args["names"],
		}

		if v, ok := args["descs"]; ok && v != nil {
			body["descs"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=design&f=batchCreate&t=json&projectID=%d&productID=%d&type=%s",
			int(args["projectID"].(float64)), int(args["productID"].(float64)), args["type"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create designs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewDesignTool := mcp.NewTool("view_design",
		mcp.WithDescription("View a design document"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
	)

	s.AddTool(viewDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		designID := int(args["designID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=view&t=json&designID=%d", designID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editDesignTool := mcp.NewTool("edit_design",
		mcp.WithDescription("Edit a design document"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
		mcp.WithString("name",
			mcp.Description("Design name"),
		),
		mcp.WithString("desc",
			mcp.Description("Design description"),
		),
		mcp.WithString("content",
			mcp.Description("Design content"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
	)

	s.AddTool(editDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["content"]; ok && v != nil {
			body["content"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=design&f=edit&t=json&designID=%d", int(args["designID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteDesignTool := mcp.NewTool("delete_design",
		mcp.WithDescription("Delete a design document"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
	)

	s.AddTool(deleteDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		designID := int(args["designID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=delete&t=json&designID=%d", designID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	assignDesignTool := mcp.NewTool("assign_design",
		mcp.WithDescription("Assign a design to a user"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Required(),
			mcp.Description("User ID to assign to"),
		),
	)

	s.AddTool(assignDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"assignedTo": int(args["assignedTo"].(float64)),
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=design&f=assignTo&t=json&designID=%d", int(args["designID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkCommitToDesignTool := mcp.NewTool("link_commit_to_design",
		mcp.WithDescription("Link a commit to a design document"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
		mcp.WithNumber("repoID",
			mcp.Required(),
			mcp.Description("Repository ID"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin commit hash"),
		),
		mcp.WithString("end",
			mcp.Description("End commit hash"),
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

	s.AddTool(linkCommitToDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}

		queryParams := fmt.Sprintf("&designID=%d&repoID=%d", int(args["designID"].(float64)), int(args["repoID"].(float64)))
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=design&f=linkCommit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link commit to design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkCommitFromDesignTool := mcp.NewTool("unlink_commit_from_design",
		mcp.WithDescription("Unlink a commit from a design document"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
		mcp.WithNumber("commitID",
			mcp.Required(),
			mcp.Description("Commit ID"),
		),
	)

	s.AddTool(unlinkCommitFromDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=unlinkCommit&t=json&designID=%d&commitID=%d",
			int(args["designID"].(float64)), int(args["commitID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink commit from design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewDesignCommitsTool := mcp.NewTool("view_design_commits",
		mcp.WithDescription("View commits linked to a design"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
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

	s.AddTool(viewDesignCommitsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("designID=%d", int(args["designID"].(float64)))
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=viewCommit&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view design commits: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getDesignSwitcherMenuTool := mcp.NewTool("get_design_switcher_menu",
		mcp.WithDescription("Get design switcher menu"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
	)

	s.AddTool(getDesignSwitcherMenuTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=ajaxSwitcherMenu&t=json&projectID=%d&productID=%d",
			int(args["projectID"].(float64)), int(args["productID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get design switcher menu: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProductStoriesForDesignTool := mcp.NewTool("get_product_stories_for_design",
		mcp.WithDescription("Get product stories for design purposes"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("status",
			mcp.Description("Story status"),
		),
		mcp.WithString("hasParent",
			mcp.Description("Has parent story"),
		),
	)

	s.AddTool(getProductStoriesForDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d&projectID=%d", int(args["productID"].(float64)), int(args["projectID"].(float64)))
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}
		if v, ok := args["hasParent"]; ok && v != nil {
			queryParams += fmt.Sprintf("&hasParent=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=ajaxGetProductStories&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get product stories for design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	confirmStoryChangeForDesignTool := mcp.NewTool("confirm_story_change_for_design",
		mcp.WithDescription("Confirm story change for design"),
		mcp.WithNumber("designID",
			mcp.Required(),
			mcp.Description("Design ID"),
		),
	)

	s.AddTool(confirmStoryChangeForDesignTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		designID := int(args["designID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=design&f=confirmStoryChange&t=json&designID=%d", designID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to confirm story change for design: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
