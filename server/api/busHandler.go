package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) handleBus(w http.ResponseWriter, r *http.Request) {
	var extractedData []struct {
		Longitude float64 `json:"longitude"`
		Latitude  float64 `json:"latitude"`
	}
	busRoute := r.FormValue("route_number")
	bus, _ := strconv.ParseInt(busRoute, 10, 32)

	_, err := getBusData(bus)
	if err != nil {
		http.Error(w, "Failed to get bus data", http.StatusInternalServerError)
		return
	}

	busData, err := consumeBusDataFromQueue()
	if err != nil {
		http.Error(w, "Failed to get bus data from queue", http.StatusInternalServerError)
		return
	}

	for _, data := range busData {
		extractedData = append(extractedData, struct {
			Longitude float64 `json:"longitude"`
			Latitude  float64 `json:"latitude"`
		}{
			Longitude: data.Longitude,
			Latitude:  data.Latitude,
		})
	}

	jsonData, err := json.Marshal(extractedData)
	errorResponse("Failed to encode", http.StatusInternalServerError, w, r, err)

	tmpl, err := template.ParseFiles("server/templates/bus.html")
	errorResponse("Failed to parse files", http.StatusInternalServerError, w, r, err)

	err = tmpl.Execute(w, string(jsonData))

	errorResponse("Failed to execute template", http.StatusInternalServerError, w, r, err)
}
