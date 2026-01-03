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
)

type ZenTaoClient struct {
	BaseURL   string
	Code      string
	Key       string
	Token     string
	Client    *http.Client
	lastTime  int64
	timeMutex sync.Mutex
}

func NewZenTaoClient(baseURL string) *ZenTaoClient {
	return &ZenTaoClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

func NewZenTaoClientWithApp(baseURL, code, key string) *ZenTaoClient {
	return &ZenTaoClient{
		BaseURL: baseURL,
		Code:    code,
		Key:     key,
		Client:  &http.Client{},
	}
}

func (c *ZenTaoClient) SetAppCredentials(code, key string) {
	c.Code = code
	c.Key = key
}

func (c *ZenTaoClient) generateToken(timestamp int64) string {
	tokenString := c.Code + c.Key + strconv.FormatInt(timestamp, 10)
	hash := md5.Sum([]byte(tokenString))
	return hex.EncodeToString(hash[:])
}

func (c *ZenTaoClient) getTimestamp() int64 {
	c.timeMutex.Lock()
	defer c.timeMutex.Unlock()

	now := time.Now().Unix()
	if now <= c.lastTime {
		now = c.lastTime + 1
	}
	c.lastTime = now
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

	return queryPath, params
}

func (c *ZenTaoClient) buildURL(path string, params map[string]string) string {
	// Just use the base URL, path will be converted in DoRequest
	baseURL := c.BaseURL

	if c.Code != "" && c.Key != "" {
		timestamp := c.getTimestamp()
		token := c.generateToken(timestamp)

		if params == nil {
			params = make(map[string]string)
		}
		params["code"] = c.Code
		params["time"] = strconv.FormatInt(timestamp, 10)
		params["token"] = token
	}

	if len(params) > 0 {
		queryValues := url.Values{}
		for k, v := range params {
			queryValues.Set(k, v)
		}
		return fmt.Sprintf("%s%s&%s", baseURL, path, queryValues.Encode())
	}

	return fmt.Sprintf("%s%s", baseURL, path)
}

func (c *ZenTaoClient) DoRequest(method, path string, body interface{}, headers map[string]string) ([]byte, error) {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	// Convert REST path to ZenTao query format
	queryPath, pathParams := c.convertRESTPath(method, path)
	requestURL := c.buildURL(queryPath, pathParams)

	req, err := http.NewRequest(method, requestURL, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute request: %w", err)
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func (c *ZenTaoClient) Get(path string) ([]byte, error) {
	return c.DoRequest(http.MethodGet, path, nil, nil)
}

func (c *ZenTaoClient) Post(path string, body interface{}) ([]byte, error) {
	return c.DoRequest(http.MethodPost, path, body, nil)
}

func (c *ZenTaoClient) Put(path string, body interface{}) ([]byte, error) {
	return c.DoRequest(http.MethodPut, path, body, nil)
}

func (c *ZenTaoClient) Delete(path string) ([]byte, error) {
	return c.DoRequest(http.MethodDelete, path, nil, nil)
}
