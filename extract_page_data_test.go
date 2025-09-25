package main

import (
	"reflect"
	"testing"
)

func TestExtractPageData(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  PageData
	}{
		{
			name:     "got all fields",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<h1>Test Title</h1>
			<p>This is the first paragraph.</p>
			<a href="/link1">Link 1</a>
			<img src="/image1.jpg" alt="Image 1">
		</body></html>`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
				ImageURLs:      []string{"https://blog.boot.dev/image1.jpg"},
			},
		},
		{
			name:     "no images",
			inputURL: "https://blog.boot.dev",
			inputBody: `<html><body>
			<h1>Test Title</h1>
			<p>This is the first paragraph.</p>
			<a href="/link1">Link 1</a>
		</body></html>`,
			expected: PageData{
				URL:            "https://blog.boot.dev",
				H1:             "Test Title",
				FirstParagraph: "This is the first paragraph.",
				OutgoingLinks:  []string{"https://blog.boot.dev/link1"},
			},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := extractPageData(tc.inputBody, tc.inputURL)

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}

		})
	}
}
