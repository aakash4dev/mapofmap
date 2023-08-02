package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateExecutionlayers{}, "mapofmap/CreateExecutionlayers", nil)
	cdc.RegisterConcrete(&MsgUpdateExecutionlayers{}, "mapofmap/UpdateExecutionlayers", nil)
	cdc.RegisterConcrete(&MsgDeleteExecutionlayers{}, "mapofmap/DeleteExecutionlayers", nil)
	cdc.RegisterConcrete(&MsgAddbatch{}, "mapofmap/Addbatch", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateExecutionlayers{},
		&MsgUpdateExecutionlayers{},
		&MsgDeleteExecutionlayers{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddbatch{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
