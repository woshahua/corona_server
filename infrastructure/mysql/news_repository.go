package mysql

import (
	"corona_server/models"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func GetNews(number int) (*[]models.News, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var news []models.News
	err := conn.Order("updated_time desc").Limit(number).Find(&news).Error
	return &news, err
}

func InsertNews(news *models.News) {
	conn, err = newConnection()
	if err != nil {}
	if news.Title != "" {
		var existNews models.News
		data := strings.Split(news.Title, " ")
		title := strings.Join(data[:len(data)-5], " ")

		// temporary handle date time string
		temp := strings.Replace(data[len(data)-3], ",", "", 1)
		dateTimeList := strings.Split(temp, "/")
		reversedTimeList := make([]string, len(dateTimeList))
		cnt := len(dateTimeList) - 1
		for cnt >= 0 {
			fmt.Println(reversedTimeList)
			reversedTimeList[len(dateTimeList)-cnt-1] = dateTimeList[cnt]
			cnt -= 1
		}

		// temporary handle clock tiem string, oh shit code, jesus forgive me

		var updateClock string
		if data[len(data)-1] == "PM" {
			updateClockList := strings.Split(data[len(data)-2], ":")
			fixHour, _ := strconv.Atoi(updateClockList[0])

			updateClockList[0] = string(fixHour + 12)
			updateClock = strings.Join(updateClockList, ":")
		} else {
			updateClock = data[len(data)-2]
		}

		updatedTime := strings.Join(reversedTimeList, "/") + " " + updateClock

		// jst time zone
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().UTC().In(jst)

		news.Title = title
		news.UpdatedTime, _ = time.Parse("2006/2/1 15:04:05", updatedTime)
		news.PassedHour = int(now.Sub(news.UpdatedTime).Hours() + 9.0)
		news.PassedDay = int((now.Sub(news.UpdatedTime).Hours() + 9.0) / 24.0)
		news.PassedMinutes = int(time.Now().Sub(news.UpdatedTime).Minutes() + 9*60)

		notExist := conn.Find(&existNews, "title = ?", news.Title).First(&existNews).RecordNotFound()
		if notExist {
			conn.NewRecord(news)
			conn.Create(&news)
			err := conn.Save(&news).Error
			if err != nil {
				log.Fatalf("failed insert patient: %v", err)
			}
		} else {
			existNews.UpdatedTime = news.UpdatedTime
			existNews.PassedDay = news.PassedDay
			existNews.PassedMinutes = news.PassedMinutes
			existNews.PassedHour = news.PassedHour

			conn.Save(&existNews)
		}
	}
}