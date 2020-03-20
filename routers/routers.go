package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/pkg/setting"
	"github.com/woshahua/corona_server/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("api")
	{
		apiv1.GET("/patient/location", api.GetPatientByLocation)
		apiv1.GET("/news", api.GetNewsData)
		apiv1.GET("/patient/daily", api.GetDailyPatient)
		apiv1.GET("/patient/dead", api.GetDeadPatient)
		apiv1.GET("/patient/current", api.GetCurrentPatient)
		apiv1.GET("/patient/period", api.GetPeriodPatient)
	}
	return r
}
