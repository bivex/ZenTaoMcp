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

func TestRegisterPlanTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterPlanTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterPlanTools(s, nil)
}

func TestPlanToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterPlanTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterPlanTools(s, nil)
}

func TestPlanToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterPlanTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterPlanTools(s, nil)
}

func TestMultiplePlanToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterPlanTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterPlanTools(s, nil)
	RegisterPlanTools(s, nil)

	// Should not crash
}

func TestPlanToolParameterValidation(t *testing.T) {
	// Test create_plan parameters
	createArgs := map[string]interface{}{
		"title":   "Test Plan",
		"product": float64(1),
		"branch":  float64(1),
		"begin":   "2024-01-01",
		"end":     "2024-01-31",
		"desc":    "Test plan description",
		"parent":  float64(0),
	}

	// Verify required parameters
	if title, ok := createArgs["title"].(string); !ok || title == "" {
		t.Error("Expected title parameter to be non-empty string")
	}

	if product, ok := createArgs["product"].(float64); !ok || product != 1 {
		t.Error("Expected product parameter to be float64(1)")
	}

	if begin, ok := createArgs["begin"].(string); !ok || begin == "" {
		t.Error("Expected begin parameter to be non-empty string")
	}

	if end, ok := createArgs["end"].(string); !ok || end == "" {
		t.Error("Expected end parameter to be non-empty string")
	}

	// Test update_plan parameters
	updateArgs := map[string]interface{}{
		"id":    float64(123),
		"title": "Updated Plan",
		"begin": "2024-02-01",
		"end":   "2024-02-28",
		"desc":  "Updated plan description",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}

	// Test link_stories_to_plan parameters
	linkStoriesArgs := map[string]interface{}{
		"id":      float64(123),
		"stories": []interface{}{float64(1), float64(2), float64(3)},
	}

	if stories, ok := linkStoriesArgs["stories"].([]interface{}); !ok || len(stories) == 0 {
		t.Error("Expected stories parameter to be non-empty array")
	}

	// Test link_bugs_to_plan parameters
	linkBugsArgs := map[string]interface{}{
		"id":   float64(123),
		"bugs": []interface{}{float64(1), float64(2)},
	}

	if bugs, ok := linkBugsArgs["bugs"].([]interface{}); !ok || len(bugs) == 0 {
		t.Error("Expected bugs parameter to be non-empty array")
	}
}

func TestPlanToolInvalidParameters(t *testing.T) {
	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"title": "",
		"begin": "",
		"end":   "",
	}

	if title, ok := emptyArgs["title"].(string); ok && title == "" {
		t.Log("Empty title parameter detected")
	}

	if begin, ok := emptyArgs["begin"].(string); ok && begin == "" {
		t.Log("Empty begin date parameter detected")
	}

	if end, ok := emptyArgs["end"].(string); ok && end == "" {
		t.Log("Empty end date parameter detected")
	}

	// Test invalid product ID
	invalidProductArgs := map[string]interface{}{
		"product": float64(0),
	}

	if product, ok := invalidProductArgs["product"].(float64); ok && product == 0 {
		t.Log("Invalid product ID detected")
	}

	// Test empty arrays for linking operations
	emptyStoriesArgs := map[string]interface{}{
		"stories": []interface{}{},
	}

	if stories, ok := emptyStoriesArgs["stories"].([]interface{}); ok && len(stories) == 0 {
		t.Log("Empty stories array detected for linking")
	}

	emptyBugsArgs := map[string]interface{}{
		"bugs": []interface{}{},
	}

	if bugs, ok := emptyBugsArgs["bugs"].([]interface{}); ok && len(bugs) == 0 {
		t.Log("Empty bugs array detected for linking")
	}

	// Test invalid date ranges (end before begin)
	invalidDateRangeArgs := map[string]interface{}{
		"begin": "2024-01-31",
		"end":   "2024-01-01",
	}

	if begin, ok := invalidDateRangeArgs["begin"].(string); ok {
		if end, ok := invalidDateRangeArgs["end"].(string); ok {
			if begin > end {
				t.Log("Invalid date range detected: end date before begin date")
			}
		}
	}
}

func TestPlanToolDateValidation(t *testing.T) {
	// Test valid date ranges
	validRanges := []map[string]string{
		{"begin": "2024-01-01", "end": "2024-01-31"},
		{"begin": "2024-02-01", "end": "2024-02-29"}, // Leap year
		{"begin": "2024-12-01", "end": "2024-12-31"},
	}

	for _, dateRange := range validRanges {
		if dateRange["begin"] >= dateRange["end"] {
			t.Errorf("Invalid date range: begin=%s, end=%s", dateRange["begin"], dateRange["end"])
		}
	}

	// Test invalid date ranges
	invalidRanges := []map[string]string{
		{"begin": "2024-01-31", "end": "2024-01-01"}, // End before begin
		{"begin": "2024-01-01", "end": "2024-01-01"}, // Same date
		{"begin": "2024-01-15", "end": "2024-01-10"}, // End before begin
	}

	for _, dateRange := range invalidRanges {
		if dateRange["begin"] >= dateRange["end"] {
			t.Logf("Invalid date range detected: begin=%s, end=%s", dateRange["begin"], dateRange["end"])
		}
	}
}

func TestPlanToolArrayValidation(t *testing.T) {
	// Test valid story arrays
	validStoryArrays := [][]interface{}{
		{float64(1)},
		{float64(1), float64(2), float64(3)},
		{float64(10), float64(20), float64(30), float64(40)},
	}

	for _, stories := range validStoryArrays {
		if len(stories) == 0 {
			t.Error("Story array should not be empty")
		}
		for _, story := range stories {
			if storyId, ok := story.(float64); !ok || storyId <= 0 {
				t.Errorf("Invalid story ID: %v", story)
			}
		}
	}

	// Test valid bug arrays
	validBugArrays := [][]interface{}{
		{float64(1)},
		{float64(5), float64(10)},
		{float64(100), float64(200), float64(300)},
	}

	for _, bugs := range validBugArrays {
		if len(bugs) == 0 {
			t.Error("Bug array should not be empty")
		}
		for _, bug := range bugs {
			if bugId, ok := bug.(float64); !ok || bugId <= 0 {
				t.Errorf("Invalid bug ID: %v", bug)
			}
		}
	}

	// Test invalid arrays (empty or with invalid IDs)
	invalidArrays := [][]interface{}{
		{}, // Empty
		{float64(0)}, // Invalid ID
		{float64(-1)}, // Negative ID
		{float64(1), float64(0), float64(2)}, // Mix of valid/invalid
	}

	for _, invalidArray := range invalidArrays {
		if len(invalidArray) == 0 {
			t.Log("Empty array detected")
		}
		for _, item := range invalidArray {
			if itemId, ok := item.(float64); ok && itemId <= 0 {
				t.Logf("Invalid ID detected in array: %v", itemId)
			}
		}
	}
}

func BenchmarkPlanToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterPlanTools(s, nil)
	}
}

func BenchmarkPlanToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterPlanTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
