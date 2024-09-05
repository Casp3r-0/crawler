package main

import (
	"fmt"
	"os"
)

func main() {
	baseURL := os.Args[1]

	if len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	const maxConcurrency = 10
	cfg, err := configure(baseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		return
	}
	fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("Pages: %d - %s\n", count, normalizedURL)
	}
}
