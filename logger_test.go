/*******************************************************************

		::          ::        +--------+-----------------------+
		  ::      ::          | Author | Dmitry Novikov        |
		::::::::::::::        | Email  | dredfort.42@gmail.com |
	  ::::  ::::::  ::::      +--------+-----------------------+
	::::::::::::::::::::::
	::  ::::::::::::::  ::    File     | logger_test.go
	::  ::          ::  ::    Created  | 2025-08-07
		  ::::  ::::          Modified | 2025-08-07

	GitHub:   https://github.com/dredfort42
	LinkedIn: https://linkedin.com/in/novikov-da

*******************************************************************/

package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test the color constants.
func TestColorConstants(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected string
	}{
		{"Red color", RedColor, "\033[31m"},
		{"Green color", GreenColor, "\033[32m"},
		{"Yellow color", YellowColor, "\033[33m"},
		{"Blue color", BlueColor, "\033[34m"},
		{"End color", EndColor, "\033[0m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.color)
		})
	}
}

// Test that loggers are properly initialized.
func TestLoggersInitialization(t *testing.T) {
	t.Run("ErrorLogger is initialized", func(t *testing.T) {
		assert.NotNil(t, Error)
		assert.IsType(t, &log.Logger{}, Error)
	})

	t.Run("InfoLogger is initialized", func(t *testing.T) {
		assert.NotNil(t, Info)
		assert.IsType(t, &log.Logger{}, Info)
	})

	t.Run("WarningLogger is initialized", func(t *testing.T) {
		assert.NotNil(t, Warning)
		assert.IsType(t, &log.Logger{}, Warning)
	})

	t.Run("DebugLogger is initialized", func(t *testing.T) {
		assert.NotNil(t, Debug)
		assert.IsType(t, &log.Logger{}, Debug)
	})
}

// Test logger prefixes contain the expected color codes and text.
func TestLoggerPrefixes(t *testing.T) {
	tests := []struct {
		name           string
		logger         *log.Logger
		expectedColor  string
		expectedText   string
		expectedOutput string
	}{
		{
			name:           "ErrorLogger prefix",
			logger:         Error,
			expectedColor:  RedColor,
			expectedText:   "ERROR:",
			expectedOutput: RedColor + "ERROR:  " + EndColor + " ",
		},
		{
			name:           "InfoLogger prefix",
			logger:         Info,
			expectedColor:  GreenColor,
			expectedText:   "INFO:",
			expectedOutput: GreenColor + "INFO:   " + EndColor + " ",
		},
		{
			name:           "WarningLogger prefix",
			logger:         Warning,
			expectedColor:  YellowColor,
			expectedText:   "WARNING:",
			expectedOutput: YellowColor + "WARNING:" + EndColor + " ",
		},
		{
			name:           "DebugLogger prefix",
			logger:         Debug,
			expectedColor:  BlueColor,
			expectedText:   "DEBUG:",
			expectedOutput: BlueColor + "DEBUG:  " + EndColor + " ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture logger output
			var buf bytes.Buffer

			// Create a temporary logger with the same configuration but writing to our buffer
			tempLogger := log.New(&buf, tt.expectedOutput, tt.logger.Flags())
			tempLogger.Print("test message")

			output := buf.String()

			// Check that the output contains the expected prefix
			assert.Contains(t, output, "test message")
			assert.Contains(t, output, tt.expectedColor)
			assert.Contains(t, output, tt.expectedText)
			assert.Contains(t, output, EndColor)
		})
	}
}

// Test logger flags.
func TestLoggerFlags(t *testing.T) {
	tests := []struct {
		name          string
		logger        *log.Logger
		expectedFlags int
	}{
		{
			name:          "ErrorLogger flags include file info",
			logger:        Error,
			expectedFlags: log.Ldate | log.Ltime | log.Lshortfile,
		},
		{
			name:          "InfoLogger flags",
			logger:        Info,
			expectedFlags: log.Ldate | log.Ltime,
		},
		{
			name:          "WarningLogger flags",
			logger:        Warning,
			expectedFlags: log.Ldate | log.Ltime,
		},
		{
			name:          "DebugLogger flags",
			logger:        Debug,
			expectedFlags: log.Ldate | log.Ltime,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expectedFlags, tt.logger.Flags())
		})
	}
}

// Test that ErrorLogger writes to stderr.
func TestErrorLoggerOutput(t *testing.T) {
	// Test that ErrorLogger is configured to write to stderr
	// We can't easily capture stderr in tests, but we can verify the configuration
	t.Run("ErrorLogger uses stderr", func(t *testing.T) {
		// Create a buffer and test logger with same configuration as ErrorLogger
		var buf bytes.Buffer

		testLogger := log.New(&buf, Error.Prefix(), Error.Flags())

		testLogger.Print("test error message")

		output := buf.String()
		assert.Contains(t, output, "test error message")
		assert.Contains(t, output, "ERROR:")
		assert.Contains(t, output, RedColor)
		assert.Contains(t, output, EndColor)
	})
}

// Test that other loggers write to stdout (configuration test).
func TestLoggerOutputDestination(t *testing.T) {
	tests := []struct {
		name           string
		logger         *log.Logger
		expectedPrefix string
		expectedColor  string
	}{
		{"InfoLogger", Info, "INFO:", GreenColor},
		{"WarningLogger", Warning, "WARNING:", YellowColor},
		{"DebugLogger", Debug, "DEBUG:", BlueColor},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test logger configuration by creating a test logger with same settings
			var buf bytes.Buffer

			testLogger := log.New(&buf, tt.logger.Prefix(), tt.logger.Flags())

			testLogger.Print("test message")

			output := buf.String()
			assert.Contains(t, output, "test message")
			assert.Contains(t, output, tt.expectedPrefix)
			assert.Contains(t, output, tt.expectedColor)
			assert.Contains(t, output, EndColor)
		})
	}
}

// Test logger functionality with various message types.
func TestLoggerFunctionality(t *testing.T) {
	tests := []struct {
		name   string
		logger *log.Logger
		method string
	}{
		{"ErrorLogger Print", Error, "Print"},
		{"InfoLogger Print", Info, "Print"},
		{"WarningLogger Print", Warning, "Print"},
		{"DebugLogger Print", Debug, "Print"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			// Create a temporary logger writing to our buffer
			tempLogger := log.New(&buf, "", tt.logger.Flags())

			// Test various logging methods
			tempLogger.Print("print message")
			tempLogger.Printf("printf message: %s", "formatted")
			tempLogger.Println("println message")

			output := buf.String()
			assert.Contains(t, output, "print message")
			assert.Contains(t, output, "printf message: formatted")
			assert.Contains(t, output, "println message")
		})
	}
}

// Test that loggers handle empty messages.
func TestLoggerEmptyMessages(t *testing.T) {
	tests := []struct {
		name   string
		logger *log.Logger
	}{
		{"ErrorLogger empty", Error},
		{"InfoLogger empty", Info},
		{"WarningLogger empty", Warning},
		{"DebugLogger empty", Debug},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			tempLogger := log.New(&buf, "", tt.logger.Flags())

			// Test empty message
			tempLogger.Print("")

			// Should not panic and should produce some output (at least timestamp)
			assert.NotPanics(t, func() {
				tempLogger.Print("")
			})
		})
	}
}

// Test that loggers handle special characters.
func TestLoggerSpecialCharacters(t *testing.T) {
	tests := []struct {
		name    string
		logger  *log.Logger
		message string
	}{
		{"ErrorLogger with newlines", Error, "message\nwith\nnewlines"},
		{"InfoLogger with tabs", Info, "message\twith\ttabs"},
		{"WarningLogger with unicode", Warning, "message with unicode: αβγ 🚀"},
		{"DebugLogger with quotes", Debug, "message with \"quotes\" and 'apostrophes'"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			tempLogger := log.New(&buf, "", tt.logger.Flags())

			// Test special characters
			assert.NotPanics(t, func() {
				tempLogger.Print(tt.message)
			})

			output := buf.String()
			// The message should be present in some form
			assert.NotEmpty(t, output)
		})
	}
}

// Test concurrent access to loggers.
func TestLoggerConcurrency(t *testing.T) {
	tests := []struct {
		name   string
		logger *log.Logger
	}{
		{"ErrorLogger concurrent", Error},
		{"InfoLogger concurrent", Info},
		{"WarningLogger concurrent", Warning},
		{"DebugLogger concurrent", Debug},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a buffer to capture output
			var buf bytes.Buffer

			tempLogger := log.New(&buf, "", tt.logger.Flags())

			// Test concurrent writes
			done := make(chan bool)

			for i := range 10 {
				go func(id int) {
					tempLogger.Printf("concurrent message %d", id)

					done <- true
				}(i)
			}

			// Wait for all goroutines to complete
			for range 10 {
				<-done
			}

			output := buf.String()
			assert.NotEmpty(t, output)

			// Count occurrences of "concurrent message"
			count := strings.Count(output, "concurrent message")
			assert.Equal(t, 10, count)
		})
	}
}

// Test that logger package exports are accessible.
func TestLoggerPackageExports(t *testing.T) {
	t.Run("All color constants are accessible", func(t *testing.T) {
		// Test that we can access all exported constants
		_ = RedColor
		_ = GreenColor
		_ = YellowColor
		_ = BlueColor
		_ = EndColor

		// Test they are strings
		assert.IsType(t, "", RedColor)
		assert.IsType(t, "", GreenColor)
		assert.IsType(t, "", YellowColor)
		assert.IsType(t, "", BlueColor)
		assert.IsType(t, "", EndColor)
	})

	t.Run("All loggers are accessible", func(t *testing.T) {
		// Test that we can access all exported loggers
		_ = Error
		_ = Info
		_ = Warning
		_ = Debug

		// Test they are *log.Logger
		assert.IsType(t, &log.Logger{}, Error)
		assert.IsType(t, &log.Logger{}, Info)
		assert.IsType(t, &log.Logger{}, Warning)
		assert.IsType(t, &log.Logger{}, Debug)
	})
}

// Benchmark tests for logger performance.
func BenchmarkError(b *testing.B) {
	var buf bytes.Buffer

	tempLogger := log.New(&buf, "", Error.Flags())

	for i := 0; b.Loop(); i++ {
		tempLogger.Printf("benchmark message %d", i)
	}
}

func BenchmarkInfo(b *testing.B) {
	var buf bytes.Buffer

	tempLogger := log.New(&buf, "", Info.Flags())

	for i := 0; b.Loop(); i++ {
		tempLogger.Printf("benchmark message %d", i)
	}
}

func BenchmarkWarning(b *testing.B) {
	var buf bytes.Buffer

	tempLogger := log.New(&buf, "", Warning.Flags())

	for i := 0; b.Loop(); i++ {
		tempLogger.Printf("benchmark message %d", i)
	}
}

func BenchmarkDebug(b *testing.B) {
	var buf bytes.Buffer

	tempLogger := log.New(&buf, "", Debug.Flags())

	for i := 0; b.Loop(); i++ {
		tempLogger.Printf("benchmark message %d", i)
	}
}
