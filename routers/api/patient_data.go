package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
)

func GetJapanesePatient(c *gin.Context) {
	data, err := models.GetJapanesePatient()
	code := e.SUCCESS
	if err != nil {
		fmt.Print("failed ", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":           code,
		"msg":            e.GetMsg(code),
		"data":           data,
		"patient_number": len(*data),
	})
}

func GetJapaneseDailyNewPatient(c *gin.Context) {
	data, err := models.GetJapaneseDailyNewPatient()
	code := e.SUCCESS
	if err != nil {
		code = e.ERROR
		fmt.Print("failed", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":           code,
		"msg":            e.GetMsg(code),
		"data":           data,
		"patient_number": len(*data),
	})
}

func GetJapanesePatientByLocation(c *gin.Context) {
	data, err := models.GetJapanesePatientByLoaction()
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
