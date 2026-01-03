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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterAuthTools(s *server.MCPServer, client *client.ZenTaoClient) {
	loginTool := mcp.NewTool("zentao_login",
		mcp.WithDescription("Login to ZenTao with app credentials (code + key)"),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("ZenTao application code"),
		),
		mcp.WithString("key",
			mcp.Required(),
			mcp.Description("ZenTao application key"),
		),
	)

	s.AddTool(loginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		code := args["code"].(string)
		key := args["key"].(string)

		client.SetAppCredentials(code, key)

		return mcp.NewToolResultText("Successfully set app credentials. The client will now use app-based authentication for all API calls."), nil
	})
}
