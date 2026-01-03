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

func RegisterMyTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Dashboard/Index
	getMyDashboardTool := mcp.NewTool("get_my_dashboard",
		mcp.WithDescription("Get user's personal dashboard in ZenTao"),
	)

	s.AddTool(getMyDashboardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=my&f=index&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get dashboard: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Score/Ranking
	getMyScoreTool := mcp.NewTool("get_my_score",
		mcp.WithDescription("Get user's score/ranking information"),
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

	s.AddTool(getMyScoreTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=score&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get score: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Calendar
	getMyCalendarTool := mcp.NewTool("get_my_calendar",
		mcp.WithDescription("Get user's calendar view"),
	)

	s.AddTool(getMyCalendarTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=my&f=calendar&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get calendar: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Work tracking
	getMyWorkTool := mcp.NewTool("get_my_work",
		mcp.WithDescription("Get user's work tracking information"),
		mcp.WithString("mode",
			mcp.Description("Work mode"),
		),
		mcp.WithString("type",
			mcp.Description("Work type"),
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

	s.AddTool(getMyWorkTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["mode"]; ok && v != nil {
			queryParams += fmt.Sprintf("&mode=%s", v)
		}
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=work&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get work: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Contributions
	getMyContributeTool := mcp.NewTool("get_my_contribute",
		mcp.WithDescription("Get user's contribution statistics"),
		mcp.WithString("mode",
			mcp.Description("Contribution mode"),
		),
		mcp.WithString("type",
			mcp.Description("Contribution type"),
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

	s.AddTool(getMyContributeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["mode"]; ok && v != nil {
			queryParams += fmt.Sprintf("&mode=%s", v)
		}
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=contribute&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get contributions: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Todos
	getMyTodosTool := mcp.NewTool("get_my_todos",
		mcp.WithDescription("Get user's personal todos"),
		mcp.WithString("type",
			mcp.Description("Todo type"),
		),
		mcp.WithNumber("userID",
			mcp.Description("User ID"),
		),
		mcp.WithString("status",
			mcp.Description("Todo status"),
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

	s.AddTool(getMyTodosTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["userID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&userID=%d", int(v.(float64)))
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=todo&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get todos: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Stories
	getMyStoriesTool := mcp.NewTool("get_my_stories",
		mcp.WithDescription("Get user's personal stories"),
		mcp.WithString("type",
			mcp.Description("Story type"),
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

	s.AddTool(getMyStoriesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=story&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get stories: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Tasks
	getMyTasksTool := mcp.NewTool("get_my_tasks",
		mcp.WithDescription("Get user's personal tasks"),
		mcp.WithString("type",
			mcp.Description("Task type"),
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

	s.AddTool(getMyTasksTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=task&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get tasks: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Bugs
	getMyBugsTool := mcp.NewTool("get_my_bugs",
		mcp.WithDescription("Get user's personal bugs"),
		mcp.WithString("type",
			mcp.Description("Bug type"),
		),
		mcp.WithString("param",
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

	s.AddTool(getMyBugsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=bug&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get bugs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Projects
	getMyProjectsTool := mcp.NewTool("get_my_projects",
		mcp.WithDescription("Get user's personal projects"),
		mcp.WithString("status",
			mcp.Description("Project status"),
			mcp.Enum("doing", "wait", "suspended", "closed", "openedbyme"),
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

	s.AddTool(getMyProjectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=project&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get projects: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Personal Executions
	getMyExecutionsTool := mcp.NewTool("get_my_executions",
		mcp.WithDescription("Get user's personal executions"),
		mcp.WithString("type",
			mcp.Description("Execution type"),
			mcp.Enum("undone", "done"),
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

	s.AddTool(getMyExecutionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=execution&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get executions: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Profile management
	getMyProfileTool := mcp.NewTool("get_my_profile",
		mcp.WithDescription("Get user's profile information"),
	)

	s.AddTool(getMyProfileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=my&f=profile&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get profile: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editMyProfileTool := mcp.NewTool("edit_my_profile",
		mcp.WithDescription("Edit user's profile information"),
		mcp.WithString("realname",
			mcp.Description("Real name"),
		),
		mcp.WithString("email",
			mcp.Description("Email address"),
		),
		mcp.WithString("mobile",
			mcp.Description("Mobile phone"),
		),
		mcp.WithString("phone",
			mcp.Description("Phone number"),
		),
		mcp.WithString("address",
			mcp.Description("Address"),
		),
		mcp.WithString("zipcode",
			mcp.Description("Zip code"),
		),
		mcp.WithString("skype",
			mcp.Description("Skype ID"),
		),
		mcp.WithString("qq",
			mcp.Description("QQ number"),
		),
		mcp.WithString("dingding",
			mcp.Description("DingTalk ID"),
		),
		mcp.WithString("weixin",
			mcp.Description("WeChat ID"),
		),
	)

	s.AddTool(editMyProfileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["realname"]; ok && v != nil {
			body["realname"] = v
		}
		if v, ok := args["email"]; ok && v != nil {
			body["email"] = v
		}
		if v, ok := args["mobile"]; ok && v != nil {
			body["mobile"] = v
		}
		if v, ok := args["phone"]; ok && v != nil {
			body["phone"] = v
		}
		if v, ok := args["address"]; ok && v != nil {
			body["address"] = v
		}
		if v, ok := args["zipcode"]; ok && v != nil {
			body["zipcode"] = v
		}
		if v, ok := args["skype"]; ok && v != nil {
			body["skype"] = v
		}
		if v, ok := args["qq"]; ok && v != nil {
			body["qq"] = v
		}
		if v, ok := args["dingding"]; ok && v != nil {
			body["dingding"] = v
		}
		if v, ok := args["weixin"]; ok && v != nil {
			body["weixin"] = v
		}

		resp, err := client.Post("/index.php?m=my&f=editProfile&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit profile: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	changeMyPasswordTool := mcp.NewTool("change_my_password",
		mcp.WithDescription("Change user's password"),
		mcp.WithString("password",
			mcp.Required(),
			mcp.Description("New password"),
		),
		mcp.WithString("passwordConfirmation",
			mcp.Required(),
			mcp.Description("Password confirmation"),
		),
	)

	s.AddTool(changeMyPasswordTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"password":             args["password"],
			"passwordConfirmation": args["passwordConfirmation"],
		}

		resp, err := client.Post("/index.php?m=my&f=changePassword&t=json", body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to change password: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Dynamic/Activity feed
	getMyDynamicTool := mcp.NewTool("get_my_dynamic",
		mcp.WithDescription("Get user's activity feed/dynamic"),
		mcp.WithString("type",
			mcp.Description("Dynamic type"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithString("date",
			mcp.Description("Date filter"),
		),
		mcp.WithString("direction",
			mcp.Description("Direction (next/prev)"),
		),
	)

	s.AddTool(getMyDynamicTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=dynamic&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get dynamic: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Team management
	getMyTeamTool := mcp.NewTool("get_my_team",
		mcp.WithDescription("Get user's team information"),
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

	s.AddTool(getMyTeamTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
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

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=my&f=team&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get team: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
