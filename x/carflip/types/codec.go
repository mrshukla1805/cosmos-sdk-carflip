package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCar{}, "carflip/CreateCar", nil)
	cdc.RegisterConcrete(&MsgUpdateCar{}, "carflip/UpdateCar", nil)
	cdc.RegisterConcrete(&MsgDeleteCar{}, "carflip/DeleteCar", nil)
	cdc.RegisterConcrete(&MsgCreateRequest{}, "carflip/CreateRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateRequest{}, "carflip/UpdateRequest", nil)
	cdc.RegisterConcrete(&MsgDeleteRequest{}, "carflip/DeleteRequest", nil)
	cdc.RegisterConcrete(&MsgUpdateRequestStatus{}, "carflip/UpdateRequestStatus", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCar{},
		&MsgUpdateCar{},
		&MsgDeleteCar{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateRequest{},
		&MsgUpdateRequest{},
		&MsgDeleteRequest{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateRequestStatus{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
