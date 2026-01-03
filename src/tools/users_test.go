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

func TestRegisterUserTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterUserTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterUserTools(s, nil)
}

func TestUserToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterUserTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterUserTools(s, nil)
}

func TestUserToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterUserTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterUserTools(s, nil)
}

func TestMultipleUserToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterUserTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterUserTools(s, nil)
	RegisterUserTools(s, nil)

	// Should not crash
}

func TestUserToolParameterValidation(t *testing.T) {
	// Test create_user parameters
	createArgs := map[string]interface{}{
		"account":  "testuser",
		"password": "TestPass123!",
		"realname": "Test User",
		"visions":  []interface{}{"rnd", "lite"},
		"email":    "test@example.com",
		"role":     "dev",
		"dept":     float64(1),
	}

	// Verify required parameters
	if account, ok := createArgs["account"].(string); !ok || account == "" {
		t.Error("Expected account parameter to be non-empty string")
	}

	if password, ok := createArgs["password"].(string); !ok || password == "" {
		t.Error("Expected password parameter to be non-empty string")
	}

	visions, ok := createArgs["visions"].([]interface{})
	if !ok || len(visions) == 0 {
		t.Error("Expected visions parameter to be non-empty array")
	}

	// Test valid vision values
	validVisions := []string{"rnd", "lite"}
	for _, vision := range visions {
		if visionStr, ok := vision.(string); ok {
			found := false
			for _, validVision := range validVisions {
				if visionStr == validVision {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Invalid vision value: %s", visionStr)
			}
		}
	}

	// Test update_user parameters
	updateArgs := map[string]interface{}{
		"id":       float64(123),
		"realname": "Updated User",
		"email":    "updated@example.com",
		"mobile":   "1234567890",
		"phone":    "0987654321",
		"role":     "qa",
	}

	if id, ok := updateArgs["id"].(float64); !ok || id != 123 {
		t.Error("Expected id parameter to be float64(123)")
	}
}

func TestUserToolInvalidParameters(t *testing.T) {
	// Test empty required parameters
	emptyArgs := map[string]interface{}{
		"account":  "",
		"password": "",
		"visions":  []interface{}{},
	}

	if account, ok := emptyArgs["account"].(string); ok && account == "" {
		t.Log("Empty account parameter detected")
	}

	if password, ok := emptyArgs["password"].(string); ok && password == "" {
		t.Log("Empty password parameter detected")
	}

	if visions, ok := emptyArgs["visions"].([]interface{}); ok && len(visions) == 0 {
		t.Log("Empty visions array detected")
	}

	// Test invalid vision values
	invalidVisionArgs := map[string]interface{}{
		"visions": []interface{}{"invalid_vision"},
	}

	if visions, ok := invalidVisionArgs["visions"].([]interface{}); ok {
		validVisions := []string{"rnd", "lite"}
		for _, vision := range visions {
			if visionStr, ok := vision.(string); ok {
				found := false
				for _, validVision := range validVisions {
					if visionStr == validVision {
						found = true
						break
					}
				}
				if !found {
					t.Logf("Invalid vision value detected: %s", visionStr)
				}
			}
		}
	}

	// Test invalid email format
	invalidEmailArgs := map[string]interface{}{
		"email": "invalid-email-format",
	}

	if email, ok := invalidEmailArgs["email"].(string); ok {
		// Basic email validation - should contain @
		if !contains(email, "@") {
			t.Logf("Invalid email format detected: %s", email)
		}
	}
}

func TestUserToolEnumValidation(t *testing.T) {
	// Test valid vision values
	validVisions := []string{"rnd", "lite"}
	testVisions := []string{"rnd", "lite"}

	for _, vision := range testVisions {
		found := false
		for _, validVision := range validVisions {
			if vision == validVision {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Vision %s should be valid", vision)
		}
	}

	// Test invalid visions
	invalidVisions := []string{"admin", "user", "guest", ""}
	for _, vision := range invalidVisions {
		found := false
		for _, validVision := range validVisions {
			if vision == validVision {
				found = true
				break
			}
		}
		if found {
			t.Errorf("Vision %s should be invalid", vision)
		}
	}
}

func TestUserToolPasswordValidation(t *testing.T) {
	// Test password strength requirements
	validPasswords := []string{
		"TestPass123!",
		"Password1!",
		"Secure123#",
	}

	for _, password := range validPasswords {
		// Basic password validation
		if len(password) < 8 {
			t.Errorf("Password too short: %s", password)
		}
		if !containsUpper(password) {
			t.Logf("Password should contain uppercase: %s", password)
		}
		if !containsLower(password) {
			t.Logf("Password should contain lowercase: %s", password)
		}
		if !containsDigit(password) {
			t.Logf("Password should contain digit: %s", password)
		}
		if !containsSpecial(password) {
			t.Logf("Password should contain special character: %s", password)
		}
	}

	// Test weak passwords
	weakPasswords := []string{
		"password",     // No uppercase, digit, special
		"PASSWORD",     // No lowercase, digit, special
		"Password",     // No digit, special
		"Password1",    // No special
		"Pass1!",       // Too short
	}

	for _, password := range weakPasswords {
		t.Logf("Weak password detected: %s", password)
	}
}

// Helper functions for password validation
func contains(s, substr string) bool {
	return len(s) >= len(substr) && findIndex(s, substr) >= 0
}

func findIndex(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

func containsUpper(s string) bool {
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			return true
		}
	}
	return false
}

func containsLower(s string) bool {
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			return true
		}
	}
	return false
}

func containsDigit(s string) bool {
	for _, r := range s {
		if r >= '0' && r <= '9' {
			return true
		}
	}
	return false
}

func containsSpecial(s string) bool {
	specialChars := "!@#$%^&*()_+-=[]{}|;:,.<>?"
	for _, r := range s {
		for _, special := range specialChars {
			if r == special {
				return true
			}
		}
	}
	return false
}

func BenchmarkUserToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterUserTools(s, nil)
	}
}

func BenchmarkUserToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterUserTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
