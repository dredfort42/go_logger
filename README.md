# Go Logger

A simple, colorized logging utility for Go applications that provides pre-configured loggers for different log levels.

## Features

- **Colorized Output**: Different colors for different log levels (Error: Red, Info: Green, Warning: Yellow, Debug: Blue)
- **Multiple Log Levels**: Error, Info, Warning, and Debug loggers
- **Smart Output**: Errors go to stderr, other logs go to stdout
- **File Information**: Error logs include file and line information
- **High Performance**: Lightweight wrapper around Go's standard `log` package
- **Thread-Safe**: Safe for concurrent use

## Installation

```bash
go get github.com/dredfort42/go_logger
```

## Quick Start

```go
package main

import "github.com/dredfort42/go_logger"

func main() {
    logger.Info.Println("Application started")
    logger.Warning.Println("This is a warning")
    logger.Error.Println("An error occurred")
    logger.Debug.Println("Debug information")
}
```

## Available Loggers

### Error

-   **Color**: Red
-   **Output**: stderr
-   **Flags**: Date, Time, Short file path
-   **Use**: For error messages

```go
logger.Error.Println("Database connection failed")
```

### Info

-   **Color**: Green
-   **Output**: stdout
-   **Flags**: Date, Time
-   **Use**: For general information

```go
logger.Info.Println("Server started on port 8080")
```

### Warning

-   **Color**: Yellow
-   **Output**: stdout
-   **Flags**: Date, Time
-   **Use**: For warning messages

```go
logger.Warning.Println("Deprecated function used")
```

### Debug

-   **Color**: Blue
-   **Output**: stdout
-   **Flags**: Date, Time
-   **Use**: For debug information

```go
logger.Debug.Printf("Processing user ID: %d", userID)
```

## Color Constants

The package also exports color constants that you can use in your own logging:

```go
fmt.Printf("%sError:%s Something went wrong\n", logger.RedColor, logger.EndColor)
```

Available constants:

-   `RedColor`
-   `GreenColor`
-   `YellowColor`
-   `BlueColor`
-   `EndColor`

## Example Output

```text
2025/08/07 10:30:45 INFO:    Application started
2025/08/07 10:30:45 WARNING: This is a warning
2025/08/07 10:30:45 logger.go:15: ERROR:   An error occurred
2025/08/07 10:30:45 DEBUG:   Debug information
```

## Performance

The logger is built on top of Go's standard `log` package and maintains excellent performance:

```text
BenchmarkError-4       862816       1210 ns/op
BenchmarkInfo-4       2721848       471.3 ns/op
BenchmarkWarning-4    1967634       535.9 ns/op
BenchmarkDebug-4      2258250       487.2 ns/op
```

## Requirements

-   Go 1.24 or later

## License

This project is licensed under the GNU General Public License v3.0 - see the [LICENSE](LICENSE) file for details.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Author

**Dmitry Novikov**

-   Email: [dredfort.42@gmail.com](mailto:dredfort.42@gmail.com)
-   GitHub: [@dredfort42](https://github.com/dredfort42)
-   LinkedIn: [novikov-da](https://linkedin.com/in/novikov-da)
