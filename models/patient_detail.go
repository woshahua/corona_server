package models

import (
	"github.com/jinzhu/gorm"
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
}

func InsertPatientDetail(patientDetail *PatientDetail) error {
	var notExist = db.Find(&patientDetail, "patient_detail.patient_number = ?", patientDetail.PatientNumber).First(&patientDetail).RecordNotFound()

	if notExist {
		db.NewRecord(patientDetail)
		db.Create(&patientDetail)
		err := db.Save(&patientDetail).Error
		return err
	} else {
		err := db.Update(&patientDetail).Error
		return err
	}
}

func GetPatientDetail() (*[]PatientDetail, error) {
	var patientDetails []PatientDetail
	err := db.Find(&patientDetails).Error

	return &patientDetails, err
}