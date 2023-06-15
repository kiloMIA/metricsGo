package main

import (
	"context"
	"fmt"
	"log"
	"time"

	bpb "github.com/kiloMIA/metricsGo/proto/buses/pb"
	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func getBusData(bus int64) (*bpb.BusResponse, error) {
	conn, err := grpc.Dial("localhost:50054", grpc.WithInsecure())
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
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
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

func publishData(routingKey string) error {
	conn, err := amqp.Dial(rbmq)
	if err != nil {
		return fmt.Errorf("could not connect to rabbitmq, %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err != nil {
		return fmt.Errorf("could not open channel, %v", err)
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
	if err != nil {

		return fmt.Errorf("could not create exchange, %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"data_direct", //exchange
		routingKey,    //routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("a"),
		})

	if err != nil {
		return fmt.Errorf("could not publish message, %v", err)
	}

	return nil
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
