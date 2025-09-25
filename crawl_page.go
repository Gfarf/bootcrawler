package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	cfg.mu.Lock()
	vPages := len(cfg.pages)
	cfg.mu.Unlock()
	fmt.Printf("starting on new URL: %s, pages visited: %d, max pages: %d\n", rawCurrentURL, vPages, cfg.maxPages)
	if vPages >= cfg.maxPages {
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		//fmt.Println(err)
		return
	}
	if cfg.baseURL.Hostname() != currentURL.Hostname() {
		//fmt.Println("not in same domain")
		return
	}
	nURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		//fmt.Println(err)
		return
	}
	cfg.mu.Lock()
	if !cfg.addPageVisit(nURL) {
		//fmt.Println("Página já mapeada")
		cfg.mu.Unlock()
		return
	}
	htmlText, err := getHTML(rawCurrentURL)
	if err != nil {
		cfg.mu.Unlock()
		//fmt.Printf("estamos na crawl pegando html, %v\n", err)
		return
	}
	cfg.pages[nURL] = extractPageData(htmlText, rawCurrentURL)
	cfg.mu.Unlock()
	//fmt.Println(htmlText)
	allurls, err := getURLsFromHTML(htmlText, cfg.baseURL)
	if err != nil {
		//fmt.Printf("estamos na crawl pegando urls, %v\n", err)
		return
	}
	for _, newUrl := range allurls {
		//fmt.Printf("estamos na recursão, com base %s, novo %s\n", rawBaseURL, newUrl)
		cfg.wg.Add(1)
		go cfg.crawlPage(newUrl)
	}
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	data, ok := cfg.pages[normalizedURL]
	if ok {
		//fmt.Println("Página já mapeada")
		data.count += 1
		cfg.pages[normalizedURL] = data
		return false
	}
	return true
}

func getH1FromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	return doc.Find("h1").First().Text()
}

func getFirstParagraphFromHTML(html string) string {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return ""
	}
	withMain := doc.Find("main").First().Find("p").First().Text()
	if withMain == "" {
		return doc.Find("p").First().Text()
	}
	return withMain
}

/*
func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {

	//fmt.Printf("Crawling: base %s, current %s\n", rawBaseURL, rawCurrentURL)

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		//fmt.Println(err)
		return
	}
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		//fmt.Println(err)
		return
	}
	if baseURL.Hostname() != currentURL.Hostname() {
		//fmt.Println("not in same domain")
		return
	}

	nURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		//fmt.Println(err)
		return
	}
	_, ok := pages[nURL]
	if ok {
		//fmt.Println("Página já mapeada")
		pages[nURL] += 1
		return
	}
	pages[nURL] = 1
	htmlText, err := getHTML(rawCurrentURL)
	if err != nil {
		//fmt.Printf("estamos na crawl pegando html, %v\n", err)
		return
	}
	//fmt.Println(htmlText)
	allurls, err := getURLsFromHTML(htmlText, baseURL)
	if err != nil {
		//fmt.Printf("estamos na crawl pegando urls, %v\n", err)
		return
	}
	for _, newUrl := range allurls {
		//fmt.Printf("estamos na recursão, com base %s, novo %s\n", rawBaseURL, newUrl)
		crawlPage(rawBaseURL, newUrl, pages)
	}
}


The no goquery working version
func getH1FromHTML(htmlBody string) string {
	hNode, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return ""
	}
	for inner := range hNode.Descendants() {
		if inner.Type == html.ElementNode && inner.Data == "h1" {
			return getTextContent(inner)
		}
	}
	return ""
}

func getFirstParagraphFromHTML(htmlBody string) string {
	hNode, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return ""
	}
	for inner := range hNode.Descendants() {
		if inner.Type == html.ElementNode && inner.Data == "p" {
			return getTextContent(inner)
		}
	}
	return ""
}

func getTextContent(n *html.Node) string {
	var result strings.Builder
	var extract func(*html.Node)
	extract = func(node *html.Node) {
		if node.Type == html.TextNode {
			result.WriteString(strings.TrimSpace(node.Data))
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			extract(c)
		}
	}
	extract(n)
	return result.String()
}*/
