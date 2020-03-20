package models

import (
	"sort"
	"strings"
)

type PatientLocation struct {
	ID       int    `gorm: "primary_key", json: "id"`
	Sum      int    `json: "sum"`
	Location string `json: "patient_location"`
}

type PatientByDate struct {
	ID        int    `gorm: "primary_key", json: "id"`
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

func GetLocationPatientData() (*[]PatientLocation, error) {
	var location []PatientLocation
	err := db.Find(&location).Error

	sort.SliceStable(location, func(i, j int) bool {
		return location[i].Sum < location[j].Sum
	})

	return &location, err
}

func GetDailyPatientData() (*DailyPatient, error) {
	var patient []PatientByDate
	err := db.Order("date desc").Limit(3).Find(&patient).Error

	var dailyPatient = DailyPatient{}
	dailyPatient.Date = patient[0].Date
	dailyPatient.Current = patient[0].Confirmed - patient[1].Confirmed
	dailyPatient.Diff = dailyPatient.Current - (patient[1].Confirmed - patient[2].Confirmed)

	return &dailyPatient, err
}

func GetDeadPatientData() (*DeadPatient, error) {
	var patient []PatientByDate
	err := db.Order("date desc").Limit(2).Find(&patient).Error

	var deadPatient = DeadPatient{}
	deadPatient.Date = patient[0].Date
	deadPatient.Current = patient[0].Dead
	deadPatient.Diff = patient[0].Dead - patient[1].Dead

	return &deadPatient, err
}

func GetCurrentPatient() (*CurrentPatient, error) {
	var patient []PatientByDate
	err := db.Order("date desc").Limit(2).Find(&patient).Error

	var currentPatient = CurrentPatient{}
	currentPatient.Date = patient[0].Date
	currentPatient.Current = patient[0].Confirmed
	currentPatient.Diff = patient[0].Confirmed - patient[1].Confirmed

	return &currentPatient, err
}

func GetPeriodPatientData() (*[]PatientByDate, error) {
	var patient []PatientByDate
	err := db.Order("date desc").Limit(5).Find(&patient).Error

	sort.SliceStable(patient, func(i, j int) bool {
		return patient[i].Date < patient[j].Date
	})

	return &patient, err
}

func GetJapanesePatientByLoaction() (*[]PatientLocation, error) {
	var locationList []PatientLocation
	err := db.Order("sum desc").Limit(5).Find(&locationList).Error

	return &locationList, err
}

func InsertPatientByDate(person *PatientByDate) error {
	var patient PatientByDate
	patient.Date = person.Date
	var notExist = db.Find(&patient, "patient_by_date.date = ?", patient.Date).First(&patient).RecordNotFound()

	if notExist {
		date := strings.Split(person.Date, "-")
		person.Date = date[1] + "." + date[2]
		db.NewRecord(person)
		db.Create(&person)
		err := db.Save(&person).Error
		return err
	} else {
		patient = *person
		date := strings.Split(person.Date, "-")
		patient.Date = date[1] + "." + date[2]
		err := db.Save(&patient).Error
		return err
	}
	return nil
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
	}
	return nil

}
