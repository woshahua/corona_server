package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/woshahua/corona_server/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json: "id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json: "modified_on"`
}

func init() {
	var (
		err                                        error
		dbType, dbName, user, password, host, port string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "fail to get section 'database': %v", err)
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	port = sec.Key("PORT").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName))

	if err != nil {
		log.Println(err)
		panic("failed to connect database")
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.AutoMigrate(&News{}, &PatientByDate{}, &PatientLocation{}, &PatientGlobal{}, &PatientGlobalByCountry{}, &PeopleLocation{}, &PatientDetail{}, &PatientTokyo{})
}

func CloseDB() {
	defer db.Close()
}
