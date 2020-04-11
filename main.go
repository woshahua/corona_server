package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"corona_server/pkg/setting"
	"corona_server/routers"
	"corona_server/service"
	"github.com/carlescere/scheduler"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	// start a new routine
	go RunCronJob()
	s.ListenAndServe()
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

		filePath = "staticFile/patientDetail.csv"
		url = "https://docs.google.com/spreadsheets/d/10MFfRQTblbOpuvOs_yjIYgntpMGBg592dL8veXoPpp4/export?format=csv&gid=0"
		err = service.DownLoadFile(filePath, url)
		if err != nil {
			log.Println("faild fetch csv file", err)
		}
	}

	importCSVDataToDB := func() {
		log.Println("Run service.Import")
		err := service.Import()
		if err != nil {
			log.Println("failed import csv file: ", err)
		}
	}

	scrapNewsData := func() {
		log.Println("Run service.fetchNews")
		service.FetchNewsData()
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
