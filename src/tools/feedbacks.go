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

func RegisterFeedbackTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createFeedbackTool := mcp.NewTool("create_feedback",
		mcp.WithDescription("Create a new feedback in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Feedback title"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithString("type",
			mcp.Description("Type (story|task|bug|todo|advice|issue|risk|opportunity)"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithNumber("public",
			mcp.Description("Public (0|1)"),
		),
		mcp.WithNumber("notify",
			mcp.Description("Notify (0|1)"),
		),
		mcp.WithString("notifyEmail",
			mcp.Description("Notify email"),
		),
		mcp.WithString("feedbackBy",
			mcp.Description("Feedback by user account"),
		),
	)

	s.AddTool(createFeedbackTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"product": int(args["product"].(float64)),
			"title":   args["title"],
		}

		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["public"]; ok && v != nil {
			body["public"] = int(v.(float64))
		}
		if v, ok := args["notify"]; ok && v != nil {
			body["notify"] = int(v.(float64))
		}
		if v, ok := args["notifyEmail"]; ok && v != nil {
			body["notifyEmail"] = v
		}
		if v, ok := args["feedbackBy"]; ok && v != nil {
			body["feedbackBy"] = v
		}

		resp, err := client.Post("/feedbacks", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create feedback: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateFeedbackTool := mcp.NewTool("update_feedback",
		mcp.WithDescription("Update an existing feedback in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Feedback ID"),
		),
		mcp.WithNumber("product",
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithString("title",
			mcp.Description("Title"),
		),
		mcp.WithString("type",
			mcp.Description("Type (story|task|bug|todo|advice|issue|risk|opportunity)"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithNumber("public",
			mcp.Description("Public (0|1)"),
		),
		mcp.WithNumber("notify",
			mcp.Description("Notify (0|1)"),
		),
		mcp.WithString("notifyEmail",
			mcp.Description("Notify email"),
		),
		mcp.WithString("feedbackBy",
			mcp.Description("Feedback by user account"),
		),
	)

	s.AddTool(updateFeedbackTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if v, ok := args["public"]; ok && v != nil {
			body["public"] = int(v.(float64))
		}
		if v, ok := args["notify"]; ok && v != nil {
			body["notify"] = int(v.(float64))
		}
		if v, ok := args["notifyEmail"]; ok && v != nil {
			body["notifyEmail"] = v
		}
		if v, ok := args["feedbackBy"]; ok && v != nil {
			body["feedbackBy"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/feedbacks/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update feedback: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	assignFeedbackTool := mcp.NewTool("assign_feedback",
		mcp.WithDescription("Assign a feedback to a user in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Feedback ID"),
		),
		mcp.WithString("assignedTo",
			mcp.Description("Assign to user account"),
		),
		mcp.WithString("comment",
			mcp.Description("Comment"),
		),
		mcp.WithString("mailto",
			mcp.Description("CC to user accounts (comma-separated)"),
		),
	)

	s.AddTool(assignFeedbackTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = v
		}
		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}
		if v, ok := args["mailto"]; ok && v != nil {
			body["mailto"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/feedbacks/%d/assign", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to assign feedback: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeFeedbackTool := mcp.NewTool("close_feedback",
		mcp.WithDescription("Close a feedback in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Feedback ID"),
		),
		mcp.WithString("closedReason",
			mcp.Description("Close reason (commented|repeat|refuse)"),
		),
		mcp.WithString("comment",
			mcp.Description("Comment"),
		),
	)

	s.AddTool(closeFeedbackTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["closedReason"]; ok && v != nil {
			body["closedReason"] = v
		}
		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/feedbacks/%d/close", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close feedback: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteFeedbackTool := mcp.NewTool("delete_feedback",
		mcp.WithDescription("Delete a feedback from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Feedback ID to delete"),
		),
	)

	s.AddTool(deleteFeedbackTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/feedbacks/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete feedback: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
