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

func (k Keeper) CarAll(goCtx context.Context, req *types.QueryAllCarRequest) (*types.QueryAllCarResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var cars []types.Car
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	carStore := prefix.NewStore(store, types.KeyPrefix(types.CarKey))

	pageRes, err := query.Paginate(carStore, req.Pagination, func(key []byte, value []byte) error {
		var car types.Car
		if err := k.cdc.Unmarshal(value, &car); err != nil {
			return err
		}

		cars = append(cars, car)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCarResponse{Car: cars, Pagination: pageRes}, nil
}

func (k Keeper) Car(goCtx context.Context, req *types.QueryGetCarRequest) (*types.QueryGetCarResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	car, found := k.GetCar(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCarResponse{Car: car}, nil
}
