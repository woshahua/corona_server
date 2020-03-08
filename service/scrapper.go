package service

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
	"github.com/woshahua/corona_server/models"
)

func ScrapNews() {

	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}
	if err := page.Navigate("https://xn--eckd2b0a6fujka.com/"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	err = page.FindByClass("tab").First("li").Click()
	if err != nil {
		log.Fatalf("Failed to fech dom contents:%v", err)
	}

	// wait for page change
	time.Sleep(3 * time.Second)

	htmlContents, err := page.HTML()
	if err != nil {
		log.Fatalf("Failed to fech dom contents:%v", err)
	}

	readerHtmlContents := strings.NewReader(htmlContents)
	contentsDom, err := goquery.NewDocumentFromReader(readerHtmlContents)
	if err != nil {
		log.Fatalf("Read dom contents from reader failed:%v", err)
	}

	var newsList []models.News

	articleList := contentsDom.Find("ul.list").First().Children()
	articleList.Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Find("a").Attr("href")
		title := s.Find("p").Text()
		news := models.News{Title: title, Link: url}
		newsList = append(newsList, news)
	})

	// reverse news list
	for i, j := 0, len(newsList)-1; i < j; i, j = i+1, j-1 {
		newsList[i], newsList[j] = newsList[j], newsList[i]
	}
	for _, news := range newsList {
		models.InsertNews(&news)
	}

	time.Sleep(1 * time.Second)
}
