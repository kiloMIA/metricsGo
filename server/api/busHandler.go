package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) handleBus(w http.ResponseWriter, r *http.Request) {
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

	// Display the bus data on the response page
	fmt.Fprintf(w, "Bus Data: %v\n", busData)
}
