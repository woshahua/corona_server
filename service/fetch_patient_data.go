package service

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"
	"strconv"

	"corona_server/models"
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

			patient := models.PatientByDate{
				Date:      line[0],
				Confirmed: confirmed,
				Recovered: recovered,
				Dead:      dead,
				Critical:  critical,
				Tested:    tested}

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

	return nil
}

func DownLoadFile(filePath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
