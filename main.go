package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/carlescere/scheduler"
	"github.com/woshahua/corona_server/pkg/setting"
	"github.com/woshahua/corona_server/routers"
	"github.com/woshahua/corona_server/service"
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

	//service.ImportLocationData()
	// start a new routine
	//go RunCronJob()
	s.ListenAndServe()
}

func RunCronJob() {
	log.Println("starting...")
	service.FetchPatientGlobalData()
	service.FetchPatientDataByCountry()

	fetchJapnesePatientCSV := func() {
		log.Println("Run service.Import")
		filePath := "staticFile/patientByDate.csv"
		url := "https://docs.google.com/spreadsheets/d/1u7aBp8XmZA28Dn6mPo8QueRdVG2a5Bu_gTpAXkAilZw/export?format=csv#gid=0"
		err := service.DownLoadFile(filePath, url)

		filePath = "staticFile/patientByLocation.csv"
		url = "https://docs.google.com/spreadsheets/d/1u7aBp8XmZA28Dn6mPo8QueRdVG2a5Bu_gTpAXkAilZw/export?format=csv&gid=428476519"
		err = service.DownLoadFile(filePath, url)

		filePath = "staticFile/patientDetail.csv"
		url = "https://docs.google.com/spreadsheets/d/10MFfRQTblbOpuvOs_yjIYgntpMGBg592dL8veXoPpp4/export?format=csv&gid=0"
		err = service.DownLoadFile(filePath, url)

		filePath = "staticFile/patientTokyo.csv"
		url = "https://docs.google.com/spreadsheets/d/1u7aBp8XmZA28Dn6mPo8QueRdVG2a5Bu_gTpAXkAilZw/export?format=csv&gid=303868583"
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

	fetchTopicData := func() {
		log.Println("Run service.fetchTopic")
		service.FetchTopicNewsData()
	}

	// fetch newest japanese patient data from:
	// https://toyokeizai.net/sp/visual/tko/covid19/csv/data.csv
	scheduler.Every(1).Hours().Run(fetchJapnesePatientCSV)

	// insert csv data to database evenry 6 hours
	scheduler.Every(1).Hours().Run(importCSVDataToDB)

	// scrap news data and import to database
	scheduler.Every(15).Minutes().Run(scrapNewsData)
	scheduler.Every(15).Minutes().Run(fetchTopicData)
	runtime.Goexit()
}
