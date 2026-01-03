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

// RegisterAiappTools registers all AI app-related tools
func RegisterAiappTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// AI App view and browse
	registerAiappViewTools(s, client)
	// Mini program chat
	registerMiniProgramTools(s, client)
	// Square browse
	registerSquareTools(s, client)
	// Models
	registerModelTools(s, client)
	// Conversation
	registerConversationTools(s, client)
}

func registerAiappViewTools(s *server.MCPServer, client *client.ZenTaoClient) {
	viewTool := mcp.NewTool("aiapp_view",
		mcp.WithDescription("View AI app"),
		mcp.WithString("id", mcp.Description("App ID")),
	)

	s.AddTool(viewTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=aiapp&f=view&t=json"

		if v, ok := args["id"]; ok && v != nil {
			url += fmt.Sprintf("&id=%v", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view AI app: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerMiniProgramTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseMiniProgramTool := mcp.NewTool("aiapp_browse_mini_program",
		mcp.WithDescription("Browse AI mini programs"),
		mcp.WithString("id", mcp.Description("ID filter")),
	)

	s.AddTool(browseMiniProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=aiapp&f=browseMiniProgram&t=json"

		if v, ok := args["id"].(string); ok {
			url += fmt.Sprintf("&id=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse mini programs: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	miniProgramChatTool := mcp.NewTool("aiapp_mini_program_chat",
		mcp.WithDescription("Mini program chat"),
		mcp.WithString("id", mcp.Description("Chat ID")),
	)

	s.AddTool(miniProgramChatTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=aiapp&f=miniProgramChat&t=json"

		if v, ok := args["id"].(string); ok {
			url += fmt.Sprintf("&id=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start mini program chat: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	collectMiniProgramTool := mcp.NewTool("aiapp_collect_mini_program",
		mcp.WithDescription("Collect mini program"),
		mcp.WithString("appID", mcp.Description("App ID")),
		mcp.WithString("delete", mcp.Description("Delete flag")),
	)

	s.AddTool(collectMiniProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=aiapp&f=collectMiniProgram&t=json"

		if v, ok := args["appID"].(string); ok {
			url += fmt.Sprintf("&appID=%s", v)
		}
		if v, ok := args["delete"].(string); ok {
			url += fmt.Sprintf("&delete=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to collect mini program: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerSquareTools(s *server.MCPServer, client *client.ZenTaoClient) {
	squareTool := mcp.NewTool("aiapp_square",
		mcp.WithDescription("Browse AI app square"),
		mcp.WithString("category", mcp.Description("Category filter")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(squareTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["category"].(string); ok {
			params["category"] = v
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

		resp, err := client.Get("/index.php?m=aiapp&f=square&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse square: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerModelTools(s *server.MCPServer, client *client.ZenTaoClient) {
	modelsTool := mcp.NewTool("aiapp_models",
		mcp.WithDescription("Get AI models"),
	)

	s.AddTool(modelsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=aiapp&f=models&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get AI models: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerConversationTools(s *server.MCPServer, client *client.ZenTaoClient) {
	conversationTool := mcp.NewTool("aiapp_conversation",
		mcp.WithDescription("AI app conversation"),
		mcp.WithString("chat",
			mcp.Required(),
			mcp.Description("Chat ID - _ will be replaced with -, if set to NEW, open a new chat"),
		),
		mcp.WithString("params",
			mcp.Description("Parameters as JSON string, base64 encoded"),
		),
	)

	s.AddTool(conversationTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["chat"] = args["chat"].(string)

		if v, ok := args["params"].(string); ok {
			params["params"] = v
		}

		resp, err := client.Post("/index.php?m=aiapp&f=conversation&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start conversation: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
