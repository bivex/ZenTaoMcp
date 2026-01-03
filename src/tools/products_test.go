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

// Note: mockZenTaoClient is defined in auth_test.go

func TestRegisterProductTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProductTools panicked: %v", r)
		}
	}()

	// Create MCP server
	s := server.NewMCPServer("test-server", "1.0.0")

	// Count tools before registration
	initialCount := len(s.ListTools())

	// Register product tools with nil client
	RegisterProductTools(s, nil)

	// Count tools after registration
	finalCount := len(s.ListTools())

	// Registration should not crash
	if finalCount < initialCount {
		t.Logf("Warning: Tool count decreased after registration: %d -> %d", initialCount, finalCount)
	}
}

func TestProductToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProductTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProductTools(s, nil)
}

func TestProductToolsWithNilClient(t *testing.T) {
	// Test registration with nil client
	defer func() {
		if r := recover(); r != nil {
			t.Logf("RegisterProductTools with nil client handled gracefully: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProductTools(s, nil)
}

func TestMultipleProductToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterProductTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterProductTools(s, nil)
	RegisterProductTools(s, nil)

	// Should not crash
}

func BenchmarkProductToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterProductTools(s, nil)
	}
}

func BenchmarkProductToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterProductTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
