# ZenTao MCP Server - Model Context Protocol Integration for ZenTao Project Management

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.23+-00ADD8?logo=go)](https://golang.org/)
[![ZenTao](https://img.shields.io/badge/ZenTao-21.7.7-blue)](https://www.zentao.net/)
[![MCP](https://img.shields.io/badge/MCP-Protocol-green)](https://modelcontextprotocol.io/)

<img width="727" height="575" alt="ZenTao MCP Server Architecture" src="https://github.com/user-attachments/assets/7c26eb36-550b-47d6-b007-2d2cb2c3aa0a" />

## Overview

**ZenTao MCP Server** is a comprehensive [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server that provides seamless integration between AI assistants and the [ZenTao](https://www.zentao.net/) project management system. Built with Go, this server exposes **400 tools**, **46 resources**, and **2 prompts** covering all major ZenTao modules including products, projects, user stories, tasks, bugs, test cases, releases, builds, and more.

### What is ZenTao?

[ZenTao](https://www.zentao.net/) is an open-source project management and bug tracking system designed for agile development teams. It supports Scrum, Kanban, and waterfall methodologies, providing comprehensive features for product management, project planning, requirement tracking, bug management, and test case management.

### What is MCP?

The [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) is an open protocol that enables AI assistants to securely access external data sources and tools. MCP servers act as bridges between AI applications and various services, allowing LLMs to interact with APIs, databases, and other systems in a standardized way.

## Key Features

- **🚀 400 MCP Tools** - Complete CRUD operations for all ZenTao entities (products, projects, stories, tasks, bugs, users, AI features, and more)
- **📦 46 MCP Resources** - URI-based data access with RESTful resource patterns
- **🔧 100 Resource Templates** - Dynamic resource access with parameterized URIs
- **💡 2 MCP Prompts** - Guided workflows for common operations (product creation, story creation)
- **🔐 Authentication Support** - App-based and session-based authentication methods
- **📊 Full API Coverage** - Supports all ZenTao modules including:
  - Product Management
  - Project Management & Executions
  - User Stories & Requirements
  - Task Management
  - Bug Tracking (Enhanced with batch operations)
  - Test Case Management (Enhanced with scenes, import/export)
  - Test Task Management (Enhanced with execution, results)
  - Case Library Management
  - Test Report Generation
  - Test Suite Management
  - Build & Release Management
  - Kanban Boards
  - AI Integration (ZAI)
  - Node Management (Zanode)
  - And much more
- **📝 Comprehensive Logging** - Detailed logging for debugging and monitoring
- **⚡ High Performance** - Built with Go for optimal performance and low resource usage

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Tools Reference](#tools-400-total)
- [Resources Reference](#resources-46-total--100-templates)
- [Prompts](#prompts-2-total)
- [Usage Examples](#usage-examples)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Installation

### Prerequisites

- [Go](https://golang.org/) 1.23 or higher
- Access to a ZenTao instance (version 21.7.7 or compatible)
- ZenTao API credentials (app code/key or session credentials)

### Build from Source

```bash
# Clone the repository
git clone https://github.com/bivex/ZenTaoMcp.git
cd ZenTaoMcp

# Build the server
cd src
go mod download
go build -o mcp-server

# Run the server
./mcp-server
```

## Quick Start

1. **Set Environment Variables:**
   ```bash
   export ZENTAO_BASE_URL="https://your-zentao-instance.com"
   export ZENTAO_AUTH_METHOD="app"
   export ZENTAO_APP_CODE="your-app-code"
   export ZENTAO_APP_KEY="your-app-key"
   ```

2. **Start the MCP Server:**
   ```bash
   ./mcp-server
   ```

3. **Connect from Your MCP Client:**
   The server communicates via stdio using the MCP protocol. Configure your MCP client to use this server.

## Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `ZENTAO_BASE_URL` | ZenTao API base URL | `http://localhost:8080` | No |
| `ZENTAO_AUTH_METHOD` | Authentication method: `app` or `session` | `app` | No |
| `ZENTAO_APP_CODE` | App code for app-based authentication | - | Yes (if using app auth) |
| `ZENTAO_APP_KEY` | App key for app-based authentication | - | Yes (if using app auth) |
| `ZENTAO_LOG_LEVEL` | Log level (debug, info, warn, error) | `info` | No |
| `ZENTAO_LOG_JSON` | Enable JSON logging format | `false` | No |

### Authentication Methods

#### App-Based Authentication (Recommended)
```bash
export ZENTAO_AUTH_METHOD="app"
export ZENTAO_APP_CODE="your-app-code"
export ZENTAO_APP_KEY="your-app-key"
```

#### Session-Based Authentication
```bash
export ZENTAO_AUTH_METHOD="session"
# Then use zentao_login_session tool with account/password
```

## Tools (400 Total)

The server provides comprehensive tools for managing all aspects of ZenTao. Here's a categorized overview:

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

### Bugs (35+ tools)
- `create_bug` - Create a new bug
- `update_bug` - Update an existing bug
- `browse_bugs` - Browse bugs with filtering and pagination
- `get_bug` - Get bug details
- `assign_bug` - Assign bug to user
- `confirm_bug` - Confirm a bug
- `resolve_bug` - Resolve a bug
- `activate_bug` - Activate a closed bug
- `close_bug` - Close a bug
- `delete_bug` - Delete a bug
- `export_bugs` - Export bugs to file
- `report_bugs` - Generate bug report
- `link_bugs` - Link related bugs
- `confirm_bug_story_change` - Confirm story change for a bug
- Batch operations: `batch_create_bugs`, `batch_edit_bugs`, `batch_change_bug_branch`, `batch_change_bug_module`, `batch_change_bug_plan`, `batch_assign_bugs`, `batch_confirm_bugs`, `batch_resolve_bugs`, `batch_close_bugs`, `batch_activate_bugs`
- And more...

### Test Cases (30+ tools)
- `create_testcase` - Create a new test case
- `update_testcase` - Update an existing test case
- `browse_testcases` - Browse test cases with filtering
- `view_testcase` - View test case details
- `delete_testcase` - Delete a test case
- `review_testcase` - Review a test case
- `create_bug_from_testcase` - Create a bug from a test case
- `link_testcases` - Link related test cases
- `link_bugs_to_testcase` - Link bugs to a test case
- Batch operations: `batch_create_testcases`, `batch_edit_testcases`, `batch_delete_testcases`, `batch_review_testcases`, `batch_change_testcase_branch`, `batch_change_testcase_module`, `batch_change_testcase_type`
- Import/Export: `export_testcases`, `export_testcase_template`, `import_testcases`, `import_testcases_from_lib`, `import_testcase_to_lib`
- Scene management: `browse_testcase_scenes`, `create_testcase_scene`, `edit_testcase_scene`, `delete_testcase_scene`
- Analysis: `group_testcases`, `get_zero_testcases`
- And more...

### Test Tasks (20+ tools)
- `create_testtask` - Create a new test task
- `get_testtasks` - Get list of test tasks
- `browse_testtasks` - Browse test tasks with filtering
- `get_testtask` - Get test task details
- `edit_testtask` - Edit an existing test task
- `start_testtask` - Start a test task
- `close_testtask` - Close a test task
- `block_testtask` - Block a test task
- `activate_testtask` - Activate a blocked test task
- `delete_testtask` - Delete a test task
- `get_testtask_cases` - Get test cases for a test task
- `get_testtask_unit_cases` - Get unit test cases for a test task
- `link_case_to_testtask` - Link test case to test task
- `unlink_case_from_testtask` - Unlink test case from test task
- `batch_unlink_cases_from_testtask` - Unlink multiple test cases
- `run_testcase` - Run a test case in test task
- `batch_run_testcases` - Run multiple test cases
- `get_testtask_results` - Get test results for a test task
- `assign_testcase` - Assign test case to user
- `batch_assign_testcases` - Assign multiple test cases
- `report_testtask` - Generate test task report
- `group_testtask_cases` - Group test cases in test task
- `browse_testtask_units` - Browse unit test tasks
- `import_unit_test_result` - Import unit test result
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

### Case Library (15 tools)
- `get_caselib_index` - Get case library index
- `create_caselib` - Create a new case library
- `edit_caselib` - Edit an existing case library
- `delete_caselib` - Delete a case library
- `browse_caselib` - Browse case library with filtering
- `view_caselib` - View case library details
- `create_caselib_case` - Create a new test case in case library
- `batch_create_caselib_cases` - Create multiple test cases
- `edit_caselib_case` - Edit a test case in case library
- `batch_edit_caselib_cases` - Edit multiple test cases
- `view_caselib_case` - View test case details in case library
- `export_caselib_template` - Export import template
- `import_caselib_cases` - Import test cases to case library
- `show_caselib_import` - Show import preview
- `export_caselib_cases` - Export test cases from case library

### QA (1 tool)
- `get_qa_index` - Get QA module index

### Test Report (5 tools)
- `browse_testreports` - Browse test reports with filtering
- `create_testreport` - Create a new test report
- `edit_testreport` - Edit an existing test report
- `view_testreport` - View test report details
- `delete_testreport` - Delete a test report

### Test Suite (9 tools)
- `get_testsuite_index` - Get test suite index
- `browse_testsuites` - Browse test suites with filtering
- `create_testsuite` - Create a new test suite
- `view_testsuite` - View test suite details
- `edit_testsuite` - Edit an existing test suite
- `delete_testsuite` - Delete a test suite
- `link_case_to_testsuite` - Link test case to test suite
- `unlink_case_from_testsuite` - Unlink test case from test suite
- `batch_unlink_cases_from_testsuite` - Unlink multiple test cases

## Resources (46 Total + 100 Templates)

Resources provide URI-based access to ZenTao data. All resources follow the pattern `zentao://{entity}/{id}` or `zentao://{entity}/{id}/{subentity}`.

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

### Case Library
- `zentao://caselibs` - List of all case libraries
- `zentao://caselibs/{libID}` - Individual case library details
- `zentao://caselibs/{libID}/cases` - Test cases in a case library
- `zentao://caselibs/{libID}/cases/{caseID}` - Individual test case in case library

### QA
- `zentao://qa` - QA module index

### Test Report
- `zentao://testreports/{reportID}` - Individual test report details
- `zentao://{objectType}/{objectID}/testreports` - Test reports for a project, execution, or product

### Test Suite
- `zentao://testsuites/{suiteID}` - Individual test suite details
- `zentao://products/{productID}/testsuites` - Test suites for a product

## Prompts (2 Total)

MCP Prompts provide guided workflows for common operations:

1. **`create_product`** - Guided workflow for creating a new product
   - Required arguments: `name`, `code`
   - Optional arguments: `program`, `line`, `PO`, `QD`, `RD`, `type`, `desc`, `acl`

2. **`create_story`** - Guided workflow for creating a new user story
   - Required arguments: `title`, `product`
   - Optional arguments: `module`, `plan`, `source`, `pri`, `estimate`, `spec`, `verify`

## Usage Examples

### Using Tools

Create a new product:
```json
{
  "name": "create_product",
  "arguments": {
    "name": "My Product",
    "code": "MP001"
  }
}
```

Create a user story:
```json
{
  "name": "create_story",
  "arguments": {
    "title": "User login functionality",
    "product": 1,
    "pri": 1,
    "spec": "Users should be able to login with email and password"
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

Get project kanban board:
```
zentao://projects/456/kanban
```

### Using Prompts

Get guided workflow for creating a product:
```
prompt://create_product?name=MyProduct&code=MP001
```

## API Documentation

Complete ZenTao REST API documentation is available in the `api_doc.txt` file, which includes:
- 82+ API endpoints
- Request/response schemas
- Required vs optional parameters
- Example requests

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

Commercial licensing available upon request.

## Contact

- **Author:** Bivex
- **Email:** support@b-b.top
- **Website:** https://contact.b-b.top
- **GitHub:** https://github.com/bivex

---

**Keywords:** ZenTao, MCP, Model Context Protocol, Project Management, Bug Tracking, Agile, Scrum, Kanban, Go, REST API, AI Integration, LLM, Claude, GPT, ZenTao API, Project Management System, Agile Development, Software Development, DevOps, CI/CD, Test Management, Requirement Management
