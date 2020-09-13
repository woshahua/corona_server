package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
	"github.com/woshahua/corona_server/service"
)

func GetPatientTokyoData(c *gin.Context) {
	data, err := models.GetPatientTokyoData()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Print("failed ", err)
	}

	updateTime := data[0].UpdateTime

	c.JSON(http.StatusOK, gin.H{
		"code":       code,
		"msg":        e.GetMsg(code),
		"data":       data,
		"updateTime": updateTime,
	})
}

func GetLatestPatient(c *gin.Context) {
	code := e.SUCCESS

	data, found := service.Cache.Get("patientJapan")
	var patient models.PatientByDate
	if found {
		patient = data.(models.PatientByDate)
	} else {
		fmt.Println("data not found")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": patient,
	})
}

func GetPeriodPatient(c *gin.Context) {
	numberQuery := c.Query("number")
	number, err := strconv.Atoi(numberQuery)
	if err != nil {
		number = 0
	}
	data, err := models.GetPeriodPatientData(number)
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Print("failed ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetDailyPatient(c *gin.Context) {
	code := e.SUCCESS

	data, found := service.Cache.Get("patientJapan")
	var patient models.PatientByDate
	if found {
		patient = data.(models.PatientByDate)
	} else {
		fmt.Println("data not found")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": patient,
	})
}

func GetDeadPatient(c *gin.Context) {
	data, err := models.GetDeadPatientData()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Print("failed", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetCurrentPatient(c *gin.Context) {
	data, err := models.GetCurrentPatient()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Print("failed", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func GetPatientByLocation(c *gin.Context) {
	data, err := models.GetLocationPatientData()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Println("failed fetch patients by location: ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
