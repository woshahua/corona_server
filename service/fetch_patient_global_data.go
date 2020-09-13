package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	cache "github.com/patrickmn/go-cache"
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
	Date               int
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

var Cache = cache.New(time.Duration(-1), time.Duration(-1))

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
	url := "https://covid19-japan-web-api.now.sh/api/v1/total?history=true"
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

	var patients []PatientDataJapan
	err = json.Unmarshal(body, &patients)
	var length = len(patients)
	log.Println(length)

	data := models.PatientByDate{}

	data.Confirmed = patients[length-1].Positive
	data.Critical = patients[length-1].Severe
	data.Symptomless = patients[length-1].Symptom_confirming
	data.Dead = patients[length-1].Death
	data.Tested = patients[length-1].Pcr
	data.Recovered = patients[length-1].Discharge

	data.NewConfirmed = patients[length-1].Positive - patients[length-2].Positive
	data.NewCritical = patients[length-1].Severe - patients[length-2].Severe
	data.NewSymptomless = patients[length-1].Symptom_confirming - patients[length-2].Symptom_confirming
	data.NewDead = patients[length-1].Death - patients[length-2].Death
	data.NewTested = patients[length-1].Pcr - patients[length-2].Pcr
	data.NewRecovered = patients[length-1].Discharge - patients[length-2].Discharge

	Cache.Set("patientJapan", data, cache.DefaultExpiration)
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
