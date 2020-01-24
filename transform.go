package main

type IncomingBus struct {
	EstimatedArrival string `json:"estimatedArrival"`
	Load string `json:"load"`
	IsWheelChairAccessible bool `json:"isWheelChairAccessible"`
	Type string `json:"type"`
}

func transform(a APIResponse) []IncomingBus {
	all := a.Services[0] // Only fetching for a single service.

	incomingBuses := make([]IncomingBus, 0)

	// Check through first NextBus
	if all.NextBus.Load == "" {
		return incomingBuses
	}
	var nextBus IncomingBus
	nextBus.EstimatedArrival = all.NextBus.EstimatedArrival
	nextBus.Load = all.NextBus.Load
	if all.NextBus.Feature == "WAB" {
		nextBus.IsWheelChairAccessible = true
	} else {
		nextBus.IsWheelChairAccessible = false
	}
	nextBus.Type = all.NextBus.Type
	incomingBuses = append(incomingBuses, nextBus)

	// Check through first NextBus2
	if all.NextBus2.Load == "" {
		return incomingBuses
	}
	var nextBus2 IncomingBus
	nextBus2.EstimatedArrival = all.NextBus2.EstimatedArrival
	nextBus2.Load = all.NextBus2.Load
	if all.NextBus2.Feature == "WAB" {
		nextBus2.IsWheelChairAccessible = true
	} else {
		nextBus2.IsWheelChairAccessible = false
	}
	nextBus2.Type = all.NextBus2.Type
	incomingBuses = append(incomingBuses, nextBus2)

	// Check through first NextBus3
	if all.NextBus3.Load == "" {
		return incomingBuses
	}
	var nextBus3 IncomingBus
	nextBus3.EstimatedArrival = all.NextBus3.EstimatedArrival
	nextBus3.Load = all.NextBus3.Load
	if all.NextBus3.Feature == "WAB" {
		nextBus3.IsWheelChairAccessible = true
	} else {
		nextBus3.IsWheelChairAccessible = false
	}
	nextBus3.Type = all.NextBus3.Type
	incomingBuses = append(incomingBuses, nextBus3)
	return incomingBuses
}
