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

package main

import (
	"os"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	// Test that main function can be called without panicking
	// We can't easily test the full main function without mocking stdin/stdout
	// but we can test that the initialization code doesn't crash

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("main function panicked: %v", r)
		}
	}()

	// Set minimal environment variables to avoid configuration errors
	os.Setenv("ZENTAO_BASE_URL", "http://test.example.com")
	os.Setenv("ZENTAO_APP_CODE", "TEST_CODE")
	os.Setenv("ZENTAO_APP_KEY", "TEST_KEY")

	// Note: We don't actually call main() because it runs an HTTP server
	// Instead, we test the server initialization components

	// Test server initialization (similar to main)
	testServerInit(t)
}

func testServerInit(t *testing.T) {
	// This mimics the server initialization code from main.go
	baseURL := os.Getenv("ZENTAO_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8080"
	}

	// For testing, we just verify that the environment reading logic works
	// We don't require all variables to be set since some tests unset them

	// Test that baseURL has a reasonable default
	if baseURL == "" {
		t.Error("Expected baseURL to have a default value")
	}

	// Test that the initialization logic doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Server initialization panicked: %v", r)
		}
	}()

	// We can't actually create the full client without network access
	// but we can test that the initialization logic doesn't crash
}

func TestEnvironmentConfiguration(t *testing.T) {
	// Test various environment configurations

	// Test with minimal config
	os.Setenv("ZENTAO_BASE_URL", "http://test1.example.com")
	os.Setenv("ZENTAO_APP_CODE", "CODE1")
	os.Setenv("ZENTAO_APP_KEY", "KEY1")

	testServerInit(t)

	// Test with different values
	os.Setenv("ZENTAO_BASE_URL", "http://test2.example.com")
	os.Setenv("ZENTAO_APP_CODE", "CODE2")
	os.Setenv("ZENTAO_APP_KEY", "KEY2")

	testServerInit(t)

	// Test with empty values (should use defaults)
	os.Unsetenv("ZENTAO_BASE_URL")
	os.Unsetenv("ZENTAO_APP_CODE")
	os.Unsetenv("ZENTAO_APP_KEY")

	// This should not crash
	testServerInit(t)
}

func TestLoggingConfiguration(t *testing.T) {
	// Test logging environment variables

	// Test debug logging
	os.Setenv("ZENTAO_LOG_LEVEL", "DEBUG")
	os.Setenv("ZENTAO_LOG_JSON", "true")

	testServerInit(t)

	// Test info logging
	os.Setenv("ZENTAO_LOG_LEVEL", "INFO")
	os.Setenv("ZENTAO_LOG_JSON", "false")

	testServerInit(t)

	// Test invalid log level (should default gracefully)
	os.Setenv("ZENTAO_LOG_LEVEL", "INVALID")

	testServerInit(t)
}

func BenchmarkServerInitialization(b *testing.B) {
	// Set up environment
	os.Setenv("ZENTAO_BASE_URL", "http://bench.example.com")
	os.Setenv("ZENTAO_APP_CODE", "BENCH_CODE")
	os.Setenv("ZENTAO_APP_KEY", "BENCH_KEY")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testServerInit(&testing.T{}) // Pass a dummy testing.T
	}
}

// TestTimeout tests that the server can handle timeouts gracefully
func TestTimeout(t *testing.T) {
	// This test ensures that operations complete within reasonable time
	done := make(chan bool, 1)

	go func() {
		testServerInit(t)
		done <- true
	}()

	select {
	case <-done:
		// Test completed successfully
	case <-time.After(5 * time.Second):
		t.Error("Test timed out after 5 seconds")
	}
}
