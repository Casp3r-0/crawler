package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	baseURL := os.Args[1]
	maxConcurrencyString := os.Args[2]
	maxPagesString := os.Args[3]

	if len(os.Args) < 4 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		os.Exit(1)
	}
	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(maxConcurrencyString)
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
		return
	}

	maxPages, err := strconv.Atoi(maxPagesString)
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
		return
	}

	cfg, err := configure(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("error - configure: %v", err)
		return
	}
	//	fmt.Printf("starting crawl of: %s\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	//	for normalizedURL, count := range cfg.pages {
	//		fmt.Printf("Pages: %d - %s\n", count, normalizedURL)
	//	}

	printReport(cfg.pages, baseURL)

}
