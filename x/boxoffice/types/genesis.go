package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		SystemInfo: SystemInfo{
			NextId: uint64(DefaultIndex),
		},
		StoredConcertList: []StoredConcert{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in storedConcert
	storedConcertIndexMap := make(map[string]struct{})

	for _, elem := range gs.StoredConcertList {
		index := string(StoredConcertKey(elem.Index))
		if _, ok := storedConcertIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for storedConcert")
		}
		storedConcertIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
