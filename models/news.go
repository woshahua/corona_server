package models

import (
	"log"

	"github.com/jinzhu/gorm"
)

type News struct {
	gorm.Model
	Title string `json: "title", gorm: "date"`
	Link  string `json: "link", gorm: "age"`
}

func GetNews(number int) (*[]News, error) {
	var news []News
	err := db.Order("id desc").Limit(number).Find(&news).Error

	return &news, err
}

func InsertNews(news *News) {
	var existNews News
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
