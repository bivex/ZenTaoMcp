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

// RegisterBiTools registers all BI (Business Intelligence) tools
func RegisterBiTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Parquet file management
	registerParquetTools(s, client)
	// DuckDB management
	registerDuckdbTools(s, client)
	// Scope and field management
	registerScopeFieldTools(s, client)
}

func registerParquetTools(s *server.MCPServer, client *client.ZenTaoClient) {
	syncParquetFileTool := mcp.NewTool("bi_sync_parquet_file",
		mcp.WithDescription("Sync Parquet file"),
	)

	s.AddTool(syncParquetFileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=bi&f=syncParquetFile&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to sync Parquet file: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	initParquetTool := mcp.NewTool("bi_init_parquet",
		mcp.WithDescription("Initialize Parquet"),
	)

	s.AddTool(initParquetTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=bi&f=initParquet&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to initialize Parquet: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDuckdbTools(s *server.MCPServer, client *client.ZenTaoClient) {
	installDuckdbTool := mcp.NewTool("bi_install_duckdb",
		mcp.WithDescription("Install DuckDB"),
	)

	s.AddTool(installDuckdbTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=bi&f=ajaxInstallDuckdb&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to install DuckDB: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	checkDuckdbTool := mcp.NewTool("bi_check_duckdb",
		mcp.WithDescription("Check DuckDB"),
	)

	s.AddTool(checkDuckdbTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=bi&f=ajaxCheckDuckdb&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to check DuckDB: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerScopeFieldTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getScopeOptionsTool := mcp.NewTool("bi_get_scope_options",
		mcp.WithDescription("Get BI scope options"),
		mcp.WithString("type", mcp.Description("Type filter")),
	)

	s.AddTool(getScopeOptionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		url := "/index.php?m=bi&f=ajaxGetScopeOptions&t=json"

		if v, ok := args["type"].(string); ok {
			url += fmt.Sprintf("&type=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get scope options: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTableFieldsMenuTool := mcp.NewTool("bi_get_table_fields_menu",
		mcp.WithDescription("Get BI table fields menu"),
	)

	s.AddTool(getTableFieldsMenuTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=bi&f=ajaxGetTableFieldsMenu&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get table fields menu: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
