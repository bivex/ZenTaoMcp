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

func RegisterUserTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createUserTool := mcp.NewTool("create_user",
		mcp.WithDescription("Create a new user in ZenTao"),
		mcp.WithString("account",
			mcp.Required(),
			mcp.Description("User account name"),
		),
		mcp.WithString("password",
			mcp.Required(),
			mcp.Description("User password"),
		),
		mcp.WithString("realname",
			mcp.Description("Real name"),
		),
		mcp.WithArray("visions",
			mcp.Description("Interface types (rnd|lite)"),
		),
	)

	s.AddTool(createUserTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"account":  args["account"],
			"password": args["password"],
		}

		if v, ok := args["realname"]; ok && v != nil {
			body["realname"] = v
		}
		if v, ok := args["visions"]; ok && v != nil {
			body["visions"] = v
		}

		resp, err := client.Post("/users", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create user: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateUserTool := mcp.NewTool("update_user",
		mcp.WithDescription("Update an existing user in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("User ID"),
		),
		mcp.WithNumber("dept",
			mcp.Description("Department ID"),
		),
		mcp.WithString("role",
			mcp.Description("Role"),
		),
		mcp.WithString("mobile",
			mcp.Description("Mobile number"),
		),
		mcp.WithString("realname",
			mcp.Description("Real name"),
		),
		mcp.WithString("email",
			mcp.Description("Email"),
		),
		mcp.WithString("phone",
			mcp.Description("Phone number"),
		),
	)

	s.AddTool(updateUserTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["dept"]; ok && v != nil {
			body["dept"] = int(v.(float64))
		}
		if v, ok := args["role"]; ok && v != nil {
			body["role"] = v
		}
		if v, ok := args["mobile"]; ok && v != nil {
			body["mobile"] = v
		}
		if v, ok := args["realname"]; ok && v != nil {
			body["realname"] = v
		}
		if v, ok := args["email"]; ok && v != nil {
			body["email"] = v
		}
		if v, ok := args["phone"]; ok && v != nil {
			body["phone"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/users/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update user: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteUserTool := mcp.NewTool("delete_user",
		mcp.WithDescription("Delete a user from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("User ID to delete"),
		),
	)

	s.AddTool(deleteUserTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/users/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete user: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get users list tool
	getUsersTool := mcp.NewTool("get_users",
		mcp.WithDescription("Get list of users in ZenTao"),
		mcp.WithString("account",
			mcp.Description("Filter by account name"),
		),
		mcp.WithString("realname",
			mcp.Description("Filter by real name"),
		),
		mcp.WithString("email",
			mcp.Description("Filter by email"),
		),
		mcp.WithString("status",
			mcp.Description("Filter by user status"),
			mcp.Enum("active", "forbidden"),
		),
		mcp.WithNumber("dept",
			mcp.Description("Filter by department ID"),
		),
		mcp.WithString("role",
			mcp.Description("Filter by role"),
		),
		mcp.WithNumber("limit",
			mcp.Description("Maximum number of users to return (default: 100)"),
		),
		mcp.WithNumber("offset",
			mcp.Description("Offset for pagination (default: 0)"),
		),
	)

	s.AddTool(getUsersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := make(map[string]string)

		// Add optional filters
		if account, ok := args["account"].(string); ok && account != "" {
			params["account"] = account
		}
		if realname, ok := args["realname"].(string); ok && realname != "" {
			params["realname"] = realname
		}
		if email, ok := args["email"].(string); ok && email != "" {
			params["email"] = email
		}
		if status, ok := args["status"].(string); ok && status != "" {
			params["status"] = status
		}
		if dept, ok := args["dept"].(float64); ok && dept > 0 {
			params["dept"] = fmt.Sprintf("%.0f", dept)
		}
		if role, ok := args["role"].(string); ok && role != "" {
			params["role"] = role
		}
		if limit, ok := args["limit"].(float64); ok && limit > 0 {
			params["limit"] = fmt.Sprintf("%.0f", limit)
		}
		if offset, ok := args["offset"].(float64); ok && offset >= 0 {
			params["offset"] = fmt.Sprintf("%.0f", offset)
		}

		path := "/users"
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
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get users: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get current user profile tool
	getMyProfileTool := mcp.NewTool("get_my_profile",
		mcp.WithDescription("Get current user's profile information"),
	)

	s.AddTool(getMyProfileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/user")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user profile: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Get user details tool
	getUserTool := mcp.NewTool("get_user",
		mcp.WithDescription("Get details of a specific user by ID"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("User ID"),
		),
	)

	s.AddTool(getUserTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/user/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get user: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
