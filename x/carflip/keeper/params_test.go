package keeper_test

import (
	"testing"

	testkeeper "carflip/testutil/keeper"
	"carflip/x/carflip/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CarflipKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
