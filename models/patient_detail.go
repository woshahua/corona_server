package models

import (
	"corona_server/library"
	"github.com/jinzhu/gorm"
    "github.com/mmcloughlin/geohash"
	"context"
)

type PatientDetail struct {
	gorm.Model
	PatientNumber          int `gorm:"unique;not null"`
	OfficialCode           string `json: "official_code"`
	OnsetDate              string `json: "onset_date"`
	ConfirmDate            string `json: "confirm_date"`
	ConsultationPrefecture string `json: "consultation_prefecture"`

	ResidentPrefecture     string `json: "resident_prefecture"`
	ResidentCity           string `json: "resident_city"`
	Age                    string `json: "age"`
	Gender                 string `json: "gender"`
	IsDischarge            string `json: "is_discharge"`

	Description            string `json: "description"`
	ActionHistory          string `json: "action_history"`
	Latitude               float64 `json: "latitude"`
	Longitude              float64 `json: "longitude"`
	GeoHash                string  `json: "geo_hash"`
}

func InsertPatientDetail(patientDetail *PatientDetail) error {
	var existed PatientDetail
	var notExist = db.Find(&patientDetail, "patient_detail.patient_number = ?", patientDetail.PatientNumber).First(&existed).RecordNotFound()

	if notExist {
		db.NewRecord(patientDetail)
		db.Create(&patientDetail)
		err := db.Save(&patientDetail).Error
		return err
	} else {
		var residentAddress  = existed.ResidentPrefecture + existed.ResidentCity
		if residentAddress != "" {
			geoInfo, err := library.GetGeoInfoFromAddress(context.Background(), residentAddress)
			if err != nil { return  err }
			patientDetail.Latitude = geoInfo.Geometry.Location.Lat
			patientDetail.Longitude = geoInfo.Geometry.Location.Lng
			patientDetail.GeoHash = geohash.Encode(geoInfo.Geometry.Location.Lat, geoInfo.Geometry.Location.Lng)
		}
		updateData := map[string]interface{}{}
		updateData["la"] = title
		updateData["body"] = body
		updateData["category"] = category
		err := db.Update(&patientDetail).Error
		return err
	}
}

func GetPatientDetail() (*[]PatientDetail, error) {
	var patientDetails []PatientDetail
	err := db.Find(&patientDetails).Error

	print("?????????", patientDetails[0].GeoHash)
	return &patientDetails, err
}