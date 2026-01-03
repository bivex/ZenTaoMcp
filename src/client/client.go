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
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/zentao/mcp-server/logger"
)

const (
	// tokenCacheDuration is how long to cache tokens before considering them stale
	// ZenTao tokens expire after 30 seconds, so we cache for 15 seconds (50% of lifetime)
	tokenCacheDuration = 15
)

type AuthMethod int

const (
	AuthNone AuthMethod = iota
	AuthApp
	AuthSession
)

type ZenTaoClient struct {
	BaseURL       string
	Code          string
	Key           string
	Token         string
	Client        *http.Client
	lastTime      int64
	timeMutex     sync.Mutex

	// Token caching for app-based auth
	cachedToken     string
	cachedTimestamp int64
	tokenMutex      sync.Mutex

	// Session-based authentication
	authMethod    AuthMethod
	sessionName   string
	sessionID     string
	sessionMutex  sync.Mutex
}

func NewZenTaoClient(baseURL string) *ZenTaoClient {
	logger.Info("client", "Creating new ZenTao client", map[string]interface{}{
		"base_url": baseURL,
		"auth_type": "none",
	})

	return &ZenTaoClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func NewZenTaoClientWithApp(baseURL, code, key string) *ZenTaoClient {
	authType := "app-based"
	if code == "" || key == "" {
		authType = "none (missing credentials)"
	}

	logger.Info("client", "Creating new ZenTao client with app authentication", map[string]interface{}{
		"base_url": baseURL,
		"auth_type": authType,
		"has_code": code != "",
		"has_key":  key != "",
	})

	return &ZenTaoClient{
		BaseURL:    baseURL,
		Code:       code,
		Key:        key,
		Client:     &http.Client{},
		authMethod: AuthApp,
	}
}

func NewZenTaoClientWithSession(baseURL string) *ZenTaoClient {
	logger.Info("client", "Creating new ZenTao client with session authentication", map[string]interface{}{
		"base_url":   baseURL,
		"auth_type": "session-based",
	})

	return &ZenTaoClient{
		BaseURL:    baseURL,
		Client:     &http.Client{},
		authMethod: AuthSession,
	}
}

func (c *ZenTaoClient) SetAppCredentials(code, key string) {
	logger.Info("client", "Setting app credentials", map[string]interface{}{
		"has_code": code != "",
		"has_key":  key != "",
	})

	c.Code = code
	c.Key = key
	c.authMethod = AuthApp
}

func (c *ZenTaoClient) generateToken(timestamp int64) string {
	tokenString := c.Code + c.Key + strconv.FormatInt(timestamp, 10)
	hash := md5.Sum([]byte(tokenString))
	token := hex.EncodeToString(hash[:])

	logger.Debug("client", "Generated authentication token", map[string]interface{}{
		"timestamp": timestamp,
		"token_hash": token[:8] + "...", // Only log first 8 chars for security
	})

	return token
}

func (c *ZenTaoClient) getTimestamp() int64 {
	c.timeMutex.Lock()
	defer c.timeMutex.Unlock()

	now := time.Now().Unix()
	originalNow := now

	if now <= c.lastTime {
		now = c.lastTime + 1
		logger.Warn("client", "Timestamp collision detected, incrementing", map[string]interface{}{
			"original_timestamp": originalNow,
			"adjusted_timestamp": now,
			"last_timestamp": c.lastTime,
		})
	}

	c.lastTime = now

	logger.Debug("client", "Generated timestamp", map[string]interface{}{
		"timestamp": now,
		"was_adjusted": now != originalNow,
	})

	return now
}

// convertRESTPath converts REST-style paths to ZenTao query parameter format
// Examples:
//
//	/products -> ?m=product&f=browse
//	/products/123 -> ?m=product&f=view&id=123
//	/product/123 (PUT) -> ?m=product&f=edit&id=123
//	/projects/123/executions -> ?m=execution&f=browse&project=123
func (c *ZenTaoClient) convertRESTPath(method, path string) (string, map[string]string) {
	params := make(map[string]string)

	// Remove leading slash
	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}

	// Check if this is already a full ZenTao URL (starts with index.php?)
	if strings.HasPrefix(path, "index.php?") {
		// Return as-is, but ensure it starts with ?
		if !strings.HasPrefix(path, "?") {
			path = "?" + strings.TrimPrefix(path, "index.php?")
		}
		logger.Debug("client", "Detected full ZenTao URL, using as-is", map[string]interface{}{
			"path": path,
		})
		return path, params
	}

	// Parse path components
	var module, function, id, subResource, originalResource string
	var parts []string

	// Split by /
	if path != "" {
		parts = strings.Split(path, "/")
	}

	// Determine module and function based on path structure
	if len(parts) > 0 {
		// Handle plurals -> singular for module name
		resource := parts[0]
		originalResource = resource
		switch resource {
		case "products":
			module = "product"
			if len(parts) > 1 {
				id = parts[1]
				if len(parts) > 2 {
					subResource = parts[2]
					// subID would be parts[3] if needed
				}
			}
		case "product":
			module = "product"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "projects":
			module = "project"
			if len(parts) > 1 {
				id = parts[1]
				if len(parts) > 2 {
					subResource = parts[2]
					// subID would be parts[3] if needed
				}
			}
		case "project":
			module = "project"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "programs":
			module = "program"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "program":
			module = "program"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "executions":
			module = "execution"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "execution":
			module = "execution"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "stories":
			module = "story"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "story":
			module = "story"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "tasks":
			module = "task"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "task":
			module = "task"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "bugs":
			module = "bug"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "bug":
			module = "bug"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "testcases":
			module = "testcase"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "testcase":
			module = "testcase"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "productplans", "plans":
			module = "productplan"
			if len(parts) > 1 {
				id = parts[1]
				if len(parts) > 2 {
					subResource = parts[2]
				}
			}
		case "productplan", "plan":
			module = "productplan"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "builds":
			module = "build"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "build":
			module = "build"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "users":
			module = "user"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "user":
			module = "user"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "feedbacks":
			module = "feedback"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "feedback":
			module = "feedback"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "tickets":
			module = "ticket"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "ticket":
			module = "ticket"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "testtasks":
			module = "testtask"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "testtask":
			module = "testtask"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "releases":
			module = "release"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "release":
			module = "release"
			if len(parts) > 1 {
				id = parts[1]
			}
		case "tokens":
			module = "tokens"
			function = "create"
		default:
			module = resource
		}
	}

	// Determine function based on method and path structure
	if function == "" {
		if subResource != "" {
			// Handle sub-resources like /projects/123/executions
			switch subResource {
			case "executions":
				module = "execution"
				function = "browse"
				params["project"] = id
			case "stories":
				module = "story"
				function = "browse"
				if originalResource == "projects" || originalResource == "project" {
					params["project"] = id
				} else {
					params["product"] = id
				}
			case "tasks":
				module = "task"
				function = "browse"
				params["execution"] = id
			case "builds":
				module = "build"
				function = "browse"
				params["project"] = id
			case "bugs":
				module = "bug"
				function = "browse"
				params["product"] = id
			case "testcases":
				module = "testcase"
				function = "browse"
				params["product"] = id
			case "plans":
				module = "productplan"
				function = "browse"
				params["product"] = id
			case "releases":
				module = "release"
				function = "browse"
				if parts[0] == "projects" || parts[0] == "project" {
					params["project"] = id
				} else {
					params["product"] = id
				}
			case "testtasks":
				module = "testtask"
				function = "browse"
				params["project"] = id
			case "linkstories":
				function = "linkstories"
				params["id"] = id
			case "unlinkstories":
				function = "unlinkstory"
				params["id"] = id
			case "linkbugs":
				function = "linkbug"
				params["id"] = id
			case "unlinkbugs":
				function = "unlinkbug"
				params["id"] = id
			case "assign":
				function = "assign"
				params["id"] = id
			case "close":
				function = "close"
				params["id"] = id
			case "change":
				function = "change"
				params["id"] = id
			default:
				function = subResource
				params["id"] = id
			}
		} else if id != "" {
			// Single resource with ID
			switch method {
			case "GET":
				function = "view"
				params["id"] = id
			case "PUT":
				function = "edit"
				params["id"] = id
			case "DELETE":
				function = "delete"
				params["id"] = id
			case "POST":
				function = "create"
			}
		} else {
			// Collection endpoint
			switch method {
			case "GET":
				function = "browse"
			case "POST":
				function = "create"
			}
		}
	}

	// Special case for user profile
	if module == "user" && function == "browse" && id == "" {
		// Could be profile or list, default to browse
		function = "browse"
	}

	// Build query string path
	queryPath := fmt.Sprintf("?m=%s&f=%s", module, function)

	logger.Debug("client", "Converted REST path to ZenTao format", map[string]interface{}{
		"original_path": "/" + path,
		"method": method,
		"module": module,
		"function": function,
		"id": id,
		"sub_resource": subResource,
		"param_count": len(params),
		"query_path": queryPath,
	})

	return queryPath, params
}

// getCachedToken returns a cached token if valid, or generates a new one
func (c *ZenTaoClient) getCachedToken() (string, int64) {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	currentTime := time.Now().Unix()

	logger.Debug("client", "Checking token cache", map[string]interface{}{
		"current_time": currentTime,
		"cached_token_exists": c.cachedToken != "",
		"cached_timestamp": c.cachedTimestamp,
		"cached_token_length": len(c.cachedToken),
	})

	// Check if we have a valid cached token
	if c.cachedToken != "" && c.cachedTimestamp > 0 {
		timeDiff := currentTime - c.cachedTimestamp
		if timeDiff < tokenCacheDuration {
			logger.Info("client", "Using cached token", map[string]interface{}{
				"token_age_seconds": timeDiff,
				"remaining_seconds": tokenCacheDuration - timeDiff,
				"cache_hit": true,
			})
			return c.cachedToken, c.cachedTimestamp
		} else {
			logger.Warn("client", "Cached token expired, generating new one", map[string]interface{}{
				"token_age_seconds": timeDiff,
				"cache_duration_seconds": tokenCacheDuration,
				"cache_hit": false,
			})
		}
	} else {
		logger.Info("client", "No valid cached token, generating new one", map[string]interface{}{
			"cached_token_empty": c.cachedToken == "",
			"cached_timestamp_zero": c.cachedTimestamp == 0,
			"cache_hit": false,
		})
	}

	// Generate new token
	timestamp := c.getTimestamp()
	token := c.generateToken(timestamp)

	// Cache the new token
	c.cachedToken = token
	c.cachedTimestamp = timestamp

	logger.Info("client", "Generated and cached new token", map[string]interface{}{
		"timestamp": timestamp,
		"token_age_seconds": 0,
	})

	return token, timestamp
}

// forceTokenRefresh clears the token cache to force generation of a new token
func (c *ZenTaoClient) forceTokenRefresh() {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	tokenAge := time.Now().Unix() - c.cachedTimestamp
	logger.Info("client", "Forcing token refresh due to expiration", map[string]interface{}{
		"token_age_seconds": tokenAge,
		"cache_duration_seconds": tokenCacheDuration,
		"reason": "Token expired during API call",
		"zenTao_token_expiry": 30, // ZenTao tokens expire after 30 seconds
	})

	c.cachedToken = ""
	c.cachedTimestamp = 0
}

// isTokenCloseToExpiry checks if the cached token is close to expiring
func (c *ZenTaoClient) isTokenCloseToExpiry() bool {
	c.tokenMutex.Lock()
	defer c.tokenMutex.Unlock()

	if c.cachedToken == "" {
		logger.Debug("client", "No cached token found", nil)
		return true
	}

	tokenAge := time.Now().Unix() - c.cachedTimestamp
	// Consider token close to expiry if it's older than 80% of cache duration
	isCloseToExpiry := tokenAge > int64(float64(tokenCacheDuration)*0.8)

	if isCloseToExpiry {
		logger.Info("client", "Token approaching expiry, will refresh proactively", map[string]interface{}{
			"token_age_seconds": tokenAge,
			"cache_duration_seconds": tokenCacheDuration,
			"threshold_percentage": 80,
		})
	}

	return isCloseToExpiry
}

// Session-based authentication methods

// GetSessionID retrieves a session from ZenTao
func (c *ZenTaoClient) GetSessionID() error {
	c.sessionMutex.Lock()
	defer c.sessionMutex.Unlock()

	logger.Debug("client", "Getting session ID", map[string]interface{}{
		"base_url": c.BaseURL,
	})

	// Make request to get session ID
	resp, err := c.doRequestSingle("GET", "?m=api&f=getSessionID&t=json", nil, nil)
	if err != nil {
		logger.Error("client", "Failed to get session ID", err, nil)
		return fmt.Errorf("failed to get session ID: %w", err)
	}

	// Parse response
	var sessionResp map[string]interface{}
	if err := json.Unmarshal(resp, &sessionResp); err != nil {
		logger.Error("client", "Failed to parse session response", err, nil)
		return fmt.Errorf("failed to parse session response: %w", err)
	}

	// Extract session data
	data, ok := sessionResp["data"].(map[string]interface{})
	if !ok {
		logger.Error("client", "Invalid session response format", nil, map[string]interface{}{
			"response": string(resp),
		})
		return fmt.Errorf("invalid session response format")
	}

	sessionName, ok := data["sessionName"].(string)
	if !ok {
		return fmt.Errorf("sessionName not found in response")
	}

	sessionID, ok := data["sessionID"].(string)
	if !ok {
		return fmt.Errorf("sessionID not found in response")
	}

	c.sessionName = sessionName
	c.sessionID = sessionID

	logger.Info("client", "Session obtained successfully", map[string]interface{}{
		"session_name": sessionName,
		"session_id_length": len(sessionID),
	})

	return nil
}

// Login performs user authentication using session
func (c *ZenTaoClient) Login(account, password string) error {
	c.sessionMutex.Lock()
	defer c.sessionMutex.Unlock()

	if c.sessionName == "" || c.sessionID == "" {
		return fmt.Errorf("session not initialized, call GetSessionID first")
	}

	logger.Info("client", "Performing session login", map[string]interface{}{
		"account": account,
		"has_session": c.sessionName != "",
	})

	// Prepare login data
	loginData := map[string]interface{}{
		"account":  account,
		"password": password,
	}

	// Make login request with session
	resp, err := c.doRequestSingle("POST", "?m=user&f=login", loginData, nil)
	if err != nil {
		logger.Error("client", "Login request failed", err, map[string]interface{}{
			"account": account,
		})
		return fmt.Errorf("login request failed: %w", err)
	}

	// Parse login response
	var loginResp map[string]interface{}
	if err := json.Unmarshal(resp, &loginResp); err != nil {
		logger.Error("client", "Failed to parse login response", err, nil)
		return fmt.Errorf("failed to parse login response: %w", err)
	}

	// Check login result
	status, ok := loginResp["status"].(string)
	if !ok || status != "success" {
		logger.Error("client", "Login failed", nil, map[string]interface{}{
			"status": status,
			"response": string(resp),
		})
		return fmt.Errorf("login failed: %v", loginResp)
	}

	logger.Info("client", "Login successful", map[string]interface{}{
		"account": account,
	})

	return nil
}

// SetSessionCredentials manually sets session data (for testing or pre-configured sessions)
func (c *ZenTaoClient) SetSessionCredentials(sessionName, sessionID string) {
	c.sessionMutex.Lock()
	defer c.sessionMutex.Unlock()

	c.sessionName = sessionName
	c.sessionID = sessionID
	c.authMethod = AuthSession

	logger.Debug("client", "Session credentials set manually", map[string]interface{}{
		"session_name": sessionName,
		"session_id_length": len(sessionID),
	})
}

// IsAuthenticated checks if the client has valid authentication
func (c *ZenTaoClient) IsAuthenticated() bool {
	switch c.authMethod {
	case AuthApp:
		c.tokenMutex.Lock()
		hasToken := c.cachedToken != "" && c.cachedTimestamp > 0
		c.tokenMutex.Unlock()
		return hasToken
	case AuthSession:
		c.sessionMutex.Lock()
		hasSession := c.sessionName != "" && c.sessionID != ""
		c.sessionMutex.Unlock()
		return hasSession
	default:
		return false
	}
}

// GetAuthMethod returns the current authentication method
func (c *ZenTaoClient) GetAuthMethod() AuthMethod {
	return c.authMethod
}

func (c *ZenTaoClient) buildURL(path string, params map[string]string) string {
	baseURL := c.BaseURL
	authAdded := false

	// Add authentication based on method
	switch c.authMethod {
	case AuthApp:
		// App-based authentication
		if c.Code != "" && c.Key != "" {
			token, timestamp := c.getCachedToken()

			logger.Debug("client", "Adding app authentication to request", map[string]interface{}{
				"path": path,
				"code": c.Code,
				"time": timestamp,
				"token_preview": func() string {
					if len(token) > 8 {
						return token[:8] + "..."
					}
					return token
				}(),
			})

			if params == nil {
				params = make(map[string]string)
			}
			params["code"] = c.Code
			params["time"] = strconv.FormatInt(timestamp, 10)
			params["token"] = token
			authAdded = true
		}
	case AuthSession:
		// Session-based authentication
		c.sessionMutex.Lock()
		if c.sessionName != "" && c.sessionID != "" {
			if params == nil {
				params = make(map[string]string)
			}
			params[c.sessionName] = c.sessionID
			authAdded = true
		}
		c.sessionMutex.Unlock()
	}

	// Ensure t=json parameter is always present for JSON API responses
	if params == nil {
		params = make(map[string]string)
	}
	if _, exists := params["t"]; !exists {
		params["t"] = "json"
	}

	finalURL := baseURL + path
	if len(params) > 0 {
		queryValues := url.Values{}
		for k, v := range params {
			queryValues.Set(k, v)
		}
		finalURL = fmt.Sprintf("%s%s&%s", baseURL, path, queryValues.Encode())
	}

	logger.Info("client", "Built request URL with authentication", map[string]interface{}{
		"base_url": baseURL,
		"path": path,
		"auth_method": c.authMethod,
		"auth_added": authAdded,
		"param_count": len(params),
		"final_url_length": len(finalURL),
		"auth_params": func() map[string]string {
			safeParams := make(map[string]string)
			for k, v := range params {
				if k == "token" && len(v) > 8 {
					safeParams[k] = v[:8] + "..."
				} else {
					safeParams[k] = v
				}
			}
			return safeParams
		}(),
	})

	return finalURL
}

func (c *ZenTaoClient) DoRequest(method, path string, body interface{}, headers map[string]string) ([]byte, error) {
	return c.doRequestWithRetry(method, path, body, headers, 2)
}

func (c *ZenTaoClient) doRequestWithRetry(method, path string, body interface{}, headers map[string]string, maxRetries int) ([]byte, error) {
	// Log token cache state at start of request
	c.tokenMutex.Lock()
	currentTime := time.Now().Unix()
	tokenAge := currentTime - c.cachedTimestamp
	hasToken := c.cachedToken != ""
	cachedTimestamp := c.cachedTimestamp
	cachedTokenPreview := ""
	if len(c.cachedToken) > 8 {
		cachedTokenPreview = c.cachedToken[:8] + "..."
	}
	c.tokenMutex.Unlock()

	// Check if this is a write operation that needs fresh tokens
	isWriteOperation := method == "POST" || method == "PUT" || method == "DELETE"
	shouldUseFreshToken := isWriteOperation || c.isTokenCloseToExpiry()

	logger.Info("client", "Starting request with current token cache state", map[string]interface{}{
		"method": method,
		"path": path,
		"has_cached_token": hasToken,
		"cached_token_preview": cachedTokenPreview,
		"cached_timestamp": cachedTimestamp,
		"current_time": currentTime,
		"token_age_seconds": tokenAge,
		"cache_duration": tokenCacheDuration,
		"is_close_to_expiry": c.isTokenCloseToExpiry(),
		"is_write_operation": isWriteOperation,
		"should_use_fresh_token": shouldUseFreshToken,
		"raw_age_calculation": fmt.Sprintf("%d - %d = %d", currentTime, cachedTimestamp, tokenAge),
	})

	var lastErr error
	var responseBody []byte

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			logger.Info("client", "Retrying request after token refresh", map[string]interface{}{
				"attempt": attempt,
				"max_retries": maxRetries,
				"method": method,
				"path": path,
			})
			// Force token refresh on retry
			c.forceTokenRefresh()
			// Small delay before retry
			time.Sleep(100 * time.Millisecond)
		} else if shouldUseFreshToken {
			// Proactively refresh token for write operations or if close to expiry
			logger.Info("client", "Proactively refreshing token before request", map[string]interface{}{
				"method": method,
				"path": path,
				"is_write_operation": isWriteOperation,
				"is_close_to_expiry": c.isTokenCloseToExpiry(),
				"reason": func() string {
					if isWriteOperation {
						return "write operation requires fresh token"
					}
					return "token close to expiry"
				}(),
			})
			c.forceTokenRefresh()
		}

		resp, err := c.doRequestSingle(method, path, body, headers)
		if err != nil {
			lastErr = err
			continue
		}

		responseBody = resp

		// Check if response indicates token expiration
		if c.isTokenExpired(responseBody) {
			if attempt < maxRetries {
				logger.Warn("client", "Token expired, will retry with fresh token", map[string]interface{}{
					"attempt": attempt + 1,
					"max_retries": maxRetries,
					"method": method,
					"path": path,
				})
				continue
			} else {
				logger.Error("client", "Token expired after maximum retries", nil, map[string]interface{}{
					"attempts": maxRetries + 1,
					"method": method,
					"path": path,
				})
				return nil, fmt.Errorf("token expired after %d attempts", maxRetries+1)
			}
		} else {
			// Response received but no token expiration detected
			logger.Debug("client", "Response received, token still valid", map[string]interface{}{
				"attempt": attempt + 1,
				"method": method,
				"path": path,
				"response_length": len(responseBody),
				"response_preview": func() string {
					previewLen := 100
					if len(responseBody) < previewLen {
						previewLen = len(responseBody)
					}
					return string(responseBody)[:previewLen]
				}(),
			})
		}

		// Success - log the successful response and return
		logger.Debug("client", "Request completed successfully", map[string]interface{}{
			"attempt": attempt + 1,
			"method": method,
			"path": path,
			"response_length": len(responseBody),
		})
		return responseBody, nil
	}

	return nil, lastErr
}

func (c *ZenTaoClient) isTokenExpired(responseBody []byte) bool {
	// Log the response for debugging token expiration issues
	responseStr := string(responseBody)
	logger.Debug("client", "Checking response for token expiration", map[string]interface{}{
		"response_length": len(responseStr),
		"response_preview": func() string {
			previewLen := 200
			if len(responseStr) < previewLen {
				previewLen = len(responseStr)
			}
			return responseStr[:previewLen]
		}(), // First 200 chars
	})

	// Check for common token expiration indicators
	var response map[string]interface{}
	if err := json.Unmarshal(responseBody, &response); err != nil {
		logger.Debug("client", "Failed to parse response as JSON, not considering token error", map[string]interface{}{
			"parse_error": err.Error(),
		})
		return false // Can't parse, assume not token error
	}

	// Check for ZenTao-specific token error codes
	if errcode, ok := response["errcode"].(float64); ok && errcode == 405 {
		logger.Warn("client", "Token expired - detected via errcode 405", map[string]interface{}{
			"errcode": errcode,
			"full_response": responseStr,
		})
		return true
	}

	// Check for error messages containing token-related text (multiple field names)
	messageFields := []string{"errmsg", "message", "error"}
	for _, field := range messageFields {
		if msg, ok := response[field].(string); ok {
			msgLower := strings.ToLower(msg)
			if strings.Contains(msgLower, "token") || strings.Contains(msgLower, "expired") {
				logger.Warn("client", "Token expired - detected via message field", map[string]interface{}{
					"field": field,
					"message": msg,
					"full_response": responseStr,
				})
				return true
			}
		}
	}

	logger.Debug("client", "Response does not indicate token expiration", map[string]interface{}{
		"response_keys": getMapKeys(response),
	})

	return false
}

// Helper function to get map keys for logging
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (c *ZenTaoClient) doRequestSingle(method, path string, body interface{}, headers map[string]string) ([]byte, error) {
	startTime := time.Now()

	logger.Debug("client", "Starting HTTP request", map[string]interface{}{
		"method": method,
		"path": path,
		"has_body": body != nil,
		"header_count": len(headers),
	})

	var reqBody io.Reader
	var bodySize int

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			logger.Error("client", "Failed to marshal request body", err, map[string]interface{}{
				"method": method,
				"path": path,
			})
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
		bodySize = len(jsonData)

		logger.Debug("client", "Request body marshaled", map[string]interface{}{
			"body_size": bodySize,
			"content_type": "application/json",
		})
	}

	// Convert REST path to ZenTao query format
	logger.Debug("client", "Converting REST path to ZenTao format", map[string]interface{}{
		"original_path": path,
		"method": method,
	})

	queryPath, pathParams := c.convertRESTPath(method, path)
	requestURL := c.buildURL(queryPath, pathParams)

	logger.LogRequest("client", method, requestURL, headers, body)

	req, err := http.NewRequest(method, requestURL, reqBody)
	if err != nil {
		logger.Error("client", "Failed to create HTTP request", err, map[string]interface{}{
			"method": method,
			"url": requestURL,
		})
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	logger.Debug("client", "Executing HTTP request", map[string]interface{}{
		"method": method,
		"url_length": len(requestURL),
		"headers_set": len(headers) + 1, // +1 for Content-Type
	})

	resp, err := c.Client.Do(req)
	if err != nil {
		duration := time.Since(startTime)
		logger.Error("client", "HTTP request failed", err, map[string]interface{}{
			"method": method,
			"url": requestURL,
			"duration_ms": duration.Milliseconds(),
		})
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	duration := time.Since(startTime)

	if err != nil {
		logger.Error("client", "Failed to read response body", err, map[string]interface{}{
			"method": method,
			"url": requestURL,
			"status_code": resp.StatusCode,
			"duration_ms": duration.Milliseconds(),
		})
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	logger.LogResponse("client", resp.StatusCode, responseBody, duration)

	logger.Debug("client", "HTTP request completed", map[string]interface{}{
		"method": method,
		"path": path,
		"status_code": resp.StatusCode,
		"response_size": len(responseBody),
		"duration_ms": duration.Milliseconds(),
		"content_type": resp.Header.Get("Content-Type"),
	})

	return responseBody, nil
}

func (c *ZenTaoClient) Get(path string) ([]byte, error) {
	logger.Debug("client", "GET request", map[string]interface{}{
		"path": path,
	})
	return c.DoRequest(http.MethodGet, path, nil, nil)
}

func (c *ZenTaoClient) Post(path string, body interface{}) ([]byte, error) {
	logger.Debug("client", "POST request", map[string]interface{}{
		"path": path,
		"has_body": body != nil,
	})
	return c.DoRequest(http.MethodPost, path, body, nil)
}

func (c *ZenTaoClient) Put(path string, body interface{}) ([]byte, error) {
	logger.Debug("client", "PUT request", map[string]interface{}{
		"path": path,
		"has_body": body != nil,
	})
	return c.DoRequest(http.MethodPut, path, body, nil)
}

func (c *ZenTaoClient) Delete(path string) ([]byte, error) {
	logger.Debug("client", "DELETE request", map[string]interface{}{
		"path": path,
	})
	return c.DoRequest(http.MethodDelete, path, nil, nil)
}
