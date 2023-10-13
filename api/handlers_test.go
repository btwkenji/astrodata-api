package handlers

import (
	"net/http"
	"net/http/httptest"

	GetLastMonthDST(c)

	if w.Code != http.StatusOK {
		t.Errorf("[!] expected status code %d, but got %d", http.StatusOK, w.Code)
	}

	expectedResponse := `{"someKey":"someValue"}`
	if w.Body.String() != expectedResponse {
		t.Errorf("[!] expected response body '%s', but got '%s'", expectedResponse, w.Body.String())
	}
}
