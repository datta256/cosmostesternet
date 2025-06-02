package types_test

import (
	"testing"

	"testernet/x/kyc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				KycAddressList: []types.KycAddress{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				KycAddressCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated kycAddress",
			genState: &types.GenesisState{
				KycAddressList: []types.KycAddress{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid kycAddress count",
			genState: &types.GenesisState{
				KycAddressList: []types.KycAddress{
					{
						Id: 1,
					},
				},
				KycAddressCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
