package main

import "net/http"

func errorResponse(message string, code int, w http.ResponseWriter, r *http.Request, err error) {
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
