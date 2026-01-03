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

// RegisterAdminTools registers all company, department, group, and user management tools
func RegisterAdminTools(s *server.MCPServer, client *client.ZenTaoClient) {
	registerCompanyTools(s, client)
	registerDepartmentTools(s, client)
	registerGroupTools(s, client)
	registerUserManagementTools(s, client)
}

func registerCompanyTools(s *server.MCPServer, client *client.ZenTaoClient) {
	companyIndexTool := mcp.NewTool("company_index",
		mcp.WithDescription("Get company index"),
	)

	s.AddTool(companyIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=company&f=index&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get company index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	companyBrowseTool := mcp.NewTool("company_browse",
		mcp.WithDescription("Browse companies"),
		mcp.WithString("browseType", mcp.Description("Browse type")),
		mcp.WithString("param", mcp.Description("Parameter value")),
		mcp.WithString("type", mcp.Description("Company type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(companyBrowseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=company&f=browse&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse companies: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	companyEditTool := mcp.NewTool("company_edit",
		mcp.WithDescription("Edit company"),
	)

	s.AddTool(companyEditTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=company&f=edit&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit company: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	companyViewTool := mcp.NewTool("company_view",
		mcp.WithDescription("View company"),
	)

	s.AddTool(companyViewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=company&f=view&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view company: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	companyDynamicTool := mcp.NewTool("company_dynamic",
		mcp.WithDescription("Get company dynamic/activity"),
		mcp.WithString("browseType", mcp.Description("Browse type")),
		mcp.WithString("param", mcp.Description("Parameter value")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithString("date", mcp.Description("Date filter")),
		mcp.WithString("direction",
			mcp.Description("Direction: next|pre"),
			mcp.Enum("next", "pre"),
		),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("productID", mcp.Description("Product ID")),
		mcp.WithString("projectID", mcp.Description("Project ID")),
		mcp.WithString("executionID", mcp.Description("Execution ID")),
	)

	s.AddTool(companyDynamicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["date"].(string); ok {
			params["date"] = v
		}
		if v, ok := args["direction"].(string); ok {
			params["direction"] = v
		}
		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["productID"]; ok && v != nil {
			params["productID"] = v
		}
		if v, ok := args["projectID"]; ok && v != nil {
			params["projectID"] = v
		}
		if v, ok := args["executionID"]; ok && v != nil {
			params["executionID"] = v
		}

		resp, err := client.Get("/index.php?m=company&f=dynamic&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get company dynamic: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	companyAjaxGetOutsideTool := mcp.NewTool("company_ajax_get_outside",
		mcp.WithDescription("Get outside companies"),
	)

	s.AddTool(companyAjaxGetOutsideTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=company&f=ajaxGetOutsideCompany&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get outside companies: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDepartmentTools(s *server.MCPServer, client *client.ZenTaoClient) {
	deptBrowseTool := mcp.NewTool("dept_browse",
		mcp.WithDescription("Browse departments"),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
	)

	s.AddTool(deptBrowseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=dept&f=browse&t=json"

		if v, ok := args["deptID"]; ok && v != nil {
			url += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse departments: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deptUpdateOrderTool := mcp.NewTool("dept_update_order",
		mcp.WithDescription("Update department order"),
	)

	s.AddTool(deptUpdateOrderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=dept&f=updateOrder&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update department order: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deptManageChildTool := mcp.NewTool("dept_manage_child",
		mcp.WithDescription("Manage child departments"),
	)

	s.AddTool(deptManageChildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=dept&f=manageChild&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage child departments: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deptEditTool := mcp.NewTool("dept_edit",
		mcp.WithDescription("Edit department"),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
	)

	s.AddTool(deptEditTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=dept&f=edit&t=json"

		if v, ok := args["deptID"]; ok && v != nil {
			url += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit department: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deptDeleteTool := mcp.NewTool("dept_delete",
		mcp.WithDescription("Delete department"),
		mcp.WithNumber("deptID",
			mcp.Required(),
			mcp.Description("Department ID"),
		),
	)

	s.AddTool(deptDeleteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		deptID := int(args["deptID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=dept&f=delete&t=json&deptID=%d", deptID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete department: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deptAjaxGetUsersTool := mcp.NewTool("dept_ajax_get_users",
		mcp.WithDescription("Get users in department"),
		mcp.WithNumber("dept", mcp.Description("Department")),
		mcp.WithString("user", mcp.Description("User account")),
		mcp.WithString("key",
			mcp.Description("Key type: id|account"),
			mcp.Enum("id", "account"),
		),
	)

	s.AddTool(deptAjaxGetUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["dept"]; ok && v != nil {
			params["dept"] = int(v.(float64))
		}
		if v, ok := args["user"].(string); ok {
			params["user"] = v
		}
		if v, ok := args["key"].(string); ok {
			params["key"] = v
		}

		resp, err := client.Get("/index.php?m=dept&f=ajaxGetUsers&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get department users: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerGroupTools(s *server.MCPServer, client *client.ZenTaoClient) {
	groupBrowseTool := mcp.NewTool("group_browse",
		mcp.WithDescription("Browse groups"),
	)

	s.AddTool(groupBrowseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=group&f=browse&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse groups: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupCreateTool := mcp.NewTool("group_create",
		mcp.WithDescription("Create a new group"),
	)

	s.AddTool(groupCreateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=group&f=create&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create group: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupEditTool := mcp.NewTool("group_edit",
		mcp.WithDescription("Edit a group"),
		mcp.WithNumber("groupID", mcp.Description("Group ID")),
	)

	s.AddTool(groupEditTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=group&f=edit&t=json"

		if v, ok := args["groupID"]; ok && v != nil {
			url += fmt.Sprintf("&groupID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit group: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupCopyTool := mcp.NewTool("group_copy",
		mcp.WithDescription("Copy a group"),
		mcp.WithNumber("groupID", mcp.Description("Group ID")),
	)

	s.AddTool(groupCopyTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=group&f=copy&t=json"

		if v, ok := args["groupID"]; ok && v != nil {
			url += fmt.Sprintf("&groupID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to copy group: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupManageViewTool := mcp.NewTool("group_manage_view",
		mcp.WithDescription("Manage group view"),
		mcp.WithNumber("groupID", mcp.Description("Group ID")),
	)

	s.AddTool(groupManageViewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=group&f=manageView&t=json"

		if v, ok := args["groupID"]; ok && v != nil {
			url += fmt.Sprintf("&groupID=%d", int(v.(float64)))
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage group view: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupManagePrivTool := mcp.NewTool("group_manage_priv",
		mcp.WithDescription("Manage group privileges"),
		mcp.WithString("type",
			mcp.Description("Type: byPackage|byGroup|byModule"),
			mcp.Enum("byPackage", "byGroup", "byModule"),
		),
		mcp.WithNumber("param", mcp.Description("Parameter")),
		mcp.WithString("nav", mcp.Description("Navigation")),
		mcp.WithString("version", mcp.Description("Version")),
	)

	s.AddTool(groupManagePrivTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["nav"].(string); ok {
			params["nav"] = v
		}
		if v, ok := args["version"].(string); ok {
			params["version"] = v
		}

		resp, err := client.Post("/index.php?m=group&f=managePriv&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage group privileges: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupManageMemberTool := mcp.NewTool("group_manage_member",
		mcp.WithDescription("Manage group members"),
		mcp.WithNumber("groupID", mcp.Description("Group ID")),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
	)

	s.AddTool(groupManageMemberTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=group&f=manageMember&t=json"

		if v, ok := args["groupID"]; ok && v != nil {
			url += fmt.Sprintf("&groupID=%d", int(v.(float64)))
		}
		if v, ok := args["deptID"]; ok && v != nil {
			url += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage group members: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupManageProjectAdminTool := mcp.NewTool("group_manage_project_admin",
		mcp.WithDescription("Manage group project admins"),
		mcp.WithNumber("groupID", mcp.Description("Group ID")),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
	)

	s.AddTool(groupManageProjectAdminTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=group&f=manageProjectAdmin&t=json"

		if v, ok := args["groupID"]; ok && v != nil {
			url += fmt.Sprintf("&groupID=%d", int(v.(float64)))
		}
		if v, ok := args["deptID"]; ok && v != nil {
			url += fmt.Sprintf("&deptID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage project admins: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupDeleteTool := mcp.NewTool("group_delete",
		mcp.WithDescription("Delete a group"),
		mcp.WithNumber("groupID",
			mcp.Required(),
			mcp.Description("Group ID"),
		),
	)

	s.AddTool(groupDeleteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		groupID := int(args["groupID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=group&f=delete&t=json&groupID=%d", groupID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete group: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupAjaxGetPrivByParentsTool := mcp.NewTool("group_ajax_get_priv_by_parents",
		mcp.WithDescription("Get group privileges by parent"),
		mcp.WithString("selectedSubset", mcp.Description("Selected subset")),
		mcp.WithString("selectedPackages", mcp.Description("Selected packages")),
	)

	s.AddTool(groupAjaxGetPrivByParentsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["selectedSubset"].(string); ok {
			params["selectedSubset"] = v
		}
		if v, ok := args["selectedPackages"].(string); ok {
			params["selectedPackages"] = v
		}

		resp, err := client.Get("/index.php?m=group&f=ajaxGetPrivByParents&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get privileges by parent: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	groupAjaxGetRelatedPrivsTool := mcp.NewTool("group_ajax_get_related_privs",
		mcp.WithDescription("Get related group privileges"),
	)

	s.AddTool(groupAjaxGetRelatedPrivsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=group&f=ajaxGetRelatedPrivs&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get related privileges: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerUserManagementTools(s *server.MCPServer, client *client.ZenTaoClient) {
	userViewTool := mcp.NewTool("admin_user_view",
		mcp.WithDescription("View user details"),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID"),
		),
	)

	s.AddTool(userViewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		userID := int(args["userID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=user&f=view&t=json&userID=%d", userID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userTodoTool := mcp.NewTool("admin_user_todo",
		mcp.WithDescription("Get user todos"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type",
			mcp.Description("Todo type: all|before|future|thisWeek|thisMonth|thisYear|assignedToOther|cycle"),
			mcp.Enum("all", "before", "future", "thisWeek", "thisMonth", "thisYear", "assignedToOther", "cycle"),
		),
		mcp.WithString("status", mcp.Description("Todo status")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userTodoTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["status"].(string); ok {
			params["status"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=todo&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user todos: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userStoryTool := mcp.NewTool("admin_user_story",
		mcp.WithDescription("Get user stories"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("storyType", mcp.Description("Story type")),
		mcp.WithString("type", mcp.Description("Type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["storyType"].(string); ok {
			params["storyType"] = v
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=story&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user stories: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userTaskTool := mcp.NewTool("admin_user_task",
		mcp.WithDescription("Get user tasks"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type", mcp.Description("Task type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userTaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=task&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user tasks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userBugTool := mcp.NewTool("admin_user_bug",
		mcp.WithDescription("Get user bugs"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type", mcp.Description("Bug type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=bug&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user bugs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userTesttaskTool := mcp.NewTool("admin_user_testtask",
		mcp.WithDescription("Get user test tasks"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userTesttaskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=testtask&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user test tasks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userTestcaseTool := mcp.NewTool("admin_user_testcase",
		mcp.WithDescription("Get user test cases"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type", mcp.Description("Test case type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userTestcaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=testcase&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user test cases: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userExecutionTool := mcp.NewTool("admin_user_execution",
		mcp.WithDescription("Get user executions"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=execution&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user executions: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userIssueTool := mcp.NewTool("admin_user_issue",
		mcp.WithDescription("Get user issues"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type", mcp.Description("Issue type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userIssueTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=issue&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user issues: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userRiskTool := mcp.NewTool("admin_user_risk",
		mcp.WithDescription("Get user risks"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("type", mcp.Description("Risk type")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(userRiskTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=user&f=risk&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user risks: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userProfileTool := mcp.NewTool("admin_user_profile",
		mcp.WithDescription("Get user profile"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
	)

	s.AddTool(userProfileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=profile&t=json"

		if v, ok := args["userID"]; ok && v != nil {
			url += fmt.Sprintf("&userID=%d", int(v.(float64)))
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user profile: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userCreateTool := mcp.NewTool("admin_user_create",
		mcp.WithDescription("Create a new user"),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
		mcp.WithString("type", mcp.Description("User type")),
	)

	s.AddTool(userCreateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["deptID"]; ok && v != nil {
			params["deptID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=create&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userBatchCreateTool := mcp.NewTool("admin_user_batch_create",
		mcp.WithDescription("Batch create users"),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
		mcp.WithString("type", mcp.Description("User type")),
	)

	s.AddTool(userBatchCreateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["deptID"]; ok && v != nil {
			params["deptID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=batchCreate&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create users: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userEditTool := mcp.NewTool("admin_user_edit",
		mcp.WithDescription("Edit user"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
	)

	s.AddTool(userEditTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=edit&t=json"

		if v, ok := args["userID"]; ok && v != nil {
			url += fmt.Sprintf("&userID=%d", int(v.(float64)))
		}

		resp, err := client.Post(url, nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userBatchEditTool := mcp.NewTool("admin_user_batch_edit",
		mcp.WithDescription("Batch edit users"),
		mcp.WithNumber("deptID", mcp.Description("Department ID")),
		mcp.WithString("type", mcp.Description("User type")),
	)

	s.AddTool(userBatchEditTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["deptID"]; ok && v != nil {
			params["deptID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=batchEdit&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch edit users: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userDeleteTool := mcp.NewTool("admin_user_delete",
		mcp.WithDescription("Delete user"),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID"),
		),
	)

	s.AddTool(userDeleteTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		userID := int(args["userID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=user&f=delete&t=json&userID=%d", userID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userUnlockTool := mcp.NewTool("admin_user_unlock",
		mcp.WithDescription("Unlock user"),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID"),
		),
	)

	s.AddTool(userUnlockTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		userID := int(args["userID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=user&f=unlock&t=json&userID=%d", userID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlock user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userUnbindTool := mcp.NewTool("admin_user_unbind",
		mcp.WithDescription("Unbind user"),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID"),
		),
	)

	s.AddTool(userUnbindTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		userID := int(args["userID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=user&f=unbind&t=json&userID=%d", userID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unbind user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userLoginTool := mcp.NewTool("admin_user_login",
		mcp.WithDescription("User login"),
		mcp.WithString("referer", mcp.Description("Referer URL")),
	)

	s.AddTool(userLoginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["referer"].(string); ok {
			params["referer"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=login&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to login: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userDenyTool := mcp.NewTool("admin_user_deny",
		mcp.WithDescription("Deny user access"),
		mcp.WithString("module", mcp.Description("Module name")),
		mcp.WithString("method", mcp.Description("Method name")),
		mcp.WithString("referer", mcp.Description("Referer URL")),
	)

	s.AddTool(userDenyTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["module"].(string); ok {
			params["module"] = v
		}
		if v, ok := args["method"].(string); ok {
			params["method"] = v
		}
		if v, ok := args["referer"].(string); ok {
			params["referer"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=deny&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to deny user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userLogoutTool := mcp.NewTool("admin_user_logout",
		mcp.WithDescription("User logout"),
		mcp.WithString("referer", mcp.Description("Referer URL")),
	)

	s.AddTool(userLogoutTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["referer"].(string); ok {
			params["referer"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=logout&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to logout: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userResetTool := mcp.NewTool("admin_user_reset",
		mcp.WithDescription("Reset user"),
	)

	s.AddTool(userResetTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=user&f=reset&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reset user: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userForgetPasswordTool := mcp.NewTool("admin_user_forget_password",
		mcp.WithDescription("Forget user password"),
	)

	s.AddTool(userForgetPasswordTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Post("/index.php?m=user&f=forgetPassword&t=json", nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to forget password: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userResetPasswordTool := mcp.NewTool("admin_user_reset_password",
		mcp.WithDescription("Reset user password"),
		mcp.WithString("code", mcp.Description("Reset code")),
	)

	s.AddTool(userResetPasswordTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["code"].(string); ok {
			params["code"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=resetPassword&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reset password: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userDynamicTool := mcp.NewTool("admin_user_dynamic",
		mcp.WithDescription("Get user dynamic"),
		mcp.WithNumber("userID", mcp.Description("User ID")),
		mcp.WithString("period", mcp.Description("Time period")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("date", mcp.Description("Date")),
		mcp.WithString("direction",
			mcp.Description("Direction: next|pre"),
			mcp.Enum("next", "pre"),
		),
	)

	s.AddTool(userDynamicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["userID"]; ok && v != nil {
			params["userID"] = int(v.(float64))
		}
		if v, ok := args["period"].(string); ok {
			params["period"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["date"]; ok && v != nil {
			params["date"] = int(v.(float64))
		}
		if v, ok := args["direction"].(string); ok {
			params["direction"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=dynamic&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user dynamic: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userCropAvatarTool := mcp.NewTool("admin_user_crop_avatar",
		mcp.WithDescription("Crop user avatar"),
		mcp.WithNumber("imageID",
			mcp.Required(),
			mcp.Description("Image ID"),
		),
	)

	s.AddTool(userCropAvatarTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		imageID := int(args["imageID"].(float64))

		resp, err := client.Post(fmt.Sprintf("/index.php?m=user&f=cropAvatar&t=json&imageID=%d", imageID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to crop avatar: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetOldContactUsersTool := mcp.NewTool("admin_user_ajax_get_old_contact_users",
		mcp.WithDescription("Get old contact users"),
		mcp.WithNumber("contactListID", mcp.Description("Contact list ID")),
		mcp.WithString("dropdownName",
			mcp.Description("Dropdown name: mailto|whitelist"),
			mcp.Enum("mailto", "whitelist"),
		),
	)

	s.AddTool(userAjaxGetOldContactUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["contactListID"]; ok && v != nil {
			params["contactListID"] = int(v.(float64))
		}
		if v, ok := args["dropdownName"].(string); ok {
			params["dropdownName"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=ajaxGetOldContactUsers&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get old contact users: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetContactUsersTool := mcp.NewTool("admin_user_ajax_get_contact_users",
		mcp.WithDescription("Get contact users"),
		mcp.WithNumber("contactListID", mcp.Description("Contact list ID")),
	)

	s.AddTool(userAjaxGetContactUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=ajaxGetContactUsers&t=json"

		if v, ok := args["contactListID"]; ok && v != nil {
			url += fmt.Sprintf("&contactListID=%d", int(v.(float64)))
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get contact users: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetContactListTool := mcp.NewTool("admin_user_ajax_get_contact_list",
		mcp.WithDescription("Get contact list"),
	)

	s.AddTool(userAjaxGetContactListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=user&f=ajaxGetContactList&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get contact list: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetOldContactListTool := mcp.NewTool("admin_user_ajax_get_old_contact_list",
		mcp.WithDescription("Get old contact list"),
		mcp.WithString("dropdownName", mcp.Description("Dropdown name")),
	)

	s.AddTool(userAjaxGetOldContactListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=ajaxGetOldContactList&t=json"

		if v, ok := args["dropdownName"].(string); ok {
			url += fmt.Sprintf("&dropdownName=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get old contact list: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetItemsTool := mcp.NewTool("admin_user_ajax_get_items",
		mcp.WithDescription("Get user items"),
		mcp.WithString("params", mcp.Description("Parameters")),
	)

	s.AddTool(userAjaxGetItemsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["params"]; ok && v != nil {
			params["params"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=ajaxGetItems&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get items: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetTemplatesTool := mcp.NewTool("admin_user_ajax_get_templates",
		mcp.WithDescription("Get user templates"),
		mcp.WithString("editor", mcp.Description("Editor type")),
		mcp.WithString("type", mcp.Description("Template type")),
	)

	s.AddTool(userAjaxGetTemplatesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["editor"].(string); ok {
			params["editor"] = v
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Get("/index.php?m=user&f=ajaxGetTemplates&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get templates: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxSaveTemplateTool := mcp.NewTool("admin_user_ajax_save_template",
		mcp.WithDescription("Save user template"),
		mcp.WithString("editor", mcp.Description("Editor type")),
		mcp.WithString("type", mcp.Description("Template type")),
	)

	s.AddTool(userAjaxSaveTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["editor"].(string); ok {
			params["editor"] = v
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=ajaxSaveTemplate&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxDeleteTemplateTool := mcp.NewTool("admin_user_ajax_delete_template",
		mcp.WithDescription("Delete user template"),
		mcp.WithNumber("templateID",
			mcp.Required(),
			mcp.Description("Template ID"),
		),
	)

	s.AddTool(userAjaxDeleteTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		templateID := int(args["templateID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=user&f=ajaxDeleteTemplate&t=json&templateID=%d", templateID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetMoreTool := mcp.NewTool("admin_user_ajax_get_more",
		mcp.WithDescription("Get more user data"),
	)

	s.AddTool(userAjaxGetMoreTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=user&f=ajaxGetMore&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get more data: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxGetGroupsTool := mcp.NewTool("admin_user_ajax_get_groups",
		mcp.WithDescription("Get user groups"),
		mcp.WithString("visions",
			mcp.Description("Interface visions: rnd|lite|rnd,lite"),
			mcp.Enum("rnd", "lite", "rnd,lite"),
		),
	)

	s.AddTool(userAjaxGetGroupsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=ajaxGetGroups&t=json"

		if v, ok := args["visions"].(string); ok {
			url += fmt.Sprintf("&visions=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get groups: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userRefreshRandomTool := mcp.NewTool("admin_user_refresh_random",
		mcp.WithDescription("Refresh random user data"),
	)

	s.AddTool(userRefreshRandomTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=user&f=refreshRandom&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to refresh random: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxPrintTemplatesTool := mcp.NewTool("admin_user_ajax_print_templates",
		mcp.WithDescription("Get print templates"),
		mcp.WithString("type", mcp.Description("Template type")),
		mcp.WithString("link", mcp.Description("Link")),
	)

	s.AddTool(userAjaxPrintTemplatesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=user&f=ajaxPrintTemplates&t=json"

		if v, ok := args["type"].(string); ok {
			url += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["link"].(string); ok {
			url += fmt.Sprintf("&link=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get print templates: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	userAjaxSaveOldTemplateTool := mcp.NewTool("admin_user_ajax_save_old_template",
		mcp.WithDescription("Save old template"),
		mcp.WithString("type", mcp.Description("Template type")),
	)

	s.AddTool(userAjaxSaveOldTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=user&f=ajaxSaveOldTemplate&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to save old template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
