package mysql

import (
	"github.com/woshahua/corona_server/models"
	"sort"
	"strings"
	"time"
)

func GetLastUpdatedTime() (models.LastUpdateTime, error) {
	conn, err = newConnection()

	var location models.PatientLocation
	var updatedOn models.LastUpdateTime
	err := conn.Find(&location).Last(&location).Error
	updateTime := models.TransferToJSTTime(time.Now())
	if err != nil {
		updatedOn.PatientDataUpdateTime = updateTime.Format("2006/1/2 15:04:05")
		return updatedOn, err
	}

	updateTime = models.TransferToJSTTime(location.UpdatedAt)
	updatedOn.PatientDataUpdateTime = updateTime.Format("2006/1/2 15:04:05")
	return updatedOn, nil
}

func GetLocationPatientData() (*[]models.PatientLocation, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var location []models.PatientLocation
	err := conn.Order("sum desc").Limit(15).Find(&location).Error

	sort.SliceStable(location, func(i, j int) bool {
		return location[i].Sum > location[j].Sum
	})

	return &location, err
}

func GetLatestPatientData() (*models.PatientByDate, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var patient models.PatientByDate

	err := conn.Order("id desc").Limit(1).Find(&patient).Error
	return &patient, err
}

func GetDailyPatientData() (*models.DailyPatient, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var patient []models.PatientByDate
	err := conn.Order("id desc").Limit(3).Find(&patient).Error

	var dailyPatient = models.DailyPatient{}
	dailyPatient.Date = patient[0].Date
	dailyPatient.Current = patient[0].Confirmed - patient[1].Confirmed
	dailyPatient.Diff = dailyPatient.Current - (patient[1].Confirmed - patient[2].Confirmed)

	return &dailyPatient, err
}

func GetDeadPatientData() (*models.DeadPatient, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var patient []models.PatientByDate
	err := conn.Order("id desc").Limit(2).Find(&patient).Error

	var deadPatient = models.DeadPatient{}
	deadPatient.Date = patient[0].Date
	deadPatient.Current = patient[0].Dead
	deadPatient.Diff = patient[0].Dead - patient[1].Dead

	return &deadPatient, err
}

func GetCurrentPatient() (*models.CurrentPatient, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var patient []models.PatientByDate
	err := conn.Order("id desc").Limit(2).Find(&patient).Error

	var currentPatient = models.CurrentPatient{}
	currentPatient.Date = patient[0].Date
	currentPatient.Current = patient[0].Confirmed
	currentPatient.Diff = patient[0].Confirmed - patient[1].Confirmed

	return &currentPatient, err
}

func GetPeriodPatientData() (*[]models.PatientByDate, error) {
	conn, err = newConnection()
	if err != nil {
		return nil, err
	}
	var patient []models.PatientByDate
	err := conn.Order("id desc").Limit(5).Find(&patient).Error

	sort.SliceStable(patient, func(i, j int) bool {
		return patient[i].Date < patient[j].Date
	})

	return &patient, err
}

func InsertPatientByDate(person *models.PatientByDate) error {
	conn, err = newConnection()
	if err != nil { return err }
	var patient models.PatientByDate
	patient.Date = person.Date
	var notExist = conn.Find(&patient, "date = ?", patient.Date).First(&patient).RecordNotFound()

	if notExist {
		date := strings.Split(person.Date, "-")
		person.Date = date[1] + "." + date[2]
		conn.NewRecord(person)
		conn.Create(&person)
		err := conn.Save(&person).Error
		return err
	} else {
		patient = *person
		date := strings.Split(person.Date, "-")
		patient.Date = date[1] + "." + date[2]
		err := conn.Save(&patient).Error
		return err
	}
	return nil
}

func UpdatePatientByLocation(location *models.PatientLocation) error {
	conn, err = newConnection()
	if err != nil { return err }
	var locationData models.PatientLocation
	locationData.Location = location.Location
	var notExist = conn.Find(&locationData, "location = ?", locationData.Location).First(&locationData).RecordNotFound()
	if notExist {
		conn.NewRecord(location)
		conn.Create(&location)
		err := conn.Save(&location).Error
		return err
	}
	return nil
}

func UpdateDataUpdatedTime(updateTime *models.LastUpdateTime) error {
	conn, err = newConnection()
	if err != nil { return err }

	var lastupdateTime models.LastUpdateTime
	err := conn.Find(&lastupdateTime).First(&lastupdateTime).Error
	if err != nil {
		return err
	}

	lastupdateTime.PatientDataUpdateTime = updateTime.PatientDataUpdateTime
	conn.Save(&lastupdateTime)
	return nil
}
