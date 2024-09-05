package main

import (
	"fmt"
	"net/url"
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
