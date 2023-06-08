package carflip

import (
	"carflip/x/carflip/keeper"
	"carflip/x/carflip/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the car
	for _, elem := range genState.CarList {
		k.SetCar(ctx, elem)
	}

	// Set car count
	k.SetCarCount(ctx, genState.CarCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.CarList = k.GetAllCar(ctx)
	genesis.CarCount = k.GetCarCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
