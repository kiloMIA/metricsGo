package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	logger *log.Logger
}

const port = 8080

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	flag.Parse()

	app := &application{
		logger: logger,
	}

	// Start the web server
	// log.Println("Web server started on localhost:8080")
	// http.ListenAndServe(":8080", nil)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting server on %s", srv.Addr)
	err := srv.ListenAndServe()

	logger.Fatal(err)
}

func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {

	// Serve the HTML file
	http.ServeFile(w, r, "server/templates/index.html")
}
