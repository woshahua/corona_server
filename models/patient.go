package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type PatientLocation struct {
	gorm.Model
	Sum      int    `json: "sum"`
	Location string `json: "patient_location"`
}

type PatientByDate struct {
	gorm.Model
	Date      string `json: "date"`
	Confirmed int    `json: "confirmed`
	Recovered int    `json: "recovered`
	Dead      int    `json: "dead`
	Critical  int    `json: "critical`
	Tested    int    `json: "tested"`
}

type DailyPatient struct {
	Date    string
	Current int
	Diff    int
}

type CurrentPatient struct {
	Date    string
	Current int
	Diff    int
}

type DeadPatient struct {
	Date    string
	Current int
	Diff    int
}

type LastUpdateTime struct {
	PatientDataUpdateTime string `json: "patient_data_updated_time"`
}

func TransferToJSTTime(utcTime time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	jstTime := utcTime.UTC().In(jst)
	return jstTime
}