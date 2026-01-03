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
	"github.com/zentao/mcp-server/logger"
)

func RegisterAuthTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Handle nil client for testing
	if client == nil {
		logger.Warn("tools", "Nil client provided to RegisterAuthTools", nil)
		return
	}

	// Check auth method to determine which tools to register
	authMethod := client.GetAuthMethod()
	toolCount := 1

	switch authMethod {
	case 1: // AuthApp
		// App-based authentication
		logger.Debug("tools", "Registering app-based auth tools", map[string]interface{}{
			"tool_count": toolCount,
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

	case 2: // AuthSession
		// Session-based authentication
		toolCount = 1
		logger.Debug("tools", "Registering session-based auth tools", map[string]interface{}{
			"tool_count": toolCount,
		})

		// Session login tool
		sessionLoginTool := mcp.NewTool("zentao_login_session",
			mcp.WithDescription("Login to ZenTao with username/password using session authentication"),
			mcp.WithString("account",
				mcp.Required(),
				mcp.Description("ZenTao username/account"),
			),
			mcp.WithString("password",
				mcp.Required(),
				mcp.Description("ZenTao password"),
			),
		)

		logger.Debug("tools", "Created zentao_login_session tool", map[string]interface{}{
			"tool_name": "zentao_login_session",
			"description": "Login to ZenTao with username/password using session authentication",
			"required_params": []string{"account", "password"},
		})

		s.AddTool(sessionLoginTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			args := request.GetArguments()
			account := args["account"].(string)
			password := args["password"].(string)

			logger.LogMCPToolCall("zentao_login_session", map[string]interface{}{
				"account": account,
				"has_password": password != "",
			})

			logger.Info("auth", "Starting session-based login", map[string]interface{}{
				"account": account,
			})

			// Step 1: Get session ID
			logger.Debug("auth", "Getting session ID", nil)
			if err := client.GetSessionID(); err != nil {
				logger.Error("auth", "Failed to get session ID", err, nil)
				return mcp.NewToolResultError(fmt.Sprintf("Failed to get session ID: %v", err)), nil
			}

			// Step 2: Login with credentials
			logger.Debug("auth", "Performing user login", map[string]interface{}{
				"account": account,
			})
			if err := client.Login(account, password); err != nil {
				logger.Error("auth", "Login failed", err, map[string]interface{}{
					"account": account,
				})
				return mcp.NewToolResultError(fmt.Sprintf("Login failed: %v", err)), nil
			}

			logger.Info("auth", "Session login successful", map[string]interface{}{
				"account": account,
				"is_authenticated": client.IsAuthenticated(),
			})

			return mcp.NewToolResultText("Session authentication successful. You can now use other ZenTao tools."), nil
		})

	default:
		// No authentication or unknown method
		logger.Warn("tools", "No authentication tools registered", map[string]interface{}{
			"auth_method": authMethod,
		})
	}

	logger.Debug("tools", "Auth tools registration completed", map[string]interface{}{
		"auth_method": authMethod,
		"tools_registered": toolCount,
	})
}
