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

func RegisterQaTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getQaIndexTool := mcp.NewTool("get_qa_index",
		mcp.WithDescription("Get QA module index"),
		mcp.WithString("locate",
			mcp.Description("Location context"),
		),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
	)

	s.AddTool(getQaIndexTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["locate"]; ok && v != nil {
			queryParams += fmt.Sprintf("&locate=%s", v)
		}
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=qa&f=index&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get QA index: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
