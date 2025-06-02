package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		KycAddressList: []KycAddress{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in kycAddress
	kycAddressIdMap := make(map[uint64]bool)
	kycAddressCount := gs.GetKycAddressCount()
	for _, elem := range gs.KycAddressList {
		if _, ok := kycAddressIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for kycAddress")
		}
		if elem.Id >= kycAddressCount {
			return fmt.Errorf("kycAddress id should be lower or equal than the last id")
		}
		kycAddressIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
