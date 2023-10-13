package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLastMonthDST(t *testing.T) {
	r := gin.New()
	r.GET("/api/dst/last-month", GetLastMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/last-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetCurrentMonthDST(t *testing.T) {
	r := gin.New()
	r.GET("/api/dst/current-month", GetCurrentMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/current-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetDSTByDate(t *testing.T) {
	r := gin.New()
	r.GET("/api/dst/by-date", GetDSTByDate)

	req := httptest.NewRequest("GET", "/api/dst/by-date?date=202309", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGet7DaysDST(t *testing.T) {
	r := gin.New()
	r.GET("/api/dst/7d", GetDST7Days)

	req := httptest.NewRequest("GET", "/api/dst/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBzDataSixHours(t *testing.T) {
	r := gin.New()
	r.GET("/api/bz/6h", GetBzDataSixHours)

	req := httptest.NewRequest("GET", "/api/bz/6h", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBzDataOneDay(t *testing.T) {
	r := gin.New()
	r.GET("/api/bz/1d", GetBzDataOneDay)

	req := httptest.NewRequest("GET", "/api/bz/1d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBzDataThreeDays(t *testing.T) {
	r := gin.New()
	r.GET("/api/bz/3d", GetBzDataThreeDays)

	req := httptest.NewRequest("GET", "/api/bz/3d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBzDataSevenDays(t *testing.T) {
	r := gin.New()
	r.GET("/api/bz/7d", GetBzDataSevenDays)

	req := httptest.NewRequest("GET", "/api/bz/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperatureRealTime(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/now", GetPlasmaTemperatureRealTime)

	req := httptest.NewRequest("GET", "/api/plasma/now", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperature2Hours(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/2h", GetPlasmaTemperature2Hours)

	req := httptest.NewRequest("GET", "/api/plasma/2h", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperature6Hours(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/6h", GetPlasmaTemperature6Hours)

	req := httptest.NewRequest("GET", "/api/plasma/6h", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperature1Day(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/1d", GetPlasmaTemperature1Day)

	req := httptest.NewRequest("GET", "/api/plasma/1d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperature3Days(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/3d", GetPlasmaTemperature3Days)

	req := httptest.NewRequest("GET", "/api/plasma/3d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPlasmaTemperature7Days(t *testing.T) {
	r := gin.New()
	r.GET("/api/plasma/7d", GetPlasmaTemperature7Days)

	req := httptest.NewRequest("GET", "/api/plasma/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPredictOneDay(t *testing.T) {
	r := gin.New()
	r.GET("/api/predict/1d", GetOneDayPredict)

	req := httptest.NewRequest("GET", "/api/predict/1d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetPredict6Hours(t *testing.T) {
	r := gin.New()
	r.GET("/api/predict/6h", GetSixHoursPredictFromNow)

	req := httptest.NewRequest("GET", "/api/predict/6h", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetDstStrength(t *testing.T) {
	r := gin.New()
	r.GET("/api/dst/now/strength", GetDSTAndStrengthNow)

	req := httptest.NewRequest("GET", "/api/dst/now/strength", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

func TestGetBzStrength(t *testing.T) {
	r := gin.New()
	r.GET("/api/bz/now/strength", GetBzAndStrengthNow)

	req := httptest.NewRequest("GET", "/api/bz/now/strength", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}
}

// Benchmarks start here
func BenchmarkGetLastMonthDST(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/last-month", GetLastMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/last-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetCurrentMonthDST(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/current-month", GetCurrentMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/current-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetDSTByDate(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/by-date", GetDSTByDate)

	req := httptest.NewRequest("GET", "/api/dst/by-date?date=202309", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGet7DaysDST(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/7d", GetDST7Days)

	req := httptest.NewRequest("GET", "/api/dst/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetBzDataSixHours(b *testing.B) {
	r := gin.New()
	r.GET("/api/bz/6h", GetBzDataSixHours)

	req := httptest.NewRequest("GET", "/api/bz/6h", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetBzDataOneDay(b *testing.B) {
	r := gin.New()
	r.GET("/api/bz/1d", GetBzDataOneDay)

	req := httptest.NewRequest("GET", "/api/bz/1d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetBzData3Days(b *testing.B) {
	r := gin.New()
	r.GET("/api/bz/3d", GetBzDataThreeDays)

	req := httptest.NewRequest("GET", "/api/bz/3d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetBzData7Days(b *testing.B) {
	r := gin.New()
	r.GET("/api/bz/7d", GetBzDataSevenDays)

	req := httptest.NewRequest("GET", "/api/bz/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetPlasmaTemperature7Days(b *testing.B) {
	r := gin.New()
	r.GET("/api/plasma/7d", GetPlasmaTemperature7Days)

	req := httptest.NewRequest("GET", "/api/plasma/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetLastMonthDSTParallel(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/last-month", GetLastMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/last-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetCurrentMonthDSTParallel(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/current-month", GetCurrentMonthDST)

	req := httptest.NewRequest("GET", "/api/dst/current-month", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetDSTByDateParallel(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/by-date", GetDSTByDate)

	req := httptest.NewRequest("GET", "/api/dst/by-date?date=202309", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGet7DaysDSTParallel(b *testing.B) {
	r := gin.New()
	r.GET("/api/dst/7d", GetDST7Days)

	req := httptest.NewRequest("GET", "/api/dst/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}

func BenchmarkGetBzData7DaysParallel(b *testing.B) {
	r := gin.New()
	r.GET("/api/bz/7d", GetBzDataSevenDays)

	req := httptest.NewRequest("GET", "/api/bz/7d", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
}
