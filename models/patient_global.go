package models

import (
	"github.com/jinzhu/gorm"
)

type PatientGlobal struct {
	gorm.Model
	Date      string
	Confirmed int `json: "confirmed"`
	Deaths    int `json: "deaths"`
	Recovered int `json: "recovered"`
	Active    int `json: "active"`
}

type PatientGlobalByCountry struct {
	gorm.Model
	Location  string
	Confirmed int     `json: "confirmed"`
	Deaths    int     `json: "deaths"`
	Recovered int     `json: "recovered"`
	Active    int     `json: "active"`
	DeathRate float32 `json: "death_rate"`
}

type PatientGrowth struct {
	NewConfirmed int `json: "new_confirmed"`
	NewDeaths    int `json: "new_deaths"`
	NewRecovered int `json: "new_recovered`
}

func GetCurrentPatientGlobalData() (*PatientGlobal, error) {
	var global PatientGlobal
	err := db.Order("id desc").Limit(1).Find(&global).Error
	return &global, err
}

func GetPatientGlobalDataNew() (*PatientGrowth, error) {
	var global []PatientGlobal
	err := db.Order("id desc").Limit(2).Find(&global).Error

	var growth PatientGrowth
	growth.NewConfirmed = global[0].Confirmed - global[1].Confirmed
	growth.NewDeaths = global[0].Deaths - global[1].Deaths
	growth.NewRecovered = global[0].Recovered - global[1].Recovered
	return &growth, err
}

func GetPatientGlobalByCountry() ([]PatientGlobalByCountry, error) {
	var countryData []PatientGlobalByCountry
	err := db.Order("id asc").Find(&countryData).Error
	return countryData, err
}

func InsertPatientGlobalData(global *PatientGlobal) error {
	globalData := PatientGlobal{}
	globalData.Date = global.Date
	var notExist = db.Find(&globalData, "patient_global.date = ?", globalData.Date).First(&globalData).RecordNotFound()

	if notExist {
		db.NewRecord(global)
		db.Create(&global)
		err := db.Save(&global).Error
		return err
	} else {
		globalData.Confirmed = global.Confirmed
		globalData.Recovered = global.Recovered
		globalData.Deaths = global.Deaths
		globalData.Active = global.Active
		err := db.Save(&globalData).Error
		return err
	}
}

func InsertPatientGlobalCountryData(globalCountry *PatientGlobalByCountry) error {
	countryData := PatientGlobalByCountry{}
	countryData.Location = globalCountry.Location
	var notExist = db.Find(&countryData, "patient_global_by_country.location = ?", countryData.Location).First(&countryData).RecordNotFound()

	if notExist {
		db.NewRecord(globalCountry)
		db.Create(&globalCountry)
		err := db.Save(&globalCountry).Error
		return err
	} else {
		countryData.Confirmed = globalCountry.Confirmed
		countryData.Active = globalCountry.Active
		countryData.Deaths = globalCountry.Deaths
		countryData.Recovered = globalCountry.Recovered
		countryData.DeathRate = globalCountry.DeathRate
		err := db.Save(&countryData).Error
		return err
	}
}
