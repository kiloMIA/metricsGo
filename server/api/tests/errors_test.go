package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	errMsg := "Some error message"
	errorResponse(errMsg, http.StatusInternalServerError, rr, req, err)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Ожидается код статуса %v, получен %v", http.StatusInternalServerError, status)
	}
}

func errorResponse(errMsg string, i int, rr *httptest.ResponseRecorder, req *http.Request, err error) {
	panic("unimplemented")
}
