package models

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/mmcloughlin/geohash"
	"github.com/woshahua/corona_server/library"
)

type PatientDetail struct {
	gorm.Model
	PatientNumber          int    `gorm:"unique;not null"`
	PatientPrefectureCode  string    `gorm:"unique;not null"`
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
		var residentAddress = patientDetail.ResidentPrefecture + patientDetail.ResidentCity
		if residentAddress != "" {
			geoInfo, err := library.GetGeoInfoFromAddress(context.Background(), residentAddress)
			if err != nil {
				return err
			}
			patientDetail.Latitude = geoInfo.Geometry.Location.Lat
			patientDetail.Longitude = geoInfo.Geometry.Location.Lng
			patientDetail.GeoHash = geohash.Encode(geoInfo.Geometry.Location.Lat, geoInfo.Geometry.Location.Lng)
		}
		err := db.Create(&patientDetail).Error
		return err
	} else {
		err := db.Find(&existed, "patient_detail.patient_prefecture_code = ?", patientDetail.PatientPrefectureCode).Update(&patientDetail).Error
		return err
	}
}

func GetAllPatientDetail() (*[]PatientDetail, error) {
	var patientDetails []PatientDetail
	err := db.Find(&patientDetails).Error
	return &patientDetails, err
}

func GetPatientDetailByGeoHash(geoHash string) (*[]PatientDetail, error) {
	var patientDetails []PatientDetail
	searchGeo := geoHash[:3]
	err := db.Where("geo_hash LIKE ?", searchGeo+"%").Find(&patientDetails).Error

	return &patientDetails, err
}
