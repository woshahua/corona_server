package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/infrastructure/mysql"
	"github.com/woshahua/corona_server/models"
	"log"
	"github.com/woshahua/corona_server/service"
	"time"
)

func FetchJapanesePatientCSV(c *gin.Context) {
	log.Println("Run service.Import")
	filePath := "staticFile/patientByDate.csv"
	url := "https://docs.google.com/spreadsheets/d/1jfB4muWkzKTR0daklmf8D5F0Uf_IYAgcx_-Ij9McClQ/export?format=csv&gid=211530313"
	err := service.DownLoadFile(filePath, url)

	filePath = "staticFile/patientByLocation.csv"
	url = "https://docs.google.com/spreadsheets/d/1jfB4muWkzKTR0daklmf8D5F0Uf_IYAgcx_-Ij9McClQ/export?format=csv&gid=1399411442"
	err = service.DownLoadFile(filePath, url)

	if err != nil {
		log.Println("faild fetch csv file", err)
	}
}

func ImportCSVDataToDB(c *gin.Context) {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().UTC().In(jst)

	var updateTime = models.LastUpdateTime{}
	updateTime.PatientDataUpdateTime = now.Format("2016/1/2 15:04:05")
	_ = mysql.UpdateDataUpdatedTime(&updateTime)

	log.Println("Run service.Import")
	err := service.Import()
	if err != nil {
		log.Println("failed import csv file: ", err)
	}
}

func FetchNewsData(c *gin.Context) {
	log.Println("Run service.ScrapNews")
	service.FetchNewsData()
}