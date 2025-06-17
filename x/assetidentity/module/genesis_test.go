package assetidentity_test

import (
	"testing"

	keepertest "testernet/testutil/keeper"
	"testernet/testutil/nullify"
	assetidentity "testernet/x/assetidentity/module"
	"testernet/x/assetidentity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AssetList: []types.Asset{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AssetCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetidentityKeeper(t)
	assetidentity.InitGenesis(ctx, k, genesisState)
	got := assetidentity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AssetList, got.AssetList)
	require.Equal(t, genesisState.AssetCount, got.AssetCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
