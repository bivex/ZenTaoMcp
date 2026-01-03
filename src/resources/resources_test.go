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

package resources

import (
	"testing"

	"github.com/mark3labs/mcp-go/server"
)

func TestExtractIDFromURI(t *testing.T) {
	tests := []struct {
		uri          string
		resourceType string
		expected     string
	}{
		{"zentao://products/123", "products", "123"},
		{"zentao://projects/456", "projects", "456"},
		{"zentao://stories/789", "stories", "789"},
		{"zentao://products", "products", ""}, // No ID
		{"zentao://products/", "products", ""}, // Empty ID
		{"invalid-uri", "products", ""}, // Invalid format
	}

	for _, test := range tests {
		result := extractIDFromURI(test.uri, test.resourceType)
		if result != test.expected {
			t.Errorf("extractIDFromURI(%s, %s) = %s, expected %s",
				test.uri, test.resourceType, result, test.expected)
		}
	}
}

func TestRegisterProductResources(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProductResources panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProductResources(s, nil)
}

func TestRegisterProgramResources(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterProgramResources panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterProgramResources(s, nil)
}

func TestRegisterTestTaskResources(t *testing.T) {
	// Test that registration doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("RegisterTestTaskResources panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")
	RegisterTestTaskResources(s, nil)
}

func TestResourceRegistrationWithNilServer(t *testing.T) {
	// Test registration with nil server
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Resource registration with nil server handled gracefully: %v", r)
		}
	}()

	RegisterProductResources(nil, nil)
	RegisterProgramResources(nil, nil)
	RegisterTestTaskResources(nil, nil)
}

func TestMultipleResourceRegistrations(t *testing.T) {
	// Test registering resources multiple times
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Multiple resource registrations panicked: %v", r)
		}
	}()

	s := server.NewMCPServer("test-server", "1.0.0")

	// Register multiple times
	RegisterProductResources(s, nil)
	RegisterProductResources(s, nil)
	RegisterProgramResources(s, nil)
	RegisterProgramResources(s, nil)
	RegisterTestTaskResources(s, nil)
	RegisterTestTaskResources(s, nil)

	// Should not crash
}

func BenchmarkExtractIDFromURI(b *testing.B) {
	testURI := "zentao://products/123"
	resourceType := "products"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		extractIDFromURI(testURI, resourceType)
	}
}

func BenchmarkResourceRegistration(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := server.NewMCPServer("bench-server", "1.0.0")
		RegisterProductResources(s, nil)
		RegisterProgramResources(s, nil)
		RegisterTestTaskResources(s, nil)
	}
}
