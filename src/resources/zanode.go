// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
//
// Licensed under the MIT License.
// Commercial licensing available upon request.

package resources

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterZanodeResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Register zanode instructions resource
	s.AddResource(mcp.NewResource(
		"zentao://zanode/instructions",
		"ZenTao Node Instructions",
		mcp.WithResourceDescription("Instructions for ZenTao Node management"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zanode&f=instruction&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get zanode instructions: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zanode/instructions",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register nodes list resource
	s.AddResource(mcp.NewResource(
		"zentao://zanode/nodes",
		"ZenTao Nodes List",
		mcp.WithResourceDescription("List of all available ZenTao nodes"),
		mcp.WithMIMEType("application/json"),
	), func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=zanode&f=ajaxGetNodes&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get nodes: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://zanode/nodes",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Register node details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/nodes/{id}", "ZenTao Node Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "zanode" || parts[1] != "nodes" {
				return nil, fmt.Errorf("invalid node URI format: %s", uri)
			}
			nodeID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=view&t=json&id=%s", nodeID))
			if err != nil {
				return nil, fmt.Errorf("failed to get node %s: %w", nodeID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register node VNC resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/nodes/{id}/vnc", "ZenTao Node VNC Access"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "nodes" || parts[3] != "vnc" {
				return nil, fmt.Errorf("invalid VNC URI format: %s", uri)
			}
			nodeID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=getVNC&t=json&nodeID=%s", nodeID))
			if err != nil {
				return nil, fmt.Errorf("failed to get VNC for node %s: %w", nodeID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register node snapshots resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/nodes/{id}/snapshots", "ZenTao Node Snapshots"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "nodes" || parts[3] != "snapshots" {
				return nil, fmt.Errorf("invalid snapshots URI format: %s", uri)
			}
			nodeID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=browseSnapshot&t=json&nodeID=%s", nodeID))
			if err != nil {
				return nil, fmt.Errorf("failed to get snapshots for node %s: %w", nodeID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register node task status resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/nodes/{id}/tasks", "ZenTao Node Tasks"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "nodes" || parts[3] != "tasks" {
				return nil, fmt.Errorf("invalid tasks URI format: %s", uri)
			}
			nodeID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetTaskStatus&t=json&nodeID=%s", nodeID))
			if err != nil {
				return nil, fmt.Errorf("failed to get tasks for node %s: %w", nodeID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register host node list resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/hosts/{hostID}/nodes", "Host Nodes List"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "hosts" || parts[3] != "nodes" {
				return nil, fmt.Errorf("invalid host nodes URI format: %s", uri)
			}
			hostID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=nodeList&t=json&hostID=%s", hostID))
			if err != nil {
				return nil, fmt.Errorf("failed to get nodes for host %s: %w", hostID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register host images resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/hosts/{hostID}/images", "Host Images List"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "hosts" || parts[3] != "images" {
				return nil, fmt.Errorf("invalid host images URI format: %s", uri)
			}
			hostID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetImages&t=json&hostID=%s", hostID))
			if err != nil {
				return nil, fmt.Errorf("failed to get images for host %s: %w", hostID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register host service status resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/hosts/{hostID}/services", "Host Service Status"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 4 || parts[0] != "zanode" || parts[1] != "hosts" || parts[3] != "services" {
				return nil, fmt.Errorf("invalid host services URI format: %s", uri)
			}
			hostID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetServiceStatus&t=json&hostID=%s", hostID))
			if err != nil {
				return nil, fmt.Errorf("failed to get service status for host %s: %w", hostID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)

	// Register image details resource template
	s.AddResourceTemplate(
		mcp.NewResourceTemplate("zentao://zanode/images/{imageID}", "ZenTao Node Image Details"),
		func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
			uri := request.Params.URI
			parts := strings.Split(strings.TrimPrefix(uri, "zentao://"), "/")
			if len(parts) < 3 || parts[0] != "zanode" || parts[1] != "images" {
				return nil, fmt.Errorf("invalid image URI format: %s", uri)
			}
			imageID := parts[2]

			resp, err := client.Get(fmt.Sprintf("/index.php?m=zanode&f=ajaxGetImage&t=json&imageID=%s", imageID))
			if err != nil {
				return nil, fmt.Errorf("failed to get image %s: %w", imageID, err)
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      request.Params.URI,
					MIMEType: "application/json",
					Text:     string(resp),
				},
			}, nil
		},
	)
}
