package testernet_test

import (
	"testing"

	keepertest "testernet/testutil/keeper"
	"testernet/testutil/nullify"
	testernet "testernet/x/testernet/module"
	"testernet/x/testernet/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TesternetKeeper(t)
	testernet.InitGenesis(ctx, k, genesisState)
	got := testernet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
