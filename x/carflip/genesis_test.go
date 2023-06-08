package carflip_test

import (
	"testing"

	keepertest "carflip/testutil/keeper"
	"carflip/testutil/nullify"
	"carflip/x/carflip"
	"carflip/x/carflip/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CarList: []types.Car{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CarCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CarflipKeeper(t)
	carflip.InitGenesis(ctx, *k, genesisState)
	got := carflip.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CarList, got.CarList)
	require.Equal(t, genesisState.CarCount, got.CarCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
