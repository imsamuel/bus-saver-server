package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (s *server) handleIncomingBusesGet (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract the busStopCode and serviceNumber params from URL.
	busStopCode := ps.ByName("busStopCode")
	serviceNumber := ps.ByName("serviceNumber")

	// Call the BusArrivalAPI using the callBusArrivalAPI function.
	a, err := callBusArrivalAPI(busStopCode, serviceNumber) // note: a is a pointer
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	transformed := transform(*a)
	transformedBytes, err := json.Marshal(transformed)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(transformedBytes)
}
