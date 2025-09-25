package main

import (
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestGetURLsFromHtml(t *testing.T) {
	var nada []string
	tests := []struct {
		name          string
		inputURL      string
		inputBody     string
		expected      []string
		errorContains string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>

					<span>Boot.dev</span>


					<span>Boot.dev</span>

			</body>
		</html>
		`,
			expected: nada,
		},
		{
			name:      "no HTML",
			inputURL:  "https://blog.boot.dev",
			inputBody: "",
			expected:  nada,
		},
		{
			name:     "absolute and relative URLs 2",
			inputURL: "https://blog.boot.dev",
			inputBody: `
		<html>
			<body>
			<p>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://blog.boot.dev/path/one">
					<span>Boot.dev</span>
				</a>
			</p>
			<p>
				<a href="/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://other.com/path/one">
					<span>Boot.dev</span>
				</a>
				<a href="https://blog.boot.dev/path/one">
					<span>Boot.dev</span>
				</a>
			</p>			
			</body>
		</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one", "https://blog.boot.dev/path/one", "https://blog.boot.dev/path/one", "https://other.com/path/one", "https://blog.boot.dev/path/one"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			baseURL, err := url.Parse(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: couldn't parse input URL: %v", i, tc.name, err)
				return
			}
			actual, err := getURLsFromHTML(tc.inputBody, baseURL)
			if err != nil && !strings.Contains(err.Error(), tc.errorContains) {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err != nil && tc.errorContains == "" {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			} else if err == nil && tc.errorContains != "" {
				t.Errorf("Test %v - '%s' FAIL: expected error containing '%v', got none.", i, tc.name, tc.errorContains)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
