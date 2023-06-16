package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleMetrics(t *testing.T) {
	app := &application{}

	reqBody := map[string]string{
		"city":   "astana",
		"metric": "temperature",
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.handleMetrics(w, r)
	}))
	defer ts.Close()

	resp, err := http.Post(ts.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusInternalServerError {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusInternalServerError)
	}

	expectedContentType := "application/json"
	actualContentType := resp.Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("unexpected Content-Type header: got %s, want %s", actualContentType, expectedContentType)
	}
}
