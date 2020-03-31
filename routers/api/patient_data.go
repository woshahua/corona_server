package api

//func GetLastestUpdateTime(c *gin.Context) {
//	data, err := models.GetLastUpdatedTime()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed ", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetLatestPatient(c *gin.Context) {
//	data, err := models.GetLatestPatientData()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed ", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetPeriodPatient(c *gin.Context) {
//	data, err := models.GetPeriodPatientData()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed ", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetDailyPatient(c *gin.Context) {
//	data, err := models.GetDailyPatientData()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed ", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetDeadPatient(c *gin.Context) {
//	data, err := models.GetDeadPatientData()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetCurrentPatient(c *gin.Context) {
//	data, err := models.GetCurrentPatient()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Print("failed", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
//
//func GetPatientByLocation(c *gin.Context) {
//	data, err := models.GetLocationPatientData()
//	code := e.SUCCESS
//	if err != nil {
//		code = e.ERROR
//		fmt.Println("failed fetch patients by location: ", err)
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}
