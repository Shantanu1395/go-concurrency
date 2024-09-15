package scraper

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
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

// recoverFromPanic abstracts panic recovery and logging
func recoverFromPanic(url string) {
	if r := recover(); r != nil {
		logger.WithFields(logrus.Fields{
			"url":   url,
			"error": r,
		}).Error("Panic occurred")
	}
}

// fetchTask represents a task that fetches data from a URL
func FetchTask(ctx context.Context, url string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("request creation failed for %s: %v", url, err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error fetching %s: %v", url, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body of %s: %v", url, err)
	}

	return fmt.Sprintf("Fetched %d bytes from %s", len(body), url), nil
}

// worker performs fetch tasks from the jobs channel and sends results to the results channel
func Worker(ctx context.Context, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer recoverFromPanic("worker")

	for url := range jobs {
		select {
		case <-ctx.Done():
			logger.Warn("Worker stopped due to context cancellation")
			return
		default:
			result, err := FetchTask(ctx, url)
			if err != nil {
				logger.Error(err)
				results <- fmt.Sprintf("Error: %v", err)
			} else {
				results <- result
			}
		}
	}
}

// rateLimitedFanOut controls the rate of task execution by limiting active workers
func RateLimitedFanOut(ctx context.Context, urls []string, workerCount int, rateLimiter chan struct{}) <-chan string {
	jobs := make(chan string, len(urls))    // Task queue
	results := make(chan string, len(urls)) // Results queue
	var wg sync.WaitGroup

	// Start workers (Fan-Out)
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go Worker(ctx, jobs, results, &wg)
	}

	// Enqueue tasks and enforce rate limit
	go func() {
		defer close(jobs)
		for _, url := range urls {
			rateLimiter <- struct{}{} // Acquire rate limit slot
			jobs <- url
			<-rateLimiter // Release rate limit slot
		}
	}()

	// Close results channel after all workers complete
	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}

// fanIn aggregates results from multiple workers into a single stream
func FanIn(results <-chan string) []string {
	var collectedResults []string
	for result := range results {
		collectedResults = append(collectedResults, result)
	}
	return collectedResults
}
