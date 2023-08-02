package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddbatch = "addbatch"

var _ sdk.Msg = &MsgAddbatch{}

func NewMsgAddbatch(creator string, chainid string, batch string) *MsgAddbatch {
	return &MsgAddbatch{
		Creator: creator,
		Chainid: chainid,
		Batch:   batch,
	}
}

func (msg *MsgAddbatch) Route() string {
	return RouterKey
}

func (msg *MsgAddbatch) Type() string {
	return TypeMsgAddbatch
}

func (msg *MsgAddbatch) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddbatch) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddbatch) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
