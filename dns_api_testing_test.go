package main

import (
	"drone-navigation-service/app/api"
	"drone-navigation-service/app/initializer"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLocation(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("GET", "http://localhost:5000/api/v1/location?sectorId=123456&companyId=atlas&x=123.12&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestGetLocationWithNotRequiredParams(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("GET", "http://localhost:5000/api/v1/location?sectorId=123456&x=123.12&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestGetLocationWithNotPermittedMethod(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("PUT", "http://localhost:5000/api/v1/location?sectorId=123456&x=123.12&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusMethodNotAllowed, response.Code)
}

func TestGetLocationWithInvalidAPI(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("GET", "http://localhost:5000/api/v1/location_test?sectorId=123456&x=123.12&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func TestGetLocationWithInvalidSectorId(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("GET", "http://localhost:5000/api/v1/location?sectorId=12345&companyId=atlas&x=123.12&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func TestGetLocationWithDecodeError(t *testing.T) {
	SetConfig()
	req, _ := http.NewRequest("GET", "http://localhost:5000/api/v1/location?sectorId=12345&companyId=atlas&x=test&y=456.56&z=789.89&vel=20.0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, response.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	api.GetRoutes().ServeHTTP(rr, req)
	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func SetConfig() {
	initializer.Initialize()

}
