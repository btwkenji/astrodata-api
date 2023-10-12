package vars

const (
	LastMonthDataUrl                 string  = "https://wdc.kugi.kyoto-u.ac.jp/dst_realtime/lastmonth/index.html"
	CurrentMonthDataUrl              string  = "https://wdc.kugi.kyoto-u.ac.jp/dst_realtime/presentmonth/index.html"
	GetDataByDateUrl                 string  = "https://wdc.kugi.kyoto-u.ac.jp/dst_realtime/%s/index.html"
	Dst7DaysDataUrl                  string  = "https://services.swpc.noaa.gov/products/kyoto-dst.json"
	RealTimeSolarWindDataSixHours    string  = "https://services.swpc.noaa.gov/products/solar-wind/mag-6-hour.json"
	RealTimeSolarWindDataOneDay      string  = "https://services.swpc.noaa.gov/products/solar-wind/mag-1-day.json"
	RealTimeSolarWindDataThreeDays   string  = "https://services.swpc.noaa.gov/products/solar-wind/mag-3-day.json"
	RealTimeSolarWindDataSevenDays   string  = "https://services.swpc.noaa.gov/products/solar-wind/mag-7-day.json"
	PlasmaTemperature1DayDataUrl     string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-1-day.json"
	PlasmaTemperature7DaysDataUrl    string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-7-day.json"
	PlasmaTemperature2hoursDataUrl   string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-2-hour.json"
	PlasmaTemperature3daysDataUrl    string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-3-day.json"
	PlasmaTemperature6hoursDataUrl   string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-6-hour.json"
	PlasmaTemperatureRealTimeDataUrl string  = "https://services.swpc.noaa.gov/products/solar-wind/plasma-5-minute.json"
	CONSTANT                         float64 = 12.48
	DataPointsPerPart                int64   = 72
)
