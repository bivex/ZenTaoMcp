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
	"regexp"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterTestReportResources(s *server.MCPServer, client *client.ZenTaoClient) {
	testReportResource := mcp.NewResourceTemplate(
		"zentao://testreports/{reportID}",
		"ZenTao Test Report",
	)

	s.AddResourceTemplate(testReportResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		reportID := extractIDFromURI(request.Params.URI, "testreports")
		if reportID == "" {
			return nil, fmt.Errorf("invalid test report URI: %s", request.Params.URI)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testreport&f=view&t=json&reportID=%s", reportID))
		if err != nil {
			return nil, fmt.Errorf("failed to get test report: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	objectTestReportsResource := mcp.NewResourceTemplate(
		"zentao://{objectType}/{objectID}/testreports",
		"ZenTao Object Test Reports",
	)

	s.AddResourceTemplate(objectTestReportsResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		// Parse URI like zentao://projects/123/testreports or zentao://executions/456/testreports
		re := regexp.MustCompile(`(projects|executions|products)/([^/]+)/testreports`)
		matches := re.FindStringSubmatch(request.Params.URI)
		if len(matches) < 3 {
			return nil, fmt.Errorf("invalid test reports URI: %s", request.Params.URI)
		}
		objectType := matches[1]
		objectID := matches[2]

		resp, err := client.Post(fmt.Sprintf("/index.php?m=testreport&f=browse&t=json&objectID=%s&objectType=%s", objectID, objectType), nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get test reports: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})
}

