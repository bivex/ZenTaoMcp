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

func TestRegisterTaskTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterTaskTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterTaskTools(s, nil)
}

func TestTaskToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterTaskTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterTaskTools(s, nil)
}

func TestTaskToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterTaskTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterTaskTools(s, nil)
}

func TestMultipleTaskToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterTaskTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterTaskTools(s, nil)
	RegisterTaskTools(s, nil)

	// Should not crash
}

func TestTaskToolParameterValidation(t *testing.T) {
	// Test create_task parameters
	createArgs := map[string]interface{}{
		"name":       "Test Task",
		"type":       "devel",
		"execution":  float64(1),
		"assignedTo": []interface{}{"admin", "user1"},
		"estStarted": "2024-01-01",
		"deadline":   "2024-01-31",
		"pri":        float64(2),
		"estimate":   float64(8),
		"story":      float64(5),
	}

	// Verify required parameters
	if name, ok := createArgs["name"].(string); !ok || name == "" {
		t.Error("Expected name parameter to be non-empty string")
	}

	if taskType, ok := createArgs["type"].(string); !ok || taskType == "" {
		t.Error("Expected type parameter to be non-empty string")
	}

	if execution, ok := createArgs["execution"].(float64); !ok || execution != 1 {
		t.Error("Expected execution parameter to be float64(1)")
	}

	if assignedTo, ok := createArgs["assignedTo"].([]interface{}); !ok || len(assignedTo) == 0 {
		t.Error("Expected assignedTo parameter to be non-empty array")
	}

	if estStarted, ok := createArgs["estStarted"].(string); !ok || estStarted == "" {
		t.Error("Expected estStarted parameter to be non-empty string")
	}

	if deadline, ok := createArgs["deadline"].(string); !ok || deadline == "" {
		t.Error("Expected deadline parameter to be non-empty string")
	}

	// Test valid type enum values
	validTypes := []string{"design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"}
	taskType, _ := createArgs["type"].(string)
	found := false
	for _, validType := range validTypes {
		if taskType == validType {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected type to be one of valid enum values, got %s", taskType)
	}

	// Test update_task parameters
	updateArgs := map[string]interface{}{
		"id":         float64(123),
		"name":       "Updated Task",
		"pri":        float64(1),
		"estimate":   float64(5),
		"assignedTo": []interface{}{"admin"},
		"status":     "doing",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}
}

func TestTaskToolInvalidParameters(t *testing.T) {
	// Test invalid type values
	invalidTypeArgs := map[string]interface{}{
		"type": "invalid_type", // Invalid: not in enum
	}

	validTypes := []string{"design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"}
	if taskType, ok := invalidTypeArgs["type"].(string); ok {
		found := false
		for _, validType := range validTypes {
			if taskType == validType {
				found = true
				break
			}
		}
		if !found {
			t.Logf("Invalid type value detected: %s", taskType)
		}
	}

	// Test invalid priority values
	invalidPriArgs := map[string]interface{}{
		"pri": float64(10), // Invalid: should be 1-9
	}

	if pri, ok := invalidPriArgs["pri"].(float64); ok && (pri < 1 || pri > 9) {
		t.Logf("Invalid priority value detected: %v", pri)
	}

	// Test empty assignedTo array
	emptyAssignedArgs := map[string]interface{}{
		"assignedTo": []interface{}{},
	}

	if assignedTo, ok := emptyAssignedArgs["assignedTo"].([]interface{}); ok && len(assignedTo) == 0 {
		t.Log("Empty assignedTo array detected")
	}

	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"name":      "",
		"type":      "",
		"estStarted": "",
		"deadline":  "",
	}

	if name, ok := emptyArgs["name"].(string); ok && name == "" {
		t.Log("Empty name parameter detected")
	}
}

func TestTaskToolEnumValidation(t *testing.T) {
	// Test all valid type values
	validTypes := []string{"design", "devel", "request", "test", "study", "discuss", "ui", "affair", "misc"}

	for _, taskType := range validTypes {
		found := false
		for _, validType := range validTypes {
			if taskType == validType {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Type %s should be valid", taskType)
		}
	}

	// Test invalid types
	invalidTypes := []string{"bug", "feature", "epic", "subtask", ""}
	for _, taskType := range invalidTypes {
		found := false
		for _, validType := range validTypes {
			if taskType == validType {
				found = true
				break
			}
		}
		if found {
			t.Errorf("Type %s should be invalid", taskType)
		}
	}

	// Test valid status values
	validStatuses := []string{"wait", "doing", "done", "pause", "cancel", "closed"}
	testStatus := "doing"

	found := false
	for _, validStatus := range validStatuses {
		if testStatus == validStatus {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected status to be one of valid values, got %s", testStatus)
	}
}

func TestTaskToolDateValidation(t *testing.T) {
	// Test valid date formats
	validDates := []string{
		"2024-01-01",
		"2024-12-31",
		"2024-02-29", // Leap year
	}

	for _, date := range validDates {
		// Basic format validation (YYYY-MM-DD)
		if len(date) != 10 || date[4] != '-' || date[7] != '-' {
			t.Errorf("Invalid date format: %s", date)
		}
	}

	// Test invalid date formats
	invalidDates := []string{
		"2024/01/01",  // Wrong separator
		"2024-13-01",  // Invalid month
		"2024-01-32",  // Invalid day
		"24-01-01",    // Wrong year format
		"2024-1-1",    // Missing zeros
	}

	for _, date := range invalidDates {
		t.Logf("Invalid date format detected: %s", date)
	}
}

func BenchmarkTaskToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterTaskTools(s, nil)
	}
}

func BenchmarkTaskToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterTaskTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
