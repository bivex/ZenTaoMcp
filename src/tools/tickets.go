package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterTicketTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTicketTool := mcp.NewTool("create_ticket",
		mcp.WithDescription("Create a new ticket in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("module",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Ticket name"),
		),
		mcp.WithString("type",
			mcp.Description("Ticket type (code|data|stuck|security|affair)"),
		),
		mcp.WithString("desc",
			mcp.Description("Ticket description"),
		),
	)

	s.AddTool(createTicketTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"product": int(args["product"].(float64)),
			"module":  int(args["module"].(float64)),
			"title":   args["title"],
		}

		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post("/tickets", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create ticket: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateTicketTool := mcp.NewTool("update_ticket",
		mcp.WithDescription("Update an existing ticket in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Ticket ID"),
		),
		mcp.WithNumber("product",
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithString("title",
			mcp.Description("Ticket name"),
		),
		mcp.WithString("type",
			mcp.Description("Ticket type (code|data|stuck|security|affair)"),
		),
		mcp.WithString("desc",
			mcp.Description("Ticket description"),
		),
	)

	s.AddTool(updateTicketTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["product"]; ok && v != nil {
			body["product"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/tickets/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update ticket: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTicketTool := mcp.NewTool("delete_ticket",
		mcp.WithDescription("Delete a ticket from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Ticket ID to delete"),
		),
	)

	s.AddTool(deleteTicketTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/tickets/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete ticket: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
