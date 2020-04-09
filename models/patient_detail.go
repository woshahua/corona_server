package models


import "github.com/jinzhu/gorm"

type PatientDetail struct {
	gorm.Model
	OfficialCode           string `json: "official_code"`
	PrefecturePatientCode  string `json: "prefectures_patient_code"`
	OnsetDate              string `json: "onset_date"`
	ConfirmDate            string `json: "confirm_date"`
	PublicDate             string `json: "public_date"`
	ConsultationPrefecture string `json: "consultation_prefecture"`

	ResidentPrefecture     string `json: "resident_prefecture"`
	ResidentCity           string `json: "resident_city"`
	Age                    string `json: "age"`
	Gender                 string `json: "gender"`
	Status                 string `json: "status"`
	isDischarge            int    `json: "is_discharge"`

	Description            string `json: "description"`
	ActionHistory          string `json: "action_history"`
}