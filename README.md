


## How to start the project
*  Clone the repository: git clone `https://github.com/kiloMIA/metricsGo.git`
*  Build via docker-compose: `docker-compose up -d --build`
*  Linux: `sudo docker compose up -d --build`
*  Windows: `docker-compose up -d --build`
*  Go to `localhost:8080` and choose the service that you want to use
    
# Architecture![Why_use_it_To_get_clarity_on_a_group's_structure_To_see_how_roles](https://github.com/kiloMIA/metricsGo/assets/97970527/606adaf3-db19-4fdd-b063-a361e2d7af7e)

# Protobuff structure
* Metrics
```
service MetricsService{
  rpc RequestTemp(TemperatureRequest) returns (TemperatureResponse) {}
  rpc RequestPol(PollutionRequest) returns (PollutionResponse) {}
}

message TemperatureResponse {
  int64 city = 1;
  string district = 2;
  double temperature = 3;
  double humidity = 4;
}

message PollutionResponse {
  int64 city = 1;
  string district = 2;
  double co2 = 3;
  double pm25 = 4;
}

message TemperatureRequest {
  int64 city = 1;
  string type = 2;
}

message PollutionRequest {
  int64 city = 1;
  string type = 2;
}
```
* Buses
```

service BusService{
  rpc RequestBus(BusRequest) returns (BusResponse) {}

}
message BusRequest {
  int64 BusNumber = 1;
}

message BusResponse {
  double longitude = 1;
  double latitude = 2;
}
```
