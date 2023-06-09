package keeper

import (
	"context"

	"carflip/x/carflip/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) RequestAll(goCtx context.Context, req *types.QueryAllRequestRequest) (*types.QueryAllRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var requests []types.Request
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	requestStore := prefix.NewStore(store, types.KeyPrefix(types.RequestKey))

	pageRes, err := query.Paginate(requestStore, req.Pagination, func(key []byte, value []byte) error {
		var request types.Request
		if err := k.cdc.Unmarshal(value, &request); err != nil {
			return err
		}

		requests = append(requests, request)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllRequestResponse{Request: requests, Pagination: pageRes}, nil
}

func (k Keeper) Request(goCtx context.Context, req *types.QueryGetRequestRequest) (*types.QueryGetRequestResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	request, found := k.GetRequest(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetRequestResponse{Request: request}, nil
}
