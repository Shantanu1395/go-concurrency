
Add testcases, benchmarks and standards to it

# Go Web Scraper with Concurrency Patterns

This project is a web scraper implemented in Go, utilizing various concurrency patterns such as Worker Pools, Fan-Out/Fan-In, Rate Limiting, and Context Management.

## Features
- **Worker Pool**: Manages a fixed number of goroutines to perform tasks concurrently.
- **Fan-Out/Fan-In**: Distributes tasks among multiple worker goroutines and consolidates results.
- **Rate Limiting**: Controls the rate of execution of goroutines to prevent resource exhaustion.
- **Context Management**: Handles goroutine cancellation gracefully.

## Setup
1. **Clone the Repository and Extract Files**:
    - Download the project zip file and extract it to your desired directory. Ensure the directory structure looks like this:
      ```
      go_web_scraper/
      ├── cmd/
      │   └── webscraper/
      │       └── main.go        # Main application code
      ├── internal/
      │   ├── scraper/
      │   │   ├── scraper.go     # Core scraper logic and concurrency patterns
      │   │   └── scraper_test.go # Tests for scraper functionality
      ├── README.md              # Project documentation
      └── DIRECTORY_STRUCTURE.md # Directory structure documentation
      ```

2. **Navigate to the Project Directory**:
    - Open your terminal and navigate to the root of the extracted project folder:
      ```bash
      cd /path/to/go_web_scraper
      ```

3. **Initialize the Go Module** (if needed):
    - If the project is not already a Go module, initialize it by running:
      ```bash
      go mod init go_concurrency
      go mod tidy
      ```
    - This will set up the module and fetch any dependencies.

## Running the Application
Navigate to the main application directory and run the following command:
```bash
go run main.go
```

## Run all testcases
### gin server
```bash
go test ./internal/server/... -v
```

## Benchmarking
To benchmark the concurrency performance, execute:
```bash
go test -bench=. ./internal/server/...
```

## Checking for Memory Leaks
The test suite includes checks for memory leaks. Run:
```bash
go test -run TestMemoryLeak ./internal/server/...
```

## Project Structure
- **cmd/webscraper/main.go**: Main application code implementing the web scraper.
- **internal/scraper/scraper.go**: Core logic for the web scraper, including concurrency patterns.
- **internal/scraper/scraper_test.go**: Test and benchmark cases for the application.
- **README.md**: Project documentation.

## Notes
This project demonstrates Go’s powerful concurrency model through practical use cases. Adjust the rate limiter and worker count based on your system's capabilities and requirements.

# Commiting changes
```bash
    git init
    git add .
    git commit -m "Initial commit"
    git remote add origin https://github.com/Shantanu1395/map_info
    git pull origin main --allow-unrelated-histories
    git merge main --allow-unrelated-histories
    git push origin main
 ```

git remote add origin https://github.com/Shantanu1395/go-concurrency.git
git remote set-url origin https://<TOKEN>@github.com/Shantanu1395/go-concurrency.git
