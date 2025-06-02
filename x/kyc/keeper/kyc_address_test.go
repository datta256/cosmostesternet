package keeper_test

import (
	"context"
	"testing"

	keepertest "testernet/testutil/keeper"
	"testernet/testutil/nullify"
	"testernet/x/kyc/keeper"
	"testernet/x/kyc/types"

	"github.com/stretchr/testify/require"
)

func createNKycAddress(keeper keeper.Keeper, ctx context.Context, n int) []types.KycAddress {
	items := make([]types.KycAddress, n)
	for i := range items {
		items[i].Id = keeper.AppendKycAddress(ctx, items[i])
	}
	return items
}

func TestKycAddressGet(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNKycAddress(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetKycAddress(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestKycAddressRemove(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNKycAddress(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveKycAddress(ctx, item.Id)
		_, found := keeper.GetKycAddress(ctx, item.Id)
		require.False(t, found)
	}
}

func TestKycAddressGetAll(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNKycAddress(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllKycAddress(ctx)),
	)
}

func TestKycAddressCount(t *testing.T) {
	keeper, ctx := keepertest.KycKeeper(t)
	items := createNKycAddress(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetKycAddressCount(ctx))
}
