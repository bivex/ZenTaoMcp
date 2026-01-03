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

func RegisterZanodeTools(s *server.MCPServer, client *client.ZenTaoClient) {
	getInstructionsTool := mcp.NewTool("get_zanode_instructions",
		mcp.WithDescription("Get instructions for ZenTao Node management"),
	)

	s.AddTool(getInstructionsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=zanode&f=instruction&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get zanode instructions: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseNodesTool := mcp.NewTool("browse_zanodes",
		mcp.WithDescription("Browse ZenTao nodes with filtering and pagination"),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithString("param",
			mcp.Description("Additional filter parameter"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID for pagination"),
		),
	)

	s.AddTool(browseNodesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=browse&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse nodes: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getNodeListTool := mcp.NewTool("get_zanode_list",
		mcp.WithDescription("Get list of nodes for a specific host"),
		mcp.WithNumber("hostID",
			mcp.Required(),
			mcp.Description("Host ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
	)

	s.AddTool(getNodeListTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&hostID=%d", int(args["hostID"].(float64)))
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=nodeList&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get node list: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createNodeTool := mcp.NewTool("create_zanode",
		mcp.WithDescription("Create a new ZenTao node"),
		mcp.WithNumber("hostID",
			mcp.Required(),
			mcp.Description("Host ID"),
		),
		mcp.WithString("node_data",
			mcp.Required(),
			mcp.Description("Node configuration data as JSON string"),
		),
	)

	s.AddTool(createNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"node_data": args["node_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=create&t=json&hostID=%d", int(args["hostID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editNodeTool := mcp.NewTool("edit_zanode",
		mcp.WithDescription("Edit an existing ZenTao node"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithString("node_data",
			mcp.Required(),
			mcp.Description("Updated node configuration data as JSON string"),
		),
	)

	s.AddTool(editNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"node_data": args["node_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=edit&t=json&id=%d", int(args["id"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewNodeTool := mcp.NewTool("view_zanode",
		mcp.WithDescription("View details of a specific ZenTao node"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(viewNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=view&t=json&id=%d", int(args["id"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	startNodeTool := mcp.NewTool("start_zanode",
		mcp.WithDescription("Start a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(startNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=start&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to start node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	closeNodeTool := mcp.NewTool("close_zanode",
		mcp.WithDescription("Close a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(closeNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=close&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to close node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	suspendNodeTool := mcp.NewTool("suspend_zanode",
		mcp.WithDescription("Suspend a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(suspendNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=suspend&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to suspend node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	rebootNodeTool := mcp.NewTool("reboot_zanode",
		mcp.WithDescription("Reboot a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(rebootNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=reboot&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to reboot node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	resumeNodeTool := mcp.NewTool("resume_zanode",
		mcp.WithDescription("Resume a suspended ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(resumeNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=resume&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to resume node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	destroyNodeTool := mcp.NewTool("destroy_zanode",
		mcp.WithDescription("Destroy a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(destroyNodeTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=destroy&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to destroy node: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getVNCTool := mcp.NewTool("get_zanode_vnc",
		mcp.WithDescription("Get VNC access information for a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
	)

	s.AddTool(getVNCTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=getVNC&t=json&nodeID=%d", int(args["nodeID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get VNC info: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createImageTool := mcp.NewTool("create_zanode_image",
		mcp.WithDescription("Create an image from a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithString("image_data",
			mcp.Description("Image configuration data as JSON string"),
		),
	)

	s.AddTool(createImageTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["image_data"]; ok && v != nil {
			body["image_data"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=createImage&t=json&nodeID=%d", int(args["nodeID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create image: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getImagesTool := mcp.NewTool("get_zanode_images",
		mcp.WithDescription("Get available images for a host"),
		mcp.WithNumber("hostID",
			mcp.Required(),
			mcp.Description("Host ID"),
		),
	)

	s.AddTool(getImagesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetImages&t=json&hostID=%d", int(args["hostID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get images: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getImageTool := mcp.NewTool("get_zanode_image",
		mcp.WithDescription("Get details of a specific image"),
		mcp.WithNumber("imageID",
			mcp.Required(),
			mcp.Description("Image ID"),
		),
	)

	s.AddTool(getImageTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetImage&t=json&imageID=%d", int(args["imageID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get image: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	updateImageTool := mcp.NewTool("update_zanode_image",
		mcp.WithDescription("Update an image configuration"),
		mcp.WithNumber("imageID",
			mcp.Required(),
			mcp.Description("Image ID"),
		),
		mcp.WithString("image_data",
			mcp.Required(),
			mcp.Description("Updated image configuration data as JSON string"),
		),
	)

	s.AddTool(updateImageTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"image_data": args["image_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=ajaxUpdateImage&t=json&imageID=%d", int(args["imageID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to update image: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	createSnapshotTool := mcp.NewTool("create_zanode_snapshot",
		mcp.WithDescription("Create a snapshot of a ZenTao node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithString("snapshot_data",
			mcp.Description("Snapshot configuration data as JSON string"),
		),
	)

	s.AddTool(createSnapshotTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["snapshot_data"]; ok && v != nil {
			body["snapshot_data"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=createSnapshot&t=json&nodeID=%d", int(args["nodeID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create snapshot: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editSnapshotTool := mcp.NewTool("edit_zanode_snapshot",
		mcp.WithDescription("Edit a snapshot configuration"),
		mcp.WithNumber("snapshotID",
			mcp.Required(),
			mcp.Description("Snapshot ID"),
		),
		mcp.WithString("snapshot_data",
			mcp.Required(),
			mcp.Description("Updated snapshot configuration data as JSON string"),
		),
	)

	s.AddTool(editSnapshotTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"snapshot_data": args["snapshot_data"],
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=editSnapshot&t=json&snapshotID=%d", int(args["snapshotID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit snapshot: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteSnapshotTool := mcp.NewTool("delete_zanode_snapshot",
		mcp.WithDescription("Delete a snapshot"),
		mcp.WithNumber("snapshotID",
			mcp.Required(),
			mcp.Description("Snapshot ID"),
		),
	)

	s.AddTool(deleteSnapshotTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=deleteSnapshot&t=json&snapshotID=%d", int(args["snapshotID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete snapshot: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseSnapshotsTool := mcp.NewTool("browse_zanode_snapshots",
		mcp.WithDescription("Browse snapshots for a node with pagination"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type filter"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Sort order"),
		),
		mcp.WithNumber("recTotal",
			mcp.Description("Total records"),
		),
		mcp.WithNumber("recPerPage",
			mcp.Description("Records per page"),
		),
		mcp.WithNumber("pageID",
			mcp.Description("Page ID for pagination"),
		),
	)

	s.AddTool(browseSnapshotsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&nodeID=%d", int(args["nodeID"].(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=browseSnapshot&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse snapshots: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	restoreSnapshotTool := mcp.NewTool("restore_zanode_snapshot",
		mcp.WithDescription("Restore a node from a snapshot"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithNumber("snapshotID",
			mcp.Required(),
			mcp.Description("Snapshot ID"),
		),
	)

	s.AddTool(restoreSnapshotTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=restoreSnapshot&t=json&nodeID=%d&snapshotID=%d", int(args["nodeID"].(float64)), int(args["snapshotID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to restore snapshot: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getTaskStatusTool := mcp.NewTool("get_zanode_task_status",
		mcp.WithDescription("Get task status for a node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithNumber("taskID",
			mcp.Description("Task ID"),
		),
		mcp.WithString("type",
			mcp.Description("Task type"),
		),
		mcp.WithString("status",
			mcp.Description("Task status filter"),
		),
	)

	s.AddTool(getTaskStatusTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := fmt.Sprintf("&nodeID=%d", int(args["nodeID"].(float64)))
		if v, ok := args["taskID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&taskID=%d", int(v.(float64)))
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["status"]; ok && v != nil {
			queryParams += fmt.Sprintf("&status=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetTaskStatus&t=json%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get task status: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getServiceStatusTool := mcp.NewTool("get_zanode_service_status",
		mcp.WithDescription("Get service status for a host"),
		mcp.WithNumber("hostID",
			mcp.Required(),
			mcp.Description("Host ID"),
		),
	)

	s.AddTool(getServiceStatusTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetServiceStatus&t=json&hostID=%d", int(args["hostID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get service status: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	installServiceTool := mcp.NewTool("install_zanode_service",
		mcp.WithDescription("Install a service on a node"),
		mcp.WithNumber("nodeID",
			mcp.Required(),
			mcp.Description("Node ID"),
		),
		mcp.WithString("service",
			mcp.Required(),
			mcp.Description("Service name"),
		),
	)

	s.AddTool(installServiceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxInstallService&t=json&nodeID=%d&service=%s", int(args["nodeID"].(float64)), args["service"]))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to install service: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getZTFScriptTool := mcp.NewTool("get_zanode_ztf_script",
		mcp.WithDescription("Get ZTF script for a node"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Script type"),
		),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
	)

	s.AddTool(getZTFScriptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetZTFScript&t=json&type=%s&objectID=%d", args["type"], int(args["objectID"].(float64))))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get ZTF script: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	runZTFScriptTool := mcp.NewTool("run_zanode_ztf_script",
		mcp.WithDescription("Run a ZTF script"),
		mcp.WithNumber("scriptID",
			mcp.Required(),
			mcp.Description("Script ID"),
		),
		mcp.WithString("script_data",
			mcp.Description("Script execution data as JSON string"),
		),
	)

	s.AddTool(runZTFScriptTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}
		if v, ok := args["script_data"]; ok && v != nil {
			body["script_data"] = v
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=zanode&f=ajaxRunZTFScript&t=json&scriptID=%d", int(args["scriptID"].(float64))), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to run ZTF script: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	getNodesTool := mcp.NewTool("get_zanodes",
		mcp.WithDescription("Get all available nodes"),
	)

	s.AddTool(getNodesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		resp, err := client.Get("/index.php?m=zanode&f=ajaxGetNodes&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get nodes: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
