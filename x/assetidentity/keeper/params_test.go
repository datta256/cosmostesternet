package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "testernet/testutil/keeper"
	"testernet/x/assetidentity/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.AssetidentityKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
