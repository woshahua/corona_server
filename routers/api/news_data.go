package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"corona_server/models"
	"corona_server/pkg/e"
	"corona_server/service"
)

func GetNewsData(c *gin.Context) {
	code := e.SUCCESS

	data, found := service.Cache.Get("news")
	var newsData []models.News
	if found {
		newsData = data.([]models.News)
	} else {
		fmt.Println("data not found")
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": newsData,
	})
}
