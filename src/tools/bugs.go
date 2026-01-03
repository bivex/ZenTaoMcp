package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterBugTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createBugTool := mcp.NewTool("create_bug",
		mcp.WithDescription("Create a new bug in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Required(),
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Required(),
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(createBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title":    args["title"],
			"severity": int(args["severity"].(float64)),
			"pri":      int(args["pri"].(float64)),
			"type":     args["type"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/bugs", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateBugTool := mcp.NewTool("update_bug",
		mcp.WithDescription("Update an existing bug in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
		mcp.WithString("title",
			mcp.Description("Bug title"),
		),
		mcp.WithNumber("severity",
			mcp.Description("Severity (1-4)"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("type",
			mcp.Description("Bug type"),
			mcp.Enum("codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("execution",
			mcp.Description("Execution ID"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
		mcp.WithString("os",
			mcp.Description("Operating system"),
		),
		mcp.WithString("browser",
			mcp.Description("Browser"),
		),
		mcp.WithString("steps",
			mcp.Description("Reproduction steps"),
		),
		mcp.WithNumber("task",
			mcp.Description("Related task ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Related story ID"),
		),
		mcp.WithString("deadline",
			mcp.Description("Deadline (YYYY-MM-DD)"),
		),
		mcp.WithArray("openedBuild",
			mcp.Description("Affected builds"),
		),
	)

	s.AddTool(updateBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["severity"]; ok && v != nil {
			body["severity"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["execution"]; ok && v != nil {
			body["execution"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}
		if v, ok := args["os"]; ok && v != nil {
			body["os"] = v
		}
		if v, ok := args["browser"]; ok && v != nil {
			body["browser"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["task"]; ok && v != nil {
			body["task"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["deadline"]; ok && v != nil {
			body["deadline"] = v
		}
		if v, ok := args["openedBuild"]; ok && v != nil {
			body["openedBuild"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/bugs/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteBugTool := mcp.NewTool("delete_bug",
		mcp.WithDescription("Delete a bug from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Bug ID to delete"),
		),
	)

	s.AddTool(deleteBugTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/bugs/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete bug: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
