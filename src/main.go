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
	"github.com/zentao/mcp-server/resources"
	"github.com/zentao/mcp-server/tools"
)

var ztClient *client.ZenTaoClient

func main() {
	baseURL := os.Getenv("ZENTAO_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	ztClient = client.NewZenTaoClient(baseURL)

	s := server.NewMCPServer(
		"ZenTao MCP Server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithResourceCapabilities(true, true),
		server.WithPromptCapabilities(true),
		server.WithRecovery(),
	)

	registerTools(s)
	registerResources(s)
	registerPrompts(s)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "Server error: %v\n", err)
		os.Exit(1)
	}
}

func registerTools(s *server.MCPServer) {
	tools.RegisterAuthTools(s, ztClient)
	tools.RegisterProductTools(s, ztClient)
	tools.RegisterProjectTools(s, ztClient)
	tools.RegisterStoryTools(s, ztClient)
	tools.RegisterTaskTools(s, ztClient)
	tools.RegisterBugTools(s, ztClient)
	tools.RegisterTestCaseTools(s, ztClient)
	tools.RegisterPlanTools(s, ztClient)
	tools.RegisterBuildTools(s, ztClient)
	tools.RegisterUserTools(s, ztClient)
	tools.RegisterFeedbackTools(s, ztClient)
	tools.RegisterTicketTools(s, ztClient)
}

func registerResources(s *server.MCPServer) {
	resources.RegisterProductResources(s, ztClient)
	resources.RegisterProjectResources(s, ztClient)
	resources.RegisterStoryResources(s, ztClient)
	resources.RegisterTaskResources(s, ztClient)
	resources.RegisterBugResources(s, ztClient)
	resources.RegisterUserResources(s, ztClient)
}

func registerPrompts(s *server.MCPServer) {
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
		return mcp.NewGetPromptResult("Create Product", []mcp.PromptMessage{
			mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent("Create a product with the specified details")),
		}), nil
	})

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
		return mcp.NewGetPromptResult("Create Story", []mcp.PromptMessage{
			mcp.NewPromptMessage(mcp.RoleUser, mcp.NewTextContent("Create a user story with the specified details")),
		}), nil
	})
}
