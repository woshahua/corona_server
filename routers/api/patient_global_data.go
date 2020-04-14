package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
)

func GetGlobalDataGrowth(c *gin.Context) {
	data, err := models.GetPatientGlobalDataNew()
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

func GetCurrentGlobalData(c *gin.Context) {
	data, err := models.GetCurrentPatientGlobalData()
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

func GetGlobalDataByCountry(c *gin.Context) {
	data, err := models.GetPatientGlobalByCountry()
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
