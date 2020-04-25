package service

import (
	"context"
	"encoding/csv"
	"github.com/mmcloughlin/geohash"
	"github.com/woshahua/corona_server/library"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/woshahua/corona_server/models"
)

func ImportLocationData() error {
	url := "https://docs.google.com/spreadsheets/d/1u7aBp8XmZA28Dn6mPo8QueRdVG2a5Bu_gTpAXkAilZw/export?format=csv&gid=47176127"
	csvPath := "staticFile/peopleSumByLocation.csv"
	err := DownLoadFile(csvPath, url)
	if err != nil {
		log.Println("failed fetch location csv file", err)
	}

	f, err := os.Open(csvPath)
	defer f.Close()
	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
		line[4] = strings.Replace(line[4], ",", "", -1)
		sum, err := strconv.Atoi(line[4])
		if err != nil {
			sum = 0
		}
		location := models.PeopleLocation{Sum: sum, Location: line[1]}
		err = models.InsertPeopleLocationSum(&location)
		if err != nil {
			return err
		}
	}
	return nil
}

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
	err = InsertPatientDetail()

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

func InsertPatientDetail() error {
	csvPath := "staticFile/patientDetail.csv"
	f, err := os.Open(csvPath)
	defer f.Close()

	if err != nil {
		return err
	}

	lines, err := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i > 0 {
			number, err := strconv.Atoi(line[0])
			if err != nil {
				number = 0
			}
			prefectureCode := line[1]
			onsetDate := line[3]
			confirmDate := line[4]
			consultationPrefecture := line[7]
			residentPrefecture := line[9]
			residentCity := line[10]
			age := line[11]
			gender := line[12]
			isDischarge := line[16]

			description := line[18]
			actionHistory := line[20]
			closeContact := line[21]
			sourceLink := line[25]

			lat := 0.0
			lng := 0.0
			geoHash := ""
			var residentAddress = residentPrefecture + residentCity
			if residentAddress != "" {
				patientDetailLocation := models.FindByAddress(residentAddress)

				if patientDetailLocation != nil {
					lat = patientDetailLocation.Latitude
					lng = patientDetailLocation.Longitude
					geoHash = patientDetailLocation.GeoHash
				} else {
					geoInfo, err := library.GetGeoInfoFromAddress(context.Background(), residentAddress)
					if err != nil {
						return err
					}
					lat = geoInfo.Geometry.Location.Lat
					lng = geoInfo.Geometry.Location.Lng
					geoHash = geohash.Encode(geoInfo.Geometry.Location.Lat, geoInfo.Geometry.Location.Lng)
					models.InsertPatientDetailLocation(&models.PatientDetailLocation{GeoHash:geoHash, Latitude: lat, Longitude: lng})
				}
			}

			patientDetail := models.PatientDetail{
				PatientNumber:          number,
				PatientPrefectureCode:  prefectureCode,
				OnsetDate:              onsetDate,
				ConfirmDate:            confirmDate,
				ConsultationPrefecture: consultationPrefecture,
				ResidentPrefecture:     residentPrefecture,
				ResidentCity:           residentCity,
				CloseContact:           closeContact,
				SourceLink:             sourceLink,
				Age:                    age,
				Gender:                 gender,
				IsDischarge:            isDischarge,
				Description:            description,
				Latitude:               lat,
				Longitude:              lng,
				GeoHash:                geoHash,
				ActionHistory:          actionHistory}

			models.InsertPatientDetail(&patientDetail)
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
