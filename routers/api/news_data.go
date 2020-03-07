package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/models"
	"github.com/woshahua/corona_server/pkg/e"
)

func GetNewsData(c *gin.Context) {
	data, err := models.GetNews(5)
	code := e.SUCCESS
	if err != nil {
		fmt.Print("failed fetch news from db", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
