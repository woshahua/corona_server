package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	cache "github.com/patrickmn/go-cache"
	"github.com/woshahua/corona_server/models"
)

type NewsData struct {
	Articles []Article `json: "articles"`
}

type Article struct {
	Source      Source    `json: "source"`
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Url         string    `json: url`
	PublishedAt time.Time `json: publishedAt`
}

type Source struct {
	Name string `json: "name"`
}

var Cache = cache.New(time.Duration(-1), time.Duration(-1))
var apiKey = "7795afb75e204bdaaf761f709ce7c48f"
var keyword = "コロナウィルス"

func FetchNewsData() {
	url := fmt.Sprintf("http://newsapi.org/v2/everything?q=%s&sortBy=publishedAt&apiKey=%s", keyword, apiKey)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var news NewsData
	err = json.Unmarshal(body, &news)
	if err != nil {
		log.Fatal(err)
	}

	var newsList []models.News
	now := transferToJSTTime(time.Now())
	// transfer parsed model to our original news model
	for _, article := range news.Articles {
		var news models.News

		news.Title = article.Title
		news.Description = article.Description
		news.UpdatedTime = transferToJSTTime(article.PublishedAt)
		news.Link = article.Url
		news.PassedHour = int(now.Sub(news.UpdatedTime).Hours())
		news.PassedDay = int((now.Sub(news.UpdatedTime).Hours()) / 24.0)
		news.PassedMinutes = int(time.Now().Sub(news.UpdatedTime).Minutes())

		newsList = append(newsList, news)
	}

	Cache.Set("news", newsList, cache.DefaultExpiration)
}

func transferToJSTTime(utcTime time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	jstTime := utcTime.UTC().In(jst)
	return jstTime
}
