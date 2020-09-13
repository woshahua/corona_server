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
	// start a new routine
	go RunCronJob()
	s.ListenAndServe()
}

func RunCronJob() {
	log.Println("starting...")

	fetchGlobalData := func() {
		service.FetchPatientGlobalData()
		service.FetchPatientDataByCountry()
	}

	fetchJapanData := func() {
		service.FetchPatientJapan()
	}

	// scrap news data and import to database
	scheduler.Every(12).Hours().Run(fetchGlobalData)
	scheduler.Every(12).Hours().Run(fetchJapanData)
	runtime.Goexit()
}
