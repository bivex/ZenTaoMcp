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

func RegisterExecutionTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseExecutionTool := mcp.NewTool("browse_execution",
		mcp.WithDescription("Browse execution details"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
	)

	s.AddTool(browseExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		executionID := int(args["executionID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=browse&t=json&executionID=%d", executionID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionTasksTool := mcp.NewTool("get_execution_tasks",
		mcp.WithDescription("Get tasks for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("status",
			mcp.Description("Task status"),
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
		mcp.WithString("from",
			mcp.Description("Source"),
		),
		mcp.WithString("blockID",
			mcp.Description("Block ID"),
		),
	)

	s.AddTool(getExecutionTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
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
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["blockID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&blockID=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=task&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution tasks: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionStoriesTool := mcp.NewTool("get_execution_stories",
		mcp.WithDescription("Get stories for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
			mcp.Enum("story", "requirement"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithString("type",
			mcp.Description("View type"),
			mcp.Enum("all", "byModule", "byProduct", "byBranch", "bySearch"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
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
		mcp.WithString("from",
			mcp.Description("Source"),
		),
		mcp.WithString("blockID",
			mcp.Description("Block ID"),
		),
	)

	s.AddTool(getExecutionStoriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
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
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["blockID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&blockID=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=story&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionBugsTool := mcp.NewTool("get_execution_bugs",
		mcp.WithDescription("Get bugs for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("productID",
			mcp.Required(),
			mcp.Description("Product ID"),
		),
		mcp.WithString("branch",
			mcp.Description("Branch"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithString("build",
			mcp.Description("Build"),
		),
		mcp.WithString("type",
			mcp.Description("Bug type"),
		),
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
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

	s.AddTool(getExecutionBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d&productID=%d", int(args["executionID"].(float64)), int(args["productID"].(float64)))
		if v, ok := args["branch"]; ok && v != nil {
			queryParams += fmt.Sprintf("&branch=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["build"]; ok && v != nil {
			queryParams += fmt.Sprintf("&build=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=bug&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionBuildsTool := mcp.NewTool("get_execution_builds",
		mcp.WithDescription("Get builds for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
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

	s.AddTool(getExecutionBuildsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=build&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution builds: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionBurnChartTool := mcp.NewTool("get_execution_burn_chart",
		mcp.WithDescription("Get burn-down chart for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("type",
			mcp.Description("Chart type"),
			mcp.Enum("noweekend", "withweekend"),
		),
		mcp.WithString("interval",
			mcp.Description("Interval"),
		),
		mcp.WithString("burnBy",
			mcp.Description("Burn by"),
			mcp.Enum("left", "estimate", "storyPoint"),
		),
	)

	s.AddTool(getExecutionBurnChartTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["interval"]; ok && v != nil {
			queryParams += fmt.Sprintf("&interval=%s", v)
		}
		if v, ok := args["burnBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&burnBy=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=burn&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution burn chart: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionCFDTool := mcp.NewTool("get_execution_cfd",
		mcp.WithDescription("Get Cumulative Flow Diagram for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("type",
			mcp.Description("CFD type"),
			mcp.Enum("story", "bug", "task"),
		),
		mcp.WithString("withWeekend",
			mcp.Description("Include weekends"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin date"),
		),
		mcp.WithString("end",
			mcp.Description("End date"),
		),
	)

	s.AddTool(getExecutionCFDTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["withWeekend"]; ok && v != nil {
			queryParams += fmt.Sprintf("&withWeekend=%s", v)
		}
		if v, ok := args["begin"]; ok && v != nil {
			queryParams += fmt.Sprintf("&begin=%s", v)
		}
		if v, ok := args["end"]; ok && v != nil {
			queryParams += fmt.Sprintf("&end=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=execution&f=cfd&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution CFD: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionKanbanTool := mcp.NewTool("get_execution_kanban",
		mcp.WithDescription("Get kanban view for an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
			mcp.Enum("all", "story", "bug", "task"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithString("groupBy",
			mcp.Description("Group by field"),
			mcp.Enum("default", "pri", "category", "module", "source", "assignedTo", "type", "story", "severity"),
		),
	)

	s.AddTool(getExecutionKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["groupBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&groupBy=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=kanban&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionTeamTool := mcp.NewTool("get_execution_team",
		mcp.WithDescription("Get execution team members"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
	)

	s.AddTool(getExecutionTeamTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		executionID := int(args["executionID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=team&t=json&executionID=%d", executionID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution team: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	manageExecutionMembersTool := mcp.NewTool("manage_execution_members",
		mcp.WithDescription("Manage execution team members"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("team2Import",
			mcp.Description("Team to import"),
		),
		mcp.WithString("dept",
			mcp.Description("Department"),
		),
		mcp.WithArray("members",
			mcp.Description("Members to add"),
		),
	)

	s.AddTool(manageExecutionMembersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["members"]; ok && v != nil {
			body["members"] = v
		}

		queryParams := fmt.Sprintf("&executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["team2Import"]; ok && v != nil {
			queryParams += fmt.Sprintf("&team2Import=%d", int(v.(float64)))
		}
		if v, ok := args["dept"]; ok && v != nil {
			queryParams += fmt.Sprintf("&dept=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=execution&f=manageMembers&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to manage execution members: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkExecutionMemberTool := mcp.NewTool("unlink_execution_member",
		mcp.WithDescription("Remove a member from execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("userID",
			mcp.Required(),
			mcp.Description("User ID to remove"),
		),
	)

	s.AddTool(unlinkExecutionMemberTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=unlinkMember&t=json&executionID=%d&userID=%d",
			int(args["executionID"].(float64)), int(args["userID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink execution member: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	linkStoryToExecutionTool := mcp.NewTool("link_story_to_execution",
		mcp.WithDescription("Link a story to an execution"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Execution ID"),
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
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra parameters"),
		),
		mcp.WithString("storyType",
			mcp.Description("Story type"),
			mcp.Enum("story", "requirement"),
		),
	)

	s.AddTool(linkStoryToExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("objectID=%d", int(args["objectID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}
		if v, ok := args["storyType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&storyType=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=execution&f=linkStory&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to link story to execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	unlinkStoryFromExecutionTool := mcp.NewTool("unlink_story_from_execution",
		mcp.WithDescription("Unlink a story from an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithNumber("storyID",
			mcp.Required(),
			mcp.Description("Story ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
		),
		mcp.WithNumber("laneID",
			mcp.Description("Lane ID"),
		),
		mcp.WithNumber("columnID",
			mcp.Description("Column ID"),
		),
	)

	s.AddTool(unlinkStoryFromExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d&storyID=%d", int(args["executionID"].(float64)), int(args["storyID"].(float64)))
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams += fmt.Sprintf("&confirm=%s", v)
		}
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}
		if v, ok := args["laneID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&laneID=%d", int(v.(float64)))
		}
		if v, ok := args["columnID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&columnID=%d", int(v.(float64)))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=unlinkStory&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to unlink story from execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	batchUnlinkStoriesFromExecutionTool := mcp.NewTool("batch_unlink_stories_from_execution",
		mcp.WithDescription("Batch unlink stories from an execution"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
	)

	s.AddTool(batchUnlinkStoriesFromExecutionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		executionID := int(args["executionID"].(float64))

		resp, err := client.Post(fmt.Sprintf("/index.php?m=execution&f=batchUnlinkStory&t=json&executionID=%d", executionID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to batch unlink stories from execution: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getExecutionDynamicTool := mcp.NewTool("get_execution_dynamic",
		mcp.WithDescription("Get execution dynamic/activity log"),
		mcp.WithNumber("executionID",
			mcp.Required(),
			mcp.Description("Execution ID"),
		),
		mcp.WithString("type",
			mcp.Description("Activity type"),
		),
		mcp.WithString("param",
			mcp.Description("Parameter"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithString("date",
			mcp.Description("Date"),
		),
		mcp.WithString("direction",
			mcp.Description("Direction"),
			mcp.Enum("next", "pre"),
		),
	)

	s.AddTool(getExecutionDynamicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("executionID=%d", int(args["executionID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%s", v)
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["date"]; ok && v != nil {
			queryParams += fmt.Sprintf("&date=%s", v)
		}
		if v, ok := args["direction"]; ok && v != nil {
			queryParams += fmt.Sprintf("&direction=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=dynamic&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get execution dynamic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	browseAllExecutionsTool := mcp.NewTool("browse_all_executions",
		mcp.WithDescription("Browse all executions"),
		mcp.WithString("status",
			mcp.Description("Execution status"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
		mcp.WithNumber("productID",
			mcp.Description("Product ID"),
		),
		mcp.WithString("param",
			mcp.Description("Parameter"),
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

	s.AddTool(browseAllExecutionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}
		if v, ok := args["productID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&productID=%d", int(v.(float64)))
		}
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=execution&f=all&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse all executions: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
