# ZenTao MCP Server

## 🔍 Comprehensive Debug Logging

The ZenTao MCP Server includes extensive debug logging capabilities for monitoring, troubleshooting, and development.

### Environment Variables

Configure logging behavior using these environment variables:

| Variable | Description | Default | Values |
|----------|-------------|---------|--------|
| `ZENTAO_LOG_LEVEL` | Logging verbosity level | `INFO` | `DEBUG`, `INFO`, `WARN`, `ERROR` |
| `ZENTAO_LOG_JSON` | Enable JSON structured logging | `false` | `true`, `false` |

### Log Levels

- **DEBUG**: Detailed diagnostic information for development
- **INFO**: General information about server operations
- **WARN**: Warning messages for potential issues
- **ERROR**: Error conditions that don't stop the server

### Log Components

The logging system categorizes messages by component:

- **server**: Main server initialization and lifecycle
- **client**: HTTP client operations and API calls
- **tools**: MCP tool registration and execution
- **resources**: MCP resource registration and access
- **prompt**: MCP prompt handling
- **api**: ZenTao API interactions
- **mcp**: MCP protocol operations
- **auth**: Authentication operations

### Human-Readable Format (Default)

```
[INFO] server/Info: ZenTao MCP Server starting up [version=1.0.0]
[DEBUG] client/Debug: Built request URL [base_url=https://zen.w-w.top/api.php, path=?m=product&f=browse, auth_added=true, param_count=3]
[WARN] server/Warn: ZENTAO_BASE_URL not set, using default [default_url=http://localhost:8080]
```

### JSON Structured Format

Set `ZENTAO_LOG_JSON=true` for machine-readable logs:

```json
{
  "timestamp": "2026-01-03T08:28:17Z",
  "level": "INFO",
  "component": "server",
  "function": "Info",
  "message": "ZenTao MCP Server starting up",
  "fields": {
    "version": "1.0.0"
  }
}
```

### Logged Operations

The logging system captures:

#### Server Lifecycle
- Configuration loading and validation
- MCP server initialization
- Tool and resource registration
- Startup/shutdown events

#### HTTP Operations
- Request URL construction
- Authentication token generation
- Request/response logging with timing
- Error handling and retries

#### MCP Operations
- Tool calls with parameters
- Resource reads with URIs
- Prompt generation
- Protocol message handling

#### API Interactions
- ZenTao API request/response cycles
- Authentication flows
- Data transformation and parsing

### Usage Examples

#### Enable Debug Logging
```bash
export ZENTAO_LOG_LEVEL=DEBUG
./mcp-server
```

#### Enable JSON Logging for Log Aggregation
```bash
export ZENTAO_LOG_LEVEL=INFO
export ZENTAO_LOG_JSON=true
./mcp-server
```

#### Monitor API Calls
```bash
export ZENTAO_LOG_LEVEL=DEBUG
export ZENTAO_LOG_JSON=false
./mcp-server 2>&1 | grep "client\|api"
```

### Performance Considerations

- **DEBUG level** generates significant log volume - use only for troubleshooting
- **JSON format** is recommended for production log aggregation systems
- **INFO level** provides good balance of information and performance
- Logs are written to stderr to avoid interfering with MCP protocol on stdout

# ZenTao MCP Server

A Model Context Protocol (MCP) server for ZenTao, enabling seamless integration between LLM applications and ZenTao project management system.

## Features

- **Authentication**: Secure login and token management
- **Product Management**: Create, update, and delete products
- **Project Management**: Manage projects and executions (sprints/iterations)
- **Story Management**: Handle user stories with CRUD operations
- **Task Management**: Create and manage tasks within executions
- **Bug Tracking**: Full bug lifecycle management
- **Test Cases**: Create and manage test cases
- **Release Planning**: Product plans and builds
- **User Management**: User administration
- **Feedback**: Customer feedback tracking
- **Tickets**: Ticket management system

## Installation

```bash
cd src
go mod download
go build
```

## Configuration

Set environment variables:

```bash
export ZENTAO_BASE_URL="http://your-zentao-instance:8080"
```

## Usage

Start the server:

```bash
./mcp-server
```

The server will communicate via stdio, following the MCP protocol.

## Tools Available

### Authentication
- `zentao_login` - Authenticate with ZenTao

### Products
- `create_product` - Create a new product
- `update_product` - Update existing product
- `delete_product` - Delete a product

### Projects
- `create_project` - Create a new project
- `update_project` - Update existing project
- `delete_project` - Delete a project
- `create_execution` - Create a new execution/sprint
- `delete_execution` - Delete an execution

### Stories
- `create_story` - Create a new user story
- `update_story` - Update existing story
- `change_story` - Change story content
- `delete_story` - Delete a story

### Tasks
- `create_task` - Create a new task
- `update_task` - Update existing task
- `delete_task` - Delete a task

### Bugs
- `create_bug` - Create a new bug
- `update_bug` - Update existing bug
- `delete_bug` - Delete a bug

### Test Cases
- `create_testcase` - Create a new test case
- `update_testcase` - Update existing test case
- `delete_testcase` - Delete a test case

### Plans
- `create_plan` - Create a new product plan
- `update_plan` - Update existing plan
- `delete_plan` - Delete a plan
- `link_stories_to_plan` - Link stories to a plan
- `unlink_stories_from_plan` - Unlink stories from a plan
- `link_bugs_to_plan` - Link bugs to a plan
- `unlink_bugs_from_plan` - Unlink bugs from a plan

### Builds
- `create_build` - Create a new build
- `update_build` - Update existing build
- `delete_build` - Delete a build

### Users
- `create_user` - Create a new user
- `update_user` - Update existing user
- `delete_user` - Delete a user

### Feedback
- `create_feedback` - Create a new feedback
- `update_feedback` - Update existing feedback
- `assign_feedback` - Assign feedback to a user
- `close_feedback` - Close feedback
- `delete_feedback` - Delete feedback

### Tickets
- `create_ticket` - Create a new ticket
- `update_ticket` - Update existing ticket
- `delete_ticket` - Delete a ticket

## Resources

The server exposes the following resources for data access:

- `zentao://products` - List all products
- `zentao://products/{id}` - Get product details
- `zentao://products/{id}/stories` - Get product stories
- `zentao://products/{id}/bugs` - Get product bugs
- `zentao://projects` - List all projects
- `zentao://projects/{id}` - Get project details
- `zentao://projects/{id}/executions` - Get project executions
- `zentao://projects/{id}/stories` - Get project stories
- `zentao://executions/{id}/tasks` - Get execution tasks
- `zentao://stories/{id}` - Get story details
- `zentao://tasks/{id}` - Get task details
- `zentao://bugs/{id}` - Get bug details
- `zentao://users/{id}` - Get user details
- `zentao://user` - Get current user profile

## Example Usage

### Login
```json
{
  "tool": "zentao_login",
  "arguments": {
    "account": "admin",
    "password": "your-password"
  }
}
```

### Create a Product
```json
{
  "tool": "create_product",
  "arguments": {
    "name": "My Product",
    "code": "PROD001",
    "desc": "A sample product"
  }
}
```

### Create a Project
```json
{
  "tool": "create_project",
  "arguments": {
    "name": "My Project",
    "code": "PROJ001",
    "begin": "2024-01-01",
    "end": "2024-12-31",
    "products": [1]
  }
}
```

### Create a Story
```json
{
  "tool": "create_story",
  "arguments": {
    "title": "User Authentication",
    "product": 1,
    "pri": 1,
    "category": "feature",
    "spec": "Implement user login and registration"
  }
}
```

## Development

To regenerate server code:
```bash
cd src
go generate ./...
```

## License

See LICENSE.md file.
