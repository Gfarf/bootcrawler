package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func getHTML(rawURL string) (string, error) {
	c := NewClient(20 * time.Second)
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error setting first request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error getting first request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("server error %d, %v", resp.StatusCode, resp.Status)
	}
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("wrong content type:  %s", contentType)
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading body data: %v", err)
	}
	return string(dat), nil
}
