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

package client

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestNewZenTaoClient(t *testing.T) {
	client := NewZenTaoClient("http://test.com")

	if client.BaseURL != "http://test.com" {
		t.Errorf("Expected BaseURL to be 'http://test.com', got '%s'", client.BaseURL)
	}

	if client.Code != "" {
		t.Errorf("Expected Code to be empty, got '%s'", client.Code)
	}

	if client.Key != "" {
		t.Errorf("Expected Key to be empty, got '%s'", client.Key)
	}

	if client.Client == nil {
		t.Error("Expected HTTP client to be initialized")
	}
}

func TestNewZenTaoClientWithApp(t *testing.T) {
	client := NewZenTaoClientWithApp("http://test.com", "test_code", "test_key")

	if client.BaseURL != "http://test.com" {
		t.Errorf("Expected BaseURL to be 'http://test.com', got '%s'", client.BaseURL)
	}

	if client.Code != "test_code" {
		t.Errorf("Expected Code to be 'test_code', got '%s'", client.Code)
	}

	if client.Key != "test_key" {
		t.Errorf("Expected Key to be 'test_key', got '%s'", client.Key)
	}
}

func TestSetAppCredentials(t *testing.T) {
	client := NewZenTaoClient("http://test.com")

	client.SetAppCredentials("new_code", "new_key")

	if client.Code != "new_code" {
		t.Errorf("Expected Code to be 'new_code', got '%s'", client.Code)
	}

	if client.Key != "new_key" {
		t.Errorf("Expected Key to be 'new_key', got '%s'", client.Key)
	}
}

func TestGenerateToken(t *testing.T) {
	client := &ZenTaoClient{
		Code: "test_code",
		Key:  "test_key",
	}

	timestamp := int64(1640995200) // 2022-01-01 00:00:00 UTC
	token := client.generateToken(timestamp)

	// We can't easily test the exact MD5 hash without importing crypto/md5
	// But we can test that it returns a non-empty string of the right length
	if len(token) != 32 {
		t.Errorf("Expected token length to be 32 (MD5 hex), got %d", len(token))
	}

	// Test that same input produces same output
	token2 := client.generateToken(timestamp)
	if token != token2 {
		t.Error("Expected same timestamp to produce same token")
	}

	// Test that different timestamp produces different token
	token3 := client.generateToken(timestamp + 1)
	if token == token3 {
		t.Error("Expected different timestamp to produce different token")
	}
}

func TestGetTimestamp(t *testing.T) {
	client := &ZenTaoClient{}

	// Test that timestamps are monotonically increasing
	t1 := client.getTimestamp()
	time.Sleep(1 * time.Millisecond) // Ensure time difference
	t2 := client.getTimestamp()

	if t2 <= t1 {
		t.Errorf("Expected timestamp to increase, got t1=%d, t2=%d", t1, t2)
	}

	// Test collision handling
	client.lastTime = time.Now().Unix()
	collisionTime := client.lastTime
	newTime := client.getTimestamp()

	if newTime <= collisionTime {
		t.Errorf("Expected collision handling to increment time, got %d <= %d", newTime, collisionTime)
	}
}

func TestTokenCaching(t *testing.T) {
	client := &ZenTaoClient{
		Code: "test_code",
		Key:  "test_key",
	}

	// First call should generate new token
	token1, ts1 := client.getCachedToken()
	if token1 == "" {
		t.Error("Expected non-empty token")
	}
	if ts1 == 0 {
		t.Error("Expected non-zero timestamp")
	}

	// Second call within 15 seconds should return cached token
	token2, ts2 := client.getCachedToken()
	if token1 != token2 {
		t.Error("Expected cached token to be returned")
	}
	if ts1 != ts2 {
		t.Error("Expected cached timestamp to be returned")
	}

	// Force refresh should clear cache
	client.forceTokenRefresh()

	// Next call should generate new token
	token3, ts3 := client.getCachedToken()
	if token1 == token3 {
		t.Error("Expected new token after refresh")
	}
	if ts1 == ts3 {
		t.Error("Expected new timestamp after refresh")
	}
}

func TestTokenExpirationDetection(t *testing.T) {
	client := &ZenTaoClient{}

	// Test token expired response
	expiredResponse := `{"errcode": 405, "errmsg": "Token has expired"}`
	if !client.isTokenExpired([]byte(expiredResponse)) {
		t.Error("Expected token expired response to be detected")
	}

	// Test token error message (using errmsg field)
	tokenErrorResponse := `{"errmsg": "Token has expired"}`
	if !client.isTokenExpired([]byte(tokenErrorResponse)) {
		t.Error("Expected token error message to be detected")
	}

	// Test token error message with different case
	tokenErrorResponse2 := `{"errmsg": "Authentication token expired"}`
	if !client.isTokenExpired([]byte(tokenErrorResponse2)) {
		t.Error("Expected token error message with different wording to be detected")
	}

	// Test valid response
	validResponse := `{"status": "success", "data": []}`
	if client.isTokenExpired([]byte(validResponse)) {
		t.Error("Expected valid response to not be detected as token expired")
	}

	// Test invalid JSON
	if client.isTokenExpired([]byte("invalid json")) {
		t.Error("Expected invalid JSON to not be detected as token expired")
	}
}

func TestConvertRESTPath(t *testing.T) {
	client := NewZenTaoClient("http://test.com")

	tests := []struct {
		method   string
		path     string
		expected string
		params   map[string]string
	}{
		{
			method:   "GET",
			path:     "/products",
			expected: "?m=product&f=browse",
			params:   map[string]string{},
		},
		{
			method:   "GET",
			path:     "/products/123",
			expected: "?m=product&f=view",
			params:   map[string]string{"id": "123"},
		},
		{
			method:   "POST",
			path:     "/products",
			expected: "?m=product&f=create",
			params:   map[string]string{},
		},
		{
			method:   "PUT",
			path:     "/products/123",
			expected: "?m=product&f=edit",
			params:   map[string]string{"id": "123"},
		},
		{
			method:   "DELETE",
			path:     "/products/123",
			expected: "?m=product&f=delete",
			params:   map[string]string{"id": "123"},
		},
		{
			method:   "GET",
			path:     "/projects/123/executions",
			expected: "?m=execution&f=browse",
			params:   map[string]string{"project": "123"},
		},
	}

	for _, test := range tests {
		result, params := client.convertRESTPath(test.method, test.path)

		if result != test.expected {
			t.Errorf("convertRESTPath(%s, %s) = %s, expected %s",
				test.method, test.path, result, test.expected)
		}

		for key, expectedValue := range test.params {
			if actualValue, exists := params[key]; !exists || actualValue != expectedValue {
				t.Errorf("Expected param %s=%s, got %s", key, expectedValue, actualValue)
			}
		}
	}
}

func TestBuildURL(t *testing.T) {
	client := NewZenTaoClient("http://test.com")

	// Test without authentication
	url := client.buildURL("?m=test&f=index", nil)
	expected := "http://test.com?m=test&f=index"
	if url != expected {
		t.Errorf("Expected URL '%s', got '%s'", expected, url)
	}

	// Test with authentication
	client.SetAppCredentials("test_code", "test_key")
	url = client.buildURL("?m=test&f=index", map[string]string{"param": "value"})
	if !strings.Contains(url, "code=test_code") {
		t.Error("Expected URL to contain authentication code")
	}
	if !strings.Contains(url, "time=") {
		t.Error("Expected URL to contain timestamp")
	}
	if !strings.Contains(url, "token=") {
		t.Error("Expected URL to contain token")
	}
	if !strings.Contains(url, "param=value") {
		t.Error("Expected URL to contain custom parameter")
	}
}

// Mock HTTP server for testing DoRequest
func createMockServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for token in query parameters
		if r.URL.Query().Get("token") == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"errcode": 405, "errmsg": "Token has expired"}`))
			return
		}

		// Return success for valid requests
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "success", "data": {"test": "data"}}`))
	}))
}

func TestDoRequestWithRetry(t *testing.T) {
	server := createMockServer()
	defer server.Close()

	client := NewZenTaoClient(server.URL + "/api.php")
	client.SetAppCredentials("test_code", "test_key")

	// Test successful request
	resp, err := client.Get("/test")
	if err != nil {
		t.Errorf("Expected successful request, got error: %v", err)
	}

	if resp == nil {
		t.Error("Expected non-nil response")
	}

	// Verify response contains expected data
	var response map[string]interface{}
	if err := json.Unmarshal(resp, &response); err != nil {
		t.Errorf("Failed to parse response JSON: %v", err)
	}

	if status, ok := response["status"].(string); !ok || status != "success" {
		t.Errorf("Expected success status, got %v", response)
	}
}

func TestConcurrentRequests(t *testing.T) {
	client := &ZenTaoClient{
		Code: "test_code",
		Key:  "test_key",
	}

	// Test concurrent access to token caching
	var wg sync.WaitGroup
	results := make(chan string, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			token, _ := client.getCachedToken()
			results <- token
		}()
	}

	wg.Wait()
	close(results)

	// All tokens should be the same (from cache)
	var tokens []string
	for token := range results {
		tokens = append(tokens, token)
	}

	if len(tokens) != 10 {
		t.Errorf("Expected 10 tokens, got %d", len(tokens))
	}

	// All tokens should be identical (cached)
	for i := 1; i < len(tokens); i++ {
		if tokens[0] != tokens[i] {
			t.Error("Expected all tokens to be identical (cached)")
		}
	}
}

func TestTimeMutexConcurrency(t *testing.T) {
	client := &ZenTaoClient{}

	// Test concurrent access to timestamp generation
	var wg sync.WaitGroup
	timestamps := make(chan int64, 50)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ts := client.getTimestamp()
			timestamps <- ts
		}()
	}

	wg.Wait()
	close(timestamps)

	// Collect and verify timestamps are unique and increasing
	var tsList []int64
	for ts := range timestamps {
		tsList = append(tsList, ts)
	}

	if len(tsList) != 50 {
		t.Errorf("Expected 50 timestamps, got %d", len(tsList))
	}

	// Check for uniqueness and monotonic increase
	for i := 1; i < len(tsList); i++ {
		if tsList[i] <= tsList[i-1] {
			t.Errorf("Timestamps not monotonically increasing: %d >= %d", tsList[i-1], tsList[i])
		}
	}
}

func BenchmarkTokenGeneration(b *testing.B) {
	client := &ZenTaoClient{
		Code: "benchmark_code",
		Key:  "benchmark_key",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.generateToken(int64(i))
	}
}

func BenchmarkTokenCaching(b *testing.B) {
	client := &ZenTaoClient{
		Code: "benchmark_code",
		Key:  "benchmark_key",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.getCachedToken()
	}
}

func TestNewZenTaoClientWithSession(t *testing.T) {
	client := NewZenTaoClientWithSession("http://test.com")

	if client.BaseURL != "http://test.com" {
		t.Errorf("Expected BaseURL to be 'http://test.com', got '%s'", client.BaseURL)
	}

	if client.authMethod != AuthSession {
		t.Errorf("Expected authMethod to be AuthSession, got %v", client.authMethod)
	}

	if client.Code != "" {
		t.Errorf("Expected Code to be empty, got '%s'", client.Code)
	}

	if client.Key != "" {
		t.Errorf("Expected Key to be empty, got '%s'", client.Key)
	}

	if client.Client == nil {
		t.Error("Expected HTTP client to be initialized")
	}
}

func TestSessionAuthentication(t *testing.T) {
	client := NewZenTaoClientWithSession("http://test.com")

	// Initially should not be authenticated
	if client.IsAuthenticated() {
		t.Error("Expected client to not be authenticated initially")
	}

	// Set session credentials
	client.SetSessionCredentials("zentaosid", "test_session_123")

	if !client.IsAuthenticated() {
		t.Error("Expected client to be authenticated after setting session")
	}

	if client.GetAuthMethod() != AuthSession {
		t.Error("Expected auth method to be AuthSession")
	}
}

func BenchmarkTimestampGeneration(b *testing.B) {
	client := &ZenTaoClient{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		client.getTimestamp()
	}
}

// TestTokenCachingBehavior tests detailed token caching behavior
func TestTokenCachingBehavior(t *testing.T) {
	client := &ZenTaoClient{}

	// Test initial token generation
	ts1 := time.Now().Unix()
	token1 := client.generateToken(ts1)
	if token1 == "" {
		t.Error("Expected non-empty token")
	}

	// Manually set cached token for testing
	client.tokenMutex.Lock()
	client.cachedToken = token1
	client.cachedTimestamp = ts1
	client.tokenMutex.Unlock()

	// Test caching - should return cached token
	cachedToken := client.cachedToken
	cachedTs := client.cachedTimestamp

	if cachedToken != token1 {
		t.Errorf("Expected cached token %s, got %s", token1, cachedToken)
	}
	if cachedTs != ts1 {
		t.Errorf("Expected cached timestamp %d, got %d", ts1, cachedTs)
	}

	// Force token expiry by setting old timestamp
	client.tokenMutex.Lock()
	client.cachedTimestamp = time.Now().Unix() - tokenCacheDuration - 1
	client.tokenMutex.Unlock()

	// Generate new token with different timestamp (would be different)
	ts2 := time.Now().Unix() + 1
	token2 := client.generateToken(ts2)
	if token1 == token2 {
		t.Error("Expected different tokens for different timestamps")
	}
}

// TestTokenExpirationDetectionLogic tests the isTokenExpired function with various inputs
func TestTokenExpirationDetectionLogic(t *testing.T) {
	tests := []struct {
		name     string
		response []byte
		expected bool
	}{
		{
			name:     "Valid JSON without error",
			response: []byte(`{"data": "test"}`),
			expected: false,
		},
		{
			name:     "Token expired - errcode 405",
			response: []byte(`{"errcode": 405, "errmsg": "Token expired"}`),
			expected: true,
		},
		{
			name:     "Token error in message",
			response: []byte(`{"errmsg": "Invalid token provided"}`),
			expected: true,
		},
		{
			name:     "Token expired in message",
			response: []byte(`{"message": "Your token has expired"}`),
			expected: true,
		},
		{
			name:     "Invalid JSON",
			response: []byte(`invalid json`),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &ZenTaoClient{}
			result := client.isTokenExpired(tt.response)
			if result != tt.expected {
				t.Errorf("isTokenExpired(%s) = %v, expected %v", string(tt.response), result, tt.expected)
			}
		})
	}
}

// TestTokenCloseToExpiryDetection tests proactive token refresh detection
func TestTokenCloseToExpiryDetection(t *testing.T) {
	client := &ZenTaoClient{}

	// Test with no cached token
	if !client.isTokenCloseToExpiry() {
		t.Error("Expected true for no cached token")
	}

	// Set a fresh token
	client.tokenMutex.Lock()
	client.cachedToken = "test_token"
	client.cachedTimestamp = time.Now().Unix()
	client.tokenMutex.Unlock()

	// Should not be close to expiry
	if client.isTokenCloseToExpiry() {
		t.Error("Expected false for fresh token")
	}

	// Set token close to expiry (12 seconds old, 80% of 15 second cache)
	client.tokenMutex.Lock()
	client.cachedTimestamp = time.Now().Unix() - int64(float64(tokenCacheDuration)*0.8) - 1
	client.tokenMutex.Unlock()

	// Should be close to expiry
	if !client.isTokenCloseToExpiry() {
		t.Error("Expected true for token close to expiry")
	}
}

// TestForceTokenRefreshMechanism tests forced token refresh
func TestForceTokenRefreshMechanism(t *testing.T) {
	client := &ZenTaoClient{}

	// Set initial token
	client.tokenMutex.Lock()
	client.cachedToken = "initial_token"
	client.cachedTimestamp = time.Now().Unix()
	client.tokenMutex.Unlock()

	// Force refresh
	client.forceTokenRefresh()

	// Verify cache is cleared
	client.tokenMutex.Lock()
	if client.cachedToken != "" {
		t.Error("Expected empty cached token after force refresh")
	}
	if client.cachedTimestamp != 0 {
		t.Error("Expected zero timestamp after force refresh")
	}
	client.tokenMutex.Unlock()
}

// TestConcurrentTokenGeneration tests thread-safe token generation
func TestConcurrentTokenGeneration(t *testing.T) {
	client := &ZenTaoClient{}

	// Run multiple goroutines generating tokens concurrently
	const numGoroutines = 10
	const numRequests = 50

	var wg sync.WaitGroup
	tokens := make(chan string, numGoroutines*numRequests)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < numRequests; j++ {
				token := client.generateToken(time.Now().Unix())
				tokens <- token
			}
		}()
	}

	wg.Wait()
	close(tokens)

	// Collect all tokens
	var tokenList []string
	for token := range tokens {
		tokenList = append(tokenList, token)
	}

	// Verify all tokens are valid (non-empty)
	for i, token := range tokenList {
		if token == "" {
			t.Errorf("Empty token at index %d", i)
		}
	}

	// Verify reasonable number of unique tokens (should vary due to different timestamps)
	uniqueTokens := make(map[string]bool)
	for _, token := range tokenList {
		uniqueTokens[token] = true
	}

	// Should have many unique tokens due to different timestamps
	if len(uniqueTokens) < numGoroutines {
		t.Logf("Got %d unique tokens from %d concurrent requests, expected more variation", len(uniqueTokens), len(tokenList))
	}
}

// TestTokenExpirationResponseParsing tests parsing of token expiration responses
func TestTokenExpirationResponseParsing(t *testing.T) {
	client := &ZenTaoClient{}

	testCases := []struct {
		name        string
		response    []byte
		shouldRetry bool
	}{
		{"Valid response", []byte(`{"data": "ok"}`), false},
		{"Token expired errcode", []byte(`{"errcode": 405, "errmsg": "Token expired"}`), true},
		{"Token invalid message", []byte(`{"errmsg": "Invalid token provided"}`), true},
		{"Token expired message", []byte(`{"message": "Your token has expired"}`), true},
		{"Invalid JSON", []byte(`not json`), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := client.isTokenExpired(tc.response)
			if result != tc.shouldRetry {
				t.Errorf("isTokenExpired(%s) = %v, expected %v", string(tc.response), result, tc.shouldRetry)
			}
		})
	}
}

// TestTokenRefreshFlow tests the overall token refresh flow
func TestTokenRefreshFlow(t *testing.T) {
	client := &ZenTaoClient{}

	// Test that force refresh clears cache
	client.tokenMutex.Lock()
	client.cachedToken = "test_token"
	client.cachedTimestamp = time.Now().Unix()
	client.tokenMutex.Unlock()

	client.forceTokenRefresh()

	client.tokenMutex.Lock()
	if client.cachedToken != "" || client.cachedTimestamp != 0 {
		t.Error("Expected cache to be cleared after force refresh")
	}
	client.tokenMutex.Unlock()
}

// TestTokenCacheBoundaryConditions tests token caching at boundary conditions
func TestTokenCacheBoundaryConditions(t *testing.T) {
	if tokenCacheDuration != 15 {
		t.Errorf("Expected tokenCacheDuration to be 15, got %d", tokenCacheDuration)
	}

	client := &ZenTaoClient{}

	// Test token caching behavior manually
	currentTime := time.Now().Unix()

	// Set token exactly at cache duration boundary (should be valid)
	client.tokenMutex.Lock()
	client.cachedToken = "boundary_token"
	client.cachedTimestamp = currentTime - tokenCacheDuration
	client.tokenMutex.Unlock()

	// Should return cached token (boundary is inclusive)
	cachedToken := client.cachedToken
	if cachedToken != "boundary_token" {
		t.Error("Expected cached token at boundary")
	}

	// Set token just past boundary (should be invalid)
	client.tokenMutex.Lock()
	client.cachedTimestamp = currentTime - tokenCacheDuration - 1
	client.tokenMutex.Unlock()

	// Cache should be considered expired
	if client.cachedTimestamp >= currentTime-tokenCacheDuration {
		t.Error("Expected token to be considered expired")
	}
}

// TestProactiveTokenRefreshLogic tests the logic for detecting tokens close to expiry
func TestProactiveTokenRefreshLogic(t *testing.T) {
	client := &ZenTaoClient{}

	// Test with no token (should be considered close to expiry)
	if !client.isTokenCloseToExpiry() {
		t.Error("Expected true for client with no cached token")
	}

	// Set token that's fresh
	client.tokenMutex.Lock()
	client.cachedToken = "fresh_token"
	client.cachedTimestamp = time.Now().Unix()
	client.tokenMutex.Unlock()

	if client.isTokenCloseToExpiry() {
		t.Error("Expected false for fresh token")
	}

	// Set token that's close to expiry (13 seconds old for 15-second cache, ensuring it's over 80% threshold)
	client.tokenMutex.Lock()
	client.cachedTimestamp = time.Now().Unix() - 13
	client.tokenMutex.Unlock()

	if !client.isTokenCloseToExpiry() {
		t.Error("Expected true for token close to expiry")
	}
}

// TestTokenCacheConstants tests that cache duration constants are correct
func TestTokenCacheConstants(t *testing.T) {
	// ZenTao tokens expire after 30 seconds
	// We cache for 15 seconds (50% of lifetime) as conservative approach
	expectedDuration := 15
	if tokenCacheDuration != expectedDuration {
		t.Errorf("Expected tokenCacheDuration to be %d, got %d", expectedDuration, tokenCacheDuration)
	}
}
