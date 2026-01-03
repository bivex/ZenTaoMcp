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

// Mock ZenTao client for testing (simplified interface)
type mockZenTaoClient struct{}

func (m *mockZenTaoClient) Get(path string) ([]byte, error) {
	return []byte(`{"status": "success"}`), nil
}

func (m *mockZenTaoClient) Post(path string, body interface{}) ([]byte, error) {
	return []byte(`{"status": "success", "id": 123}`), nil
}

func (m *mockZenTaoClient) Put(path string, body interface{}) ([]byte, error) {
	return []byte(`{"status": "success"}`), nil
}

func (m *mockZenTaoClient) Delete(path string) ([]byte, error) {
	return []byte(`{"status": "success"}`), nil
}

func TestRegisterAuthTools(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterAuthTools panicked: %v", r)
		}
	}()

	// Create MCP server
	s := server.NewMCPServer("test-server", "1.0.0")

	// Count tools before registration
	initialCount := len(s.ListTools())

	// For now, just test that the function can be called
	// We can't easily pass a mock client that implements the full interface
	// So we'll test with nil to ensure no panics
	RegisterAuthTools(s, nil)

	// Count tools after registration
	finalCount := len(s.ListTools())

	// Registration should not crash, even with nil client
	if finalCount < initialCount {
		t.Logf("Warning: Tool count decreased after registration: %d -> %d", initialCount, finalCount)
	}
}

func TestAuthToolRegistration(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterAuthTools panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterAuthTools(s, nil)
}

func TestMultipleAuthToolRegistrations(t *testing.T) {
	// Test registering tools multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple RegisterAuthTools calls panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterAuthTools(s, nil)
	RegisterAuthTools(s, nil)

	// Should not crash
}

func TestAuthToolsWithNilClient(t *testing.T) {
	// Test registration with nil client (should handle gracefully)
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterAuthTools with nil client panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterAuthTools(s, nil)
}

func BenchmarkAuthToolRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterAuthTools(s, nil)
	}
}

func BenchmarkAuthToolList(b *testing.B) {
	s := server.NewMCPServer("bench-server", "1.0.0")
	RegisterAuthTools(s, nil)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = s.ListTools()
	}
}
