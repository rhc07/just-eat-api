package main

import (
	"testing"
)

//type App struct {
//	Router *mux.Router
//}

//func TestGetProduct(t *testing.T) {
//
//	req, _ := http.NewRequest("GET", "/", nil)
//	response := executeRequest(req)
//
//	checkResponseCode(t, http.StatusOK, response.Code)
//}
//
//func executeRequest(req *http.Request) *httptest.ResponseRecorder {
//	rr := httptest.NewRecorder()
//	App{}
//	a.Router.ServeHTTP(rr, req)
//
//	return rr
//}
//
//func executeRequest(req *http.Request) {
//
//}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
