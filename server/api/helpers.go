package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	bpb "github.com/kiloMIA/metricsGo/proto/buses/pb"
	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func getBusData(bus int64) (*bpb.BusResponse, error) {
	conn, err := grpc.Dial("buses:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := bpb.NewBusServiceClient(conn)
	bus_req := &bpb.BusRequest{
		BusNumber: bus,
	}
	res, err := client.RequestBus(context.Background(), bus_req)
	if err != nil {
		log.Fatalf("failed to get bus data: %v", err)
	}
	return res, nil
}

func getMetricsData(city int64, reqType string) (interface{}, error) {
	conn, err := grpc.Dial("metrics:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
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
			log.Fatalf("failed to get temperature data: %v", err)
		}
		return temperatureRes, nil
	case "pollution":
		pollutionRes, err := client.RequestPol(context.Background(), polreq)
		if err != nil {
			log.Fatalf("failed to get pollution data: %v", err)
		}
		return pollutionRes, nil
	default:
		return nil, fmt.Errorf("invalid request type: %s", reqType)
	}
}

func consumeBusDataFromQueue() (*bpb.BusResponse, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"bus_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	msg := <-msgs
	busResponse := &bpb.BusResponse{}
	err = json.Unmarshal(msg.Body, busResponse)
	if err != nil {
		log.Fatalf("Failed to unmarshal message body: %v", err)
	}

	return busResponse, nil
}

func chooseCity(city string) int64 {
	var id int64
	switch city {
	case "astana":
		id = 1
	case "oskemen":
		id = 2
	case "atyrau":
		id = 4
	case "semey":
		id = 3
	}

	return id
}
