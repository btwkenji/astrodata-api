package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kenjitheman/magneconn_api/core"
	"github.com/kenjitheman/magneconn_api/vars"
	"net/http"
)

func GetLastMonthDST(c *gin.Context) {
	data, err := core.Parser(vars.LastMonthDataUrl)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetCurrentMonthDST(c *gin.Context) {
	data, err := core.Parser(vars.CurrentMonthDataUrl)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetDSTByDate(c *gin.Context) {
	date := c.DefaultQuery("date", "")
	if date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date parameter is missing!"})
		return
	}
	data, err := core.Parser(fmt.Sprintf(vars.GetDataByDateUrl, date))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
	if len(date) != 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBzDataSixHours(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataSixHours)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBzDataOneDay(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataOneDay)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBzDataThreeDays(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataThreeDays)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetBzDataSevenDays(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataSevenDays)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetSixHoursPredictFromNow(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataSixHours)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, core.CalculateDst(data))
}

func GetOneDayPredict(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataOneDay)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	parts := core.SplitData(data)
	dstValues := core.CalculateDstForParts(parts)
	c.JSON(http.StatusOK, dstValues)
}

func GetBzNow(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataSixHours)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, data[len(data)-1])
}

func GetDSTNow(c *gin.Context) {
	data, err := core.Parser(vars.CurrentMonthDataUrl)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	if len(data) > 0 && data[len(data)-1] > 0 {
		c.JSON(http.StatusOK, data[len(data)-1])
	} else {
		c.JSON(http.StatusOK, 0)
	}
}

func GetBzAndStrengthNow(c *gin.Context) {
	data, err := core.ParseBzGsmData(vars.RealTimeSolarWindDataSixHours)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stormStrength := ""
	if data[len(data)-1] >= -200 && data[len(data)-1] < -7 {
		stormStrength = "strong"
	} else if data[len(data)-1] >= -7 && data[len(data)-1] < -3 {
		stormStrength = "medium"
	} else if data[len(data)-1] >= -3 && data[len(data)-1] < 10 {
		stormStrength = "weak"
	}

	response := BzResponse{
		Bz:            data[len(data)-1],
		StormStrength: stormStrength,
	}

	c.JSON(http.StatusOK, response)
}

func GetDSTAndStrengthNow(c *gin.Context) {
	data, err := core.Parser(vars.CurrentMonthDataUrl)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	stormStrength := ""
	if len(data) > 0 && data[len(data)-1] >= -7 && data[len(data)-1] < -5 {
		stormStrength = "strong"
	} else if len(data) > 0 && data[len(data)-1] >= -4 && data[len(data)-1] < -2 {
		stormStrength = "medium"
	} else if len(data) > 0 && data[len(data)-1] >= -1 {
		stormStrength = "weak"
	}

	response := DSTResponse{
		DST:           data[len(data)-1],
		StormStrength: stormStrength,
	}

	c.JSON(http.StatusOK, response)
}

func GetDST7Days(c *gin.Context) {
  data, err := core.ParseDst7Days(vars.Dst7DaysDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperature1Day(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperature1DayDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperature7Days(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperature7DaysDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperature2Hours(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperature2hoursDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperature3Days(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperature3daysDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperature6Hours(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperature6hoursDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}

func GetPlasmaTemperatureRealTime(c *gin.Context) {
  data, err := core.ParsePlasmaTemperatureData(vars.PlasmaTemperatureRealTimeDataUrl)
  if err != nil {
    c.AbortWithError(http.StatusBadRequest, err)
    return
  }
  c.JSON(http.StatusOK, data)
}
