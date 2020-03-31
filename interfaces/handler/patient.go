package handler

import (
	"corona_server/infrastructure/mysql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
)

func GetLastestUpdateTime(c *gin.Context) {
	data, err := mysql.GetLastUpdatedTime()
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

func GetLatestPatient(c *gin.Context) {
	data, err := mysql.GetLatestPatientData()
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

func GetPeriodPatient(c *gin.Context) {
	data, err := mysql.GetPeriodPatientData()
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
	data, err := mysql.GetDailyPatientData()
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
	data, err := mysql.GetCurrentPatient()
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
	data, err := mysql.GetLocationPatientData()
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

