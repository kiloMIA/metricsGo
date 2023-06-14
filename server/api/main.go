package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	bpb "github.com/kiloMIA/metricsGo/proto/buses/pb"
	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	"google.golang.org/grpc"
)

type application struct {
	logger *log.Logger
}

const port = "5000"

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

func getMetricsData(city int64, reqType string) (interface{}, error) {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := metricspb.NewMetricsServiceClient(conn)
	tempreq := &metricspb.TemperatureRequest{
		City: city,
		Type: reqType,
	}
	polreq := &metricspb.PollutionRequest{
		City: city,
		Type: reqType,
	}
	switch reqType {
	case "temperature":
		temperatureRes, err := client.RequestTemp(context.Background(), tempreq)
		if err != nil {
			log.Fatalf("Failed to get temperature data: %v", err)
		}
		return temperatureRes, nil
	case "pollution":
		pollutionRes, err := client.RequestPol(context.Background(), polreq)
		if err != nil {
			log.Fatalf("Failed to get pollution data: %v", err)
		}
		return pollutionRes, nil
	default:
		return nil, fmt.Errorf("Invalid request type: %s", reqType)
	}
}

func getBusData(bus int64) (*bpb.BusResponse, error) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := bpb.NewBusServiceClient(conn)
	bus_req := &bpb.BusRequest{
		BusNumber: bus,
	}
	res, err := client.RequestBus(context.Background(), bus_req)
	if err != nil {
		log.Fatalf("Failed to get bus data: %v", err)
	}
	return res, nil
}

func (app *application) handleIndex(w http.ResponseWriter, r *http.Request) {

	// Serve the HTML file
	http.ServeFile(w, r, "templates/index.html")
}
