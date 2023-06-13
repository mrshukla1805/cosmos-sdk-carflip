package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUpdateRequestStatus = "update_request_status"

var _ sdk.Msg = &MsgUpdateRequestStatus{}

func NewMsgUpdateRequestStatus(creator string, status string, id uint64) *MsgUpdateRequestStatus {
	return &MsgUpdateRequestStatus{
		Creator: creator,
		Status:  status,
		Id:      id,
	}
}

func (msg *MsgUpdateRequestStatus) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRequestStatus) Type() string {
	return TypeMsgUpdateRequestStatus
}

func (msg *MsgUpdateRequestStatus) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRequestStatus) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRequestStatus) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
