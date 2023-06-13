package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"

	"carflip/x/carflip/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetRecievedRequest(goCtx context.Context, req *types.QueryGetRecievedRequestRequest) (*types.QueryGetRecievedRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))
	iterator := sdk.KVStorePrefixIterator(requestStore, []byte{})

	defer iterator.Close()
	var list []types.Request
	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.RequestTo == req.Creator {
			list = append(list, val)
		}
	}
	return &types.QueryGetRecievedRequestResponse{
		Request: list,
	}, nil

}
