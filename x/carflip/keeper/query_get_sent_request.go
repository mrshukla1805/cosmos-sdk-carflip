package keeper

import (
	"context"
	"fmt"

	"carflip/x/carflip/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GetSentRequest(goCtx context.Context, req *types.QueryGetSentRequestRequest) (*types.QueryGetSentRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))
	iterator := sdk.KVStorePrefixIterator(requestStore, []byte{})

	defer iterator.Close()
	var list []types.Request
	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		if val.Creator == req.Creator {
			list = append(list, val)
		}
	}
	fmt.Println(list)
	return &types.QueryGetSentRequestResponse{
		Request: list,
	}, nil

}
