package examples

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/log"
)

// ==============================================================================
// EASY EXAMPLES - Basic logging concepts
// ==============================================================================

// SimpleLogExample demonstrates basic logging levels
// Concept: Understanding different log levels and their use cases
func SimpleLogExample() {
	fmt.Println("\n=== EASY: Basic Log Levels ===")

	// Debug: Detailed information for diagnosing problems
	log.Debug("This is a debug message - useful during development")

	// Info: General informational messages about application progress
	log.Info("Application started successfully")

	// Warn: Warning messages for potentially harmful situations
	log.Warn("Memory usage is high, consider optimizing")

	// Error: Error messages for serious problems
	log.Error("Failed to connect to database")

	// Fatal: Very severe errors that will cause the application to exit
	// Note: We're NOT calling Fatal() here as it would exit the program
	fmt.Println("(Fatal level exists but would exit the program)")
}

// LogWithFieldsExample shows how to add structured context to logs
// Concept: Structured logging with key-value pairs
func LogWithFieldsExample() {
	fmt.Println("\n=== EASY: Logging with Fields ===")

	// Single field logging
	// Fields provide context without cluttering the message
	log.Info("User logged in", "username", "alice")

	// Multiple fields
	// Fields are specified as key, value pairs
	log.Info("Request processed",
		"method", "GET",
		"path", "/api/users",
		"status", 200,
		"duration", "45ms")

	// Fields with different types
	log.Info("Transaction completed",
		"id", 12345,
		"amount", 99.99,
		"success", true,
		"timestamp", time.Now())
}

// LogFormattingExample demonstrates different output formats
// Concept: Customizing how log messages appear
func LogFormattingExample() {
	fmt.Println("\n=== EASY: Log Formatting ===")

	// Create a custom logger to avoid affecting global settings
	logger := log.New(os.Stderr)

	// Text format (default) - human-readable colored output
	logger.Info("This is text format (default)")

	// Set a custom time format
	logger.SetTimeFormat("15:04:05") // HH:MM:SS format
	logger.Info("Custom time format", "time_format", "HH:MM:SS")

	// You can also use other formats like time.Kitchen
	logger.SetTimeFormat(time.Kitchen)
	logger.Info("Kitchen time format", "example", "3:04PM")
}

// ==============================================================================
// MEDIUM EXAMPLES - Structured logging and configuration
// ==============================================================================

// LogLevelsExample demonstrates log level filtering
// Concept: Controlling which log messages are displayed
func LogLevelsExample() {
	fmt.Println("\n=== MEDIUM: Log Level Filtering ===")

	// Create a new logger instance
	logger := log.New(os.Stderr)

	// Set to DebugLevel - shows all messages
	logger.SetLevel(log.DebugLevel)
	logger.Debug("Debug level: This will be shown")
	logger.Info("Debug level: This will be shown")

	// Set to InfoLevel - hides debug messages
	logger.SetLevel(log.InfoLevel)
	logger.Debug("Info level: This will NOT be shown")
	logger.Info("Info level: This will be shown")

	// Set to WarnLevel - only shows warnings and errors
	logger.SetLevel(log.WarnLevel)
	logger.Info("Warn level: This will NOT be shown")
	logger.Warn("Warn level: This will be shown")

	// Set to ErrorLevel - only shows errors
	logger.SetLevel(log.ErrorLevel)
	logger.Warn("Error level: This will NOT be shown")
	logger.Error("Error level: This will be shown")
}

// SubLoggerExample demonstrates creating contextual loggers
// Concept: Creating child loggers with persistent context
func SubLoggerExample() {
	fmt.Println("\n=== MEDIUM: Sub-Loggers with Context ===")

	// Create a base logger
	baseLogger := log.New(os.Stderr)

	// Create a sub-logger with a prefix
	// This adds context that will appear in all messages from this logger
	userLogger := baseLogger.With("module", "user-service")

	// All messages from userLogger will include the module field
	userLogger.Info("User authentication started")
	userLogger.Info("Password verified", "username", "alice")
	userLogger.Info("Session created", "session_id", "abc123")

	// Create another sub-logger with different context
	dbLogger := baseLogger.With("module", "database", "host", "localhost")
	dbLogger.Info("Connection pool initialized", "size", 10)
	dbLogger.Info("Query executed", "table", "users", "rows", 42)

	// You can chain With() calls to add more context
	transactionLogger := dbLogger.With("transaction_id", "tx_789")
	transactionLogger.Info("Transaction started")
	transactionLogger.Info("Transaction committed")
}

// StructuredDataExample shows logging complex data structures
// Concept: Logging maps, slices, and structured data
func StructuredDataExample() {
	fmt.Println("\n=== MEDIUM: Logging Structured Data ===")

	logger := log.New(os.Stderr)

	// Logging with a map
	userInfo := map[string]interface{}{
		"id":    123,
		"name":  "Alice Johnson",
		"email": "alice@example.com",
		"role":  "admin",
	}
	logger.Info("User details", "user", userInfo)

	// Logging with nested structures
	config := map[string]interface{}{
		"server": map[string]interface{}{
			"host": "localhost",
			"port": 8080,
		},
		"database": map[string]interface{}{
			"host":     "db.example.com",
			"port":     5432,
			"database": "myapp",
		},
	}
	logger.Info("Configuration loaded", "config", config)

	// Logging arrays
	activeUsers := []string{"alice", "bob", "charlie"}
	logger.Info("Active users", "count", len(activeUsers), "users", activeUsers)
}

// LoggerOptionsExample demonstrates various logger configuration options
// Concept: Customizing logger behavior and appearance
func LoggerOptionsExample() {
	fmt.Println("\n=== MEDIUM: Logger Configuration Options ===")

	// Create logger with custom options
	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportCaller:    true,           // Include file and line number
		ReportTimestamp: true,           // Include timestamp
		TimeFormat:      time.TimeOnly,  // Custom time format
		Level:           log.DebugLevel, // Set initial log level
		Prefix:          "MyApp 🚀",      // Add a prefix to all messages
	})

	logger.Debug("Debug message with caller info")
	logger.Info("Info message with custom configuration")
	logger.Warn("Warning message with prefix")
}

// ==============================================================================
// HARD EXAMPLES - Advanced logging patterns and techniques
// ==============================================================================

// ApplicationLoggerExample demonstrates a production-ready logging setup
// Concept: Creating a comprehensive logging system for an application
func ApplicationLoggerExample() {
	fmt.Println("\n=== HARD: Production Application Logger ===")

	// Create different loggers for different parts of the application
	// Main application logger
	appLogger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Level:           log.InfoLevel,
		Prefix:          "APP",
	})

	// HTTP server logger
	httpLogger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Level:           log.DebugLevel,
		Prefix:          "HTTP",
	})

	// Database logger
	dbLogger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Level:           log.WarnLevel,
		Prefix:          "DB",
	})

	// Simulate application startup
	appLogger.Info("Starting application", "version", "1.0.0", "env", "production")

	// Simulate HTTP requests
	httpLogger.Info("Request received",
		"method", "POST",
		"path", "/api/users",
		"remote_addr", "192.168.1.100",
		"user_agent", "Mozilla/5.0")

	httpLogger.Debug("Request headers",
		"content_type", "application/json",
		"authorization", "Bearer ***")

	// Simulate database operations
	dbLogger.Info("Query executed",
		"query", "SELECT * FROM users WHERE active = true",
		"duration_ms", 23,
		"rows", 145)

	dbLogger.Warn("Slow query detected",
		"query", "SELECT * FROM logs WHERE date > '2024-01-01'",
		"duration_ms", 1534,
		"threshold_ms", 1000)

	// Simulate error scenarios
	appLogger.Error("Failed to process request",
		"error", "database connection timeout",
		"retry_count", 3,
		"will_retry", false)
}

// PerformanceLoggingExample shows how to log performance metrics
// Concept: Measuring and logging execution time and performance data
func PerformanceLoggingExample() {
	fmt.Println("\n=== HARD: Performance Logging ===")

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.TimeOnly,
		Prefix:          "PERF",
	})

	// Function to simulate work and log performance
	measureAndLog := func(operation string, work func()) {
		start := time.Now()

		// Do the work
		work()

		// Calculate duration
		duration := time.Since(start)

		// Log with performance metrics
		logger.Info("Operation completed",
			"operation", operation,
			"duration_ms", duration.Milliseconds(),
			"duration_us", duration.Microseconds())

		// Warn if operation is slow
		if duration.Milliseconds() > 100 {
			logger.Warn("Slow operation detected",
				"operation", operation,
				"duration_ms", duration.Milliseconds(),
				"threshold_ms", 100)
		}
	}

	// Simulate different operations
	measureAndLog("database_query", func() {
		time.Sleep(50 * time.Millisecond)
	})

	measureAndLog("api_call", func() {
		time.Sleep(150 * time.Millisecond)
	})

	measureAndLog("cache_lookup", func() {
		time.Sleep(5 * time.Millisecond)
	})

	// Log aggregate performance stats
	logger.Info("Performance summary",
		"total_operations", 3,
		"avg_duration_ms", 68,
		"max_duration_ms", 150,
		"min_duration_ms", 5)
}

// ErrorTrackingExample demonstrates advanced error logging
// Concept: Comprehensive error tracking with context and stack information
func ErrorTrackingExample() {
	fmt.Println("\n=== HARD: Error Tracking and Context ===")

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		ReportCaller:    true,
		TimeFormat:      time.RFC3339,
		Prefix:          "ERROR_TRACKER",
	})

	// Simulate a function that might error
	processUser := func(userID int) error {
		// Simulate an error condition
		if userID == 0 {
			return fmt.Errorf("invalid user ID: %d", userID)
		}
		return nil
	}

	// Example 1: Simple error logging
	if err := processUser(0); err != nil {
		logger.Error("Failed to process user",
			"error", err.Error(),
			"user_id", 0)
	}

	// Example 2: Error with full context
	userID := 123
	if err := processUser(userID); err != nil {
		logger.Error("User processing failed",
			"error", err.Error(),
			"user_id", userID,
			"timestamp", time.Now(),
			"severity", "critical",
			"retry_possible", true)
	} else {
		logger.Info("User processed successfully", "user_id", userID)
	}

	// Example 3: Error chain logging
	simulateErrorChain := func() error {
		// Simulate nested errors
		err := fmt.Errorf("database connection failed")
		err = fmt.Errorf("failed to fetch user: %w", err)
		err = fmt.Errorf("API request failed: %w", err)
		return err
	}

	if err := simulateErrorChain(); err != nil {
		logger.Error("Error chain detected",
			"error", err.Error(),
			"error_type", fmt.Sprintf("%T", err),
			"component", "user_service")
	}
}

// AuditLogExample demonstrates creating audit trails
// Concept: Structured logging for security and compliance
func AuditLogExample() {
	fmt.Println("\n=== HARD: Audit Logging ===")

	// Create a specialized audit logger
	auditLogger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.RFC3339,
		Level:           log.InfoLevel,
		Prefix:          "AUDIT",
	})

	// Helper function to create audit entries
	logAuditEvent := func(action, actor, resource string, details map[string]interface{}) {
		// Combine all audit information
		fields := []interface{}{
			"action", action,
			"actor", actor,
			"resource", resource,
			"timestamp", time.Now().Format(time.RFC3339),
		}

		// Add additional details
		for k, v := range details {
			fields = append(fields, k, v)
		}

		auditLogger.Info("Audit event", fields...)
	}

	// Example audit events
	logAuditEvent("LOGIN", "alice@example.com", "system", map[string]interface{}{
		"ip_address":  "192.168.1.100",
		"user_agent":  "Mozilla/5.0",
		"success":     true,
		"auth_method": "password",
	})

	logAuditEvent("FILE_ACCESS", "bob@example.com", "documents/confidential.pdf", map[string]interface{}{
		"action_type":    "download",
		"file_size":      "2.4MB",
		"classification": "confidential",
		"approved":       true,
	})

	logAuditEvent("PERMISSION_CHANGE", "admin@example.com", "user:charlie", map[string]interface{}{
		"old_role":    "user",
		"new_role":    "admin",
		"reason":      "promotion",
		"approved_by": "manager@example.com",
	})

	logAuditEvent("DATA_EXPORT", "alice@example.com", "customer_data", map[string]interface{}{
		"record_count": 1500,
		"format":       "CSV",
		"purpose":      "reporting",
		"encrypted":    true,
	})
}

// DistributedTracingExample shows how to log with tracing context
// Concept: Logging that supports distributed system tracing
func DistributedTracingExample() {
	fmt.Println("\n=== HARD: Distributed Tracing Context ===")

	logger := log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.TimeOnly,
		Prefix:          "TRACE",
	})

	// Simulate a distributed request with trace IDs
	traceID := "trace-abc-123-xyz"
	spanID := "span-001"

	// Service A: API Gateway
	serviceA := logger.With(
		"service", "api-gateway",
		"trace_id", traceID,
		"span_id", spanID,
	)
	serviceA.Info("Request received", "path", "/api/order", "method", "POST")

	// Service B: Order Service
	spanID = "span-002"
	serviceB := logger.With(
		"service", "order-service",
		"trace_id", traceID,
		"span_id", spanID,
		"parent_span", "span-001",
	)
	serviceB.Info("Processing order", "order_id", "order-789")
	serviceB.Debug("Validating order items", "item_count", 3)

	// Service C: Payment Service
	spanID = "span-003"
	serviceC := logger.With(
		"service", "payment-service",
		"trace_id", traceID,
		"span_id", spanID,
		"parent_span", "span-002",
	)
	serviceC.Info("Processing payment", "amount", 99.99, "currency", "USD")
	serviceC.Info("Payment authorized", "transaction_id", "txn-456")

	// Service D: Inventory Service
	spanID = "span-004"
	serviceD := logger.With(
		"service", "inventory-service",
		"trace_id", traceID,
		"span_id", spanID,
		"parent_span", "span-002",
	)
	serviceD.Info("Updating inventory", "sku", "PROD-123", "quantity", -1)
	serviceD.Info("Inventory updated successfully")

	// Back to Service B
	serviceB.Info("Order completed successfully", "order_id", "order-789")

	// Back to Service A
	serviceA.Info("Request completed",
		"status", 200,
		"duration_ms", 234,
		"trace_id", traceID)
}

// RunAllLogExamples executes all log examples
func RunAllLogExamples() {
	fmt.Println("\n" + strings.Repeat("=", 70))
	fmt.Println("LOG EXAMPLES - Structured Logging")
	fmt.Println(strings.Repeat("=", 70))

	// Easy examples
	SimpleLogExample()
	LogWithFieldsExample()
	LogFormattingExample()

	// Medium examples
	LogLevelsExample()
	SubLoggerExample()
	StructuredDataExample()
	LoggerOptionsExample()

	// Hard examples
	ApplicationLoggerExample()
	PerformanceLoggingExample()
	ErrorTrackingExample()
	AuditLogExample()
	DistributedTracingExample()

	fmt.Println("\n" + strings.Repeat("=", 70))
}
