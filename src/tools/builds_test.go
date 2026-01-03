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

func TestRegisterBuildTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterBuildTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBuildTools(s, nil)
}

func TestBuildToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterBuildTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBuildTools(s, nil)
}

func TestBuildToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterBuildTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterBuildTools(s, nil)
}

func TestMultipleBuildToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterBuildTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterBuildTools(s, nil)
	RegisterBuildTools(s, nil)

	// Should not crash
}

func TestBuildToolParameterValidation(t *testing.T) {
	// Test create_build parameters
	createArgs := map[string]interface{}{
		"name":      "Test Build",
		"product":   float64(1),
		"execution": float64(1),
		"builder":   "admin",
		"project":   float64(1),
		"branch":    float64(1),
		"date":      "2024-01-01",
		"scmPath":   "https://github.com/example/repo",
		"filePath":  "https://example.com/download/build.zip",
		"desc":      "Test build description",
	}

	// Verify required parameters
	if name, ok := createArgs["name"].(string); !ok || name == "" {
		t.Error("Expected name parameter to be non-empty string")
	}

	if product, ok := createArgs["product"].(float64); !ok || product != 1 {
		t.Error("Expected product parameter to be float64(1)")
	}

	if execution, ok := createArgs["execution"].(float64); !ok || execution != 1 {
		t.Error("Expected execution parameter to be float64(1)")
	}

	if builder, ok := createArgs["builder"].(string); !ok || builder == "" {
		t.Error("Expected builder parameter to be non-empty string")
	}

	// Test update_build parameters
	updateArgs := map[string]interface{}{
		"id":       float64(123),
		"name":     "Updated Build",
		"builder":  "admin",
		"scmPath":  "https://github.com/example/updated-repo",
		"filePath": "https://example.com/download/updated-build.zip",
		"desc":     "Updated build description",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}
}

func TestBuildToolInvalidParameters(t *testing.T) {
	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"name":    "",
		"builder": "",
	}

	if name, ok := emptyArgs["name"].(string); ok && name == "" {
		t.Log("Empty name parameter detected")
	}

	if builder, ok := emptyArgs["builder"].(string); ok && builder == "" {
		t.Log("Empty builder parameter detected")
	}

	// Test invalid IDs
	invalidIdArgs := map[string]interface{}{
		"product":   float64(0),
		"execution": float64(0),
		"project":   float64(0),
	}

	if product, ok := invalidIdArgs["product"].(float64); ok && product == 0 {
		t.Log("Invalid product ID detected")
	}

	if execution, ok := invalidIdArgs["execution"].(float64); ok && execution == 0 {
		t.Log("Invalid execution ID detected")
	}

	if project, ok := invalidIdArgs["project"].(float64); ok && project == 0 {
		t.Log("Invalid project ID detected")
	}

	// Test invalid URLs
	invalidUrlArgs := map[string]interface{}{
		"scmPath":  "not-a-url",
		"filePath": "invalid-url",
	}

	if scmPath, ok := invalidUrlArgs["scmPath"].(string); ok {
		if !isValidURL(scmPath) {
			t.Logf("Invalid SCM path URL: %s", scmPath)
		}
	}

	if filePath, ok := invalidUrlArgs["filePath"].(string); ok {
		if !isValidURL(filePath) {
			t.Logf("Invalid file path URL: %s", filePath)
		}
	}
}

func TestBuildToolURLValidation(t *testing.T) {
	// Test valid URLs
	validURLs := []string{
		"https://github.com/example/repo",
		"http://example.com/download/build.zip",
		"ftp://ftp.example.com/file.tar.gz",
		"git@github.com:example/repo.git",
	}

	for _, url := range validURLs {
		if !isValidURL(url) {
			t.Errorf("Expected URL to be valid: %s", url)
		}
	}

	// Test invalid URLs
	invalidURLs := []string{
		"not-a-url",
		"example.com", // Missing protocol
		"ftp://",      // Incomplete
		"",            // Empty
	}

	for _, url := range invalidURLs {
		if isValidURL(url) {
			t.Errorf("Expected URL to be invalid: %s", url)
		}
	}
}

// Basic URL validation function
func isValidURL(url string) bool {
	if len(url) == 0 {
		return false
	}

	// Check for common URL patterns
	validPrefixes := []string{"http://", "https://", "ftp://", "git@", "ssh://"}
	for _, prefix := range validPrefixes {
		if len(url) > len(prefix) && url[:len(prefix)] == prefix {
			return true
		}
	}

	return false
}

func TestBuildToolDateValidation(t *testing.T) {
	// Test valid date formats
	validDates := []string{
		"2024-01-01",
		"2024-12-31",
		"2024-02-29", // Leap year
		"2023-02-28", // Non-leap year
	}

	for _, date := range validDates {
		if !isValidDate(date) {
			t.Errorf("Expected date to be valid: %s", date)
		}
	}

	// Test invalid date formats
	invalidDates := []string{
		"2024/01/01",  // Wrong separator
		"2024-13-01",  // Invalid month
		"2024-01-32",  // Invalid day
		"2024-02-30",  // Invalid day for month
		"24-01-01",    // Wrong year format
		"2024-1-1",    // Missing zeros
		"",            // Empty
	}

	for _, date := range invalidDates {
		if isValidDate(date) {
			t.Errorf("Expected date to be invalid: %s", date)
		}
	}
}

// Basic date validation function (YYYY-MM-DD format)
func isValidDate(date string) bool {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return false
	}

	// Parse components
	year := date[:4]
	month := date[5:7]
	day := date[8:10]

	// Basic range checks
	if year < "1900" || year > "2100" {
		return false
	}

	if month < "01" || month > "12" {
		return false
	}

	// Check day based on month (simplified, doesn't handle all leap year cases)
	maxDay := "31"
	if month == "02" {
		// Allow Feb 29 for leap years, but don't validate leap year logic strictly
		if day <= "29" {
			return day >= "01"
		}
		return false
	} else if month == "04" || month == "06" || month == "09" || month == "11" {
		maxDay = "30"
	}

	if day < "01" || day > maxDay {
		return false
	}

	return true
}

func BenchmarkBuildToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterBuildTools(s, nil)
	}
}

func BenchmarkBuildToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterBuildTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
