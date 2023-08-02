package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateExecutionlayers = "create_executionlayers"
	TypeMsgUpdateExecutionlayers = "update_executionlayers"
	TypeMsgDeleteExecutionlayers = "delete_executionlayers"
)

var _ sdk.Msg = &MsgCreateExecutionlayers{}

func NewMsgCreateExecutionlayers(
	creator string,
	index string,
	name string,
	info string,
	chainid string,

) *MsgCreateExecutionlayers {
	return &MsgCreateExecutionlayers{
		Creator: creator,
		Index:   index,
		Name:    name,
		Info:    info,
		Chainid: chainid,
	}
}

func (msg *MsgCreateExecutionlayers) Route() string {
	return RouterKey
}

func (msg *MsgCreateExecutionlayers) Type() string {
	return TypeMsgCreateExecutionlayers
}

func (msg *MsgCreateExecutionlayers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateExecutionlayers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateExecutionlayers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateExecutionlayers{}

func NewMsgUpdateExecutionlayers(
	creator string,
	index string,
	name string,
	info string,
	chainid string,

) *MsgUpdateExecutionlayers {
	return &MsgUpdateExecutionlayers{
		Creator: creator,
		Index:   index,
		Name:    name,
		Info:    info,
		Chainid: chainid,
	}
}

func (msg *MsgUpdateExecutionlayers) Route() string {
	return RouterKey
}

func (msg *MsgUpdateExecutionlayers) Type() string {
	return TypeMsgUpdateExecutionlayers
}

func (msg *MsgUpdateExecutionlayers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateExecutionlayers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateExecutionlayers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteExecutionlayers{}

func NewMsgDeleteExecutionlayers(
	creator string,
	index string,

) *MsgDeleteExecutionlayers {
	return &MsgDeleteExecutionlayers{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteExecutionlayers) Route() string {
	return RouterKey
}

func (msg *MsgDeleteExecutionlayers) Type() string {
	return TypeMsgDeleteExecutionlayers
}

func (msg *MsgDeleteExecutionlayers) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteExecutionlayers) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteExecutionlayers) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
