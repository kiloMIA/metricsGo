package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/streadway/amqp"
)

type application struct {
	logger *log.Logger
}

const (
	port = 8080
	rbmq = "amqp://guest:guest@localhost:5672/"
)

func main() {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	flag.Parse()

	app := &application{
		logger: logger,
	}

	// Start the web server
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

	conn, err := amqp.Dial(rbmq)
	if err != nil {
		log.Panicf("Could not connect to RabbitMQ")
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		log.Panicf("Could not open channel")
	}
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"data_direct", //exchnage name
		"direct",      //type
		true,          //durable
		false,         //auto-deleted
		false,         //internal
		false,         //no-wait
		nil,           //arguments
	)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
}

func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {

	// Serve the HTML file
	http.ServeFile(w, r, "server/templates/index.html")
}
