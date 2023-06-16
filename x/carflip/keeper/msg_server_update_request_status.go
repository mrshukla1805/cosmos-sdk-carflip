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
	if msg.Creator != val.RequestTo {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	if val.Status == "complete" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "request already completed cannot change the status now")
	}

	if val.Status == "rejected" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "request already rejected cannot change the status now")
	}

	if msg.Status == "rejected" {
		val.Status = "rejected"
		//Create a MsgUpdateRequest
		tempMsg := types.MsgUpdateRequest{
			Creator:   val.Creator,
			Id:        val.Id,
			RequestTo: val.RequestTo,
			Status:    val.Status,
			Number:    val.Number,
		}
		k.UpdateRequest(ctx, &tempMsg)
		return &types.MsgUpdateRequestStatusResponse{}, nil
	}

	if msg.Status == "complete" {
		// Here we check if we have the required number of cars and then accordingly complete the transaction or give an error
		//First we have to get the car data of the user
		query := types.QueryGetCarByUserRequest{
			Creator: val.RequestTo,
		}
		userGivingCar, _ := k.GetCarByUser(ctx, &query)
		//Then check
		if userGivingCar.Car.Number < val.Number {
			// Then raise error
			return nil, sdkerrors.Wrap(sdkerrors.ErrNotSupported, "The user does not have as many cars as the request requires")
		} else {
			// Decrease the Number of Cars of the user who got the request
			userGivingCar.Car.Number -= val.Number
			// Now save this car
			updateCarMsg := types.MsgUpdateCar{
				Creator: userGivingCar.Car.Creator,
				Id:      userGivingCar.Car.Id,
				Name:    userGivingCar.Car.Name,
				Number:  userGivingCar.Car.Number,
				Price:   userGivingCar.Car.Price,
			}
			k.UpdateCar(ctx, &updateCarMsg)
			//Now increase the count of car of the user who created the request
			query := types.QueryGetCarByUserRequest{
				Creator: val.Creator,
			}
			userTakingCar, _ := k.GetCarByUser(ctx, &query)
			userTakingCar.Car.Number += val.Number
			// Now save this car
			updateCarMsg = types.MsgUpdateCar{
				Creator: userTakingCar.Car.Creator,
				Id:      userTakingCar.Car.Id,
				Name:    userTakingCar.Car.Name,
				Number:  userTakingCar.Car.Number,
				Price:   userTakingCar.Car.Price,
			}
			k.UpdateCar(ctx, &updateCarMsg)

		}

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
