package keeper

import (
	"context"
	"fmt"

	"carflip/x/carflip/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateRequestStatus(goCtx context.Context, msg *types.MsgUpdateRequestStatus) (*types.MsgUpdateRequestStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx
	val, found := k.GetRequest(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if val.Status == "complete" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "request already completed cannot change the status now")
	}

	if msg.Status == "complete" {

	}

	var request = types.Request{
		Creator:   val.Creator,
		Id:        val.Id,
		RequestTo: val.RequestTo,
		Status:    msg.Status,
		Number:    val.Number,
	}

	k.SetRequest(ctx, request)

	return &types.MsgUpdateRequestStatusResponse{}, nil
}
