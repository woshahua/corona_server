package models

import (
	"log"
	"sort"
)

type Patient struct {
	ID              int    `gorm: "primary_key", json: "id"`
	Date            string `json: "date", gorm: "date"`
	Age             string `json: "age", gorm: "age"`
	PatientLocation string `json: "patient_location"`
}

type Japanese struct {
	Patient
}

type PatientLocation struct {
	Location   string
	PatientSum int
}

func GetJapanesePatient() (*[]Japanese, error) {
	var patient []Japanese
	err := db.Find(&patient).Error

	// sort by patients by date
	sort.SliceStable(patient, func(i, j int) bool {
		return patient[i].ID < patient[j].ID
	})

	return &patient, err
}

func GetJapaneseDailyNewPatient() (*[]Japanese, error) {
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

	return &newPatient, err
}

func GetJapanesePatientByLoaction() (*[]PatientLocation, error) {
	var patient []Japanese
	var patientLocationList []PatientLocation
	err := db.Find(&patient).Error

	m := make(map[string]int)
	for _, person := range patient {
		if val, ok := m[person.PatientLocation]; ok {
			m[person.PatientLocation] = val + 1
		} else {
			m[person.PatientLocation] = 1
		}
	}

	for key, value := range m {
		patientLoc := PatientLocation{Location: key, PatientSum: value}
		patientLocationList = append(patientLocationList, patientLoc)
	}

	sort.SliceStable(patientLocationList, func(i, j int) bool {
		return patientLocationList[i].PatientSum > patientLocationList[j].PatientSum
	})

	return &patientLocationList, err
}

func InsertJapanesePatient(person *Japanese) {
	db.NewRecord(person)
	db.Create(&person)
	err := db.Save(&person).Error
	if err != nil {
		log.Fatalf("failed insert patient: %v", err)
	}
}
