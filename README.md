# ZenTao MCP Server

<img width="727" height="575" alt="image" src="https://github.com/user-attachments/assets/7c26eb36-550b-47d6-b007-2d2cb2c3aa0a" />

A comprehensive Model Context Protocol (MCP) server for integrating with ZenTao project management system. This server provides **304 tools**, **42 resources**, and **2 prompts** covering all major ZenTao modules.

## Features

- **304 MCP Tools** - Complete CRUD operations for all ZenTao entities
- **42 MCP Resources** - Data access via URI-based resources
- **96 Resource Templates** - Dynamic resource access with parameters
- **2 MCP Prompts** - Guided workflows for common operations
- **Full API Coverage** - Supports products, projects, stories, tasks, bugs, users, AI features, and more
- **Authentication Support** - App-based and session-based authentication
- **Comprehensive Logging** - Detailed logging for debugging and monitoring

## Installation

```bash
cd src
go mod download
go build
./mcp-server
```

## Configuration

Set the following environment variables:

- `ZENTAO_BASE_URL` - ZenTao API base URL (default: `http://localhost:8080`)
- `ZENTAO_AUTH_METHOD` - Authentication method: `app` or `session` (default: `app`)
- `ZENTAO_APP_CODE` - App code for app-based authentication
- `ZENTAO_APP_KEY` - App key for app-based authentication
- `ZENTAO_LOG_LEVEL` - Log level (optional)
- `ZENTAO_LOG_JSON` - Enable JSON logging (optional)

## Tools (304 Total)

### Authentication (2 tools)
- `zentao_login_app` - Login with app credentials
- `zentao_login_session` - Login with session credentials

### Products (8 tools)
- `create_product` - Create a new product
- `edit_product` - Edit an existing product
- `view_product` - View product details
- `delete_product` - Delete a product
- `browse_products` - Browse products with filters
- `get_product_index` - Get product index/dashboard
- `get_product_projects` - Get projects for a product
- `get_product_stories` - Get stories for a product

### Projects (15+ tools)
- `create_project` - Create a new project
- `edit_project` - Edit an existing project
- `view_project` - View project details
- `get_project_index` - Get project index
- `browse_projects` - Browse projects with filters
- `get_project_kanban` - Get project kanban board
- `get_project_team` - Get project team members
- `manage_project_members` - Manage project members
- `get_project_executions` - Get project executions
- `get_project_stories` - Get project stories
- `get_project_bugs` - Get project bugs
- `get_project_testcases` - Get project test cases
- `get_project_builds` - Get project builds
- `get_project_releases` - Get project releases
- `update_execution` - Update execution details

### Stories (20+ tools)
- `create_story` - Create a new user story
- `edit_story` - Edit an existing story
- `view_story` - View story details
- `delete_story` - Delete a story
- `browse_stories` - Browse stories with filters
- `review_story` - Review a story
- `close_story` - Close a story
- `assign_story` - Assign story to user
- `link_story` - Link related stories
- And more...

### Tasks (15+ tools)
- `create_task` - Create a new task
- `edit_task` - Edit an existing task
- `view_task` - View task details
- `start_task` - Start a task
- `finish_task` - Finish a task
- `close_task` - Close a task
- `assign_task` - Assign task to user
- And more...

### Bugs (15+ tools)
- `create_bug` - Create a new bug
- `edit_bug` - Edit an existing bug
- `view_bug` - View bug details
- `resolve_bug` - Resolve a bug
- `close_bug` - Close a bug
- `assign_bug` - Assign bug to user
- And more...

### Test Cases (10+ tools)
- `create_testcase` - Create a new test case
- `edit_testcase` - Edit an existing test case
- `view_testcase` - View test case details
- `delete_testcase` - Delete a test case
- And more...

### Test Tasks (3 tools)
- `get_testtasks` - Get list of test tasks
- `get_testtask` - Get test task details
- `get_project_testtasks` - Get test tasks for a project

### Plans (10+ tools)
- `create_plan` - Create a new plan
- `edit_plan` - Edit an existing plan
- `view_plan` - View plan details
- `delete_plan` - Delete a plan
- And more...

### Builds (15+ tools)
- `create_build` - Create a new build
- `edit_build` - Edit an existing build
- `view_build` - View build details
- `delete_build` - Delete a build
- `get_product_builds` - Get builds for a product
- `get_project_builds` - Get builds for a project
- `get_execution_builds` - Get builds for an execution
- `link_story_to_build` - Link story to build
- `link_bug_to_build` - Link bug to build
- And more...

### Releases (2 tools)
- `get_project_releases` - Get releases for a project
- `get_product_releases` - Get releases for a product

### Users (5+ tools)
- `create_user` - Create a new user
- `edit_user` - Edit an existing user
- `get_user` - Get user details
- `delete_user` - Delete a user
- `browse_users` - Browse users with filters

### Feedback (5+ tools)
- `create_feedback` - Create a new feedback
- `edit_feedback` - Edit an existing feedback
- `view_feedback` - View feedback details
- `delete_feedback` - Delete a feedback
- `assign_feedback` - Assign feedback to user

### Tickets (5+ tools)
- `create_ticket` - Create a new ticket
- `edit_ticket` - Edit an existing ticket
- `view_ticket` - View ticket details
- `delete_ticket` - Delete a ticket
- `browse_tickets` - Browse tickets with filters

### Programs (15+ tools)
- `create_program` - Create a new program
- `edit_program` - Edit an existing program
- `view_program` - View program details
- `close_program` - Close a program
- `start_program` - Start a program
- `get_program_products` - Get products in a program
- `get_program_projects` - Get projects in a program
- `get_program_stakeholders` - Get program stakeholders
- And more...

### My Module (15+ tools)
- `get_my_dashboard` - Get user dashboard
- `get_my_profile` - Get user profile
- `edit_my_profile` - Edit user profile
- `change_my_password` - Change user password
- `get_my_todos` - Get user todos
- `get_my_stories` - Get user stories
- `get_my_tasks` - Get user tasks
- `get_my_bugs` - Get user bugs
- `get_my_projects` - Get user projects
- `get_my_executions` - Get user executions
- `get_my_team` - Get user team
- `get_my_dynamic` - Get user activity feed
- And more...

### Todos (20+ tools)
- `create_todo` - Create a new todo
- `edit_todo` - Edit an existing todo
- `start_todo` - Start a todo
- `finish_todo` - Finish a todo
- `close_todo` - Close a todo
- `assign_todo` - Assign todo to user
- `batch_create_todos` - Create multiple todos
- `batch_edit_todos` - Edit multiple todos
- `batch_finish_todos` - Finish multiple todos
- And more...

### Personnel (5 tools)
- `get_accessible_personnel` - Get accessible personnel
- `get_personnel_invest` - Get personnel investment
- `get_personnel_whitelist` - Get personnel whitelist
- `add_personnel_whitelist` - Add to whitelist
- `unbind_personnel_whitelist` - Remove from whitelist

### Stakeholders (12 tools)
- `browse_stakeholders` - Browse stakeholders
- `create_stakeholder` - Create a new stakeholder
- `batch_create_stakeholders` - Create multiple stakeholders
- `edit_stakeholder` - Edit an existing stakeholder
- `delete_stakeholder` - Delete a stakeholder
- `view_stakeholder` - View stakeholder details
- `communicate_stakeholder` - Communicate with stakeholder
- `expect_stakeholder` - Set stakeholder expectations
- And more...

### Branches (8 tools)
- `manage_branches` - Manage product branches
- `create_branch` - Create a new branch
- `edit_branch` - Edit an existing branch
- `close_branch` - Close a branch
- `activate_branch` - Activate a branch
- `sort_branches` - Sort branches
- `get_branches` - Get branches list
- `merge_branch` - Merge branches

### Designs (15+ tools)
- `browse_designs` - Browse design documents
- `create_design` - Create a new design
- `batch_create_designs` - Create multiple designs
- `view_design` - View design details
- `edit_design` - Edit an existing design
- `delete_design` - Delete a design
- `assign_design` - Assign design to user
- `link_commit_to_design` - Link commit to design
- And more...

### Project Builds (12 tools)
- `browse_project_builds` - Browse project builds
- `create_project_build` - Create a new project build
- `edit_project_build` - Edit an existing build
- `view_project_build` - View build details
- `delete_project_build` - Delete a build
- `link_story_to_project_build` - Link story to build
- `link_bug_to_project_build` - Link bug to build
- And more...

### Executions (20+ tools)
- `browse_execution` - Browse executions
- `get_execution_tasks` - Get execution tasks
- `get_execution_stories` - Get execution stories
- `get_execution_bugs` - Get execution bugs
- `get_execution_builds` - Get execution builds
- `get_execution_burn_chart` - Get burn chart
- `get_execution_cfd` - Get cumulative flow diagram
- `get_execution_kanban` - Get kanban board
- `get_execution_team` - Get execution team
- `manage_execution_members` - Manage team members
- `link_story_to_execution` - Link story to execution
- And more...

### Kanban (30+ tools)
- Space management: `create_space`, `edit_space`, `delete_space`
- Board management: `create_kanban`, `edit_kanban`, `view_kanban`
- Region management: `create_region`, `edit_region`, `delete_region`
- Lane management: `create_lane`, `sort_lane`, `delete_lane`
- Column management: `create_column`, `split_column`, `archive_column`
- Card management: `create_card`, `edit_card`, `move_card`, `finish_card`
- Import/Export: `import_card`, `import_plan`, `import_release`, `import_build`, `import_execution`
- And more...

### Epics (20+ tools)
- `create_epic` - Create a new epic
- `edit_epic` - Edit an existing epic
- `view_epic` - View epic details
- `delete_epic` - Delete an epic
- `link_story_to_epic` - Link story to epic
- `link_requirement_to_epic` - Link requirement to epic
- `batch_operations` - Batch epic operations
- And more...

### Requirements (20+ tools)
- `create_requirement` - Create a new requirement
- `edit_requirement` - Edit an existing requirement
- `view_requirement` - View requirement details
- `delete_requirement` - Delete a requirement
- `link_story_to_requirement` - Link story to requirement
- `link_requirement` - Link related requirements
- `batch_operations` - Batch requirement operations
- And more...

### Spaces (3 tools)
- `browse_spaces` - Browse spaces
- `create_application_space` - Create application space
- `get_store_app_info` - Get store app information

### Transfer (6 tools)
- `export_data` - Export data from a module
- `export_template` - Export import template
- `import_data` - Import data to a module
- `get_import_table_body` - Get import preview table
- `get_import_options` - Get import field options
- `validate_import_data` - Validate import data

### ZAI - ZenTao AI (6 tools)
- `get_zai_settings` - Get ZAI module settings
- `update_zai_settings` - Update ZAI settings
- `get_zai_token` - Get AI authentication token
- `get_vectorization_status` - Get vectorization status
- `enable_vectorization` - Enable vectorization
- `sync_vectorization` - Sync vectorization data

### AI - Artificial Intelligence (23 tools)
- Mini Programs: `get_mini_programs`, `publish_mini_program`, `unpublish_mini_program`, `import_mini_program`
- Prompts: `get_prompts`, `create_prompt`, `edit_prompt`, `delete_prompt`, `execute_prompt`, `publish_prompt`
- Prompt Configuration: `assign_prompt_role`, `select_prompt_data_source`, `set_prompt_purpose`, `set_prompt_target_form`, `finalize_prompt`
- Prompt Execution: `reset_prompt_execution`, `audit_prompt`, `get_testing_location`
- Role Templates: `get_role_templates`
- And more...

### Zanode - Node Management (27 tools)
- Node Management: `browse_zanodes`, `create_zanode`, `edit_zanode`, `view_zanode`, `get_zanodes`
- Lifecycle: `start_zanode`, `close_zanode`, `suspend_zanode`, `reboot_zanode`, `resume_zanode`, `destroy_zanode`
- VNC: `get_zanode_vnc` - Get VNC access
- Images: `create_zanode_image`, `get_zanode_images`, `get_zanode_image`, `update_zanode_image`
- Snapshots: `create_zanode_snapshot`, `edit_zanode_snapshot`, `delete_zanode_snapshot`, `browse_zanode_snapshots`, `restore_zanode_snapshot`
- Tasks & Services: `get_zanode_task_status`, `get_zanode_service_status`, `install_zanode_service`
- ZTF Scripts: `get_zanode_ztf_script`, `run_zanode_ztf_script`
- Instructions: `get_zanode_instructions`

## Resources (42 Total + 96 Templates)

### Products
- `zentao://products` - List of all products
- `zentao://products/{id}` - Individual product details
- `zentao://products/{id}/projects` - Product projects
- `zentao://products/{id}/stories` - Product stories
- `zentao://products/{id}/plans` - Product plans
- `zentao://products/{id}/releases` - Product releases

### Projects
- `zentao://projects` - List of all projects
- `zentao://projects/{id}` - Individual project details
- `zentao://projects/{id}/executions` - Project executions
- `zentao://projects/{id}/stories` - Project stories
- `zentao://projects/{id}/bugs` - Project bugs
- `zentao://projects/{id}/builds` - Project builds
- `zentao://projects/{id}/releases` - Project releases
- `zentao://projects/{id}/team` - Project team

### Programs
- `zentao://programs` - List of all programs
- `zentao://programs/{id}` - Individual program details
- `zentao://programs/{id}/products` - Program products
- `zentao://programs/{id}/projects` - Program projects
- `zentao://programs/{id}/stakeholders` - Program stakeholders

### Stories
- `zentao://stories/{id}` - Individual story details
- `zentao://stories/{id}/tasks` - Story tasks
- `zentao://stories/{id}/bugs` - Story bugs

### Tasks
- `zentao://tasks/{id}` - Individual task details

### Bugs
- `zentao://bugs/{id}` - Individual bug details

### Users
- `zentao://users` - List of all users
- `zentao://users/{id}` - Individual user details

### Test Tasks
- `zentao://testtasks` - List of test tasks
- `zentao://testtasks/{id}` - Individual test task details

### Builds
- `zentao://builds/{id}` - Individual build details
- `zentao://products/{id}/builds` - Product builds
- `zentao://projects/{id}/builds` - Project builds
- `zentao://executions/{id}/builds` - Execution builds

### Plans
- `zentao://plans/{id}` - Individual plan details

### Releases
- `zentao://projects/{id}/releases` - Project releases
- `zentao://products/{id}/releases` - Product releases

### API Libraries
- `zentao://api-libs` - List of API libraries
- `zentao://api-libs/{libId}/releases` - Library releases
- `zentao://api-libs/{libId}/structs` - Library structures
- `zentao://api-libs/{libId}/apis` - Library APIs
- `zentao://api-libs/{libId}/apis/{apiId}` - Individual API details

### Entries
- `zentao://entries` - List of entries
- `zentao://entries/{id}` - Individual entry details
- `zentao://entries/{id}/log` - Entry log

### My Module
- `zentao://my/dashboard` - User dashboard
- `zentao://my/profile` - User profile
- `zentao://my/calendar` - User calendar
- `zentao://my/work/{mode}` - User work items
- `zentao://my/todos` - User todos
- `zentao://my/stories` - User stories
- `zentao://my/tasks` - User tasks
- `zentao://my/bugs` - User bugs
- `zentao://my/projects` - User projects
- `zentao://my/executions` - User executions
- `zentao://my/team` - User team
- `zentao://my/dynamic` - User activity feed

### Todos
- `zentao://todos` - List of todos
- `zentao://todos/{id}` - Individual todo details
- `zentao://todos/{id}/detail` - Todo detailed view
- `zentao://todos/status/{status}` - Todos by status
- `zentao://todos/type/{type}` - Todos by type

### Personnel
- `zentao://personnel/accessible` - Accessible personnel
- `zentao://personnel/invest/{programID}` - Personnel investment
- `zentao://personnel/whitelist/{objectID}/{module}/{objectType}` - Personnel whitelist

### Stakeholders
- `zentao://stakeholders` - List of stakeholders
- `zentao://stakeholders/{projectID}` - Project stakeholders
- `zentao://stakeholders/{stakeholderID}` - Individual stakeholder details

### Executions
- `zentao://executions` - List of executions
- `zentao://executions/{id}` - Individual execution details
- `zentao://executions/{id}/tasks` - Execution tasks
- `zentao://executions/{id}/stories` - Execution stories
- `zentao://executions/{id}/bugs` - Execution bugs
- `zentao://executions/{id}/team` - Execution team
- `zentao://executions/{id}/burn` - Burn chart
- `zentao://executions/{id}/cfd` - Cumulative flow diagram
- `zentao://executions/{id}/kanban` - Kanban board

### Kanban
- `zentao://kanban/spaces` - List of kanban spaces
- `zentao://kanban/spaces/{spaceID}` - Space details
- `zentao://kanban/{kanbanID}` - Kanban board details
- `zentao://kanban/{kanbanID}/regions/{regionID}` - Region details
- `zentao://kanban/{kanbanID}/regions/{regionID}/lanes/{laneID}` - Lane details
- `zentao://kanban/{kanbanID}/regions/{regionID}/lanes/{laneID}/columns/{columnID}` - Column details
- `zentao://kanban/cards/{cardID}` - Card details
- `zentao://kanban/{kanbanID}/regions/{regionID}/archived-cards` - Archived cards
- `zentao://kanban/{kanbanID}/regions/{regionID}/archived-columns` - Archived columns

### Epics
- `zentao://epics/{id}` - Individual epic details
- `zentao://epics/{id}/stories` - Epic stories
- `zentao://epics/{id}/requirements` - Epic requirements
- `zentao://products/{id}/epics` - Product epics
- `zentao://projects/{id}/epics` - Project epics
- `zentao://executions/{id}/epics` - Execution epics

### Requirements
- `zentao://requirements/{id}` - Individual requirement details
- `zentao://requirements/{id}/stories` - Requirement stories
- `zentao://requirements/{id}/linked-requirements` - Linked requirements
- `zentao://products/{id}/requirements` - Product requirements
- `zentao://projects/{id}/requirements` - Project requirements
- `zentao://executions/{id}/requirements` - Execution requirements

### Spaces
- `zentao://spaces` - List of spaces
- `zentao://spaces/{id}` - Individual space details
- `zentao://spaces/{id}/applications/{appID}` - Space application details

### Transfer
- `zentao://transfer/modules` - Supported transfer modules
- `zentao://transfer/{module}/export-template` - Export template for module
- `zentao://transfer/{module}/import-status` - Import status for module
- `zentao://transfer/exports` - Export history
- `zentao://transfer/imports` - Import history

### ZAI
- `zentao://zai/settings` - ZAI module settings
- `zentao://zai/token` - AI authentication token
- `zentao://zai/vectorization/status` - Vectorization status

### AI
- `zentao://ai/admin` - AI module admin overview
- `zentao://ai/mini-programs` - List of mini programs
- `zentao://ai/prompts` - List of prompts
- `zentao://ai/role-templates` - Role templates
- `zentao://ai/prompts/{id}` - Individual prompt details
- `zentao://ai/mini-programs/{appID}` - Mini program details
- `zentao://ai/prompts/published` - Published prompts
- `zentao://ai/prompts/testing` - Testing prompts
- `zentao://ai/mini-programs/published` - Published mini programs

### Zanode
- `zentao://zanode/instructions` - Node management instructions
- `zentao://zanode/nodes` - List of all nodes
- `zentao://zanode/nodes/{id}` - Individual node details
- `zentao://zanode/nodes/{id}/vnc` - Node VNC access
- `zentao://zanode/nodes/{id}/snapshots` - Node snapshots
- `zentao://zanode/nodes/{id}/tasks` - Node tasks
- `zentao://zanode/hosts/{hostID}/nodes` - Host nodes list
- `zentao://zanode/hosts/{hostID}/images` - Host images list
- `zentao://zanode/hosts/{hostID}/services` - Host service status
- `zentao://zanode/images/{imageID}` - Image details

## Prompts (2 Total)

1. **`create_product`** - Guided workflow for creating a new product
   - Required arguments: `name`, `code`
   - Optional arguments: `program`, `line`, `PO`, `QD`, `RD`, `type`, `desc`, `acl`

2. **`create_story`** - Guided workflow for creating a new user story
   - Required arguments: `title`, `product`
   - Optional arguments: `module`, `plan`, `source`, `pri`, `estimate`, `spec`, `verify`

## Usage Examples

### Using Tools

```json
{
  "name": "create_product",
  "arguments": {
    "name": "My Product",
    "code": "MP001"
  }
}
```

### Using Resources

Access product details:
```
zentao://products/123
```

Access user's todos:
```
zentao://my/todos
```

### Using Prompts

Get guided workflow for creating a product:
```
prompt://create_product?name=MyProduct&code=MP001
```

## License

MIT License. Commercial licensing available upon request.

## Contact

- Author: Bivex
- Contact: support@b-b.top
- GitHub: https://github.com/bivex
