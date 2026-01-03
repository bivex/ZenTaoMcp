# Agent Guide for ZenTao MCP Server

This document helps AI agents work effectively with this ZenTao MCP Server codebase.

## Project Overview

This is a Go-based **Model Context Protocol (MCP) server** that provides integration between LLM applications and the ZenTao project management system. The server communicates via stdio and exposes MCP tools and resources for managing products, projects, stories, tasks, bugs, test cases, and more.

- **Language**: Go 1.23.2
- **Core Library**: `github.com/mark3labs/mcp-go` v0.4.0
- **Communication**: stdio (MCP protocol)
- **Purpose**: Bridge LLMs to ZenTao REST API

## Essential Commands

### Build and Run

```bash
# From project root
cd src
go mod download      # Download dependencies
go build              # Build the binary
./mcp-server          # Run the server
```

### Environment Variables

- `ZENTAO_BASE_URL` (optional): Default is `http://localhost:8080`

### Current Build Issues

**IMPORTANT**: The codebase has build errors with the current MCP library version (v0.4.0). The errors indicate API incompatibilities:

- `mcp.NewResource` is undefined
- `mcp.WithResourceDescription` and `mcp.WithMIMEType` are undefined
- `s.AddResource` fails - MCPServer is a pointer to interface, not interface
- `mcp.TextResourceContents` struct cannot be used as `mcp.ResourceContents` value
- `mcp.NewTool` is undefined in tools/

These suggest the MCP library API has changed. Check the library documentation for the correct API before attempting to fix.

## Project Structure

```
/Volumes/External/Code/ZenTaoMcp
├── src/
│   ├── main.go              # Entry point - server initialization
│   ├── client/
│   │   └── client.go        # ZenTao HTTP client with auth
│   ├── tools/               # MCP tools (one file per entity)
│   │   ├── auth.go          # Authentication tools
│   │   ├── products.go      # Product CRUD tools
│   │   ├── projects.go      # Project and execution tools
│   │   ├── stories.go       # User story tools
│   │   ├── tasks.go         # Task tools
│   │   ├── bugs.go          # Bug tracking tools
│   │   ├── testcases.go     # Test case tools
│   │   ├── plans.go         # Release planning tools
│   │   ├── builds.go        # Build/release tools
│   │   ├── users.go         # User management tools
│   │   ├── feedbacks.go     # Customer feedback tools
│   │   └── tickets.go       # Ticket management tools
│   ├── resources/           # MCP resources (data access)
│   │   ├── entities.go      # Resource URI helper
│   │   ├── products.go      # Product resources
│   │   └── (other resources)
│   ├── go.mod
│   └── go.sum
├── api_doc.txt              # Complete ZenTao API documentation
├── README.md                # User-facing documentation
└── LICENSE.md               # MIT License
```

## Code Conventions and Patterns

### Copyright Header

Every source file MUST include this copyright header:

```go
// Copyright (c) 2026 Bivex
//
// Author: Bivex
// Contact: support@b-b.top
//
// For up-to-date contact information:
// https://github.com/bivex
//
//
// Licensed under the MIT License.
// Commercial licensing available upon request.
```

Note: Some files have slight variations ("Licensed under MIT License" vs "Licensed under the MIT License"), but maintain the pattern.

### Tool Registration Pattern

Tools are organized by entity type. Each file has a `Register[Entity]Tools` function:

```go
func RegisterProductTools(s *server.MCPServer, client *client.ZenTaoClient) {
    createProductTool := mcp.NewTool("create_product",
        mcp.WithDescription("Create a new product in ZenTao"),
        mcp.WithString("name",
            mcp.Required(),
            mcp.Description("Product name"),
        ),
        // ... more parameters
    )

    s.AddTool(createProductTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
        args := request.GetArguments()
        // Process arguments
        resp, err := client.Post("/products", body)
        if err != nil {
            return mcp.NewToolResultError(fmt.Sprintf("Failed to create product: %v", err)), nil
        }
        return mcp.NewToolResultText(string(resp)), nil
    })
}
```

### Tool Parameter Handling

All parameters from MCP come as `map[string]interface{}`. Types need explicit conversion:

```go
args := request.GetArguments()

// String parameter
name := args["name"].(string)

// Number parameter (MCP sends numbers as float64)
id := int(args["id"].(float64))

// Optional parameters - check existence
if v, ok := args["program"]; ok && v != nil {
    body["program"] = int(v.(float64))
}
```

### Resource Registration Pattern

Resources expose data via URIs. Each has a URI pattern and handler:

```go
func RegisterProductResources(s *server.MCPServer, client *client.ZenTaoClient) {
    productListResource := mcp.NewResource(
        "zentao://products",
        "ZenTao Product List",
        mcp.WithResourceDescription("List of all products in ZenTao"),
        mcp.WithMIMEType("application/json"),
    )

    s.AddResource(productListResource, func(ctx context.Context, request mcp.ReadResourceRequest) ([]mcp.ResourceContents, error) {
        resp, err := client.Get("/products")
        if err != nil {
            return nil, fmt.Errorf("failed to get products: %w", err)
        }
        return []mcp.ResourceContents{
            mcp.TextResourceContents{
                Uri:      "zentao://products",
                MimeType: "application/json",
                Text:     string(resp),
            },
        }, nil
    })
}
```

### URI Pattern Extraction

Resources use `{id}` placeholders. The `extractIDFromURI` helper in `resources/entities.go` extracts IDs:

```go
func extractIDFromURI(uri, resourceType string) string {
    re := regexp.MustCompile(fmt.Sprintf("%s/([^/]+)", resourceType))
    matches := re.FindStringSubmatch(uri)
    if len(matches) > 1 {
        return matches[1]
    }
    return ""
}
```

Usage:
```go
id := extractIDFromURI(request.Params.Uri, "products")
resp, err := client.Get(fmt.Sprintf("/products/%s", id))
```

### Client API Usage

The `ZenTaoClient` in `client/client.go` handles HTTP requests with automatic token injection:

```go
// Authentication
token, err := client.GetToken(account, password)

// GET request
resp, err := client.Get("/products")

// POST request
resp, err := client.Post("/products", body)

// PUT request
resp, err := client.Put(fmt.Sprintf("/product/%d", id), body)

// DELETE request
resp, err := client.Delete(fmt.Sprintf("/product/%d", id))
```

All requests automatically include the `Token` header if authenticated.

## ZenTao API Integration

### API Documentation

The complete ZenTao REST API is documented in `api_doc.txt`. This contains:
- 82 API endpoints
- All request/response schemas
- Required vs optional parameters
- Example requests

Refer to this file when adding new tools or resources.

### Authentication Flow

1. Client calls `/tokens` endpoint with account/password
2. Response contains a token
3. Token is stored in `ZenTaoClient.Token`
4. All subsequent requests include `Token` header

### Common API Patterns

- **Lists**: `GET /{entity}` - Returns array
- **Details**: `GET /{entity}/{id}` - Returns single object
- **Create**: `POST /{entity}` - Body with fields
- **Update**: `PUT /{entity}/{id}` - Partial updates accepted
- **Delete**: `DELETE /{entity}/{id}`

## Adding New Features

### Adding a New Tool

1. Create or edit file in `tools/` directory
2. Follow naming: `Register[Entity]Tools(s, client)`
3. Define tool with `mcp.NewTool()`, parameters with `mcp.WithString/Number/etc.`
4. Add handler function that:
   - Extracts arguments
   - Builds request body
   - Calls `client.Get/Post/Put/Delete()`
   - Returns `mcp.NewToolResultText()` or `mcp.NewToolResultError()`
5. Register in `main.go` in `registerTools()` function

### Adding a New Resource

1. Create or edit file in `resources/` directory
2. Follow naming: `Register[Entity]Resources(s, client)`
3. Define resource with `mcp.NewResource(uri, name, options)`
4. Add handler that:
   - Extracts ID from URI if needed
   - Calls `client.Get()`
   - Returns `[]mcp.ResourceContents` with `mcp.TextResourceContents`
5. Register in `main.go` in `registerResources()` function

### API Endpoints to Consider

Based on `api_doc.txt`, potential additions include:
- Programs (project sets)
- Test tasks
- Releases
- More query parameters for filtering

## Important Gotchas

1. **Type Conversions**: MCP always sends numbers as `float64`. Always cast to `int` for IDs.
2. **Optional Parameters**: Use pattern `if v, ok := args["param"]; ok && v != nil` to check optional params.
3. **URI IDs**: String IDs from URIs need to be used as strings in API paths (`/products/%s`), not integers.
4. **Response Formatting**: All responses are returned as string via `mcp.NewToolResultText(string(resp))`.
5. **Error Handling**: Use `fmt.Errorf("failed to X: %w", err)` for resource errors, and `mcp.NewToolResultError(fmt.Sprintf("Failed to X: %v", err))` for tool errors.
6. **Enum Values**: Use `mcp.Enum("value1", "value2")` to restrict parameter values.
7. **Array Parameters**: Use `mcp.WithArray()` for array-type parameters.
8. **Resource URIs**: Follow pattern `zentao://{entity}/{id}` or `zentao://{entity}/{id}/{subentity}`.
9. **Build Errors**: Current code doesn't build due to MCP library API changes. Don't attempt to run tests without fixing the build first.

## Testing

**Note**: No test files (`*_test.go`) exist in the codebase. No CI configuration found.

When adding tests:
- Place test files next to source (e.g., `tools/products_test.go`)
- Use Go's `testing` package
- Mock the `ZenTaoClient` for HTTP calls

## Dependencies

From `go.mod`:
- `github.com/mark3labs/mcp-go` v0.4.0 - MCP protocol implementation
- Standard library only for everything else (net/http, encoding/json, fmt, etc.)

No third-party HTTP libraries - uses `net/http` directly.

## File Organization Summary

| Directory | Purpose |
|-----------|---------|
| `src/client/` | HTTP client with automatic auth token management |
| `src/tools/` | MCP tools (one file per domain entity) |
| `src/resources/` | MCP resources (data access via URIs) |
| `src/main.go` | Server initialization, tool/resource registration |
| `api_doc.txt` | ZenTao REST API reference (82 endpoints) |
| `README.md` | User documentation |

## Common Workflows

### Adding CRUD for New Entity

1. Check `api_doc.txt` for endpoint details
2. Create `src/tools/newentity.go` with `RegisterNewEntityTools()`
3. Create `src/resources/newentity.go` with `RegisterNewEntityResources()`
4. Add both registrations in `src/main.go`
5. Add documentation in `src/README.md`

### Fixing Build Errors

1. Check MCP library version and documentation
2. Identify changed APIs (likely `NewTool`, `NewResource`, `AddResource`, `AddTool`)
3. Update all tool and resource definitions to match new API
4. Test build with `go build`

### Debugging

Since server communicates via stdio, debugging requires:
- Add logging to see what's happening
- Use MCP client to send requests
- Check responses in stderr/stdout
- Ensure ZenTao instance is running and accessible

## License

MIT License. Commercial licensing available upon request.
