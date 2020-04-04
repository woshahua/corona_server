package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/woshahua/corona_server/interfaces/handler"
	"github.com/gin-gonic/gin"
	"google.golang.org/appengine/log"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	apiv1 := r.Group("api")
	{
		apiv1.GET("/patient/location", handler.GetPatientByLocation)
		apiv1.GET("/news", handler.GetNewsData)
		apiv1.GET("/patient/daily", handler.GetDailyPatient)
		apiv1.GET("/patient/dead", handler.GetDeadPatient)
		apiv1.GET("/patient/current", handler.GetCurrentPatient)
		apiv1.GET("/patient/period", handler.GetPeriodPatient)
		apiv1.GET("/patient/latest", handler.GetLatestPatient)
		apiv1.GET("/patient/updateTime", handler.GetLastestUpdateTime)
	}

	cronRouter := r.Group("/cron")
	{
		cronRouter.Use(isAppEngineCron())
		cronRouter.GET("/fetch_japanese_patient_csv", handler.FetchJapanesePatientCSV)
		cronRouter.GET("/import_csv_data", handler.ImportCSVDataToDB)
		cronRouter.GET("/scrap_news_data", handler.ScrapNewsData)
	}
	return r
}

func isAppEngineCron() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header.Get("X-Appengine-Cron") != "true" {
			log.Warningf(c, "This request is not permitted : %s")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"status":  http.StatusForbidden,
				"message": "Permission denied",
			})
			return
		}
	}
}