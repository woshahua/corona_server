package main

import (
	"corona_server/environment"
	"corona_server/infrastructure/mysql"
	"corona_server/routers"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/carlescere/scheduler"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/service"
)

func main() {
	//router := routers.InitRouter()
	//
	//s := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	//	Handler:        router,
	//	ReadTimeout:    setting.ReadTimeout,
	//	WriteTimeout:   setting.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//// start a new routine
	//go RunCronJob()
	//s.ListenAndServe()

	if !environment.IsAppEngine() {
		if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/dev_credential.json"); err != nil {
			log.Fatal(err)
		}
	}

	if err := mysql.Connection(); err != nil {
		log.Fatal(err)
	}

	if err := routers.InitRouter().Run(environment.GetSharedEnvironments().IPAddress + ":" + environment.GetSharedEnvironments().Port); err != nil {
		log.Fatal(err)
	}
}

func RunCronJob() {
	log.Println("starting...")

	fetchJapnesePatientCSV := func() {
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

	importCSVDataToDB := func() {
		jst := time.FixedZone("Asia/Tokyo", 9*60*60)
		now := time.Now().UTC().In(jst)

		var updateTime = models.LastUpdateTime{}
		updateTime.PatientDataUpdateTime = now.Format("2016/1/2 15:04:05")
		_ = models.UpdateDataUpdatedTime(&updateTime)

		log.Println("Run service.Import")
		err := service.Import()
		if err != nil {
			log.Println("failed import csv file: ", err)
		}
	}

	scrapNewsData := func() {
		log.Println("Run service.ScrapNews")
		service.ScrapNews()
	}

	// fetch newest japanese patient data from:
	// https://toyokeizai.net/sp/visual/tko/covid19/csv/data.csv
	scheduler.Every(3).Hours().Run(fetchJapnesePatientCSV)

	// insert csv data to database evenry 6 hours
	scheduler.Every(12).Hours().Run(importCSVDataToDB)

	// scrap news data and import to database
	scheduler.Every(15).Minutes().Run(scrapNewsData)
	runtime.Goexit()
}
