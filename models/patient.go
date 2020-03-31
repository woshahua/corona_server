package models

type PatientLocation struct {
	ID       int    `json: "id"`
	Sum      int    `json: "sum"`
	Location string `json: "patient_location"`
}

type PatientByDate struct {
	ID        int    `json: "id"`
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

type LastUpdateTime struct {
	PatientDataUpdateTime string `json: "patient_data_updated_time"`
}