package response

import "github.com/woshahua/corona_server/models"

type GetGetPatientDetails struct {
	PatientDetails []models.PatientDetail `json:"patientDetails"`
	Lat     float64     `json:"lat"`
	Lng     float64     `json:"lng"`
	Count   int         `json:"count"`
}
