package main

import (
	"fmt"
	"sync"
)

// --- Part 1: The Safe Cache (Simulating Shared Memory) ---
// In Lab 1, this represents the Coordinator's state.
type SafeCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

// Helper to check and mark a URL as visited safely.
// Returns true if ALREADY visited, false if NEW.
func (c *SafeCache) CheckAndMark(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.visited[url] {
		return true
	}
	c.visited[url] = true
	return false
}

// --- Part 2: The Crawler (Simulating a Worker) ---
type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

// Note: We pass *sync.WaitGroup and *SafeCache pointers so all threads share the same instance.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup, cache *SafeCache) {
	// TODO: 1. Ensure we signal "Done" when this function exits (use defer)
	// hint: defer wg.Done()

	// TODO: 2. Check if depth is <= 0. If so, return.
	
	// TODO: 3. Check if the URL has already been visited (thread-safe).
	// hint: if cache.CheckAndMark(url) { return }

	// TODO: 4. Fetch the data
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	// TODO: 5. Loop through the found URLs and start new goroutines
	// For every URL found:
	//   a. Increment the WaitGroup counter
	//   b. Launch a goroutine to crawl that URL
	for _, u := range urls {
		// hint: wg.Add(1)
		// hint: go Crawl(u, depth-1, fetcher, wg, cache)
	}
}

// --- Part 3: Main (Simulating the System Start) ---
func main() {
	// Initialize shared state
	cache := SafeCache{visited: make(map[string]bool)}
	var wg sync.WaitGroup

	// Kick off the first task
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg, &cache)

	// TODO: 6. Block here until the entire tree is explored
	// hint: wg.Wait()
	
	// If you don't wait, the program exits immediately!
	fmt.Println("Done crawling!")
}

// --- Part 4: The Fake Fetcher (Do not modify) ---
// This mocks a real network call.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// dummy data
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
