package main

import (
	"fmt"
	"time"

	logger "github.com/dredfort42/go_logger"
)

func main() {
	fmt.Println("=== Go Logger Example ===")

	// Basic usage examples
	logger.Info.Println("Application started successfully")
	logger.Info.Printf("Current time: %s", time.Now().Format("15:04:05"))

	logger.Warning.Println("This is a warning message")
	logger.Warning.Printf("Memory usage: %.1f%%", 85.7)

	logger.Debug.Println("Debug information")
	logger.Debug.Printf("Processing %d items", 42)

	// Simulate an error
	logger.Error.Println("This is an error message")
	logger.Error.Printf("Failed to connect to database: %s", "connection timeout")

	fmt.Println()
	fmt.Println("=== Custom colored output ===")
	fmt.Printf("%sCustom red message%s\n", logger.RedColor, logger.EndColor)
	fmt.Printf("%sCustom green message%s\n", logger.GreenColor, logger.EndColor)
	fmt.Printf("%sCustom yellow message%s\n", logger.YellowColor, logger.EndColor)
	fmt.Printf("%sCustom blue message%s\n", logger.BlueColor, logger.EndColor)

	fmt.Println()
	fmt.Println("=== Demonstration complete ===")
}
