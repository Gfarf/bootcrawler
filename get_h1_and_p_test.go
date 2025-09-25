package main

import (
	"reflect"
	"testing"
)

func TestGetH1andP(t *testing.T) {
	tests := []struct {
		name       string
		inputURL   string
		inputBody  string
		expectedH1 string
		expectedP  string
	}{
		{
			name:     "has both",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<h1>Welcome to Boot.dev</h1>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>`,
			expectedH1: "Welcome to Boot.dev",
			expectedP:  "Learn to code by building real projects.",
		},
		{
			name:     "only h1",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<h1>Welcome to Boot.dev</h1>
				<main>
				</main>
			</body>
			</html>`,
			expectedH1: "Welcome to Boot.dev",
			expectedP:  "",
		},
		{
			name:     "only p",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>`,
			expectedH1: "",
			expectedP:  "Learn to code by building real projects.",
		},
		{
			name:     "p before main",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<p>This is the before paragraph.</p>
				<main>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
				</main>
			</body>
			</html>`,
			expectedH1: "",
			expectedP:  "Learn to code by building real projects.",
		},
		{
			name:     "p but no  main",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
			<body>
				<p>This is the before paragraph.</p>
				<p>Learn to code by building real projects.</p>
				<p>This is the second paragraph.</p>
			</body>
			</html>`,
			expectedH1: "",
			expectedP:  "This is the before paragraph.",
		},
	}
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actualH1 := getH1FromHTML(tc.inputBody)
			actualP := getFirstParagraphFromHTML(tc.inputBody)

			if !reflect.DeepEqual(actualH1, tc.expectedH1) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expectedH1, actualH1)
			}
			if !reflect.DeepEqual(actualP, tc.expectedP) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expectedP, actualP)
			}
		})
	}
}
