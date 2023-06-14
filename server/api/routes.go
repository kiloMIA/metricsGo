package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	router.HandlerFunc(http.MethodPost, "/metrics", app.handleMetrics)
	router.HandlerFunc(http.MethodGet, "/", app.handleIndex)
	router.HandlerFunc(http.MethodPost, "/bus", app.handleBus)

	return router
}
