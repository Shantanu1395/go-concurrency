
# Go Web Scraper Project Documentation

This document provides an overview of the Go Web Scraper project, including the execution order of files, their roles, and a detailed listing of the project directory structure with the contents of each file.

## Execution Order of Files

### 1. `cmd/webscraper/main.go`
- **Role**: Entry point of the application.
- **Execution**: Initializes the application, sets up context, configures concurrency settings, and invokes the scraper logic using various concurrency patterns.

**Contents**:
- Initializes a structured logger.
- Defines a list of URLs to scrape.
- Sets up a context with a timeout to manage goroutine lifetimes.
- Configures rate limiting using a buffered channel.
- Calls `RateLimitedFanOut` to start the worker pool and manage tasks.
- Aggregates results using `FanIn` and prints them.

### 2. `internal/scraper/scraper.go`
- **Role**: Core logic and concurrency implementation of the web scraper.
- **Execution**: Called from `main.go` to perform the actual scraping tasks and manage concurrency through worker pools, fan-out, and fan-in patterns.

**Contents**:
- `FetchTask`: Fetches data from a given URL.
- `Worker`: Executes tasks fetched from the `jobs` channel and sends results to the `results` channel.
- `RateLimitedFanOut`: Manages the distribution of tasks to workers and controls concurrency with rate limiting.
- `FanIn`: Aggregates results from multiple workers into a single stream.

### 3. `internal/scraper/scraper_test.go`
- **Role**: Contains test cases to verify the functionality and performance of the scraper.
- **Execution**: Run to ensure the scraper works correctly, handles concurrency as expected, and checks for memory leaks.

**Contents**:
- `TestFetchTask`: Tests the `FetchTask` function for successful data retrieval.
- `TestRateLimitedFanOut`: Tests the `RateLimitedFanOut` function for correct task distribution and result aggregation.
- `BenchmarkRateLimitedFanOut`: Benchmarks the `RateLimitedFanOut` function to measure performance under load.
- `TestMemoryLeak`: Checks for potential memory leaks by running multiple instances of the scraper in parallel.

### 4. `internal/scraper/scraper_leak_test.go`
- **Role**: Specifically designed to demonstrate a visible memory leak.
- **Execution**: Simulates a scenario where goroutines are not properly terminated, leading to increased memory usage and potential leaks.

**Contents**:
- `leakFunction`: A function that simulates a memory leak by continuously allocating memory without cleanup.
- `TestVisibleMemoryLeak`: Executes `leakFunction` in multiple goroutines and monitors for persistent memory usage to detect leaks.

### 5. Documentation Files
- **`README.md`**
  - **Role**: Provides comprehensive project documentation, including setup, execution instructions, testing commands, and troubleshooting steps.
- **`DIRECTORY_STRUCTURE.md`**
  - **Role**: Documents the directory layout of the project, providing an overview of where each file is located and its purpose.

## New Directory Structure

```
go_web_scraper/
│
├── cmd/
│   └── webscraper/
│       └── main.go        # Main application code
│
├── internal/
│   ├── scraper/
│   │   ├── scraper.go     # Core scraper logic and concurrency patterns
│   │   ├── scraper_test.go # Tests for scraper functionality
│   │   └── scraper_leak_test.go # Test to demonstrate a visible memory leak
│
├── README.md              # Project documentation
├── DIRECTORY_STRUCTURE.md # Directory structure documentation
└── DOCUMENTATION.md       # Detailed execution order and contents of files
```

## Summary
- **Execution Flow**: The application starts with `main.go`, which orchestrates the scraper by calling functions in `scraper.go`. Test files (`scraper_test.go` and `scraper_leak_test.go`) validate and benchmark the code.
- **Concurrency Management**: Core functionality like worker pools and task distribution are handled in `scraper.go`.
- **Testing and Validation**: Tests ensure the application works correctly and is free from memory leaks, while benchmarks help gauge performance under various conditions.
