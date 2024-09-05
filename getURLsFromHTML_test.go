package main

import (
	"net/url"
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name         string
		htmlBody     string
		rawBaseURL   string
		expectedURLs []string
		expectError  bool
	}{
		{
			name: "Valid HTML with absolute and relative URLs",
			htmlBody: `
				<html>
					<body>
						<a href="https://example.com">Example</a>
						<a href="/relative/path">Relative</a>
						<a href="https://google.com">Google</a>
					</body>
				</html>
			`,
			rawBaseURL: "https://mywebsite.com",
			expectedURLs: []string{
				"https://example.com",
				"https://mywebsite.com/relative/path",
				"https://google.com",
			},
			expectError: false,
		},
		{
			name: "HTML with no links",
			htmlBody: `
				<html>
					<body>
						<p>No links here</p>
					</body>
				</html>
			`,
			rawBaseURL:   "https://mywebsite.com",
			expectedURLs: []string{},
			expectError:  false,
		},
		{
			name:         "Invalid HTML",
			htmlBody:     "This is not valid HTML",
			rawBaseURL:   "https://mywebsite.com",
			expectedURLs: nil,
			expectError:  false,
		},
		{
			name: "HTML with invalid href",
			htmlBody: `
				<html>
					<body>
						<a href="https://valid.com">Valid</a>
						<a href="://invalid">Invalid</a>
					</body>
				</html>
			`,
			rawBaseURL: "https://mywebsite.com",
			expectedURLs: []string{
				"https://valid.com",
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseURL, _ := url.Parse("www.example.com")
			urls, err := getURLsFromHTML(tt.htmlBody, baseURL)

			if tt.expectError {
				if err == nil {
					t.Errorf("Expected an error, but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}

			if len(urls) == 0 && len(tt.expectedURLs) == 0 {
				return
			}
			if !reflect.DeepEqual(urls, tt.expectedURLs) {
				t.Errorf("Expected URLs %v, but got %v", tt.expectedURLs, urls)
			}
		})
	}
}
