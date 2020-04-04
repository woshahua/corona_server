package main

import (
	"github.com/woshahua/corona_server/environment"
	"github.com/woshahua/corona_server/infrastructure/mysql"
	"log"
	"github.com/woshahua/corona_server/routers"
	"os"
)

func main() {
	if environment.IsProd() {
		if err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./config/pro_credential.json"); err != nil {
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