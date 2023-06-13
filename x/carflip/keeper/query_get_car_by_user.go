package keeper

import (
	"context"

	"carflip/x/carflip/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetCarByUser(goCtx context.Context, req *types.QueryGetCarByUserRequest) (*types.QueryGetCarByUserResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx
	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.CarKey))
	iterator := sdk.KVStorePrefixIterator(requestStore, []byte{})
	defer iterator.Close()
	var data types.Car
	for ; iterator.Valid(); iterator.Next() {
		var val types.Car
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Creator == req.Creator {
			data = val
		}
	}
	return &types.QueryGetCarByUserResponse{
		Car: &data,
	}, nil
}
