package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"mapofmap/testutil/sample"
)

func TestMsgCreateExecutionlayers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateExecutionlayers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateExecutionlayers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateExecutionlayers{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateExecutionlayers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateExecutionlayers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateExecutionlayers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateExecutionlayers{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteExecutionlayers_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteExecutionlayers
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteExecutionlayers{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteExecutionlayers{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
