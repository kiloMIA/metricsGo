package main

import (
	"fmt"
	"net/http"
)

func (app *application) handleMetrics(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	reqType := r.FormValue("metric")
	num := chooseCity(city)

	_, err := getMetricsData(num, reqType)
	if err != nil {
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	var metricData map[string]interface{}
	switch reqType {
	case "temperature":
		metricData, err = consumeFromQueue("temperature_queue")
		if err != nil {
			http.Error(w, "Failed to get temperature data from queue", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Temperature Data: %v\n", metricData)
	case "pollution":
		metricData, err = consumeFromQueue("pollution_queue")
		if err != nil {
			http.Error(w, "Failed to get pollution data from queue", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Pollution Data: %v\n", metricData)
	default:
		http.Error(w, "Unknown metric type", http.StatusBadRequest)
		return
	}
}
