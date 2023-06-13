package main

import (
	"context"
	"flag"
	"log"
	"net"

	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	metricspb.UnimplementedMetricsServiceServer
}

func (s *server) RequestTemp(ctx context.Context, req *metricspb.TemperatureRequest) (*metricspb.TemperatureResponse, error) {
	log.Printf("Received: %v", req.GetCity())
	return &metricspb.TemperatureResponse{City: req.City, Temperature: 150000, District: "14", Humidity: 100}, nil
}

// might be implemented in python

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	metricspb.RegisterMetricsServiceServer(s, &server{})
	log.Println("Metric Service server started on localhost:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
