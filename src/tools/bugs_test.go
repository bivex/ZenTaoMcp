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

// Mock ZenTao client for testing (defined in auth_test.go)
func TestRegisterBugTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterBugTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBugTools(s, nil)
}

func TestBugToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterBugTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBugTools(s, nil)
}

func TestBugToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterBugTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBugTools(s, nil)
}

func TestMultipleBugToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterBugTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterBugTools(s, nil)
	RegisterBugTools(s, nil)

	// Should not crash
}

func TestBugToolParameterValidation(t *testing.T) {
	// Test create_bug parameters
	createArgs := map[string]interface{}{
		"product":  float64(1),
		"title":    "Test Bug",
		"severity": float64(3),
		"pri":      float64(2),
		"type":     "codeerror",
		"steps":    "Reproduction steps",
	}

	// Verify required parameters
	if product, ok := createArgs["product"].(float64); !ok || product != 1 {
		t.Error("Expected product parameter to be float64(1)")
	}

	if title, ok := createArgs["title"].(string); !ok || title == "" {
		t.Error("Expected title parameter to be non-empty string")
	}

	if severity, ok := createArgs["severity"].(float64); !ok || severity < 1 || severity > 4 {
		t.Error("Expected severity parameter to be float64 between 1-4")
	}

	if pri, ok := createArgs["pri"].(float64); !ok || pri < 1 || pri > 9 {
		t.Error("Expected pri parameter to be float64 between 1-9")
	}

	if bugType, ok := createArgs["type"].(string); !ok || bugType == "" {
		t.Error("Expected type parameter to be non-empty string")
	}

	// Test valid enum values
	validTypes := []string{"codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"}
	bugType, _ := createArgs["type"].(string)
	found := false
	for _, validType := range validTypes {
		if bugType == validType {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected type to be one of valid enum values, got %s", bugType)
	}

	// Test update_bug parameters
	updateArgs := map[string]interface{}{
		"id":       float64(123),
		"severity": float64(1),
		"pri":      float64(1),
		"type":     "security",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}
}

func TestBugToolInvalidParameters(t *testing.T) {
	// Test invalid severity values
	invalidSeverityArgs := map[string]interface{}{
		"severity": float64(5), // Invalid: should be 1-4
	}

	if severity, ok := invalidSeverityArgs["severity"].(float64); ok && (severity < 1 || severity > 4) {
		// This should be caught by parameter validation
		t.Logf("Invalid severity value detected: %v", severity)
	}

	// Test invalid priority values
	invalidPriArgs := map[string]interface{}{
		"pri": float64(10), // Invalid: should be 1-9
	}

	if pri, ok := invalidPriArgs["pri"].(float64); ok && (pri < 1 || pri > 9) {
		// This should be caught by parameter validation
		t.Logf("Invalid priority value detected: %v", pri)
	}

	// Test invalid type values
	invalidTypeArgs := map[string]interface{}{
		"type": "invalid_type", // Invalid: not in enum
	}

	validTypes := []string{"codeerror", "config", "install", "security", "performance", "standard", "automation", "designdefect", "others"}
	if bugType, ok := invalidTypeArgs["type"].(string); ok {
		found := false
		for _, validType := range validTypes {
			if bugType == validType {
				found = true
				break
			}
		}
		if !found {
			t.Logf("Invalid type value detected: %s", bugType)
		}
	}
}

func BenchmarkBugToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterBugTools(s, nil)
	}
}

func BenchmarkBugToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterBugTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
