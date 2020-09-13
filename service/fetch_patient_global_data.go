package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/woshahua/corona_server/models"
)

type PatientGlobalData struct {
	Data PatientData `json: "data"`
	Dt   string      `json: "dt"`
}

type PatientData struct {
	Confirmed int `json: "confirmed"`
	Deaths    int `json: "deaths"`
	Recovered int `json: "recovered"`
	Active    int `json: "active"`
}

type PatientDataJapan struct {
	Date               string
	Pcr                int
	Hospitalize        int
	Positive           int
	Severe             int
	Discharge          int
	Death              int
	Symptom_confirming int
}

type PatientGlobalDataByCountry struct {
	Data []PatientDataByCountry `json: "data"`
}
type PatientDataByCountry struct {
	Location  string `json: "location"`
	Confirmed int    `json: "confirmed"`
	Deaths    int    `json: "deaths"`
	Recovered int    `json: "recovered"`
	Active    int    `json: "active"`
}

func FetchPatientGlobalData() {
	url := "https://covid2019-api.herokuapp.com/v2/total"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var global PatientGlobalData
	err = json.Unmarshal(body, &global)

	if err != nil {
		log.Fatal(err)
	}

	dateString := strings.Split(global.Dt, " ")[0]

	var globalDataModel models.PatientGlobal
	globalDataModel.Active = global.Data.Active
	globalDataModel.Recovered = global.Data.Recovered
	globalDataModel.Deaths = global.Data.Deaths
	globalDataModel.Confirmed = global.Data.Confirmed
	globalDataModel.Date = dateString

	models.InsertPatientGlobalData(&globalDataModel)
}

func FetchPatientJapan() {
	url := "https://covid19-japan-web-api.now.sh/api/v1/total"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	var patient PatientDataJapan
	err = json.Unmarshal(body, &patient)
	log.Println(patient)
}

func FetchPatientDataByCountry() {
	url := "https://covid2019-api.herokuapp.com/v2/current"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var global PatientGlobalDataByCountry
	err = json.Unmarshal(body, &global)

	if err != nil {
		log.Fatal(err)
	}

	for _, data := range global.Data {
		var countryData models.PatientGlobalByCountry
		countryData.Active = data.Active
		countryData.Recovered = data.Recovered
		countryData.Deaths = data.Deaths
		countryData.Confirmed = data.Confirmed
		countryData.Location = data.Location
		countryData.DeathRate = float32(countryData.Deaths) / float32(countryData.Confirmed)
		models.InsertPatientGlobalCountryData(&countryData)
	}
}
