package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Defines the schema of the root of the API response.
type APIResponse struct {
	Services []Service
}

// Defines the incoming bus of a single service at a bus stop.
type Service struct {
	NextBus NextBus
	NextBus2 NextBus
	NextBus3 NextBus
}

// Defines a NextBus field in the API response.
type NextBus struct {
	EstimatedArrival string `json:"EstimatedArrival"`
	Load string `json:"Load"`
	Feature string `json:"Feature"`
	Type string `json:"Type"`
}

const baseURL = "http://datamall2.mytransport.sg/ltaodataservice/BusArrivalv2"

var accountKey = os.Getenv("ACCOUNT_KEY")

func callBusArrivalAPI(busStopCode string, serviceNumber string) (*APIResponse, error) {
	// Create a new request with URL set to the BusArrival API endpoint + query params.
	req, err := http.NewRequest("GET", fmt.Sprintf("%s?BusStopCode=%s&ServiceNo=%s", baseURL, busStopCode, serviceNumber), nil)
	if err != nil {
		return nil, err
	}

	// Attach the required key for accessing the Bus Arrival API.
	req.Header.Set("AccountKey", accountKey)

	// Make the API call.
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// Decode response body into APIResponse struct.
	a := &APIResponse{}
	if err := json.NewDecoder(res.Body).Decode(a); err != nil {
		return nil, err
	}

	return a, nil
}