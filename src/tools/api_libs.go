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
	"net/url"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

func RegisterApiLibTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createApiLibTool := mcp.NewTool("create_api_lib",
		mcp.WithDescription("Create a new API library in ZenTao"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Library type"),
			mcp.Enum("project", "product"),
		),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID (project or product ID)"),
		),
		mcp.WithString("name",
			mcp.Description("Library name"),
		),
		mcp.WithString("desc",
			mcp.Description("Library description"),
		),
	)

	s.AddTool(createApiLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{
			"type":     args["type"],
			"objectID": int(args["objectID"].(float64)),
		}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("type=%s&objectID=%d", args["type"], int(args["objectID"].(float64)))
		if v, ok := args["name"]; ok && v != nil {
			queryParams += fmt.Sprintf("&name=%s", v)
		}
		if v, ok := args["desc"]; ok && v != nil {
			queryParams += fmt.Sprintf("&desc=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=createLib&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create API library: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editApiLibTool := mcp.NewTool("edit_api_lib",
		mcp.WithDescription("Edit an existing API library in ZenTao"),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithString("name",
			mcp.Description("Library name"),
		),
		mcp.WithString("desc",
			mcp.Description("Library description"),
		),
	)

	s.AddTool(editApiLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("id=%d", int(args["id"].(float64)))
		if v, ok := args["name"]; ok && v != nil {
			queryParams += fmt.Sprintf("&name=%s", v)
		}
		if v, ok := args["desc"]; ok && v != nil {
			queryParams += fmt.Sprintf("&desc=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=editLib&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit API library: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteApiLibTool := mcp.NewTool("delete_api_lib",
		mcp.WithDescription("Delete an API library from ZenTao"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID to delete"),
		),
	)

	s.AddTool(deleteApiLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=deleteLib&t=json&libID=%d", libID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete API library: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getApiLibReleasesTool := mcp.NewTool("get_api_lib_releases",
		mcp.WithDescription("Get releases for an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by field"),
		),
	)

	s.AddTool(getApiLibReleasesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		queryParams := fmt.Sprintf("libID=%d", libID)
		if v, ok := args["orderBy"]; ok && v != nil {
			queryParams += fmt.Sprintf("&orderBy=%s", v)
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=releases&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get API library releases: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createApiLibReleaseTool := mcp.NewTool("create_api_lib_release",
		mcp.WithDescription("Create a new release for an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithString("name",
			mcp.Description("Release name"),
		),
		mcp.WithString("desc",
			mcp.Description("Release description"),
		),
		mcp.WithString("version",
			mcp.Description("Release version"),
		),
	)

	s.AddTool(createApiLibReleaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["version"]; ok && v != nil {
			body["version"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("libID=%d", int(args["libID"].(float64)))
		if v, ok := args["name"]; ok && v != nil {
			queryParams += fmt.Sprintf("&name=%s", v)
		}
		if v, ok := args["desc"]; ok && v != nil {
			queryParams += fmt.Sprintf("&desc=%s", v)
		}
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=createRelease&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create API library release: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteApiLibReleaseTool := mcp.NewTool("delete_api_lib_release",
		mcp.WithDescription("Delete a release from an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("id",
			mcp.Required(),
			mcp.Description("Release ID to delete"),
		),
	)

	s.AddTool(deleteApiLibReleaseTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))
		id := int(args["id"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=deleteRelease&t=json&libID=%d&id=%d", libID, id))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete API library release: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getApiLibStructsTool := mcp.NewTool("get_api_lib_structs",
		mcp.WithDescription("Get structures for an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("releaseID",
			mcp.Description("Release ID"),
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

	s.AddTool(getApiLibStructsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		queryParams := fmt.Sprintf("libID=%d", libID)
		if v, ok := args["releaseID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&releaseID=%d", int(v.(float64)))
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

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=struct&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get API library structures: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createApiLibStructTool := mcp.NewTool("create_api_lib_struct",
		mcp.WithDescription("Create a new structure for an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithString("name",
			mcp.Description("Structure name"),
		),
		mcp.WithString("type",
			mcp.Description("Structure type"),
		),
		mcp.WithString("desc",
			mcp.Description("Structure description"),
		),
		mcp.WithString("content",
			mcp.Description("Structure content"),
		),
	)

	s.AddTool(createApiLibStructTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["content"]; ok && v != nil {
			body["content"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("libID=%d", int(args["libID"].(float64)))
		if v, ok := args["name"]; ok && v != nil {
			queryParams += fmt.Sprintf("&name=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["desc"]; ok && v != nil {
			queryParams += fmt.Sprintf("&desc=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=createStruct&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create API library structure: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editApiLibStructTool := mcp.NewTool("edit_api_lib_struct",
		mcp.WithDescription("Edit an existing structure in an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("structID",
			mcp.Required(),
			mcp.Description("Structure ID"),
		),
		mcp.WithString("name",
			mcp.Description("Structure name"),
		),
		mcp.WithString("type",
			mcp.Description("Structure type"),
		),
		mcp.WithString("desc",
			mcp.Description("Structure description"),
		),
		mcp.WithString("content",
			mcp.Description("Structure content"),
		),
	)

	s.AddTool(editApiLibStructTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["name"]; ok && v != nil {
			body["name"] = v
		}
		if v, ok := args["type"]; ok && v != nil {
			body["type"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["content"]; ok && v != nil {
			body["content"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("libID=%d&structID=%d", int(args["libID"].(float64)), int(args["structID"].(float64)))
		if v, ok := args["name"]; ok && v != nil {
			queryParams += fmt.Sprintf("&name=%s", v)
		}
		if v, ok := args["type"]; ok && v != nil {
			queryParams += fmt.Sprintf("&type=%s", v)
		}
		if v, ok := args["desc"]; ok && v != nil {
			queryParams += fmt.Sprintf("&desc=%s", v)
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=editStruct&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit API library structure: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteApiLibStructTool := mcp.NewTool("delete_api_lib_struct",
		mcp.WithDescription("Delete a structure from an API library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("structID",
			mcp.Required(),
			mcp.Description("Structure ID to delete"),
		),
	)

	s.AddTool(deleteApiLibStructTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))
		structID := int(args["structID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=deleteStruct&t=json&libID=%d&structID=%d", libID, structID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete API library structure: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	createApiTool := mcp.NewTool("create_api",
		mcp.WithDescription("Create a new API in ZenTao"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithString("space",
			mcp.Description("Space type"),
			mcp.Enum("api", "project", "product"),
		),
		mcp.WithString("title",
			mcp.Description("API title"),
		),
		mcp.WithString("path",
			mcp.Description("API path"),
		),
		mcp.WithString("method",
			mcp.Description("HTTP method"),
			mcp.Enum("GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"),
		),
		mcp.WithString("requestType",
			mcp.Description("Request type"),
		),
		mcp.WithString("desc",
			mcp.Description("API description"),
		),
		mcp.WithString("params",
			mcp.Description("API parameters"),
		),
		mcp.WithString("response",
			mcp.Description("API response"),
		),
	)

	s.AddTool(createApiTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["path"]; ok && v != nil {
			body["path"] = v
		}
		if v, ok := args["method"]; ok && v != nil {
			body["method"] = v
		}
		if v, ok := args["requestType"]; ok && v != nil {
			body["requestType"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["params"]; ok && v != nil {
			body["params"] = v
		}
		if v, ok := args["response"]; ok && v != nil {
			body["response"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("libID=%d", int(args["libID"].(float64)))
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["space"]; ok && v != nil {
			queryParams += fmt.Sprintf("&space=%s", v)
		}
		if v, ok := args["title"]; ok && v != nil {
			queryParams += fmt.Sprintf("&title=%s", url.QueryEscape(v.(string)))
		}
		if v, ok := args["path"]; ok && v != nil {
			queryParams += fmt.Sprintf("&path=%s", url.QueryEscape(v.(string)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=create&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create API: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	editApiTool := mcp.NewTool("edit_api",
		mcp.WithDescription("Edit an existing API in ZenTao"),
		mcp.WithNumber("apiID",
			mcp.Required(),
			mcp.Description("API ID"),
		),
		mcp.WithString("title",
			mcp.Description("API title"),
		),
		mcp.WithString("path",
			mcp.Description("API path"),
		),
		mcp.WithString("method",
			mcp.Description("HTTP method"),
			mcp.Enum("GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"),
		),
		mcp.WithString("requestType",
			mcp.Description("Request type"),
		),
		mcp.WithString("desc",
			mcp.Description("API description"),
		),
		mcp.WithString("params",
			mcp.Description("API parameters"),
		),
		mcp.WithString("response",
			mcp.Description("API response"),
		),
	)

	s.AddTool(editApiTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		body := map[string]interface{}{}

		if v, ok := args["title"]; ok && v != nil {
			body["title"] = v
		}
		if v, ok := args["path"]; ok && v != nil {
			body["path"] = v
		}
		if v, ok := args["method"]; ok && v != nil {
			body["method"] = v
		}
		if v, ok := args["requestType"]; ok && v != nil {
			body["requestType"] = v
		}
		if v, ok := args["desc"]; ok && v != nil {
			body["desc"] = v
		}
		if v, ok := args["params"]; ok && v != nil {
			body["params"] = v
		}
		if v, ok := args["response"]; ok && v != nil {
			body["response"] = v
		}

		// Build query string for API endpoint
		queryParams := fmt.Sprintf("apiID=%d", int(args["apiID"].(float64)))
		if v, ok := args["title"]; ok && v != nil {
			queryParams += fmt.Sprintf("&title=%s", url.QueryEscape(v.(string)))
		}
		if v, ok := args["path"]; ok && v != nil {
			queryParams += fmt.Sprintf("&path=%s", url.QueryEscape(v.(string)))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=api&f=edit&t=json&%s", queryParams), body)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit API: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteApiTool := mcp.NewTool("delete_api",
		mcp.WithDescription("Delete an API from ZenTao"),
		mcp.WithNumber("apiID",
			mcp.Required(),
			mcp.Description("API ID to delete"),
		),
	)

	s.AddTool(deleteApiTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		apiID := int(args["apiID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=delete&t=json&apiID=%d", apiID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete API: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getApiTool := mcp.NewTool("get_api",
		mcp.WithDescription("Get API details from ZenTao"),
		mcp.WithNumber("libID",
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("apiID",
			mcp.Description("API ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Version"),
		),
		mcp.WithNumber("release",
			mcp.Description("Release"),
		),
	)

	s.AddTool(getApiTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["libID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&libID=%d", int(v.(float64)))
		}
		if v, ok := args["apiID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&apiID=%d", int(v.(float64)))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["release"]; ok && v != nil {
			queryParams += fmt.Sprintf("&release=%d", int(v.(float64)))
		}

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=view&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get API: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})

	getApisTool := mcp.NewTool("get_apis",
		mcp.WithDescription("Get list of APIs from ZenTao"),
		mcp.WithNumber("libID",
			mcp.Description("Library ID"),
		),
		mcp.WithNumber("moduleID",
			mcp.Description("Module ID"),
		),
		mcp.WithNumber("apiID",
			mcp.Description("API ID"),
		),
		mcp.WithNumber("version",
			mcp.Description("Version"),
		),
		mcp.WithNumber("release",
			mcp.Description("Release"),
		),
		mcp.WithString("browseType",
			mcp.Description("Browse type"),
		),
		mcp.WithString("params",
			mcp.Description("Parameters"),
		),
		mcp.WithString("orderBy",
			mcp.Description("Order by"),
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
		mcp.WithString("mode",
			mcp.Description("Mode"),
		),
		mcp.WithString("search",
			mcp.Description("Search term"),
		),
	)

	s.AddTool(getApisTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		queryParams := ""
		if v, ok := args["libID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&libID=%d", int(v.(float64)))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&moduleID=%d", int(v.(float64)))
		}
		if v, ok := args["apiID"]; ok && v != nil {
			queryParams += fmt.Sprintf("&apiID=%d", int(v.(float64)))
		}
		if v, ok := args["version"]; ok && v != nil {
			queryParams += fmt.Sprintf("&version=%d", int(v.(float64)))
		}
		if v, ok := args["release"]; ok && v != nil {
			queryParams += fmt.Sprintf("&release=%d", int(v.(float64)))
		}
		if v, ok := args["browseType"]; ok && v != nil {
			queryParams += fmt.Sprintf("&browseType=%s", v)
		}
		if v, ok := args["params"]; ok && v != nil {
			queryParams += fmt.Sprintf("&params=%s", v)
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
		if v, ok := args["mode"]; ok && v != nil {
			queryParams += fmt.Sprintf("&mode=%s", v)
		}
		if v, ok := args["search"]; ok && v != nil {
			queryParams += fmt.Sprintf("&search=%s", url.QueryEscape(v.(string)))
		}

		if queryParams != "" {
			queryParams = queryParams[1:] // Remove leading &
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=api&f=index&t=json&%s", queryParams))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get APIs: %v", err)), nil
		}

		return mcp.NewToolResultText(string(resp)), nil
	})
}
