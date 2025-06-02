package kyc_test

import (
	"testing"

	keepertest "testernet/testutil/keeper"
	"testernet/testutil/nullify"
	kyc "testernet/x/kyc/module"
	"testernet/x/kyc/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		KycAddressList: []types.KycAddress{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		KycAddressCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KycKeeper(t)
	kyc.InitGenesis(ctx, k, genesisState)
	got := kyc.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.KycAddressList, got.KycAddressList)
	require.Equal(t, genesisState.KycAddressCount, got.KycAddressCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
