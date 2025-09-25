package main

import (
	"encoding/csv"
	"os"
	"strings"
)

//page_url, h1, first_paragraph, outgoing_link_urls, image_urls

func writeCSVReport(pages map[string]PageData, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(f)
	err = writer.Write([]string{"page_url", "h1", "first_paragraph", "outgoing_link_urls", "image_urls"})
	if err != nil {
		return err
	}
	for _, values := range pages {
		err = writer.Write([]string{
			values.URL,
			values.H1,
			values.FirstParagraph,
			strings.Join(values.OutgoingLinks, ";"),
			strings.Join(values.ImageURLs, ";")})
		if err != nil {
			return err
		}
	}
	return nil
}
