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

func RegisterProgramTools(s *server.MCPServer, client *client.ZenTaoClient) {
	browseProgramsTool := mcp.NewTool("browse_programs",
		mcp.WithDescription("Browse programs in ZenTao"),
		mcp.WithString("status",
			mcp.Description("Program status"),
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
		mcp.WithNumber("param",
			mcp.Description("Parameter value"),
		),
	)

	s.AddTool(browseProgramsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		if v, ok := args["param"]; ok && v != nil {
			queryParams += fmt.Sprintf("&param=%d", int(v.(float64)))
		}

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=browse&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse programs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProgramKanbanTool := mcp.NewTool("get_program_kanban",
		mcp.WithDescription("Get program kanban view"),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
	)

	s.AddTool(getProgramKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams = fmt.Sprintf("browseType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=kanban&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get program kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProgramProductsTool := mcp.NewTool("get_program_products",
		mcp.WithDescription("Get products for a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
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

	s.AddTool(getProgramProductsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("programID=%d", int(args["programID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=product&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get program products: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createProgramTool := mcp.NewTool("create_program",
		mcp.WithDescription("Create a new program"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Program name"),
		),
		mcp.WithString("code",
			mcp.Required(),
			mcp.Description("Program code"),
		),
		mcp.WithNumber("parentProgramID",
			mcp.Description("Parent program ID"),
		),
		mcp.WithNumber("charterID",
			mcp.Description("Charter ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra information"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("End date (YYYY-MM-DD)"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithNumber("budget",
			mcp.Description("Budget"),
		),
	)

	s.AddTool(createProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
			"code": args["code"],
		}

		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["budget"]; ok && v != nil {
			body["budget"] = int(v.(float64))
		}

		queryParams := ""
		if v, ok := args["parentProgramID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&parentProgramID=%d", int(v.(float64)))
		}
		if v, ok := args["charterID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&charterID=%d", int(v.(float64)))
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=create&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editProgramTool := mcp.NewTool("edit_program",
		mcp.WithDescription("Edit an existing program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("name",
			mcp.Description("Program name"),
		),
		mcp.WithString("code",
			mcp.Description("Program code"),
		),
		mcp.WithNumber("parentProgramID",
			mcp.Description("Parent program ID"),
		),
		mcp.WithString("begin",
			mcp.Description("Begin date (YYYY-MM-DD)"),
		),
		mcp.WithString("end",
			mcp.Description("End date (YYYY-MM-DD)"),
		),
		mcp.WithString("desc",
			mcp.Description("Description"),
		),
		mcp.WithNumber("budget",
			mcp.Description("Budget"),
		),
	)

	s.AddTool(editProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["code"]; ok && v != nil {
			body["code"] = v
		}
		if v, ok := args["begin"]; ok && v != nil {
			body["begin"] = v
		}
		if v, ok := args["end"]; ok && v != nil {
			body["end"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["budget"]; ok && v != nil {
			body["budget"] = int(v.(float64))
		}

		queryParams := fmt.Sprintf("&programID=%d", int(args["programID"].(float64)))
		if v, ok := args["parentProgramID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&parentProgramID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=edit&t=json%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeProgramTool := mcp.NewTool("close_program",
		mcp.WithDescription("Close a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("comment",
			mcp.Description("Close comment"),
		),
	)

	s.AddTool(closeProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=close&t=json&programID=%d", int(args["programID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	startProgramTool := mcp.NewTool("start_program",
		mcp.WithDescription("Start a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("comment",
			mcp.Description("Start comment"),
		),
	)

	s.AddTool(startProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=start&t=json&programID=%d", int(args["programID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateProgramTool := mcp.NewTool("activate_program",
		mcp.WithDescription("Activate a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("comment",
			mcp.Description("Activate comment"),
		),
	)

	s.AddTool(activateProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=activate&t=json&programID=%d", int(args["programID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	suspendProgramTool := mcp.NewTool("suspend_program",
		mcp.WithDescription("Suspend a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("comment",
			mcp.Description("Suspend comment"),
		),
	)

	s.AddTool(suspendProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["comment"]; ok && v != nil {
			body["comment"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=suspend&t=json&programID=%d", int(args["programID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to suspend program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteProgramTool := mcp.NewTool("delete_program",
		mcp.WithDescription("Delete a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("confirm",
			mcp.Description("Confirmation"),
			mcp.Enum("yes", "no"),
		),
	)

	s.AddTool(deleteProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		programID := int(args["programID"].(float64))

		queryParams := ""
		if v, ok := args["confirm"]; ok && v != nil {
			queryParams = fmt.Sprintf("&confirm=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=delete&t=json&programID=%d%s", programID, queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProgramProjectsTool := mcp.NewTool("get_program_projects",
		mcp.WithDescription("Get projects for a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
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

	s.AddTool(getProgramProjectsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("programID=%d", int(args["programID"].(float64)))
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=project&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get program projects: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getProgramStakeholdersTool := mcp.NewTool("get_program_stakeholders",
		mcp.WithDescription("Get stakeholders for a program"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
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

	s.AddTool(getProgramStakeholdersTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("programID=%d", int(args["programID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=stakeholder&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get program stakeholders: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewProgramTool := mcp.NewTool("view_program",
		mcp.WithDescription("View program details"),
		mcp.WithNumber("programID",
			mcp.Required(),
			mcp.Description("Program ID"),
		),
	)

	s.AddTool(viewProgramTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		programID := int(args["programID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=program&f=view&t=json&programID=%d", programID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view program: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	exportProgramsTool := mcp.NewTool("export_programs",
		mcp.WithDescription("Export programs to file"),
		mcp.WithString("status",
			mcp.Description("Program status filter"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
	)

	s.AddTool(exportProgramsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=program&f=export&t=json%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to export programs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	refreshProgramStatsTool := mcp.NewTool("refresh_program_stats",
		mcp.WithDescription("Refresh program statistics"),
	)

	s.AddTool(refreshProgramStatsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=program&f=refreshStats&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to refresh program stats: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}