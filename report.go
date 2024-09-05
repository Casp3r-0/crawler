package main

import (
	"fmt"
	"sort"
)

type result struct {
	urlString string
	linkCount int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Printf("\n=============================\n  REPORT for %s\n=============================\n", baseURL)

	var structSlice []result
	for key, value := range pages {
		structSlice = append(structSlice, result{urlString: key, linkCount: value})
	}
	sort.Slice(structSlice, func(i, j int) bool { return structSlice[i].linkCount > structSlice[j].linkCount })

	for i := 0; i < len(structSlice); i++ {
		fmt.Printf("\nFound %d internal links to %s\n", structSlice[i].linkCount, structSlice[i].urlString)
	}

}
