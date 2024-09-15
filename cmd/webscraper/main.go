package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"go_concurrency/internal/scraper"
)

// Initialize a structured logger
var logger = initLogger()

func initLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	return logger
}

func main() {
	start := time.Now()
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}

	// Context with timeout for managing the lifetime of goroutines
	goRuntime, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Rate limiter with a buffered channel to control concurrency rate
	rateLimiter := make(chan struct{}, 5) // Limits to 5 concurrent tasks

	// Apply Fan-Out pattern with rate limiting and worker pool
	results := scraper.RateLimitedFanOut(goRuntime, urls, 5, rateLimiter)

	// Apply Fan-In pattern to aggregate results
	collectedResults := scraper.FanIn(results)

	// Print all collected results
	for _, result := range collectedResults {
		fmt.Println(result)
	}

	logger.Infof("Elapsed time: %v", time.Since(start))
}
