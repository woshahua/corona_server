package api

import (
	"fmt"
	response2 "github.com/woshahua/corona_server/routers/api/response"
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

	data := []*models.PatientDetail{}
	code := e.SUCCESS
	if locationParams.Lat != 0 && locationParams.Lng != 0 {
		geohash := geohash.Encode(locationParams.Lat, locationParams.Lng)
		d, err := models.GetPatientDetailByGeoHash(geohash)
		data = d

		code = e.SUCCESS
		if err != nil {
			code = e.ERROR
			fmt.Println("failed fetch patients details: ", err)
		}
	} else {
		d, err := models.GetAllPatientDetail()
		data = d
		code = e.SUCCESS
		if err != nil {
			code = e.ERROR
			fmt.Println("failed fetch patients details: ", err)
		}
	}

	// group by
	collections := make(map[string][]models.PatientDetail)
	for _, p := range data {
		collections[p.GeoHash] = append(collections[p.GeoHash], *p)
	}

	response := []response2.GetGetPatientDetails{}

	for _, c := range collections {
		response = append(response, response2.GetGetPatientDetails{
			PatientDetails: c,
			Lat: c[0].Latitude,
			Lng: c[0].Longitude,
			Count: len(c)})
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": response,
	})
}
