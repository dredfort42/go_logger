# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-08-07

### Added

-   Initial release of the Go Logger package
-   Four pre-configured loggers: Error, Info, Warning, Debug
-   Colorized output with different colors for each log level
-   Smart output routing (errors to stderr, others to stdout)
-   Thread-safe concurrent logging support
-   Color constants for custom usage
-   Comprehensive test suite with 100% coverage
-   Performance benchmarks
-   Example usage program
-   Complete documentation

### Features

-   **Error**: Red colored, outputs to stderr, includes file information
-   **Info**: Green colored, outputs to stdout
-   **Warning**: Yellow colored, outputs to stdout
-   **Debug**: Blue colored, outputs to stdout
-   **Color Constants**: RedColor, GreenColor, YellowColor, BlueColor, EndColor

### Performance

-   High performance logging built on Go's standard log package
-   Benchmark results:
    -   Error: ~1210 ns/op
    -   Info: ~471 ns/op
    -   Warning: ~536 ns/op
    -   Debug: ~487 ns/op
