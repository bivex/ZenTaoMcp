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
	"github.com/zentao/mcp-server/logger"
)

func RegisterAuthTools(s *server.MCPServer, client *client.ZenTaoClient) {
	logger.Debug("tools", "Registering auth tools", map[string]interface{}{
		"tool_count": 1,
	})

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

	logger.Debug("tools", "Created zentao_login tool", map[string]interface{}{
		"tool_name": "zentao_login",
		"description": "Login to ZenTao with app credentials (code + key)",
		"required_params": []string{"code", "key"},
	})

	s.AddTool(loginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		code := args["code"].(string)
		key := args["key"].(string)

		logger.LogMCPToolCall("zentao_login", map[string]interface{}{
			"has_code": code != "",
			"has_key":  key != "",
			"code_length": len(code),
			"key_length": len(key),
		})

		logger.Info("auth", "Setting ZenTao app credentials", map[string]interface{}{
			"code_provided": code != "",
			"key_provided": key != "",
		})

		client.SetAppCredentials(code, key)

		logger.Info("auth", "App credentials set successfully", map[string]interface{}{
			"client_has_code": client.Code != "",
			"client_has_key": client.Key != "",
		})

		result := mcp.NewToolResultText("Successfully set app credentials. The client will now use app-based authentication for all API calls.")

		logger.Debug("auth", "Auth tool completed", map[string]interface{}{
			"result_type": "text",
		})

		return result, nil
	})

	logger.Debug("tools", "Successfully registered zentao_login tool", nil)
}
