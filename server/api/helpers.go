package main

import (
	"context"
	"fmt"
	"log"

	bpb "github.com/kiloMIA/metricsGo/proto/buses/pb"
	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	"google.golang.org/grpc"
)

func getBusData(bus int64) (*bpb.BusResponse, error) {
	conn, err := grpc.Dial("buses:50054", grpc.WithInsecure())
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

func getMetricsData(city int64, reqType string) (interface{}, error) {
	conn, err := grpc.Dial("metrics:50052", grpc.WithInsecure())
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
