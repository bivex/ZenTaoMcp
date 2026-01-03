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
