package models

import (
	"sort"
	"time"

	"github.com/jinzhu/gorm"
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

func GetLastUpdatedTime() (string, error) {
	var location PatientLocation
	err := db.Find(&location).Last(&location).Error
	if err != nil {
		return time.Now().Format("2006/1/2 15:04:05"), err
	}

	return location.UpdatedAt.Format("2006/1/2 15:04:05"), nil
}

func GetLocationPatientData() (*[]PatientLocation, error) {
	var location []PatientLocation
	err := db.Order("sum desc").Limit(15).Find(&location).Error

	sort.SliceStable(location, func(i, j int) bool {
		return location[i].Sum > location[j].Sum
	})

	return &location, err
}

func GetLatestPatientData() (*PatientByDate, error) {
	var patient PatientByDate

	err := db.Order("id desc").Limit(1).Find(&patient).Error
	return &patient, err
}

func GetDailyPatientData() (*DailyPatient, error) {
	var patient []PatientByDate
	err := db.Order("id desc").Limit(3).Find(&patient).Error

	var dailyPatient = DailyPatient{}
	dailyPatient.Date = patient[0].Date
	dailyPatient.Current = patient[0].Confirmed - patient[1].Confirmed
	dailyPatient.Diff = dailyPatient.Current - (patient[1].Confirmed - patient[2].Confirmed)

	return &dailyPatient, err
}

func GetDeadPatientData() (*DeadPatient, error) {
	var patient []PatientByDate
	err := db.Order("id desc").Limit(2).Find(&patient).Error

	var deadPatient = DeadPatient{}
	deadPatient.Date = patient[0].Date
	deadPatient.Current = patient[0].Dead
	deadPatient.Diff = patient[0].Dead - patient[1].Dead

	return &deadPatient, err
}

func GetCurrentPatient() (*CurrentPatient, error) {
	var patient []PatientByDate
	err := db.Order("id desc").Limit(2).Find(&patient).Error

	var currentPatient = CurrentPatient{}
	currentPatient.Date = patient[0].Date
	currentPatient.Current = patient[0].Confirmed
	currentPatient.Diff = patient[0].Confirmed - patient[1].Confirmed

	return &currentPatient, err
}

func GetPeriodPatientData() (*[]PatientByDate, error) {
	var patient []PatientByDate
	err := db.Order("id desc").Limit(5).Find(&patient).Error

	sort.SliceStable(patient, func(i, j int) bool {
		return patient[i].Date < patient[j].Date
	})

	return &patient, err
}

func InsertPatientByDate(person *PatientByDate) error {
	var patient PatientByDate
	patient.Date = person.Date
	var notExist = db.Find(&patient, "patient_by_date.date = ?", patient.Date).First(&patient).RecordNotFound()

	if notExist {
		db.NewRecord(person)
		db.Create(&person)
		err := db.Save(&person).Error
		return err
	} else {
		patient.Confirmed = person.Confirmed
		patient.Recovered = person.Recovered
		patient.Dead = person.Dead
		patient.Critical = person.Critical
		patient.Tested = person.Tested
		err := db.Save(&patient).Error
		return err
	}
}

func UpdatePatientByLocation(location *PatientLocation) error {
	var locationData PatientLocation
	locationData.Location = location.Location
	var notExist = db.Find(&locationData, "patient_location.location = ?", locationData.Location).First(&locationData).RecordNotFound()
	if notExist {
		db.NewRecord(location)
		db.Create(&location)
		err := db.Save(&location).Error
		return err
	} else {
		locationData.Location = location.Location
		locationData.Sum = location.Sum
		db.Save(&locationData)
		return nil
	}
}
