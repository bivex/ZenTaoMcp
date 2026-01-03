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

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterTestSuiteResources(s *server.MCPServer, client *client.ZenTaoClient) {
	testSuiteResource := mcp.NewResourceTemplate(
		"zentao://testsuites/{suiteID}",
		"ZenTao Test Suite",
	)

	s.AddResourceTemplate(testSuiteResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		suiteID := extractIDFromURI(request.Params.URI, "testsuites")
		if suiteID == "" {
			return nil, fmt.Errorf("invalid test suite URI: %s", request.Params.URI)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=view&t=json&suiteID=%s", suiteID))
		if err != nil {
			return nil, fmt.Errorf("failed to get test suite: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	productTestSuitesResource := mcp.NewResourceTemplate(
		"zentao://products/{productID}/testsuites",
		"ZenTao Product Test Suites",
	)

	s.AddResourceTemplate(productTestSuitesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		productID := extractIDFromURI(request.Params.URI, "products")
		if productID == "" {
			return nil, fmt.Errorf("invalid product test suites URI: %s", request.Params.URI)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=testsuite&f=browse&t=json&productID=%s", productID))
		if err != nil {
			return nil, fmt.Errorf("failed to get product test suites: %w", err)
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

