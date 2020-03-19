package service

import (
	"encoding/csv"
	"io"
	"net/http"
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
			patient := models.PatientByDate{
				Date:      line[0],
				Confirmed: strconv.Atoi(line[1]),
				Recoverd:  strconv.Atoi(line[2]),
				Dead:      strconv.Atoi(line[3]),
				Critical:  strconv.Atoi(line[4]),
				Tested:    strconv.Atoi(line[5])
			}

			models.InsertPatientByDate(&patient)
		}
	}

	csvPath := "staticFile/patientByLocation.csv"
	f, err := os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i > 1 {
			location := models.PatientLocation{ Sum: strconv.Atoi(line[1]), Location: line[2]}
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
