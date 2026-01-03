package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterTestCaseTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTestCaseTool := mcp.NewTool("create_testcase",
		mcp.WithDescription("Create a new test case in ZenTao"),
		mcp.WithNumber("product",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("title",
			mcp.Required(),
			mcp.Description("Test case title"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Test case type"),
			mcp.Enum("feature", "performance", "config", "install", "security", "interface", "unit", "other"),
		),
		mcp.WithArray("steps",
			mcp.Required(),
			mcp.Description("Test case steps"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Story ID"),
		),
		mcp.WithString("stage",
			mcp.Description("Stage"),
			mcp.Enum("unittest", "feature", "intergrate", "system", "smoke", "bvt"),
		),
		mcp.WithString("precondition",
			mcp.Description("Precondition"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(createTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		productID := int(args["product"].(float64))

		body := map[string]interface{}{
			"title": args["title"],
			"type":  args["type"],
			"steps": args["steps"],
		}

		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["stage"]; ok && v != nil {
			body["stage"] = v
		}
		if v, ok := args["precondition"]; ok && v != nil {
			body["precondition"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/products/%d/testcases", productID), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	updateTestCaseTool := mcp.NewTool("update_testcase",
		mcp.WithDescription("Update an existing test case in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test case ID"),
		),
		mcp.WithString("title",
			mcp.Description("Test case title"),
		),
		mcp.WithString("type",
			mcp.Description("Test case type"),
			mcp.Enum("feature", "performance", "config", "install", "security", "interface", "unit", "other"),
		),
		mcp.WithArray("steps",
			mcp.Description("Test case steps"),
		),
		mcp.WithNumber("branch",
			mcp.Description("Branch ID"),
		),
		mcp.WithNumber("module",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("story",
			mcp.Description("Story ID"),
		),
		mcp.WithString("stage",
			mcp.Description("Stage"),
			mcp.Enum("unittest", "feature", "intergrate", "system", "smoke", "bvt"),
		),
		mcp.WithString("precondition",
			mcp.Description("Precondition"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority (1-9)"),
		),
		mcp.WithString("keywords",
			mcp.Description("Keywords"),
		),
	)

	s.AddTool(updateTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		body := make(map[string]interface{})

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["steps"]; ok && v != nil {
			body["steps"] = v
		}
		if v, ok := args["branch"]; ok && v != nil {
			body["branch"] = int(v.(float64))
		}
		if v, ok := args["module"]; ok && v != nil {
			body["module"] = int(v.(float64))
		}
		if v, ok := args["story"]; ok && v != nil {
			body["story"] = int(v.(float64))
		}
		if v, ok := args["stage"]; ok && v != nil {
			body["stage"] = v
		}
		if v, ok := args["precondition"]; ok && v != nil {
			body["precondition"] = v
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["keywords"]; ok && v != nil {
			body["keywords"] = v
		}

		resp, err := client.Put(fmt.Sprintf("/testcases/%d", id), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTestCaseTool := mcp.NewTool("delete_testcase",
		mcp.WithDescription("Delete a test case from ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Test case ID to delete"),
		),
	)

	s.AddTool(deleteTestCaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		id := int(args["id"].(float64))

		resp, err := client.Delete(fmt.Sprintf("/testcases/%d", id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete test case: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
