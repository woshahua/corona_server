package models

import "sort"

type NewPatient struct {
	Sum       int `json "Sum"`
	NewGrowth int `json "NewGrowth`
}

type DeadPatient struct {
	Sum       int
	NewGrowth int
}

type PatientSummary struct {
	PatientSum  int `json: "PatientSum"`
	NewPatient  NewPatient
	DeadPatient DeadPatient
}

func GetJapanesePatientSummary() (*PatientSummary, error) {
	var summary PatientSummary
	var patient []Japanese
	err := db.Find(&patient).Error

	length := len(patient)

	sort.SliceStable(patient, func(i, j int) bool {
		return patient[i].ID < patient[j].ID
	})

	curDate := patient[length-1].Date
	var newPatient []Japanese
	for _, person := range patient {
		if person.Date == curDate {
			newPatient = append(newPatient, person)
		}
	}

	summary.PatientSum = len(patient)
	summary.NewPatient = NewPatient{Sum: len(newPatient), NewGrowth: 0}

	return &summary, err
}
