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

func (c *ZenTaoClient) GetToken(account, password string) (string, error) {
	type TokenRequest struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	reqBody := TokenRequest{
		Account:  account,
		Password: password,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.Client.Post(
		fmt.Sprintf("%s/tokens", c.BaseURL),
		"application/json",
		bytes.NewBuffer(jsonBody),
	)
	if err != nil {
		return "", fmt.Errorf("failed to get token: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if token, ok := result["token"].(string); ok {
		c.Token = token
		return token, nil
	}

	return "", fmt.Errorf("token not found in response")
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

func (c *ZenTaoClient) buildURL(path string, params map[string]string) string {
	baseURL := fmt.Sprintf("%s%s", c.BaseURL, path)

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
		return fmt.Sprintf("%s?%s", baseURL, queryValues.Encode())
	}

	return baseURL
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

	requestURL := c.buildURL(path, nil)

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
