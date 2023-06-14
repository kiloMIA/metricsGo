package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	bpb "github.com/kiloMIA/metricsGo/proto/buses/pb"
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

func handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Handle the user request
		city := r.FormValue("cityName")
		reqType := r.FormValue("reqType")
		num, _ := strconv.ParseInt(city, 10, 32)
		responseData, err := getMetricsData(num, reqType)
		if err != nil {
			http.Error(w, "Failed to get data", http.StatusInternalServerError)
			return
		}

		switch data := responseData.(type) {
		case *metricspb.TemperatureResponse:
			// Handle temperature response
			fmt.Fprintf(w, "Temperature Data: City %d, District %s, Temperature %d, Humidity %d\n",
				data.City, data.District, data.Temperature, data.Humidity)
		case *metricspb.PollutionResponse:
			// Handle pollution response
			fmt.Fprintf(w, "Pollution Data: City %d, District %s, CO2 %d, PM2.5 %d\n",
				data.City, data.District, data.Co2, data.Pm25)
		default:
			// Handle unknown response type
			http.Error(w, "Unknown response type", http.StatusInternalServerError)
			return
		}
		busRoute := r.FormValue("busRoute")
		bus, _ := strconv.ParseInt(busRoute, 10, 32)
		busData, err := getBusData(bus)
		if err != nil {
			http.Error(w, "Failed to get bus data", http.StatusInternalServerError)
			return
		}

		// Display the temperature data on the response page
		fmt.Fprintf(w, "Bus Data: %v\n", busData)

	} else {
		// Serve the HTML file
		http.ServeFile(w, r, "templates/index.html")
	}
}
