package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetHealth(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	GetHealth(c)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"status":"OK"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("[!] expected response body '%s', but got '%s'", expectedResponse, w.Body.String())
	}
}
