package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
)

func TestRoutes(t *testing.T) {
	app := &application{}

	router := app.routes()

	ts := httptest.NewServer(router)
	defer ts.Close()

	testCases := []struct {
		method   string
		path     string
		expected int
	}{
		{http.MethodGet, "/static/css/style.css", http.StatusOK},
		{http.MethodPost, "/metrics", http.StatusOK},
		{http.MethodGet, "/", http.StatusOK},
		{http.MethodPost, "/bus", http.StatusOK},
	}

	for _, tc := range testCases {
		t.Run(tc.method+" "+tc.path, func(t *testing.T) {
			req, err := http.NewRequest(tc.method, ts.URL+tc.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tc.expected {
				t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, tc.expected)
			}
		})
	}
}
