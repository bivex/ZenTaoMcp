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

func RegisterBuildTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createBuildTool := mcp.NewTool("create_build",
		mcp.WithDescription("Create a new build"),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Build name"),
		),
		mcp.WithString("builder",
			mcp.Description("Builder"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithString("scmPath",
			mcp.Description("SCM path"),
		),
		mcp.WithString("filePath",
			mcp.Description("File path"),
		),
		mcp.WithString("date",
			mcp.Description("Build date"),
		),
	)

	s.AddTool(createBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["builder"]; ok && v != nil {
			body["builder"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["scmPath"]; ok && v != nil {
			body["scmPath"] = v
		}
		if v, ok := args["filePath"]; ok && v != nil {
			body["filePath"] = v
		}
		if v, ok := args["date"]; ok && v != nil {
			body["date"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=create&t=json&executionID=%d&productID=%d&projectID=%d",
			int(args["executionID"].(float64)), int(args["productID"].(float64)), int(args["projectID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editBuildTool := mcp.NewTool("edit_build",
		mcp.WithDescription("Edit an existing build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("name",
			mcp.Description("Build name"),
		),
		mcp.WithString("builder",
			mcp.Description("Builder"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithString("scmPath",
			mcp.Description("SCM path"),
		),
		mcp.WithString("filePath",
			mcp.Description("File path"),
		),
		mcp.WithString("date",
			mcp.Description("Build date"),
		),
	)

	s.AddTool(editBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["builder"]; ok && v != nil {
			body["builder"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["scmPath"]; ok && v != nil {
			body["scmPath"] = v
		}
		if v, ok := args["filePath"]; ok && v != nil {
			body["filePath"] = v
		}
		if v, ok := args["date"]; ok && v != nil {
			body["date"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=edit&t=json&buildID=%d", int(args["buildID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewBuildTool := mcp.NewTool("view_build",
		mcp.WithDescription("View build details"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("type",
			mcp.Description("View type"),
		),
		mcp.WithString("link",
			mcp.Description("Link"),
		),
		mcp.WithString("param",
			mcp.Description("Parameter"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(viewBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("buildID=%d", int(args["buildID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["link"]; ok && v != nil {
			queryParams += fmt.Sprintf("&link=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteBuildTool := mcp.NewTool("delete_build",
		mcp.WithDescription("Delete a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("execution", "project"),
		),
	)

	s.AddTool(deleteBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("buildID=%d", int(args["buildID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=delete&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProductBuildsTool := mcp.NewTool("get_product_builds",
		mcp.WithDescription("Get product builds"),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("varName",
			mcp.Description("Variable name"),
		),
		mcp.WithString("build",
			mcp.Description("Build filter"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("type",
			mcp.Description("Build type"),
		),
	)

	s.AddTool(getProductBuildsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("productID=%d", int(args["productID"].(float64)))
		if v, ok := args["varName"]; ok && v != nil {
			queryParams += fmt.Sprintf("&varName=%s", v)
		}
		if v, ok := args["build"]; ok && v != nil {
			queryParams += fmt.Sprintf("&build=%s", v)
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetProductBuilds&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get product builds: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProjectBuildsTool := mcp.NewTool("get_project_builds",
		mcp.WithDescription("Get project builds"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithString("varName",
			mcp.Description("Variable name"),
		),
		mcp.WithString("build",
			mcp.Description("Build filter"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("needCreate",
			mcp.Description("Need create option"),
		),
		mcp.WithString("type",
			mcp.Description("Build type"),
		),
		mcp.WithString("system",
			mcp.Description("System"),
		),
	)

	s.AddTool(getProjectBuildsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("projectID=%d", int(args["projectID"].(float64)))
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["varName"]; ok && v != nil {
			queryParams += fmt.Sprintf("&varName=%s", v)
		}
		if v, ok := args["build"]; ok && v != nil {
			queryParams += fmt.Sprintf("&build=%s", v)
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["needCreate"]; ok && v != nil {
			queryParams += fmt.Sprintf("&needCreate=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["system"]; ok && v != nil {
			queryParams += fmt.Sprintf("&system=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetProjectBuilds&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get project builds: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionBuildsTool := mcp.NewTool("get_execution_builds",
		mcp.WithDescription("Get execution builds"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithString("varName",
			mcp.Description("Variable name"),
		),
		mcp.WithString("build",
			mcp.Description("Build filter"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("needCreate",
			mcp.Description("Need create option"),
		),
		mcp.WithString("type",
			mcp.Description("Build type"),
		),
	)

	s.AddTool(getExecutionBuildsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["varName"]; ok && v != nil {
			queryParams += fmt.Sprintf("&varName=%s", v)
		}
		if v, ok := args["build"]; ok && v != nil {
			queryParams += fmt.Sprintf("&build=%s", v)
		}
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["needCreate"]; ok && v != nil {
			queryParams += fmt.Sprintf("&needCreate=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetExecutionBuilds&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution builds: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getLastBuildTool := mcp.NewTool("get_last_build",
		mcp.WithDescription("Get last build for project/execution"),
		mcp.WithNumber("projectID",
			mcp.Description("Project ID"),
		),
		mcp.WithNumber("executionID",
			mcp.Description("Execution ID"),
		),
	)

	s.AddTool(getLastBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["projectID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&projectID=%d", int(v.(float64)))
		}
		if v, ok := args["executionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&executionID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=ajaxGetLastBuild&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get last build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkStoryToBuildTool := mcp.NewTool("link_story_to_build",
		mcp.WithDescription("Link a story to a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(linkStoryToBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("buildID=%d", int(args["buildID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=linkStory&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link story to build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkStoryFromBuildTool := mcp.NewTool("unlink_story_from_build",
		mcp.WithDescription("Unlink a story from a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
	)

	s.AddTool(unlinkStoryFromBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=unlinkStory&t=json&buildID=%d&storyID=%d",
			int(args["buildID"].(float64)), int(args["storyID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink story from build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkStoriesFromBuildTool := mcp.NewTool("batch_unlink_stories_from_build",
		mcp.WithDescription("Batch unlink stories from a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
	)

	s.AddTool(batchUnlinkStoriesFromBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		buildID := int(args["buildID"].(float64))

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=batchUnlinkStory&t=json&buildID=%d", buildID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink stories from build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkBugToBuildTool := mcp.NewTool("link_bug_to_build",
		mcp.WithDescription("Link a bug to a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
	)

	s.AddTool(linkBugToBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("buildID=%d", int(args["buildID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=linkBug&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link bug to build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkBugFromBuildTool := mcp.NewTool("unlink_bug_from_build",
		mcp.WithDescription("Unlink a bug from a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
	)

	s.AddTool(unlinkBugFromBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=build&f=unlinkBug&t=json&buildID=%d&bugID=%d",
			int(args["buildID"].(float64)), int(args["bugID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink bug from build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkBugsFromBuildTool := mcp.NewTool("batch_unlink_bugs_from_build",
		mcp.WithDescription("Batch unlink bugs from a build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
	)

	s.AddTool(batchUnlinkBugsFromBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		buildID := int(args["buildID"].(float64))

		resp, err := client.Post(fmt.Sprintf("/index.php?m=build&f=batchUnlinkBug&t=json&buildID=%d", buildID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink bugs from build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}