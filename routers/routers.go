package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/woshahua/corona_server/pkg/setting"
	"github.com/woshahua/corona_server/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("api")
	{
		apiv1.GET("/patient/japanese/location", api.GetJapanesePatientByLocation)
		apiv1.GET("/patient/news", api.GetNewsData)
		apiv1.GET("/patient/japanese/summary", api.GetJapanesePatientSummary)
	}
	return r
}
