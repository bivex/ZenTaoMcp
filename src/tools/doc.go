// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
//
//
// Licensed under MIT License.
// Commercial licensing available upon request.

package tools

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/zentao/mcp-server/client"
)

// RegisterDocTools registers all documentation-related tools
func RegisterDocTools(s *server.MCPServer, client *client.ZenTaoClient) {
	// Space Management
	registerSpaceTools(s, client)
	// Library Management
	registerLibTools(s, client)
	// Document Management
	registerDocTools(s, client)
	// Template Management
	registerTemplateTools(s, client)
	// Browse/View Tools
	registerBrowseTools(s, client)
	// File Management
	registerFileTools(s, client)
	// Catalog Management
	registerCatalogTools(s, client)
}

func registerSpaceTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createSpaceTool := mcp.NewTool("doc_create_space",
		mcp.WithDescription("Create a new documentation space"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Space type"),
		),
	)

	s.AddTool(createSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}

		resp, err := client.Post("/index.php?m=doc&f=createSpace&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editSpaceTool := mcp.NewTool("doc_edit_space",
		mcp.WithDescription("Edit a documentation space"),
		mcp.WithNumber("spaceID",
			mcp.Required(),
			mcp.Description("Space ID"),
		),
	)

	s.AddTool(editSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		spaceID := int(args["spaceID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=editSpace&t=json&spaceID=%d", spaceID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteSpaceTool := mcp.NewTool("doc_delete_space",
		mcp.WithDescription("Delete a documentation space"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
	)

	s.AddTool(deleteSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=deleteSpace&t=json&libID=%d", libID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerLibTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createLibTool := mcp.NewTool("doc_create_lib",
		mcp.WithDescription("Create a new documentation library"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Library type: api|project|product|execution|custom|mine"),
			mcp.Enum("api", "project", "product", "execution", "custom", "mine"),
		),
		mcp.WithNumber("objectID", mcp.Description("Object ID")),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
	)

	s.AddTool(createLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["objectID"]; ok && v != nil {
			params["objectID"] = int(v.(float64))
		}
		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=doc&f=createLib&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editLibTool := mcp.NewTool("doc_edit_lib",
		mcp.WithDescription("Edit a documentation library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
	)

	s.AddTool(editLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=editLib&t=json&libID=%d", libID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteLibTool := mcp.NewTool("doc_delete_lib",
		mcp.WithDescription("Delete a documentation library"),
		mcp.WithNumber("libID",
			mcp.Required(),
			mcp.Description("Library ID"),
		),
	)

	s.AddTool(deleteLibTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		libID := int(args["libID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=deleteLib&t=json&libID=%d", libID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete library: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerDocTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createDocTool := mcp.NewTool("doc_create",
		mcp.WithDescription("Create a new document"),
		mcp.WithString("objectType",
			mcp.Required(),
			mcp.Description("Object type: product|project|execution|custom"),
			mcp.Enum("product", "project", "execution", "custom"),
		),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithString("docType",
			mcp.Required(),
			mcp.Description("Document type: html|word|ppt|excel"),
			mcp.Enum("html", "word", "ppt", "excel"),
		),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithNumber("appendLib", mcp.Description("Append library")),
	)

	s.AddTool(createDocTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["objectType"] = args["objectType"].(string)
		params["objectID"] = int(args["objectID"].(float64))
		params["docType"] = args["docType"].(string)

		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["appendLib"]; ok && v != nil {
			params["appendLib"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=doc&f=create&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create document: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editDocTool := mcp.NewTool("doc_edit",
		mcp.WithDescription("Edit a document"),
		mcp.WithNumber("docID",
			mcp.Required(),
			mcp.Description("Document ID"),
		),
		mcp.WithBoolean("comment", mcp.Description("Include comments")),
		mcp.WithNumber("appendLib", mcp.Description("Append library")),
	)

	s.AddTool(editDocTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docID := int(args["docID"].(float64))

		params := make(map[string]interface{})
		if v, ok := args["comment"]; ok && v != nil {
			params["comment"] = v.(bool)
		}
		if v, ok := args["appendLib"]; ok && v != nil {
			params["appendLib"] = int(v.(float64))
		}

		resp, err := client.Post(fmt.Sprintf("/index.php?m=doc&f=edit&t=json&docID=%d", docID), params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit document: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteDocTool := mcp.NewTool("doc_delete",
		mcp.WithDescription("Delete a document"),
		mcp.WithNumber("docID",
			mcp.Required(),
			mcp.Description("Document ID"),
		),
	)

	s.AddTool(deleteDocTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docID := int(args["docID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=delete&t=json&docID=%d", docID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete document: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	viewDocTool := mcp.NewTool("doc_view",
		mcp.WithDescription("View a document"),
		mcp.WithNumber("docID",
			mcp.Required(),
			mcp.Description("Document ID"),
		),
		mcp.WithNumber("version", mcp.Description("Document version")),
	)

	s.AddTool(viewDocTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docID := args["docID"].(float64)

		url := fmt.Sprintf("/index.php?m=doc&f=view&t=json&docID=%v", int(docID))

		if v, ok := args["version"]; ok && v != nil {
			url += fmt.Sprintf("&version=%d", int(v.(float64)))
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to view document: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	uploadDocsTool := mcp.NewTool("doc_upload_docs",
		mcp.WithDescription("Upload documents"),
		mcp.WithString("objectType",
			mcp.Required(),
			mcp.Description("Object type: product|project|execution|custom"),
			mcp.Enum("product", "project", "execution", "custom"),
		),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithString("docType",
			mcp.Required(),
			mcp.Description("Document type: html|word|ppt|excel|attachment"),
			mcp.Enum("html", "word", "ppt", "excel", "attachment"),
		),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
	)

	s.AddTool(uploadDocsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["objectType"] = args["objectType"].(string)
		params["objectID"] = int(args["objectID"].(float64))
		params["docType"] = args["docType"].(string)

		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}

		resp, err := client.Post("/index.php?m=doc&f=uploadDocs&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to upload documents: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerTemplateTools(s *server.MCPServer, client *client.ZenTaoClient) {
	createTemplateTool := mcp.NewTool("doc_create_template",
		mcp.WithDescription("Create a new document template"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
	)

	s.AddTool(createTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		params["moduleID"] = int(args["moduleID"].(float64))

		resp, err := client.Post("/index.php?m=doc&f=createTemplate&t=json", params)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to create template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	editTemplateTool := mcp.NewTool("doc_edit_template",
		mcp.WithDescription("Edit a document template"),
		mcp.WithNumber("docID",
			mcp.Required(),
			mcp.Description("Template/Document ID"),
		),
	)

	s.AddTool(editTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docID := int(args["docID"].(float64))

		resp, err := client.Post(fmt.Sprintf("/index.php?m=doc&f=editTemplate&t=json&docID=%d", docID), nil)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteTemplateTool := mcp.NewTool("doc_delete_template",
		mcp.WithDescription("Delete a document template"),
		mcp.WithNumber("templateID",
			mcp.Required(),
			mcp.Description("Template ID"),
		),
	)

	s.AddTool(deleteTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		templateID := int(args["templateID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=deleteTemplate&t=json&templateID=%d", templateID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete template: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	browseTemplateTool := mcp.NewTool("doc_browse_template",
		mcp.WithDescription("Browse document templates"),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithString("type", mcp.Description("Template type")),
		mcp.WithNumber("docID", mcp.Description("Document ID")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(browseTemplateTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["docID"]; ok && v != nil {
			params["docID"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get("/index.php?m=doc&f=browseTemplate&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse templates: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerBrowseTools(s *server.MCPServer, client *client.ZenTaoClient) {
	mySpaceTool := mcp.NewTool("doc_my_space",
		mcp.WithDescription("Browse my documentation space"),
		mcp.WithNumber("objectID", mcp.Description("Object ID")),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithString("browseType",
			mcp.Description("Browse type: all|draft|bysearch"),
			mcp.Enum("all", "draft", "bysearch"),
		),
		mcp.WithNumber("param", mcp.Description("Parameter value")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
		mcp.WithNumber("docID", mcp.Description("Document ID")),
		mcp.WithString("search", mcp.Description("Search term")),
	)

	s.AddTool(mySpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["objectID"]; ok && v != nil {
			params["objectID"] = int(v.(float64))
		}
		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}
		if v, ok := args["docID"]; ok && v != nil {
			params["docID"] = v
		}
		if v, ok := args["search"].(string); ok {
			params["search"] = v
		}

		resp, err := client.Get("/index.php?m=doc&f=mySpace&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse my space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	productSpaceTool := mcp.NewTool("doc_product_space",
		mcp.WithDescription("Browse product documentation space"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithString("browseType",
			mcp.Description("Browse type: all|draft|bysearch"),
			mcp.Enum("all", "draft", "bysearch"),
		),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("param", mcp.Description("Parameter value")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
		mcp.WithNumber("docID", mcp.Description("Document ID")),
		mcp.WithString("search", mcp.Description("Search term")),
	)

	s.AddTool(productSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		objectID := int(args["objectID"].(float64))

		params := make(map[string]interface{})

		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}
		if v, ok := args["docID"]; ok && v != nil {
			params["docID"] = v
		}
		if v, ok := args["search"].(string); ok {
			params["search"] = v
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=productSpace&t=json&objectID=%d", objectID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse product space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	projectSpaceTool := mcp.NewTool("doc_project_space",
		mcp.WithDescription("Browse project documentation space"),
		mcp.WithNumber("objectID",
			mcp.Required(),
			mcp.Description("Object ID"),
		),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithString("browseType",
			mcp.Description("Browse type: all|draft|bysearch"),
			mcp.Enum("all", "draft", "bysearch"),
		),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("param", mcp.Description("Parameter value")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
		mcp.WithNumber("docID", mcp.Description("Document ID")),
		mcp.WithString("search", mcp.Description("Search term")),
	)

	s.AddTool(projectSpaceTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		objectID := int(args["objectID"].(float64))

		params := make(map[string]interface{})

		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}
		if v, ok := args["docID"]; ok && v != nil {
			params["docID"] = v
		}
		if v, ok := args["search"].(string); ok {
			params["search"] = v
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=projectSpace&t=json&objectID=%d", objectID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to browse project space: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	tableContentsTool := mcp.NewTool("doc_table_contents",
		mcp.WithDescription("Get table of contents for documentation"),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Type: custom|product|project|execution|doctemplate"),
			mcp.Enum("custom", "product", "project", "execution", "doctemplate"),
		),
		mcp.WithNumber("objectID", mcp.Description("Object ID")),
		mcp.WithNumber("libID", mcp.Description("Library ID")),
		mcp.WithNumber("moduleID", mcp.Description("Module ID")),
		mcp.WithString("browseType",
			mcp.Description("Browse type: all|draft|bysearch"),
			mcp.Enum("all", "draft", "bysearch"),
		),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("param", mcp.Description("Parameter value")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
	)

	s.AddTool(tableContentsTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docType := args["type"].(string)

		params := make(map[string]interface{})

		if v, ok := args["objectID"]; ok && v != nil {
			params["objectID"] = int(v.(float64))
		}
		if v, ok := args["libID"]; ok && v != nil {
			params["libID"] = int(v.(float64))
		}
		if v, ok := args["moduleID"]; ok && v != nil {
			params["moduleID"] = int(v.(float64))
		}
		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=tableContents&t=json&type=%s", docType))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to get table contents: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerFileTools(s *server.MCPServer, client *client.ZenTaoClient) {
	showFilesTool := mcp.NewTool("doc_show_files",
		mcp.WithDescription("Show files in documentation"),
		mcp.WithString("type", mcp.Description("Type")),
		mcp.WithNumber("objectID", mcp.Description("Object ID")),
		mcp.WithString("viewType", mcp.Description("View type")),
		mcp.WithString("browseType", mcp.Description("Browse type")),
		mcp.WithNumber("param", mcp.Description("Parameter value")),
		mcp.WithString("orderBy", mcp.Description("Order by field")),
		mcp.WithNumber("recTotal", mcp.Description("Total records")),
		mcp.WithNumber("recPerPage", mcp.Description("Records per page")),
		mcp.WithNumber("pageID", mcp.Description("Page ID")),
		mcp.WithString("searchTitle", mcp.Description("Search title")),
	)

	s.AddTool(showFilesTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		params := make(map[string]interface{})

		if v, ok := args["type"].(string); ok {
			params["type"] = v
		}
		if v, ok := args["objectID"]; ok && v != nil {
			params["objectID"] = int(v.(float64))
		}
		if v, ok := args["viewType"].(string); ok {
			params["viewType"] = v
		}
		if v, ok := args["browseType"].(string); ok {
			params["browseType"] = v
		}
		if v, ok := args["param"]; ok && v != nil {
			params["param"] = v
		}
		if v, ok := args["orderBy"].(string); ok {
			params["orderBy"] = v
		}
		if v, ok := args["recTotal"]; ok && v != nil {
			params["recTotal"] = int(v.(float64))
		}
		if v, ok := args["recPerPage"]; ok && v != nil {
			params["recPerPage"] = int(v.(float64))
		}
		if v, ok := args["pageID"]; ok && v != nil {
			params["pageID"] = int(v.(float64))
		}
		if v, ok := args["searchTitle"].(string); ok {
			params["searchTitle"] = v
		}

		resp, err := client.Get("/index.php?m=doc&f=showFiles&t=json")
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to show files: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteFileTool := mcp.NewTool("doc_delete_file",
		mcp.WithDescription("Delete a file from document"),
		mcp.WithNumber("docID",
			mcp.Required(),
			mcp.Description("Document ID"),
		),
		mcp.WithNumber("fileID",
			mcp.Required(),
			mcp.Description("File ID"),
		),
		mcp.WithString("confirm", mcp.Description("Confirmation string")),
	)

	s.AddTool(deleteFileTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		docID := int(args["docID"].(float64))
		fileID := int(args["fileID"].(float64))

		url := fmt.Sprintf("/index.php?m=doc&f=deleteFile&t=json&docID=%d&fileID=%d", docID, fileID)

		if v, ok := args["confirm"].(string); ok {
			url += fmt.Sprintf("&confirm=%s", v)
		}

		resp, err := client.Get(url)
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete file: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}

func registerCatalogTools(s *server.MCPServer, client *client.ZenTaoClient) {
	editCatalogTool := mcp.NewTool("doc_edit_catalog",
		mcp.WithDescription("Edit a document catalog"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
		mcp.WithString("type",
			mcp.Required(),
			mcp.Description("Type: doc|api"),
			mcp.Enum("doc", "api"),
		),
	)

	s.AddTool(editCatalogTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		moduleID := int(args["moduleID"].(float64))
		docType := args["type"].(string)

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=editCatalog&t=json&moduleID=%d&type=%s", moduleID, docType))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to edit catalog: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})

	deleteCatalogTool := mcp.NewTool("doc_delete_catalog",
		mcp.WithDescription("Delete a document catalog"),
		mcp.WithNumber("moduleID",
			mcp.Required(),
			mcp.Description("Module ID"),
		),
	)

	s.AddTool(deleteCatalogTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()
		moduleID := int(args["moduleID"].(float64))

		resp, err := client.Get(fmt.Sprintf("/index.php?m=doc&f=deleteCatalog&t=json&moduleID=%d", moduleID))
		if err != nil {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to delete catalog: %v", err)), nil
		}
		return mcp.NewToolResultText(string(resp)), nil
	})
}
