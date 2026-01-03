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

func RegisterCaseLibResources(s *server.MCPServer, client *client.ZenTaoClient) {
	// Case library list resource
	caseLibListResource := mcp.NewResource(
		"zentao://caselibs",
		"ZenTao Case Library List",
		mcp.WithResourceDescription("List of all case libraries in ZenTao"),
		mcp.WithMIMEType("application/json"),
	)

	s.AddResource(caseLibListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		resp, err := client.Get("/index.php?m=caselib&f=index&t=json")
		if err != nil {
			return nil, fmt.Errorf("failed to get case libraries: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      "zentao://caselibs",
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Individual case library resource
	caseLibResource := mcp.NewResourceTemplate(
		"zentao://caselibs/{libID}",
		"ZenTao Case Library",
	)

	s.AddResourceTemplate(caseLibResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		libID := extractIDFromURI(request.Params.URI, "caselibs")
		if libID == "" {
			return nil, fmt.Errorf("invalid case library URI: %s", request.Params.URI)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=view&t=json&libID=%s", libID))
		if err != nil {
			return nil, fmt.Errorf("failed to get case library: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Case library cases resource
	caseLibCasesResource := mcp.NewResourceTemplate(
		"zentao://caselibs/{libID}/cases",
		"ZenTao Case Library Cases",
	)

	s.AddResourceTemplate(caseLibCasesResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		libID := extractIDFromURI(request.Params.URI, "caselibs")
		if libID == "" {
			return nil, fmt.Errorf("invalid case library URI: %s", request.Params.URI)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=browse&t=json&libID=%s", libID))
		if err != nil {
			return nil, fmt.Errorf("failed to get case library cases: %w", err)
		}
		return []mcp.ResourceContents{
			mcp.TextResourceContents{
				URI:      request.Params.URI,
				MIMEType: "application/json",
				Text:     string(resp),
			},
		}, nil
	})

	// Individual case in library resource
	caseLibCaseResource := mcp.NewResourceTemplate(
		"zentao://caselibs/{libID}/cases/{caseID}",
		"ZenTao Case Library Case",
	)

	s.AddResourceTemplate(caseLibCaseResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
		re := regexp.MustCompile(`caselibs/([^/]+)/cases/([^/]+)`)
		matches := re.FindStringSubmatch(request.Params.URI)
		if len(matches) < 3 {
			return nil, fmt.Errorf("invalid case library case URI: %s", request.Params.URI)
		}
		caseID := matches[2]

		resp, err := client.Get(fmt.Sprintf("/index.php?m=caselib&f=viewCase&t=json&caseID=%s", caseID))
		if err != nil {
			return nil, fmt.Errorf("failed to get case library case: %w", err)
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

