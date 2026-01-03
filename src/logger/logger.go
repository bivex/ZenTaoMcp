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

package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

type Logger struct {
	level      LogLevel
	enableJSON bool
}

var defaultLogger *Logger

func init() {
	// Initialize default logger
	defaultLogger = &Logger{
		level:      INFO,
		enableJSON: false,
	}

	// Check environment variables
	if level := os.Getenv("ZENTAO_LOG_LEVEL"); level != "" {
		switch strings.ToUpper(level) {
		case "DEBUG":
			defaultLogger.level = DEBUG
		case "INFO":
			defaultLogger.level = INFO
		case "WARN":
			defaultLogger.level = WARN
		case "ERROR":
			defaultLogger.level = ERROR
		}
	}

	if os.Getenv("ZENTAO_LOG_JSON") == "true" {
		defaultLogger.enableJSON = true
	}
}

type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Component string                 `json:"component"`
	Function  string                 `json:"function"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Error     string                 `json:"error,omitempty"`
}

func (l *Logger) log(level LogLevel, component string, message string, fields map[string]interface{}, err error) {
	if level < l.level {
		return
	}

	// Get caller information
	pc, _, _, ok := runtime.Caller(2)
	funcName := "unknown"
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			funcName = fn.Name()
			// Extract just the function name
			if lastSlash := strings.LastIndex(funcName, "/"); lastSlash >= 0 {
				funcName = funcName[lastSlash+1:]
			}
			if lastDot := strings.LastIndex(funcName, "."); lastDot >= 0 {
				funcName = funcName[lastDot+1:]
			}
		}
	}

	entry := LogEntry{
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Component: component,
		Function:  funcName,
		Fields:    fields,
	}

	switch level {
	case DEBUG:
		entry.Level = "DEBUG"
	case INFO:
		entry.Level = "INFO"
	case WARN:
		entry.Level = "WARN"
	case ERROR:
		entry.Level = "ERROR"
	}

	if err != nil {
		entry.Error = err.Error()
	}

	if l.enableJSON {
		entry.Message = message
		if jsonData, err := json.Marshal(entry); err == nil {
			fmt.Fprintln(os.Stderr, string(jsonData))
		} else {
			fmt.Fprintf(os.Stderr, "[%s] %s: %s %s\n", entry.Level, component, message, entry.Error)
		}
	} else {
		logMsg := fmt.Sprintf("[%s] %s/%s: %s", entry.Level, component, funcName, message)

		if fields != nil && len(fields) > 0 {
			fieldStrs := make([]string, 0, len(fields))
			for k, v := range fields {
				fieldStrs = append(fieldStrs, fmt.Sprintf("%s=%v", k, v))
			}
			logMsg += " [" + strings.Join(fieldStrs, ", ") + "]"
		}

		if err != nil {
			logMsg += " error=" + err.Error()
		}

		fmt.Fprintln(os.Stderr, logMsg)
	}
}

func (l *Logger) Debug(component, message string, fields map[string]interface{}) {
	l.log(DEBUG, component, message, fields, nil)
}

func (l *Logger) Info(component, message string, fields map[string]interface{}) {
	l.log(INFO, component, message, fields, nil)
}

func (l *Logger) Warn(component, message string, fields map[string]interface{}) {
	l.log(WARN, component, message, fields, nil)
}

func (l *Logger) Error(component, message string, err error, fields map[string]interface{}) {
	l.log(ERROR, component, message, fields, err)
}

// Global logger functions
func Debug(component, message string, fields map[string]interface{}) {
	defaultLogger.Debug(component, message, fields)
}

func Info(component, message string, fields map[string]interface{}) {
	defaultLogger.Info(component, message, fields)
}

func Warn(component, message string, fields map[string]interface{}) {
	defaultLogger.Warn(component, message, fields)
}

func Error(component, message string, err error, fields map[string]interface{}) {
	defaultLogger.Error(component, message, err, fields)
}

// Convenience functions for common patterns
func LogRequest(component, method, url string, headers map[string]string, body interface{}) {
	fields := map[string]interface{}{
		"method":  method,
		"url":     url,
		"headers": headers,
	}
	if body != nil {
		fields["body"] = fmt.Sprintf("%v", body)
	}
	Debug(component, "Making HTTP request", fields)
}

func LogResponse(component string, statusCode int, responseBody []byte, duration time.Duration) {
	fields := map[string]interface{}{
		"status_code": statusCode,
		"duration_ms": duration.Milliseconds(),
		"body_size":   len(responseBody),
	}

	// Try to parse as JSON for better logging
	var jsonData interface{}
	if json.Unmarshal(responseBody, &jsonData) == nil {
		fields["response"] = jsonData
	} else {
		// If not JSON, truncate the response body for logging
		bodyStr := string(responseBody)
		if len(bodyStr) > 500 {
			bodyStr = bodyStr[:500] + "..."
		}
		fields["response"] = bodyStr
	}

	Debug(component, "Received HTTP response", fields)
}

func LogAPIRequest(module, function string, params map[string]string, auth map[string]interface{}) {
	fields := map[string]interface{}{
		"module":   module,
		"function": function,
		"params":   params,
		"auth":     auth,
	}
	Debug("api", "ZenTao API request", fields)
}

func LogAPIResponse(module, function string, response []byte, err error) {
	fields := map[string]interface{}{
		"module":   module,
		"function": function,
		"response_size": len(response),
	}

	if err != nil {
		Error("api", "ZenTao API error", err, fields)
		return
	}

	// Try to parse response for structured logging
	var jsonData interface{}
	if json.Unmarshal(response, &jsonData) == nil {
		fields["response"] = jsonData
		Debug("api", "ZenTao API response", fields)
	} else {
		Debug("api", "ZenTao API response (non-JSON)", fields)
	}
}

func LogMCPToolCall(toolName string, args map[string]interface{}) {
	fields := map[string]interface{}{
		"tool":      toolName,
		"arguments": args,
	}
	Info("mcp", "Tool call received", fields)
}

func LogMCPResourceRead(uri string, params map[string]interface{}) {
	fields := map[string]interface{}{
		"uri":    uri,
		"params": params,
	}
	Info("mcp", "Resource read request", fields)
}
