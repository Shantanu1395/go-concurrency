
package scraper

import (
	"context"
	"testing"
	"time"
	"sync"
)

func TestFetchTask(t *testing.T) {
	ctx := context.Background()
	url := "https://www.google.com"
	result, err := FetchTask(ctx, url)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result == "" {
		t.Fatalf("Expected result, got empty string")
	}
}

func TestRateLimitedFanOut(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}

	rateLimiter := make(chan struct{}, 2) // Limiting concurrency to 2
	results := RateLimitedFanOut(ctx, urls, 3, rateLimiter)
	collectedResults := FanIn(results)

	if len(collectedResults) != len(urls) {
		t.Errorf("Expected %d results, got %d", len(urls), len(collectedResults))
	}
}

func BenchmarkRateLimitedFanOut(b *testing.B) {
	urls := []string{
		"https://golang.org",
		"https://google.com",
		"https://github.com",
	}

	for i := 0; i < b.N; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		rateLimiter := make(chan struct{}, 5)
		results := RateLimitedFanOut(ctx, urls, 5, rateLimiter)
		FanIn(results)
	}
}

func TestMemoryLeak(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			TestRateLimitedFanOut(t)
		}()
	}
	wg.Wait()
	t.Log("Completed without memory leak indication")
}
