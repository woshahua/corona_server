package models

import (
	"log"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title         string    `json: "title"`
	Link          string    `json: "link"`
	Description   string    `josn: "description"`
	UpdatedTime   time.Time `json: "updated_time"`
	PassedHour    int       `json: "passed_hour`
	PassedMinutes int       `json: "passed_minutes`
	PassedDay     int       `json: "passed_day`
}

func GetNews(number int) (*[]News, error) {
	var news []News
	err := db.Order("updated_time desc").Limit(number).Find(&news).Error

	return &news, err
}

func InsertNews(news *News) {
	if news.Title != "" {
		var existNews News
		data := strings.Split(news.Title, " ")
		title := strings.Join(data[:len(data)-4], " ")
		updatedTime := strings.Join(data[len(data)-2:], " ")

		// jst time zone
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().UTC().In(jst)

		news.Title = title
		news.UpdatedTime, _ = time.Parse("2006/1/2 15:04:05", updatedTime)
		news.PassedHour = int(now.Sub(news.UpdatedTime).Hours() + 9.0)
		news.PassedDay = int((now.Sub(news.UpdatedTime).Hours() + 9.0) / 24.0)
		news.PassedMinutes = int(time.Now().Sub(news.UpdatedTime).Minutes() + 9*60)

		notExist := db.Find(&existNews, "title = ?", news.Title).RecordNotFound()
		if notExist {
			db.NewRecord(news)
			db.Create(&news)
			err := db.Save(&news).Error
			if err != nil {
				log.Fatalf("failed insert patient: %v", err)
			}
		}
	}
}
