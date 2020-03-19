package service

import (
	"encoding/csv"
	"io"
	"net/http"
	"os"

	"github.com/woshahua/corona_server/models"
)

func Import() error {
	csvPath := "staticFile/data.csv"
	f, err := os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i > 0 {
			patient := models.Japanese{Patient: models.Patient{ID: i, Date: line[], Age: line[3], PatientLocation: line[5]}}
			models.InsertJapanesePatient(&patient)
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
