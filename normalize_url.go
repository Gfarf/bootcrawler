package main

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(s string) (string, error) {
	rurl, err := url.Parse(s)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	res := rurl.Host + rurl.Path
	res = strings.ToLower(res)
	res = strings.TrimSuffix(res, "/")
	if strings.Contains(res, " ") {
		return "", errors.New("not a URL")
	}
	return res, nil
}
