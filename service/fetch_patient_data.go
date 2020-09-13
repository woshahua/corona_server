package service

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/woshahua/corona_server/models"
)

func Import() error {
	csvPath := "staticFile/patientByDate.csv"

	f, err := os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i > 0 {
			confirmed, err := strconv.Atoi(line[1])
			if err != nil {
				confirmed = 0
			}

			recovered, err := strconv.Atoi(line[2])
			if err != nil {
				recovered = 0
			}

			dead, err := strconv.Atoi(line[3])
			if err != nil {
				dead = 0
			}

			critical, err := strconv.Atoi(line[4])
			if err != nil {
				critical = 0
			}

			tested, err := strconv.Atoi(line[5])
			if err != nil {
				tested = 0
			}

			symptomless, err := strconv.Atoi(line[6])
			if err != nil {
				symptomless = 0
			}

			patient := models.PatientByDate{
				Date:        line[0],
				Confirmed:   confirmed,
				Recovered:   recovered,
				Dead:        dead,
				Critical:    critical,
				Tested:      tested,
				Symptomless: symptomless,
			}

			models.InsertPatientByDate(&patient)
		}
	}

	csvPath = "staticFile/patientByLocation.csv"
	f, err = os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	lines, err = csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i > 1 {
			sum, err := strconv.Atoi(line[2])
			if err != nil {
				sum = 0
			}

			location := models.PatientLocation{Sum: sum, Location: line[1]}
			models.UpdatePatientByLocation(&location)
		}
	}

	// insert patient gis location data
	// err = InsertPatientDetail()

	if err != nil {
		return err
	}

	csvPath = "staticFile/patientTokyo.csv"
	f, err = os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	var updateTime string
	lines, err = csv.NewReader(f).ReadAll()
	for _, line := range lines {
		if line[0] != "updateTime" {
			sum, err := strconv.Atoi(line[1])
			if err != nil {
				sum = 0
			}

			location := models.PatientTokyo{Sum: sum, Location: line[0], UpdateTime: updateTime}
			models.InsertPatientTokyo(&location)
		} else {
			updateTime = line[1]
		}
	}

	return nil
}
