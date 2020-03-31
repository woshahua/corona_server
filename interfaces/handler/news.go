package handler

import (
	"github.com/woshahua/corona_server/infrastructure/mysql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/pkg/e"
	"net/http"
	"strconv"
)

func GetNewsData(c *gin.Context) {
	code := e.SUCCESS

	num, err := strconv.Atoi(c.Query("number"))
	if err != nil {
		code = e.ERROR
		fmt.Print("convert string to int failed", err)
	}
	data, err := mysql.GetNews(num)
	if err != nil {
		code = e.ERROR
		fmt.Print("failed fetch news from db", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
