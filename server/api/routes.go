package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()

	staticHandler := http.FileServer(http.Dir("server/templates/static"))
	router.Handler("GET", "/static/*filepath", http.StripPrefix("/static/", staticHandler))

	router.HandlerFunc(http.MethodPost, "/metrics", app.handleMetrics)
	router.HandlerFunc(http.MethodGet, "/", app.handleIndex)
	router.HandlerFunc(http.MethodPost, "/bus", app.handleBus)

	return router
}
