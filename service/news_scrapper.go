package service

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func ScapNewsSummary(url string) (string, error) {
	description := ""

	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Println("error scraping open site", err)
		return "", err
	}

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		if name, _ := s.Attr("name"); name == "description" {
			description, _ = s.Attr("content")
		}
	})

	return description, nil
}