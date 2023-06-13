package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	"google.golang.org/grpc"
)

func main() {
	flag.Parse()

	// Set up routes
	http.HandleFunc("/", handleIndex)

	// Start the web server
	log.Println("Web server started on localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func getMetricsData(city int64) (*metricspb.TemperatureResponse, error) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := metricspb.NewMetricsServiceClient(conn)
	req := &metricspb.TemperatureRequest{
		City: city,
	}

	res, err := client.RequestTemp(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed to get temprature data: %v", err)
	}
	return res, nil
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle the user request
		city := r.FormValue("city")
		num, _ := strconv.ParseInt(city, 10, 64)
		temperatureData, err := getMetricsData(num)
		if err != nil {
			http.Error(w, "Failed to get temperature data", http.StatusInternalServerError)
			return
		}

		// Display the temperature data on the response page
		fmt.Fprintf(w, "Temperature Data: %v\n", temperatureData)
	} else {
		// Serve the HTML file
		http.ServeFile(w, r, "server/templates/index.html")

	}
}
