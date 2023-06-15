package main

import (
	"fmt"
	"net/http"

	metricspb "github.com/kiloMIA/metricsGo/proto/metrics/pb"
)

func (app *application) handleMetrics(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")

	reqType := r.FormValue("metric")
	//num, _ := strconv.ParseInt(i, 10, 32)
	num := chooseCity(city)

	responseData, err := getMetricsData(num, reqType)
	if err != nil {
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	metricData, err := consumeBusDataFromQueue()
	if err != nil {
		http.Error(w, "Failed to get bus data from queue", http.StatusInternalServerError)
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

	fmt.Fprintf(w, "Bus Data: %v\n", metricData)

}
