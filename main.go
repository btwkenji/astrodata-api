package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kenjitheman/magneconn_api/api"
	"github.com/kenjitheman/magneconn_api/api/health"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/health", health.GetHealth)
		api.GET("/dst/last-month", handlers.GetLastMonthDST)
		api.GET("/dst/current-month", handlers.GetCurrentMonthDST)
    api.GET("/dst/7d", handlers.GetDST7Days)
		api.GET("/dst/by-date", handlers.GetDSTByDate)
    api.GET("/dst/now", handlers.GetDSTNow)
    api.GET("/dst/now/strength", handlers.GetDSTAndStrengthNow)
		api.GET("/predict/6h", handlers.GetSixHoursPredictFromNow)
    api.GET("/predict/1d", handlers.GetOneDayPredict)
		api.GET("/bz/6h", handlers.GetBzDataSixHours)
		api.GET("/bz/1d", handlers.GetBzDataOneDay)
		api.GET("/bz/3d", handlers.GetBzDataThreeDays)
		api.GET("/bz/7d", handlers.GetBzDataSevenDays)
    api.GET("/bz/now", handlers.GetBzNow)
    api.GET("/bz/now/strength", handlers.GetBzAndStrengthNow)
    api.GET("plasma/now", handlers.GetPlasmaTemperatureRealTime)
    api.GET("plasma/2h", handlers.GetPlasmaTemperature2Hours)
    api.GET("plasma/6h", handlers.GetPlasmaTemperature6Hours)
    api.GET("plasma/1d", handlers.GetPlasmaTemperature1Day)
    api.GET("plasma/3d", handlers.GetPlasmaTemperature3Days)
    api.GET("plasma/7d", handlers.GetPlasmaTemperature7Days)
	}

	r.Run()
}
