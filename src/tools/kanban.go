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

func RegisterKanbanTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Kanban Space Management
	getKanbanSpacesTool := mcp.NewTool("get_kanban_spaces",
		mcp.WithDescription("Get kanban spaces"),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
			mcp.Enum("involved", "cooperation", "public", "private"),
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

	s.AddTool(getKanbanSpacesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=space&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get kanban spaces: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createKanbanSpaceTool := mcp.NewTool("create_kanban_space",
		mcp.WithDescription("Create a new kanban space"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Space type"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Space name"),
		),
		mcp.WithString("desc",
			mcp.Description("Space description"),
		),
	)

	s.AddTool(createKanbanSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
			"type": args["type"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=createSpace&t=json&type=%s", args["type"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban space: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editKanbanSpaceTool := mcp.NewTool("edit_kanban_space",
		mcp.WithDescription("Edit a kanban space"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
		mcp.WithString("name",
			mcp.Description("Space name"),
		),
		mcp.WithString("desc",
			mcp.Description("Space description"),
		),
	)

	s.AddTool(editKanbanSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=editSpace&t=json&spaceID=%d", int(args["spaceID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit kanban space: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateKanbanSpaceTool := mcp.NewTool("activate_kanban_space",
		mcp.WithDescription("Activate a kanban space"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
	)

	s.AddTool(activateKanbanSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=activateSpace&t=json&spaceID=%d", int(args["spaceID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate kanban space: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	closeKanbanSpaceTool := mcp.NewTool("close_kanban_space",
		mcp.WithDescription("Close a kanban space"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
	)

	s.AddTool(closeKanbanSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=closeSpace&t=json&spaceID=%d", int(args["spaceID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close kanban space: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanSpaceTool := mcp.NewTool("delete_kanban_space",
		mcp.WithDescription("Delete a kanban space"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
	)

	s.AddTool(deleteKanbanSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=deleteSpace&t=json&spaceID=%d", int(args["spaceID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban space: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Kanban Board Management
	createKanbanTool := mcp.NewTool("create_kanban",
		mcp.WithDescription("Create a new kanban board"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
		mcp.WithString("type",
			mcp.Description("Kanban type"),
		),
		mcp.WithNumber("copyKanbanID",
			mcp.Description("Copy from kanban ID"),
		),
		mcp.WithString("extra",
			mcp.Description("Extra parameters"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Kanban name"),
		),
		mcp.WithString("desc",
			mcp.Description("Kanban description"),
		),
	)

	s.AddTool(createKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("spaceID=%d", int(args["spaceID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["copyKanbanID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&copyKanbanID=%d", int(v.(float64)))
		}
		if v, ok := args["extra"]; ok && v != nil {
			queryParams += fmt.Sprintf("&extra=%s", v)
		}

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=create&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editKanbanTool := mcp.NewTool("edit_kanban",
		mcp.WithDescription("Edit a kanban board"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithString("name",
			mcp.Description("Kanban name"),
		),
		mcp.WithString("desc",
			mcp.Description("Kanban description"),
		),
	)

	s.AddTool(editKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=edit&t=json&kanbanID=%d", int(args["kanbanID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewKanbanTool := mcp.NewTool("view_kanban",
		mcp.WithDescription("View a kanban board"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithString("regionID",
			mcp.Description("Region ID"),
		),
	)

	s.AddTool(viewKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("kanbanID=%d", int(args["kanbanID"].(float64)))
		if v, ok := args["regionID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&regionID=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanTool := mcp.NewTool("delete_kanban",
		mcp.WithDescription("Delete a kanban board"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
	)

	s.AddTool(deleteKanbanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=delete&t=json&kanbanID=%d", int(args["kanbanID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Kanban Region Management
	createKanbanRegionTool := mcp.NewTool("create_kanban_region",
		mcp.WithDescription("Create a kanban region"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("kanban", "execution"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Region name"),
		),
		mcp.WithString("desc",
			mcp.Description("Region description"),
		),
	)

	s.AddTool(createKanbanRegionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("kanbanID=%d", int(args["kanbanID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=createRegion&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban region: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editKanbanRegionTool := mcp.NewTool("edit_kanban_region",
		mcp.WithDescription("Edit a kanban region"),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithString("name",
			mcp.Description("Region name"),
		),
		mcp.WithString("desc",
			mcp.Description("Region description"),
		),
	)

	s.AddTool(editKanbanRegionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=editRegion&t=json&regionID=%d", int(args["regionID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit kanban region: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanRegionTool := mcp.NewTool("delete_kanban_region",
		mcp.WithDescription("Delete a kanban region"),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
	)

	s.AddTool(deleteKanbanRegionTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=deleteRegion&t=json&regionID=%d", int(args["regionID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban region: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Kanban Lane Management
	createKanbanLaneTool := mcp.NewTool("create_kanban_lane",
		mcp.WithDescription("Create a kanban lane"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithString("from",
			mcp.Description("Source"),
			mcp.Enum("kanban", "execution"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Lane name"),
		),
		mcp.WithString("color",
			mcp.Description("Lane color"),
		),
	)

	s.AddTool(createKanbanLaneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("kanbanID=%d&regionID=%d", int(args["kanbanID"].(float64)), int(args["regionID"].(float64)))
		if v, ok := args["from"]; ok && v != nil {
			queryParams += fmt.Sprintf("&from=%s", v)
		}

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["color"]; ok && v != nil {
			body["color"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=createLane&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban lane: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanLaneTool := mcp.NewTool("delete_kanban_lane",
		mcp.WithDescription("Delete a kanban lane"),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithNumber("laneID",
			mcp.Required(),
			mcp.Description("Lane ID"),
		),
	)

	s.AddTool(deleteKanbanLaneTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=deleteLane&t=json&regionID=%d&laneID=%d",
			int(args["regionID"].(float64)), int(args["laneID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban lane: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Kanban Column Management
	createKanbanColumnTool := mcp.NewTool("create_kanban_column",
		mcp.WithDescription("Create a kanban column"),
		mcp.WithNumber("fromColumnID",
			mcp.Required(),
			mcp.Description("Source column ID"),
		),
		mcp.WithString("position",
			mcp.Required(),
			mcp.Description("Position"),
			mcp.Enum("left", "right"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Column name"),
		),
		mcp.WithNumber("limit",
			mcp.Description("WIP limit"),
		),
	)

	s.AddTool(createKanbanColumnTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name":     args["name"],
			"position": args["position"],
		}

		if v, ok := args["limit"]; ok && v != nil {
			body["limit"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=createColumn&t=json&fromColumnID=%d&position=%s",
			int(args["fromColumnID"].(float64)), args["position"]), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban column: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	splitKanbanColumnTool := mcp.NewTool("split_kanban_column",
		mcp.WithDescription("Split a kanban column"),
		mcp.WithNumber("columnID",
			mcp.Required(),
			mcp.Description("Column ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("New column name"),
		),
	)

	s.AddTool(splitKanbanColumnTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=splitColumn&t=json&columnID=%d", int(args["columnID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to split kanban column: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	archiveKanbanColumnTool := mcp.NewTool("archive_kanban_column",
		mcp.WithDescription("Archive a kanban column"),
		mcp.WithNumber("columnID",
			mcp.Required(),
			mcp.Description("Column ID"),
		),
	)

	s.AddTool(archiveKanbanColumnTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=archiveColumn&t=json&columnID=%d", int(args["columnID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to archive kanban column: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanColumnTool := mcp.NewTool("delete_kanban_column",
		mcp.WithDescription("Delete a kanban column"),
		mcp.WithNumber("columnID",
			mcp.Required(),
			mcp.Description("Column ID"),
		),
	)

	s.AddTool(deleteKanbanColumnTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=deleteColumn&t=json&columnID=%d", int(args["columnID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban column: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Kanban Card Management
	createKanbanCardTool := mcp.NewTool("create_kanban_card",
		mcp.WithDescription("Create a kanban card"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithNumber("groupID",
			mcp.Required(),
			mcp.Description("Group ID"),
		),
		mcp.WithNumber("columnID",
			mcp.Required(),
			mcp.Description("Column ID"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Card name"),
		),
		mcp.WithString("desc",
			mcp.Description("Card description"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithString("color",
			mcp.Description("Card color"),
		),
	)

	s.AddTool(createKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"name": args["name"],
		}

		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["color"]; ok && v != nil {
			body["color"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=createCard&t=json&kanbanID=%d&regionID=%d&groupID=%d&columnID=%d",
			int(args["kanbanID"].(float64)), int(args["regionID"].(float64)), int(args["groupID"].(float64)), int(args["columnID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editKanbanCardTool := mcp.NewTool("edit_kanban_card",
		mcp.WithDescription("Edit a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
		mcp.WithString("name",
			mcp.Description("Card name"),
		),
		mcp.WithString("desc",
			mcp.Description("Card description"),
		),
		mcp.WithNumber("assignedTo",
			mcp.Description("Assigned to user ID"),
		),
		mcp.WithNumber("pri",
			mcp.Description("Priority"),
		),
		mcp.WithString("color",
			mcp.Description("Card color"),
		),
	)

	s.AddTool(editKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["assignedTo"]; ok && v != nil {
			body["assignedTo"] = int(v.(float64))
		}
		if v, ok := args["pri"]; ok && v != nil {
			body["pri"] = int(v.(float64))
		}
		if v, ok := args["color"]; ok && v != nil {
			body["color"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=editCard&t=json&cardID=%d", int(args["cardID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	viewKanbanCardTool := mcp.NewTool("view_kanban_card",
		mcp.WithDescription("View a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
	)

	s.AddTool(viewKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=viewCard&t=json&cardID=%d", int(args["cardID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	moveKanbanCardTool := mcp.NewTool("move_kanban_card",
		mcp.WithDescription("Move a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
		mcp.WithNumber("fromColID",
			mcp.Required(),
			mcp.Description("From column ID"),
		),
		mcp.WithNumber("toColID",
			mcp.Required(),
			mcp.Description("To column ID"),
		),
		mcp.WithNumber("fromLaneID",
			mcp.Required(),
			mcp.Description("From lane ID"),
		),
		mcp.WithNumber("toLaneID",
			mcp.Required(),
			mcp.Description("To lane ID"),
		),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithString("showModal",
			mcp.Description("Show modal"),
		),
	)

	s.AddTool(moveKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		queryParams := fmt.Sprintf("cardID=%d&fromColID=%d&toColID=%d&fromLaneID=%d&toLaneID=%d&kanbanID=%d",
			int(args["cardID"].(float64)), int(args["fromColID"].(float64)), int(args["toColID"].(float64)),
			int(args["fromLaneID"].(float64)), int(args["toLaneID"].(float64)), int(args["kanbanID"].(float64)))

		if v, ok := args["showModal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&showModal=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=moveCard&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to move kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	finishKanbanCardTool := mcp.NewTool("finish_kanban_card",
		mcp.WithDescription("Finish a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
	)

	s.AddTool(finishKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=finishCard&t=json&cardID=%d", int(args["cardID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to finish kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	activateKanbanCardTool := mcp.NewTool("activate_kanban_card",
		mcp.WithDescription("Activate a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
	)

	s.AddTool(activateKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=activateCard&t=json&cardID=%d", int(args["cardID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to activate kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	archiveKanbanCardTool := mcp.NewTool("archive_kanban_card",
		mcp.WithDescription("Archive a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
	)

	s.AddTool(archiveKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=archiveCard&t=json&cardID=%d", int(args["cardID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to archive kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteKanbanCardTool := mcp.NewTool("delete_kanban_card",
		mcp.WithDescription("Delete a kanban card"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
	)

	s.AddTool(deleteKanbanCardTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=deleteCard&t=json&cardID=%d", int(args["cardID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete kanban card: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// Import tools - focused on key ones
	importCardsFromPlanTool := mcp.NewTool("import_cards_from_plan",
		mcp.WithDescription("Import cards from a plan"),
		mcp.WithNumber("kanbanID",
			mcp.Required(),
			mcp.Description("Kanban ID"),
		),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithNumber("groupID",
			mcp.Required(),
			mcp.Description("Group ID"),
		),
		mcp.WithNumber("columnID",
			mcp.Required(),
			mcp.Description("Column ID"),
		),
		mcp.WithNumber("selectedProductID",
			mcp.Required(),
			mcp.Description("Selected product ID"),
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

	s.AddTool(importCardsFromPlanTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("kanbanID=%d&regionID=%d&groupID=%d&columnID=%d&selectedProductID=%d",
			int(args["kanbanID"].(float64)), int(args["regionID"].(float64)), int(args["groupID"].(float64)),
			int(args["columnID"].(float64)), int(args["selectedProductID"].(float64)))

		if v, ok := args["recTotal"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recTotal=%d", int(v.(float64)))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			queryParams += fmt.Sprintf("&recPerPage=%d", int(v.(float64)))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageID=%d", int(v.(float64)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=kanban&f=importPlan&t=json&%s", queryParams), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to import cards from plan: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	setCardColorTool := mcp.NewTool("set_kanban_card_color",
		mcp.WithDescription("Set kanban card color"),
		mcp.WithNumber("cardID",
			mcp.Required(),
			mcp.Description("Card ID"),
		),
		mcp.WithString("color",
			mcp.Required(),
			mcp.Description("Color"),
		),
	)

	s.AddTool(setCardColorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=setCardColor&t=json&cardID=%d&color=%s",
			int(args["cardID"].(float64)), args["color"]))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to set kanban card color: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	// AJAX helpers
	getKanbanLanesTool := mcp.NewTool("get_kanban_lanes",
		mcp.WithDescription("Get kanban lanes"),
		mcp.WithNumber("regionID",
			mcp.Required(),
			mcp.Description("Region ID"),
		),
		mcp.WithString("type",
			mcp.Description("Type"),
			mcp.Enum("all", "story", "task", "bug"),
		),
		mcp.WithString("field",
			mcp.Description("Field"),
		),
		mcp.WithString("pageType",
			mcp.Description("Page type"),
		),
	)

	s.AddTool(getKanbanLanesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("regionID=%d", int(args["regionID"].(float64)))
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["field"]; ok && v != nil {
			queryParams += fmt.Sprintf("&field=%s", v)
		}
		if v, ok := args["pageType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&pageType=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=ajaxGetLanes&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get kanban lanes: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getKanbanColumnsTool := mcp.NewTool("get_kanban_columns",
		mcp.WithDescription("Get kanban columns"),
		mcp.WithNumber("laneID",
			mcp.Required(),
			mcp.Description("Lane ID"),
		),
	)

	s.AddTool(getKanbanColumnsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=kanban&f=ajaxGetColumns&t=json&laneID=%d", int(args["laneID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get kanban columns: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
