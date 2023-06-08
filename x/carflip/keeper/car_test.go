package keeper_test

import (
	"testing"

	keepertest "carflip/testutil/keeper"
	"carflip/testutil/nullify"
	"carflip/x/carflip/keeper"
	"carflip/x/carflip/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNCar(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Car {
	items := make([]types.Car, n)
	for i := range items {
		items[i].Id = keeper.AppendCar(ctx, items[i])
	}
	return items
}

func TestCarGet(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	items := createNCar(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetCar(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestCarRemove(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	items := createNCar(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveCar(ctx, item.Id)
		_, found := keeper.GetCar(ctx, item.Id)
		require.False(t, found)
	}
}

func TestCarGetAll(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	items := createNCar(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllCar(ctx)),
	)
}

func TestCarCount(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	items := createNCar(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetCarCount(ctx))
}
