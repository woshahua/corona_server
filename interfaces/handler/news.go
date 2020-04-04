package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
	"github.com/woshahua/corona_server/service"
	"net/http"
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
