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

func RegisterStakeholderTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseStakeholdersTool := mcp.NewTool("browse_stakeholders",
		mcp.WithDescription("Browse stakeholders for a project"),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
			mcp.Enum("all", "inside", "outside", "key"),
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

	s.AddTool(browseStakeholdersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=browse&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse stakeholders: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createStakeholderTool := mcp.NewTool("create_stakeholder",
		mcp.WithDescription("Create a new stakeholder"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID (program/project)"),
		),
		mcp.WithNumber("user",
			mcp.Required(),
			mcp.Description("User ID"),
		),
		mcp.WithString("role",
			mcp.Description("Stakeholder role"),
		),
		mcp.WithString("type",
			mcp.Description("Stakeholder type"),
			mcp.Enum("inside", "outside"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
		mcp.WithString("name",
			mcp.Description("Stakeholder name"),
		),
		mcp.WithString("company",
			mcp.Description("Company"),
		),
		mcp.WithString("phone",
			mcp.Description("Phone"),
		),
		mcp.WithString("email",
			mcp.Description("Email"),
		),
		mcp.WithString("qq",
			mcp.Description("QQ"),
		),
		mcp.WithString("weixin",
			mcp.Description("WeChat"),
		),
	)

	s.AddTool(createStakeholderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"user": int(args["user"].(float64)),
		}

		if v, ok := args["role"]; ok && v != nil {
			body["role"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["from"]; ok && v != nil {
			body["from"] = v
		}
		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["company"]; ok && v != nil {
			body["company"] = v
		}
		if v, ok := args["phone"]; ok && v != nil {
			body["phone"] = v
		}
		if v, ok := args["email"]; ok && v != nil {
			body["email"] = v
		}
		if v, ok := args["qq"]; ok && v != nil {
			body["qq"] = v
		}
		if v, ok := args["weixin"]; ok && v != nil {
			body["weixin"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=stakeholder&f=create&t=json&objectID=%d", int(args["objectID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create stakeholder: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchCreateStakeholdersTool := mcp.NewTool("batch_create_stakeholders",
		mcp.WithDescription("Create multiple stakeholders at once"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("dept",
			mcp.Description("Department"),
		),
		mcp.WithNumber("parentID",
			mcp.Description("Parent stakeholder ID"),
		),
		mcp.WithArray("users",
			mcp.Required(),
			mcp.Description("User IDs to add as stakeholders"),
		),
	)

	s.AddTool(batchCreateStakeholdersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"users": args["users"],
		}

		queryParams := fmt.Sprintf("&projectID=%d", int(args["projectID"].(float64)))
		if v, ok := args["dept"]; ok && v != nil {
			queryParams += fmt.Sprintf("&dept=%s", v)
		}
		if v, ok := args["parentID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&parentID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=stakeholder&f=batchCreate&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch create stakeholders: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editStakeholderTool := mcp.NewTool("edit_stakeholder",
		mcp.WithDescription("Edit an existing stakeholder"),
		mcp.WithNumber("stakeholderID",
			mcp.Required(),
			mcp.Description("Stakeholder ID"),
		),
		mcp.WithString("role",
			mcp.Description("Stakeholder role"),
		),
		mcp.WithString("type",
			mcp.Description("Stakeholder type"),
			mcp.Enum("inside", "outside"),
		),
		mcp.WithString("name",
			mcp.Description("Stakeholder name"),
		),
		mcp.WithString("company",
			mcp.Description("Company"),
		),
		mcp.WithString("phone",
			mcp.Description("Phone"),
		),
		mcp.WithString("email",
			mcp.Description("Email"),
		),
		mcp.WithString("qq",
			mcp.Description("QQ"),
		),
		mcp.WithString("weixin",
			mcp.Description("WeChat"),
		),
		mcp.WithString("key",
			mcp.Description("Key stakeholder flag"),
		),
	)

	s.AddTool(editStakeholderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["role"]; ok && v != nil {
			body["role"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["company"]; ok && v != nil {
			body["company"] = v
		}
		if v, ok := args["phone"]; ok && v != nil {
			body["phone"] = v
		}
		if v, ok := args["email"]; ok && v != nil {
			body["email"] = v
		}
		if v, ok := args["qq"]; ok && v != nil {
			body["qq"] = v
		}
		if v, ok := args["weixin"]; ok && v != nil {
			body["weixin"] = v
		}
		if v, ok := args["key"]; ok && v != nil {
			body["key"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=stakeholder&f=edit&t=json&stakeholderID=%d", int(args["stakeholderID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit stakeholder: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getStakeholderMembersTool := mcp.NewTool("get_stakeholder_members",
		mcp.WithDescription("Get stakeholder members for program/project"),
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getStakeholderMembersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["programID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&programID=%d", int(v.(float64)))
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=ajaxGetMembers&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get stakeholder members: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getCompanyUsersTool := mcp.NewTool("get_company_users",
		mcp.WithDescription("Get company users for stakeholder management"),
		mcp.WithNumber("programID",
			mcp.Description("Program ID"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getCompanyUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["programID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&programID=%d", int(v.(float64)))
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=ajaxGetCompanyUser&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get company users: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getOutsideUsersTool := mcp.NewTool("get_outside_users",
		mcp.WithDescription("Get outside users for stakeholder management"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
	)

	s.AddTool(getOutsideUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		objectID := int(args["objectID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=ajaxGetOutsideUser&t=json&objectID=%d", objectID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get outside users: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteStakeholderTool := mcp.NewTool("delete_stakeholder",
		mcp.WithDescription("Delete a stakeholder"),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID to remove as stakeholder"),
		),
	)

	s.AddTool(deleteStakeholderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		userID := int(args["userID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=delete&t=json&userID=%d", userID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete stakeholder: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewStakeholderTool := mcp.NewTool("view_stakeholder",
		mcp.WithDescription("View stakeholder details"),
		mcp.WithNumber("stakeholderID",
			mcp.Required(),
			mcp.Description("Stakeholder ID"),
		),
	)

	s.AddTool(viewStakeholderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		stakeholderID := int(args["stakeholderID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=view&t=json&stakeholderID=%d", stakeholderID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view stakeholder: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	communicateStakeholderTool := mcp.NewTool("communicate_stakeholder",
		mcp.WithDescription("Record communication with stakeholder"),
		mcp.WithNumber("stakeholderID",
			mcp.Required(),
			mcp.Description("Stakeholder ID"),
		),
		mcp.WithString("mode",
			mcp.Description("Communication mode"),
		),
		mcp.WithString("content",
			mcp.Description("Communication content"),
		),
		mcp.WithString("date",
			mcp.Description("Communication date"),
		),
		mcp.WithString("contactedBy",
			mcp.Description("Contacted by"),
		),
		mcp.WithString("feedback",
			mcp.Description("Stakeholder feedback"),
		),
	)

	s.AddTool(communicateStakeholderTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["mode"]; ok && v != nil {
			body["mode"] = v
		}
		if v, ok := args["content"]; ok && v != nil {
			body["content"] = v
		}
		if v, ok := args["date"]; ok && v != nil {
			body["date"] = v
		}
		if v, ok := args["contactedBy"]; ok && v != nil {
			body["contactedBy"] = v
		}
		if v, ok := args["feedback"]; ok && v != nil {
			body["feedback"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=stakeholder&f=communicate&t=json&stakeholderID=%d", int(args["stakeholderID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to communicate with stakeholder: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	stakeholderExpectTool := mcp.NewTool("stakeholder_expect",
		mcp.WithDescription("Record stakeholder expectations"),
		mcp.WithNumber("stakeholderID",
			mcp.Required(),
			mcp.Description("Stakeholder ID"),
		),
		mcp.WithString("expect",
			mcp.Description("Stakeholder expectations"),
		),
		mcp.WithString("keyNote",
			mcp.Description("Key notes"),
		),
	)

	s.AddTool(stakeholderExpectTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["expect"]; ok && v != nil {
			body["expect"] = v
		}
		if v, ok := args["keyNote"]; ok && v != nil {
			body["keyNote"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=stakeholder&f=expect&t=json&stakeholderID=%d", int(args["stakeholderID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to record stakeholder expectations: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getStakeholderIssuesTool := mcp.NewTool("get_stakeholder_issues",
		mcp.WithDescription("Get issues related to stakeholder"),
		mcp.WithNumber("stakeholderID",
			mcp.Required(),
			mcp.Description("Stakeholder ID"),
		),
	)

	s.AddTool(getStakeholderIssuesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		stakeholderID := int(args["stakeholderID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=stakeholder&f=userIssue&t=json&stakeholderID=%d", stakeholderID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get stakeholder issues: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
