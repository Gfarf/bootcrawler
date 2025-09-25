package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(html string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return []string{""}, fmt.Errorf("error parsing html Body, %v", err)
	}
	var res []string
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		src, ok := s.Attr("src")
		if !ok || strings.TrimSpace(src) == "" {
			return
		}

		u, err := url.Parse(src)
		if err != nil {
			fmt.Printf("couldn't parse href %q: %v\n", src, err)
			return
		}

		resolved := baseURL.ResolveReference(u)
		res = append(res, resolved.String())
	})

	return res, nil
}
