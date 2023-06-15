package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *application) handleMetrics(w http.ResponseWriter, r *http.Request) {
	city := r.FormValue("city")
	reqType := r.FormValue("metric")
	num := chooseCity(city)

	_, err := getMetricsData(num, reqType)

	var metricData []map[string]interface{}
	switch reqType {
	case "temperature":
		metricData, err = consumeFromQueue("temperature_queue")

		errorResponse("Failed to get temperature data from queue", http.StatusInternalServerError, w, r, err)

	case "pollution":
		metricData, err = consumeFromQueue("pollution_queue")

		errorResponse("Failed to get pollution data from queue", http.StatusInternalServerError, w, r, err)

	default:
		errorResponse("Unknown request type", http.StatusBadGateway, w, r, err)

	}
	jsonData, err := json.MarshalIndent(metricData, "", "    ")
	errorResponse("Failed To Convert To JSON Data", http.StatusInternalServerError, w, r, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
	log.Print(metricData)
}
