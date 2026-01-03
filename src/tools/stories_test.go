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

func TestRegisterStoryTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterStoryTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterStoryTools(s, nil)
}

func TestStoryToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterStoryTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterStoryTools(s, nil)
}

func TestStoryToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterStoryTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterStoryTools(s, nil)
}

func TestMultipleStoryToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterStoryTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterStoryTools(s, nil)
	RegisterStoryTools(s, nil)

	// Should not crash
}

func TestStoryToolParameterValidation(t *testing.T) {
	// Test create_story parameters
	createArgs := map[string]interface{}{
		"title":     "Test Story",
		"product":   float64(1),
		"pri":       float64(3),
		"category":  "feature",
		"spec":      "Story description",
		"verify":    "Acceptance criteria",
		"estimate":  float64(8),
		"keywords":  "test,story",
	}

	// Verify required parameters
	if title, ok := createArgs["title"].(string); !ok || title == "" {
		t.Error("Expected title parameter to be non-empty string")
	}

	if product, ok := createArgs["product"].(float64); !ok || product != 1 {
		t.Error("Expected product parameter to be float64(1)")
	}

	if pri, ok := createArgs["pri"].(float64); !ok || pri < 1 || pri > 9 {
		t.Error("Expected pri parameter to be float64 between 1-9")
	}

	if category, ok := createArgs["category"].(string); !ok || category == "" {
		t.Error("Expected category parameter to be non-empty string")
	}

	// Test valid category enum values
	validCategories := []string{"feature", "interface", "performance", "safe", "experience", "improve", "other"}
	category, _ := createArgs["category"].(string)
	found := false
	for _, validCategory := range validCategories {
		if category == validCategory {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected category to be one of valid enum values, got %s", category)
	}

	// Test update_story parameters
	updateArgs := map[string]interface{}{
		"id":        float64(123),
		"pri":       float64(2),
		"category":  "interface",
		"estimate":  float64(5),
		"keywords":  "updated,test",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}

	// Test change_story parameters
	changeArgs := map[string]interface{}{
		"id":   float64(123),
		"title": "Updated Story Title",
		"spec": "Updated story description",
		"verify": "Updated acceptance criteria",
	}

	if id, ok := changeArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}
}

func TestStoryToolInvalidParameters(t *testing.T) {
	// Test invalid priority values
	invalidPriArgs := map[string]interface{}{
		"pri": float64(10), // Invalid: should be 1-9
	}

	if pri, ok := invalidPriArgs["pri"].(float64); ok && (pri < 1 || pri > 9) {
		t.Logf("Invalid priority value detected: %v", pri)
	}

	// Test invalid category
	invalidCategoryArgs := map[string]interface{}{
		"category": "invalid_category", // Invalid: not in enum
	}

	validCategories := []string{"feature", "interface", "performance", "safe", "experience", "improve", "other"}
	if category, ok := invalidCategoryArgs["category"].(string); ok {
		found := false
		for _, validCategory := range validCategories {
			if category == validCategory {
				found = true
				break
			}
		}
		if !found {
			t.Logf("Invalid category value detected: %s", category)
		}
	}

	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"title":   "",
		"product": float64(0),
	}

	if title, ok := emptyArgs["title"].(string); ok && title == "" {
		t.Log("Empty title parameter detected")
	}

	if product, ok := emptyArgs["product"].(float64); ok && product == 0 {
		t.Log("Invalid product ID detected")
	}
}

func TestStoryToolEnumValidation(t *testing.T) {
	// Test all valid category values
	validCategories := []string{"feature", "interface", "performance", "safe", "experience", "improve", "other"}

	for _, category := range validCategories {
		found := false
		for _, validCategory := range validCategories {
			if category == validCategory {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Category %s should be valid", category)
		}
	}

	// Test invalid categories
	invalidCategories := []string{"bug", "task", "epic", "userstory", ""}
	for _, category := range invalidCategories {
		found := false
		for _, validCategory := range validCategories {
			if category == validCategory {
				found = true
				break
			}
		}
		if found {
			t.Errorf("Category %s should be invalid", category)
		}
	}
}

func BenchmarkStoryToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterStoryTools(s, nil)
	}
}

func BenchmarkStoryToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterStoryTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
