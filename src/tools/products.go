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
)

func RegisterProductTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createProductTool := mcp.NewTool("create_product",
		mcp.WithDescription("Create a new product in ZenTao"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Product name"),
		),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("Product code"),
		),
		mcp.WithNumber("program",
			mcp.Description("Program ID"),
		),
		mcp.WithNumber("line",
			mcp.Description("Product line"),
		),
		mcp.WithNumber("PO",
			mcp.Description("Product Owner ID"),
		),
		mcp.WithNumber("QD",
			mcp.Description("Quality Director ID"),
		),
		mcp.WithNumber("RD",
			mcp.Description("Release Director ID"),
		),
		mcp.WithString("type",
			mcp.Description("Product type (normal|branch)"),
			mcp.Enum("normal", "branch"),
		),
		mcp.WithString("desc",
			mcp.Description("Product description"),
		),
		mcp.WithString("acl",
			mcp.Description("Access control (open|private)"),
			mcp.Enum("open", "private"),
		),
	)

	s.AddTool(createProductTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
			"code": args["code"],
		}

		if v, ok := args["program"]; ok && v != nil {
			body["program"] = int(v.(float64))
		}
		if v, ok := args["line"]; ok && v != nil {
			body["line"] = int(v.(float64))
		}
		if v, ok := args["PO"]; ok && v != nil {
			body["PO"] = int(v.(float64))
		}
		if v, ok := args["QD"]; ok && v != nil {
			body["QD"] = int(v.(float64))
		}
		if v, ok := args["RD"]; ok && v != nil {
			body["RD"] = int(v.(float64))
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["acl"]; ok && v != nil {
			body["acl"] = v
		}

		resp, err := client.Post("/products", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create product: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateProductTool := mcp.NewTool("update_product",
		mcp.WithDescription("Update an existing product in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("name",
			mcp.Description("Product name"),
		),
		mcp.WithString("code",
			mcp.Description("Product code"),
		),
		mcp.WithString("type",
			mcp.Description("Product type (normal|branch|platform)"),
			mcp.Enum("normal", "branch", "platform"),
		),
		mcp.WithNumber("line",
			mcp.Description("Product line"),
		),
		mcp.WithNumber("program",
			mcp.Description("Program"),
		),
		mcp.WithString("status",
			mcp.Description("Product status (normal|closed)"),
			mcp.Enum("normal", "closed"),
		),
		mcp.WithString("desc",
			mcp.Description("Product description"),
		),
	)

	s.AddTool(updateProductTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["line"]; ok && v != nil {
			body["line"] = int(v.(float64))
		}
		if v, ok := args["program"]; ok && v != nil {
			body["program"] = int(v.(float64))
		}
		if v, ok := args["status"]; ok && v != nil {
			body["status"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/product/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update product: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteProductTool := mcp.NewTool("delete_product",
		mcp.WithDescription("Delete a product from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Product ID to delete"),
		),
	)

	s.AddTool(deleteProductTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/product/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete product: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
