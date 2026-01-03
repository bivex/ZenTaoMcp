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
}
