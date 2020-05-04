package models

import (
	"github.com/jinzhu/gorm"
)

type PatientDetail struct {
	gorm.Model
	PatientNumber          int    `json: "patient_number"`
	PatientPrefectureCode  string `gorm:"index:patient_prefecture_code;unique;not null"`
	OnsetDate              string `json: "onset_date"`
	ConfirmDate            string `json: "confirm_date"`
	ConsultationPrefecture string `json: "consultation_prefecture"`

	ResidentPrefecture string `json: "resident_prefecture"`
	ResidentCity       string `json: "resident_city"`
	Age                string `json: "age"`
	Gender             string `json: "gender"`
	IsDischarge        string `json: "is_discharge"`
	CloseContact       string `json: "close_contact"`
	SourceLink         string `json: "source_link"`

	Description   string  `json: "description"`
	ActionHistory string  `json: "action_history"`
	Latitude      float64 `json: "latitude"`
	Longitude     float64 `json: "longitude"`
	GeoHash       string  `json: "geo_hash"`
}

func InsertPatientDetail(patientDetail *PatientDetail) error {
	var existed PatientDetail
	var notExist = db.Find(&existed, "patient_detail.patient_prefecture_code = ?", patientDetail.PatientPrefectureCode).First(&existed).RecordNotFound()

	if notExist {
		err := db.Create(&patientDetail).Error
		return err
	} else {
		err := db.Find(&existed, "patient_detail.patient_prefecture_code = ?", patientDetail.PatientPrefectureCode).Update(&patientDetail).Error
		return err
	}
}

func GetAllPatientDetail() ([]*PatientDetail, error) {
	results := []*PatientDetail{}
	err := db.Find(&results).Error
	return results, err
}

func GetPatientDetailByGeoHash(geoHash string, matchingNum int) ([]*PatientDetail, error) {
	results := []*PatientDetail{}
	searchGeo := geoHash[:matchingNum]
	err := db.Where("geo_hash LIKE ?", searchGeo+"%").Find(&results).Error

	return results, err
}
