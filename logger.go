/*******************************************************************

		::          ::        +--------+-----------------------+
		  ::      ::          | Author | Dmitry Novikov        |
		::::::::::::::        | Email  | dredfort.42@gmail.com |
	  ::::  ::::::  ::::      +--------+-----------------------+
	::::::::::::::::::::::
	::  ::::::::::::::  ::    File     | logger.go
	::  ::          ::  ::    Created  | 2025-08-07
		  ::::  ::::          Modified | 2025-08-07

	GitHub:   https://github.com/dredfort42
	LinkedIn: https://linkedin.com/in/novikov-da

*******************************************************************/

// Package logger provides a simple, colorized logging utility for Go applications.
//
// This package offers pre-configured loggers for different log levels (Error, Info, Warning, Debug)
// with colored output and smart routing (errors to stderr, others to stdout).
//
// Basic usage:
//
//	import "github.com/dredfort42/go_logger"
//
//	logger.InfoLogger.Println("Application started")
//	logger.WarningLogger.Println("This is a warning")
//	logger.ErrorLogger.Println("An error occurred")
//	logger.DebugLogger.Println("Debug information")
//
// All loggers are thread-safe and can be used concurrently from multiple goroutines.
package logger

import (
	"log"
	"os"
)

// Color codes for terminal output.
const (
	RedColor    = "\033[31m"
	GreenColor  = "\033[32m"
	YellowColor = "\033[33m"
	BlueColor   = "\033[34m"
	EndColor    = "\033[0m"
)

var (
	// Error logs errors.
	Error = log.New(
		os.Stderr,
		RedColor+"ERROR:  "+EndColor+" ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	// Info logs general information.
	Info = log.New(
		os.Stdout,
		GreenColor+"INFO:   "+EndColor+" ",
		log.Ldate|log.Ltime,
	)

	// Warning logs warnings.
	Warning = log.New(
		os.Stdout,
		YellowColor+"WARNING:"+EndColor+" ",
		log.Ldate|log.Ltime,
	)

	// Debug logs debug information.
	Debug = log.New(
		os.Stdout,
		BlueColor+"DEBUG:  "+EndColor+" ",
		log.Ldate|log.Ltime,
	)
)
