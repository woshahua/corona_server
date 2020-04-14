package api

import (
	"fmt"
	"net/http"

	"github.com/woshahua/corona_server/pkg/e"

	"github.com/gin-gonic/gin"
	"github.com/mmcloughlin/geohash"
	"github.com/woshahua/corona_server/models"
)

type LocationParams struct {
	Lat float64 `form:"lat"`
	Lng float64 `form:"lng"`
}

func GetPatientDetails(c *gin.Context) {
	var locationParams LocationParams
	c.ShouldBindQuery(&locationParams)

	if locationParams.Lat != 0 && locationParams.Lng != 0 {
		geohash := geohash.Encode(locationParams.Lat, locationParams.Lng)
		data, err := models.GetPatientDetailByGeoHash(geohash)

		code := e.SUCCESS
		if err != nil {
			code = e.ERROR
			fmt.Println("failed fetch patients details: ", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	} else {
		data, err := models.GetAllPatientDetail()
		code := e.SUCCESS
		if err != nil {
			code = e.ERROR
			fmt.Println("failed fetch patients details: ", err)
		}

		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
}
