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

func RegisterZaiTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getZaiSettingsTool := mcp.NewTool("get_zai_settings",
		mcp.WithDescription("Get ZenTao AI (ZAI) module settings"),
		mcp.WithString("mode",
			mcp.Description("Settings mode (basic, advanced, etc.)"),
		),
	)

	s.AddTool(getZaiSettingsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["mode"]; ok && v != nil {
			queryParams = fmt.Sprintf("&mode=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zai&f=setting&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get ZAI settings: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateZaiSettingsTool := mcp.NewTool("update_zai_settings",
		mcp.WithDescription("Update ZenTao AI (ZAI) module settings"),
		mcp.WithString("mode",
			mcp.Required(),
			mcp.Description("Settings mode"),
		),
		mcp.WithString("settings",
			mcp.Required(),
			mcp.Description("Settings configuration as JSON string"),
		),
	)

	s.AddTool(updateZaiSettingsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"mode":     args["mode"],
			"settings": args["settings"],
		}

		resp, err := client.Post("/index.php?m=zai&f=setting&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update ZAI settings: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getZaiTokenTool := mcp.NewTool("get_zai_token",
		mcp.WithDescription("Get authentication token for ZenTao AI services"),
	)

	s.AddTool(getZaiTokenTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=zai&f=ajaxGetToken&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get ZAI token: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getVectorizationStatusTool := mcp.NewTool("get_vectorization_status",
		mcp.WithDescription("Get current vectorization status in ZenTao AI"),
	)

	s.AddTool(getVectorizationStatusTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=zai&f=vectorized&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get vectorization status: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	enableVectorizationTool := mcp.NewTool("enable_vectorization",
		mcp.WithDescription("Enable vectorization for ZenTao AI"),
		mcp.WithString("config",
			mcp.Description("Vectorization configuration as JSON string"),
		),
	)

	s.AddTool(enableVectorizationTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["config"]; ok && v != nil {
			body["config"] = v
		}

		resp, err := client.Post("/index.php?m=zai&f=ajaxEnableVectorization&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to enable vectorization: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	syncVectorizationTool := mcp.NewTool("sync_vectorization",
		mcp.WithDescription("Synchronize vectorization data in ZenTao AI"),
		mcp.WithString("sync_options",
			mcp.Description("Sync options as JSON string (full, incremental, specific_modules)"),
		),
	)

	s.AddTool(syncVectorizationTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["sync_options"]; ok && v != nil {
			body["sync_options"] = v
		}

		resp, err := client.Post("/index.php?m=zai&f=ajaxSyncVectorization&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to sync vectorization: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getAiModelsTool := mcp.NewTool("get_ai_models",
		mcp.WithDescription("Get available AI models in ZenTao AI"),
	)

	s.AddTool(getAiModelsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=zai&f=models&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get AI models: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	testAiConnectionTool := mcp.NewTool("test_ai_connection",
		mcp.WithDescription("Test connection to AI services"),
		mcp.WithString("model",
			mcp.Description("Specific AI model to test"),
		),
	)

	s.AddTool(testAiConnectionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["model"]; ok && v != nil {
			queryParams = fmt.Sprintf("&model=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zai&f=testConnection&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to test AI connection: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
