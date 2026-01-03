package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterStoryTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createStoryTool := mcp.NewTool("create_story",
		mcp.WithDescription("Create a new user story in ZenTao"),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Story title"),
		),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("pri",
			mcp.Required(),
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("category",
			mcp.Required(),
			mcp.Description("Story category"),
		),
		mcp.WithString("spec",
			mcp.Description("Story description"),
		),
		mcp.WithString("verify",
			mcp.Description("Acceptance criteria"),
		),
		mcp.WithString("source",
			mcp.Description("Source"),
		),
		mcp.WithString("sourceNote",
			mcp.Description("Source note"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(createStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"title":    args["title"],
			"product":  int(args["product"].(float64)),
			"pri":      int(args["pri"].(float64)),
			"category": args["category"],
		}

		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["verify"]; ok && v != nil {
			body["verify"] = v
		}
		if v, ok := args["source"]; ok && v != nil {
			body["source"] = v
		}
		if v, ok := args["sourceNote"]; ok && v != nil {
			body["sourceNote"] = v
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Post("/stories", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateStoryTool := mcp.NewTool("update_story",
		mcp.WithDescription("Update an existing story in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithString("source",
			mcp.Description("Source"),
		),
		mcp.WithString("sourceNote",
			mcp.Description("Source note"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("category",
			mcp.Description("Type (feature|interface|performance|safe|experience|improve|other)"),
			mcp.Enum("feature", "interface", "performance", "safe", "experience", "improve", "other"),
		),
		mcp.WithNumber("estimate",
			mcp.Description("Estimated hours"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(updateStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["source"]; ok && v != nil {
			body["source"] = v
		}
		if v, ok := args["sourceNote"]; ok && v != nil {
			body["sourceNote"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["category"]; ok && v != nil {
			body["category"] = v
		}
		if v, ok := args["estimate"]; ok && v != nil {
			body["estimate"] = v.(float64)
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/stories/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	changeStoryTool := mcp.NewTool("change_story",
		mcp.WithDescription("Change story content"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
		mcp.WithString("title",
			mcp.Description("Story title"),
		),
		mcp.WithString("spec",
			mcp.Description("Story description"),
		),
		mcp.WithString("verify",
			mcp.Description("Acceptance criteria"),
		),
	)

	s.AddTool(changeStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["spec"]; ok && v != nil {
			body["spec"] = v
		}
		if v, ok := args["verify"]; ok && v != nil {
			body["verify"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/stories/%d/change", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to change story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteStoryTool := mcp.NewTool("delete_story",
		mcp.WithDescription("Delete a story from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Story ID to delete"),
		),
	)

	s.AddTool(deleteStoryTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/stories/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete story: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
