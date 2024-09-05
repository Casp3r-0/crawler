package main

import (
	//	"net/url"
	"fmt"
	"regexp"
	"strings"
	// "errors"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

func normalizeURL(url string) (string, error) {

	if url == "" {
		return "", fmt.Errorf("Invalid URL")
	}
	// remove Scheme
	re := regexp.MustCompile(`.*\/\/`)
	s := re.ReplaceAllString(url, "")

	// remove trailing /
	p := strings.TrimSuffix(s, "/")

	return p, nil

}
