package identity_test

import (
	"testing"

	keepertest "testernet/testutil/keeper"
	"testernet/testutil/nullify"
	identity "testernet/x/identity/module"
	"testernet/x/identity/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		IdentityList: []types.Identity{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		IdentityCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IdentityKeeper(t)
	identity.InitGenesis(ctx, k, genesisState)
	got := identity.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.IdentityList, got.IdentityList)
	require.Equal(t, genesisState.IdentityCount, got.IdentityCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
