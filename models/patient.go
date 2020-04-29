package models

import (
	"sort"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type PeopleLocation struct {
	gorm.Model
	Sum      int
	Location string
}
type PatientLocation struct {
	gorm.Model
	Sum           int     `json: "sum"`
	Location      string  `json: "patient_location"`
	InfectionRate float32 `json: "infection_rate"`
}

type PatientByDate struct {
	gorm.Model
	Date           string `json: "date"`
	Confirmed      int    `json: "confirmed`
	NewConfirmed   int    `json: "new_confirmed`
	Recovered      int    `json: "recovered`
	NewRecovered   int    `json: "new_recovered`
	Dead           int    `json: "dead`
	NewDead        int    `json: "new_dead`
	Critical       int    `json: "critical`
	NewCritical    int    `json: "new_critical`
	Tested         int    `json: "tested"`
	NewTested      int    `json: "new_tested"`
	Symptomless    int    `json: "symptomless"`
	NewSymptomless int    `json: "new_symptomless"`
}

type PatientTokyo struct {
	gorm.Model
	Sum        int    `json "sum"`
	Location   string `json: "location"`
	UpdateTime string `json: "update_time"`
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

func GetLastUpdatedTime() (LastUpdateTime, error) {
	var location PatientLocation
	var updatedOn LastUpdateTime
	err := db.Find(&location).Last(&location).Error
	updateTime := TransferToJSTTime(time.Now())
	if err != nil {
		updatedOn.PatientDataUpdateTime = updateTime.Format("2006/1/2 15:04:05")
		return updatedOn, err
	}

	updateTime = TransferToJSTTime(location.UpdatedAt)
	updatedOn.PatientDataUpdateTime = updateTime.Format("2006/1/2 15:04:05")
	return updatedOn, nil
}

func GetLocationPatientData() (*[]PatientLocation, error) {
	var location []PatientLocation
	err := db.Order("sum desc").Find(&location).Error

	sort.SliceStable(location, func(i, j int) bool {
		return location[i].Sum > location[j].Sum
	})

	return &location, err
}

func GetLatestPatientData() (*PatientByDate, error) {
	var patient []PatientByDate

	err := db.Order("id desc").Limit(2).Find(&patient).Error

	var latestPatient PatientByDate
	latestPatient.Confirmed = patient[0].Confirmed
	latestPatient.Critical = patient[0].Critical
	latestPatient.Symptomless = patient[0].Symptomless
	latestPatient.Tested = patient[0].Tested
	latestPatient.Dead = patient[0].Dead
	latestPatient.Recovered = patient[0].Recovered
	latestPatient.NewConfirmed = patient[0].Confirmed - patient[1].Confirmed
	latestPatient.NewCritical = patient[0].Critical - patient[1].Critical
	latestPatient.NewDead = patient[0].Dead - patient[1].Dead
	latestPatient.NewRecovered = patient[0].Recovered - patient[1].Recovered
	latestPatient.NewSymptomless = patient[0].Symptomless - patient[1].Symptomless
	latestPatient.NewTested = patient[0].Tested - patient[1].Tested
	return &latestPatient, err
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
	err := db.Order("id desc").Find(&patient).Error

	var patientFiltered []PatientByDate
	for _, data := range patient {
		if data.Date > "02-26" {
			patientFiltered = append(patientFiltered, data)
		}
	}

	sort.SliceStable(patientFiltered, func(i, j int) bool {
		return patientFiltered[i].Date < patientFiltered[j].Date
	})

	return &patientFiltered, err
}

func GetPatientTokyoData() ([]PatientTokyo, error) {
	var tokyoLocationList []PatientTokyo
	err := db.Find(&tokyoLocationList).Error

	return tokyoLocationList, err
}

func InsertPatientTokyo(tokyoLocation *PatientTokyo) error {
	var location PatientTokyo
	location.Location = tokyoLocation.Location

	notExist := db.Find(&location, "patient_tokyo.location = ?", location.Location).First(&location).RecordNotFound()
	if notExist {
		db.NewRecord(tokyoLocation)
		db.Create(&tokyoLocation)
		err := db.Save(&tokyoLocation).Error
		return err
	} else {
		location.Sum = tokyoLocation.Sum
		location.UpdateTime = tokyoLocation.UpdateTime
		err := db.Save(&location).Error
		return err
	}
}

func InsertPatientByDate(person *PatientByDate) error {
	var patient PatientByDate
	date := strings.Split(person.Date, "-")
	person.Date = date[1] + "/" + date[2]

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
		patient.Symptomless = person.Symptomless
		err := db.Save(&patient).Error
		return err
	}
}

func UpdatePatientByLocation(location *PatientLocation) error {
	var locationData PatientLocation
	locationData.Location = location.Location

	peopleSum := PeopleLocation{}
	peopleSum.Location = location.Location
	db.Find(&peopleSum, "people_location.location = ?", peopleSum.Location).First(&peopleSum)

	var notExist = db.Find(&locationData, "patient_location.location = ?", locationData.Location).First(&locationData).RecordNotFound()
	if notExist {
		if location.Sum == 0 {
			location.InfectionRate = float32(0.0)
		} else {
			location.InfectionRate = float32(location.Sum) / float32(peopleSum.Sum)
		}
		db.NewRecord(location)
		db.Create(&location)
		err := db.Save(&location).Error
		return err
	} else {
		if location.Sum == 0 {
			locationData.InfectionRate = float32(0.0)
		} else {
			locationData.InfectionRate = float32(location.Sum) / float32(peopleSum.Sum)
		}
		locationData.Location = location.Location
		locationData.Sum = location.Sum
		db.Save(&locationData)
		return nil
	}
}

func InsertPeopleLocationSum(location *PeopleLocation) error {
	var locationData PeopleLocation
	locationData.Location = location.Location
	var notExist = db.Find(&locationData, "people_location.location = ?", locationData.Location).First(&locationData).RecordNotFound()
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

func TransferToJSTTime(utcTime time.Time) time.Time {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	jstTime := utcTime.UTC().In(jst)
	return jstTime
}
