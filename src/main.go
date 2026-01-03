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

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
	"github.com/zentao/mcp-server/logger"
	"github.com/zentao/mcp-server/resources"
	"github.com/zentao/mcp-server/tools"
)

var ztClient *client.ZenTaoClient

func main() {
	logger.Info("server", "ZenTao MCP Server starting up", map[string]interface{}{
		"version": "1.0.0",
	})

	// Read configuration
	baseURL := os.Getenv("ZENTAO_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
		logger.Warn("server", "ZENTAO_BASE_URL not set, using default", map[string]interface{}{
			"default_url": baseURL,
		})
	}

	// Determine authentication method
	authMethod := os.Getenv("ZENTAO_AUTH_METHOD")
	if authMethod == "" {
		authMethod = "app" // Default to app-based auth
	}

	code := os.Getenv("ZENTAO_APP_CODE")
	key := os.Getenv("ZENTAO_APP_KEY")

	logger.Info("server", "Configuration loaded", map[string]interface{}{
		"base_url":        baseURL,
		"auth_method":     authMethod,
		"has_app_code":    code != "",
		"has_app_key":     key != "",
		"log_level":       os.Getenv("ZENTAO_LOG_LEVEL"),
		"log_json":        os.Getenv("ZENTAO_LOG_JSON"),
	})

	// Initialize ZenTao client based on auth method
	logger.Debug("server", "Initializing ZenTao client", map[string]interface{}{
		"base_url":    baseURL,
		"auth_method": authMethod,
		"app_code":    code,
		"app_key":     key != "",
	})

	switch authMethod {
	case "app":
		ztClient = client.NewZenTaoClientWithApp(baseURL, code, key)
	case "session":
		ztClient = client.NewZenTaoClientWithSession(baseURL)
		logger.Info("server", "Session-based authentication enabled", map[string]interface{}{
			"note": "Use zentao_login_session tool for authentication",
		})
	default:
		logger.Warn("server", "Unknown auth method, defaulting to app-based", map[string]interface{}{
			"provided_method": authMethod,
			"default_method":  "app",
		})
		ztClient = client.NewZenTaoClientWithApp(baseURL, code, key)
	}

	// Initialize MCP server
	logger.Info("server", "Initializing MCP server", map[string]interface{}{
		"name":    "ZenTao MCP Server",
		"version": "1.0.0",
	})
	s := server.NewMCPServer(
		"ZenTao MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithRecovery(),
	)

	// Register components
	logger.Info("server", "Registering tools", nil)
	registerTools(s)

	logger.Info("server", "Registering resources", nil)
	registerResources(s)

	logger.Info("server", "Registering prompts", nil)
	registerPrompts(s)

	logger.Info("server", "Starting MCP server on stdio", nil)

	// Start server
	if err := server.ServeStdio(s); err != nil {
		logger.Error("server", "Server failed to start", err, map[string]interface{}{
			"transport": "stdio",
		})
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func registerTools(s *server.MCPServer) {
	logger.Debug("server", "Registering auth tools", nil)
	tools.RegisterAuthTools(s, ztClient)

	logger.Debug("server", "Registering product tools", nil)
	tools.RegisterProductTools(s, ztClient)

	logger.Debug("server", "Registering project tools", nil)
	tools.RegisterProjectTools(s, ztClient)

	logger.Debug("server", "Registering story tools", nil)
	tools.RegisterStoryTools(s, ztClient)

	logger.Debug("server", "Registering task tools", nil)
	tools.RegisterTaskTools(s, ztClient)

	logger.Debug("server", "Registering bug tools", nil)
	tools.RegisterBugTools(s, ztClient)

	logger.Debug("server", "Registering test case tools", nil)
	tools.RegisterTestCaseTools(s, ztClient)

	logger.Debug("server", "Registering plan tools", nil)
	tools.RegisterPlanTools(s, ztClient)

	logger.Debug("server", "Registering build tools", nil)
	tools.RegisterBuildTools(s, ztClient)

	logger.Debug("server", "Registering user tools", nil)
	tools.RegisterUserTools(s, ztClient)

	logger.Debug("server", "Registering feedback tools", nil)
	tools.RegisterFeedbackTools(s, ztClient)

	logger.Debug("server", "Registering ticket tools", nil)
	tools.RegisterTicketTools(s, ztClient)

	logger.Debug("server", "Registering program tools", nil)
	tools.RegisterProgramTools(s, ztClient)

	logger.Debug("server", "Registering test task tools", nil)
	tools.RegisterTestTaskTools(s, ztClient)

	logger.Debug("server", "Registering release tools", nil)
	tools.RegisterReleaseTools(s, ztClient)

	logger.Info("server", "All tool registrations completed", map[string]interface{}{
		"total_tools": 16,
	})
}

func registerResources(s *server.MCPServer) {
	logger.Debug("server", "Registering product resources", nil)
	resources.RegisterProductResources(s, ztClient)

	logger.Debug("server", "Registering project resources", nil)
	resources.RegisterProjectResources(s, ztClient)

	logger.Debug("server", "Registering program resources", nil)
	resources.RegisterProgramResources(s, ztClient)

	logger.Debug("server", "Registering story resources", nil)
	resources.RegisterStoryResources(s, ztClient)

	logger.Debug("server", "Registering task resources", nil)
	resources.RegisterTaskResources(s, ztClient)

	logger.Debug("server", "Registering bug resources", nil)
	resources.RegisterBugResources(s, ztClient)

	logger.Debug("server", "Registering user resources", nil)
	resources.RegisterUserResources(s, ztClient)

	logger.Debug("server", "Registering test task resources", nil)
	resources.RegisterTestTaskResources(s, ztClient)

	logger.Debug("server", "Registering build resources", nil)
	resources.RegisterBuildResources(s, ztClient)

	logger.Debug("server", "Registering plan resources", nil)
	resources.RegisterPlanResources(s, ztClient)

	logger.Debug("server", "Registering build resources", nil)
	resources.RegisterBuildResources(s, ztClient)

	logger.Debug("server", "Registering plan resources", nil)
	resources.RegisterPlanResources(s, ztClient)

	logger.Debug("server", "Registering release resources", nil)
	resources.RegisterReleaseResources(s, ztClient)

	logger.Info("server", "All resource registrations completed", map[string]interface{}{
		"total_resources": 10,
		"templates_registered": 21,
		"note": "List resources + resource templates for individual/scoped access",
	})
}

func registerPrompts(s *server.MCPServer) {
	logger.Debug("server", "Registering create_product prompt", nil)
	s.AddPrompt(mcp.NewPrompt("create_product",
		mcp.WithPromptDescription("Create a new product in ZenTao"),
		mcp.WithArgument("name",
			mcp.ArgumentDescription("Product name"),
			mcp.RequiredArgument(),
		),
		mcp.WithArgument("code",
			mcp.ArgumentDescription("Product code"),
			mcp.RequiredArgument(),
		),
	), func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		logger.LogMCPResourceRead("prompt://create_product", map[string]interface{}{
			"arguments": request.Params.Arguments,
		})

		result := mcp.NewGetPromptResult("Create Product", []mcp.PromptMessage{
			mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent("Create a product with the specified details")),
		})

		logger.Debug("prompt", "Generated create_product prompt", map[string]interface{}{
			"message_count": len(result.Messages),
		})

		return result, nil
	})

	logger.Debug("server", "Successfully registered create_product prompt", nil)

	logger.Debug("server", "Registering create_story prompt", nil)
	s.AddPrompt(mcp.NewPrompt("create_story",
		mcp.WithPromptDescription("Create a new user story"),
		mcp.WithArgument("title",
			mcp.ArgumentDescription("Story title"),
			mcp.RequiredArgument(),
		),
		mcp.WithArgument("product",
			mcp.ArgumentDescription("Product ID"),
			mcp.RequiredArgument(),
		),
	), func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		logger.LogMCPResourceRead("prompt://create_story", map[string]interface{}{
			"arguments": request.Params.Arguments,
		})

		result := mcp.NewGetPromptResult("Create Story", []mcp.PromptMessage{
			mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent("Create a user story with the specified details")),
		})

		logger.Debug("prompt", "Generated create_story prompt", map[string]interface{}{
			"message_count": len(result.Messages),
		})

		return result, nil
	})

	logger.Debug("server", "Successfully registered create_story prompt", nil)

	logger.Info("server", "All prompt registrations completed", map[string]interface{}{
		"total_prompts": 2,
	})
}
