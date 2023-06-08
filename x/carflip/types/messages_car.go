package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateCar = "create_car"
	TypeMsgUpdateCar = "update_car"
	TypeMsgDeleteCar = "delete_car"
)

var _ sdk.Msg = &MsgCreateCar{}

func NewMsgCreateCar(creator string, name string, number int32, price sdk.Coin) *MsgCreateCar {
	return &MsgCreateCar{
		Creator: creator,
		Name:    name,
		Number:  number,
		Price:   price,
	}
}

func (msg *MsgCreateCar) Route() string {
	return RouterKey
}

func (msg *MsgCreateCar) Type() string {
	return TypeMsgCreateCar
}

func (msg *MsgCreateCar) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCar) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCar) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCar{}

func NewMsgUpdateCar(creator string, id uint64, name string, number int32, price sdk.Coin) *MsgUpdateCar {
	return &MsgUpdateCar{
		Id:      id,
		Creator: creator,
		Name:    name,
		Number:  number,
		Price:   price,
	}
}

func (msg *MsgUpdateCar) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCar) Type() string {
	return TypeMsgUpdateCar
}

func (msg *MsgUpdateCar) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCar) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCar) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCar{}

func NewMsgDeleteCar(creator string, id uint64) *MsgDeleteCar {
	return &MsgDeleteCar{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteCar) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCar) Type() string {
	return TypeMsgDeleteCar
}

func (msg *MsgDeleteCar) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCar) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCar) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
