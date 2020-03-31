package routers

import (
	"github.com/woshahua/corona_server/interfaces/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	//r.Use(cors.Default())
	//gin.SetMode(setting.RunMode)

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
	return r
}
