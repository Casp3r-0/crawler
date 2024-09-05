package main

import (
	"fmt"
	"net/url"
	"os"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	parsedBaseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("Invalid Base URL: %v", err)
		return
	}
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Invalid Current URL: %v", err)
		return
	}

	if parsedBaseURL.Hostname() != parsedCurrentURL.Hostname() {
		fmt.Printf("Current URL out of domain: %s", parsedCurrentURL)
		return
	}

	normalizedCurrentPage, _ := normalizeURL(rawCurrentURL)

	if _, exists := pages[normalizedCurrentPage]; exists {
		pages[normalizedCurrentPage]++
		return
	}
	pages[normalizedCurrentPage] = 1

	fmt.Printf("crawling %s\n", rawCurrentURL)

	body, _ := getHTML(rawCurrentURL)

	links, _ := getURLsFromHTML(body, rawBaseURL)

	for _, link := range links {
		crawlPage(rawBaseURL, link, pages)
	}

	return
}

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

	fmt.Printf("starting crawl of: %s\n", baseURL)

	pages := make(map[string]int)
	crawlPage(baseURL, baseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("Pages: %d - %s\n", count, normalizedURL)
	}
}
