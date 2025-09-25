package main

import (
	"fmt"
	"log"
	"net/url"
)

type PageData struct {
	URL            string
	H1             string
	FirstParagraph string
	OutgoingLinks  []string
	ImageURLs      []string
	count          int
}

func extractPageData(html, pageURL string) PageData {
	data := PageData{}
	rurl, err := url.Parse(pageURL)
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't parse URL: %w", err))
	}
	links, err := getURLsFromHTML(html, rurl)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting images from URL: %w", err))
	}
	imgs, err := getImagesFromHTML(html, rurl)
	if err != nil {
		log.Fatal(fmt.Errorf("error getting links from URL: %w", err))
	}
	data.URL = pageURL
	data.H1 = getH1FromHTML(html)
	data.FirstParagraph = getFirstParagraphFromHTML(html)
	data.OutgoingLinks = links
	data.ImageURLs = imgs
	data.count = 1
	return data
}
