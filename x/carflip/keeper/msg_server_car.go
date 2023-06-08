package keeper

import (
	"context"
	"fmt"

	"carflip/x/carflip/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCar(goCtx context.Context, msg *types.MsgCreateCar) (*types.MsgCreateCarResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var car = types.Car{
		Creator: msg.Creator,
		Name:    msg.Name,
		Number:  msg.Number,
		Price:   msg.Price,
	}

	id := k.AppendCar(
		ctx,
		car,
	)

	return &types.MsgCreateCarResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateCar(goCtx context.Context, msg *types.MsgUpdateCar) (*types.MsgUpdateCarResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var car = types.Car{
		Creator: msg.Creator,
		Id:      msg.Id,
		Name:    msg.Name,
		Number:  msg.Number,
		Price:   msg.Price,
	}

	// Checks that the element exists
	val, found := k.GetCar(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetCar(ctx, car)

	return &types.MsgUpdateCarResponse{}, nil
}

func (k msgServer) DeleteCar(goCtx context.Context, msg *types.MsgDeleteCar) (*types.MsgDeleteCarResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetCar(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveCar(ctx, msg.Id)

	return &types.MsgDeleteCarResponse{}, nil
}
