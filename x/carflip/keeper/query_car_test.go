package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "carflip/testutil/keeper"
	"carflip/testutil/nullify"
	"carflip/x/carflip/types"
)

func TestCarQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCar(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetCarRequest
		response *types.QueryGetCarResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetCarRequest{Id: msgs[0].Id},
			response: &types.QueryGetCarResponse{Car: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetCarRequest{Id: msgs[1].Id},
			response: &types.QueryGetCarResponse{Car: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetCarRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Car(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestCarQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CarflipKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCar(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCarRequest {
		return &types.QueryAllCarRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CarAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Car), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Car),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CarAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Car), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Car),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CarAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Car),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CarAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
