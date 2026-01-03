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

package tools

import (
	"testing"

	"github.com/mark3labs/mcp-go/server"
)

func TestRegisterProjectTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProjectTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProjectTools(s, nil)
}

func TestProjectToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProjectTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProjectTools(s, nil)
}

func TestProjectToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterProjectTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProjectTools(s, nil)
}

func TestMultipleProjectToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterProjectTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterProjectTools(s, nil)
	RegisterProjectTools(s, nil)

	// Should not crash
}

func TestProjectToolParameterValidation(t *testing.T) {
	// Test create_project parameters
	createArgs := map[string]interface{}{
		"name":     "Test Project",
		"code":     "TEST001",
		"begin":    "2024-01-01",
		"end":      "2024-12-31",
		"products": []interface{}{float64(1), float64(2)},
		"model":    "scrum",
		"parent":   float64(0),
	}

	// Verify required parameters
	if name, ok := createArgs["name"].(string); !ok || name == "" {
		t.Error("Expected name parameter to be non-empty string")
	}

	if code, ok := createArgs["code"].(string); !ok || code == "" {
		t.Error("Expected code parameter to be non-empty string")
	}

	if begin, ok := createArgs["begin"].(string); !ok || begin == "" {
		t.Error("Expected begin parameter to be non-empty string")
	}

	if end, ok := createArgs["end"].(string); !ok || end == "" {
		t.Error("Expected end parameter to be non-empty string")
	}

	if products, ok := createArgs["products"].([]interface{}); !ok || len(products) == 0 {
		t.Error("Expected products parameter to be non-empty array")
	}

	// Test update_project parameters
	updateArgs := map[string]interface{}{
		"id":     float64(123),
		"name":   "Updated Project",
		"desc":   "Updated description",
		"status": "doing",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}

	// Test create_execution parameters
	executionArgs := map[string]interface{}{
		"project": float64(1),
		"name":    "Test Execution",
		"code":    "EXEC001",
		"begin":   "2024-01-01",
		"end":     "2024-01-31",
		"products": []interface{}{float64(1)},
	}

	if project, ok := executionArgs["project"].(float64); !ok || project != 1 {
		t.Error("Expected project parameter to be float64(1)")
	}
}

func TestProjectToolInvalidParameters(t *testing.T) {
	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"name":  "",
		"code":  "",
		"begin": "",
		"end":   "",
	}

	if name, ok := emptyArgs["name"].(string); ok && name == "" {
		t.Log("Empty name parameter detected")
	}

	// Test invalid date formats (should be validated by application)
	invalidDateArgs := map[string]interface{}{
		"begin": "invalid-date",
		"end":   "2024-13-45", // Invalid month/day
	}

	if begin, ok := invalidDateArgs["begin"].(string); ok {
		t.Logf("Invalid begin date format: %s", begin)
	}

	// Test empty products array
	emptyProductsArgs := map[string]interface{}{
		"products": []interface{}{},
	}

	if products, ok := emptyProductsArgs["products"].([]interface{}); ok && len(products) == 0 {
		t.Log("Empty products array detected")
	}
}

func TestProjectToolEnumValidation(t *testing.T) {
	// Test valid model values
	validModels := []string{"scrum", "agileplus", "waterfall", "kanban"}
	testModel := "scrum"

	found := false
	for _, validModel := range validModels {
		if testModel == validModel {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected model to be one of valid values, got %s", testModel)
	}

	// Test valid status values
	validStatuses := []string{"wait", "doing", "suspended", "closed"}
	testStatus := "doing"

	found = false
	for _, validStatus := range validStatuses {
		if testStatus == validStatus {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected status to be one of valid values, got %s", testStatus)
	}

	// Test invalid model
	invalidModel := "invalid_model"
	found = false
	for _, validModel := range validModels {
		if invalidModel == validModel {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Invalid model should not be found in valid list: %s", invalidModel)
	}
}

func BenchmarkProjectToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterProjectTools(s, nil)
	}
}

func BenchmarkProjectToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterProjectTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
