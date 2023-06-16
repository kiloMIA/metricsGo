package main

import (
	"errors"
	"net/http/httptest"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	tests := []struct {
		message     string
		code        int
		expected    int
		expectedErr error
	}{
		{"Error message", 400, 500, errors.New("test error")},
		{"Another error message", 500, 500, nil},
	}

	for _, test := range tests {
		// Create a new HTTP request and response recorder
		req := httptest.NewRequest("GET", "/", nil)
		res := httptest.NewRecorder()

		// Call the errorResponse function with the test parameters
		errorResponse(test.message, test.code, res, req, test.expectedErr)

		// Check the response status code
		if res.Code != test.expected {
			t.Errorf("errorResponse returned status code %d, expected %d", res.Code, test.expected)
		}
	}
}
