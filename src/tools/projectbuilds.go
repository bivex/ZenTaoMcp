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

func RegisterProjectBuildTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseProjectBuildsTool := mcp.NewTool("browse_project_builds",
		mcp.WithDescription("Browse builds for a project"),
		mcp.WithNumber("projectID",
			mcp.Required(),
			mcp.Description("Project ID"),
		),
		mcp.WithString("type",
			mcp.Description("Build type"),
			mcp.Enum("all", "product", "bysearch"),
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

	s.AddTool(browseProjectBuildsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("projectID=%d", int(args["projectID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=browse&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse project builds: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createProjectBuildTool := mcp.NewTool("create_project_build",
		mcp.WithDescription("Create a new build for a project"),
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

	s.AddTool(createProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=projectbuild&f=create&t=json&projectID=%d", int(args["projectID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editProjectBuildTool := mcp.NewTool("edit_project_build",
		mcp.WithDescription("Edit an existing project build"),
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

	s.AddTool(editProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Post(fmt.Sprintf("/index.php?m=projectbuild&f=edit&t=json&buildID=%d", int(args["buildID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewProjectBuildTool := mcp.NewTool("view_project_build",
		mcp.WithDescription("View project build details"),
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

	s.AddTool(viewProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteProjectBuildTool := mcp.NewTool("delete_project_build",
		mcp.WithDescription("Delete a project build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
	)

	s.AddTool(deleteProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		buildID := int(args["buildID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=delete&t=json&buildID=%d", buildID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkStoryToProjectBuildTool := mcp.NewTool("link_story_to_project_build",
		mcp.WithDescription("Link a story to a project build"),
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

	s.AddTool(linkStoryToProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=linkStory&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link story to project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkStoryFromProjectBuildTool := mcp.NewTool("unlink_story_from_project_build",
		mcp.WithDescription("Unlink a story from a project build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
	)

	s.AddTool(unlinkStoryFromProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=unlinkStory&t=json&buildID=%d&storyID=%d",
			int(args["buildID"].(float64)), int(args["storyID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink story from project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkStoriesFromProjectBuildTool := mcp.NewTool("batch_unlink_stories_from_project_build",
		mcp.WithDescription("Batch unlink stories from a project build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
	)

	s.AddTool(batchUnlinkStoriesFromProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		buildID := int(args["buildID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=batchUnlinkStory&t=json&buildID=%d", buildID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink stories from project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkBugToProjectBuildTool := mcp.NewTool("link_bug_to_project_build",
		mcp.WithDescription("Link a bug to a project build"),
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

	s.AddTool(linkBugToProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=linkBug&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link bug to project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkBugFromProjectBuildTool := mcp.NewTool("unlink_bug_from_project_build",
		mcp.WithDescription("Unlink a bug from a project build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
		mcp.WithNumber("bugID",
			mcp.Required(),
			mcp.Description("Bug ID"),
		),
	)

	s.AddTool(unlinkBugFromProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=unlinkBug&t=json&buildID=%d&bugID=%d",
			int(args["buildID"].(float64)), int(args["bugID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink bug from project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkBugsFromProjectBuildTool := mcp.NewTool("batch_unlink_bugs_from_project_build",
		mcp.WithDescription("Batch unlink bugs from a project build"),
		mcp.WithNumber("buildID",
			mcp.Required(),
			mcp.Description("Build ID"),
		),
	)

	s.AddTool(batchUnlinkBugsFromProjectBuildTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		buildID := int(args["buildID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=projectbuild&f=batchUnlinkBug&t=json&buildID=%d", buildID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink bugs from project build: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
