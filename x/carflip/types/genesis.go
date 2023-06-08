package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		CarList:     []Car{},
		RequestList: []Request{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in car
	carIdMap := make(map[uint64]bool)
	carCount := gs.GetCarCount()
	for _, elem := range gs.CarList {
		if _, ok := carIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for car")
		}
		if elem.Id >= carCount {
			return fmt.Errorf("car id should be lower or equal than the last id")
		}
		carIdMap[elem.Id] = true
	}
	// Check for duplicated ID in request
	requestIdMap := make(map[uint64]bool)
	requestCount := gs.GetRequestCount()
	for _, elem := range gs.RequestList {
		if _, ok := requestIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for request")
		}
		if elem.Id >= requestCount {
			return fmt.Errorf("request id should be lower or equal than the last id")
		}
		requestIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
