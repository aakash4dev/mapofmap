package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	keepertest "mapofmap/testutil/keeper"
	"mapofmap/x/mapofmap/keeper"
	"mapofmap/x/mapofmap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestExecutionlayersMsgServerCreate(t *testing.T) {
	k, ctx := keepertest.MapofmapKeeper(t)
	srv := keeper.NewMsgServerImpl(*k)
	wctx := sdk.WrapSDKContext(ctx)
	creator := "A"
	for i := 0; i < 5; i++ {
		expected := &types.MsgCreateExecutionlayers{Creator: creator,
			Index: strconv.Itoa(i),
		}
		_, err := srv.CreateExecutionlayers(wctx, expected)
		require.NoError(t, err)
		rst, found := k.GetExecutionlayers(ctx,
			expected.Index,
		)
		require.True(t, found)
		require.Equal(t, expected.Creator, rst.Creator)
	}
}

func TestExecutionlayersMsgServerUpdate(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgUpdateExecutionlayers
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgUpdateExecutionlayers{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgUpdateExecutionlayers{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgUpdateExecutionlayers{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MapofmapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)
			expected := &types.MsgCreateExecutionlayers{Creator: creator,
				Index: strconv.Itoa(0),
			}
			_, err := srv.CreateExecutionlayers(wctx, expected)
			require.NoError(t, err)

			_, err = srv.UpdateExecutionlayers(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				rst, found := k.GetExecutionlayers(ctx,
					expected.Index,
				)
				require.True(t, found)
				require.Equal(t, expected.Creator, rst.Creator)
			}
		})
	}
}

func TestExecutionlayersMsgServerDelete(t *testing.T) {
	creator := "A"

	tests := []struct {
		desc    string
		request *types.MsgDeleteExecutionlayers
		err     error
	}{
		{
			desc: "Completed",
			request: &types.MsgDeleteExecutionlayers{Creator: creator,
				Index: strconv.Itoa(0),
			},
		},
		{
			desc: "Unauthorized",
			request: &types.MsgDeleteExecutionlayers{Creator: "B",
				Index: strconv.Itoa(0),
			},
			err: sdkerrors.ErrUnauthorized,
		},
		{
			desc: "KeyNotFound",
			request: &types.MsgDeleteExecutionlayers{Creator: creator,
				Index: strconv.Itoa(100000),
			},
			err: sdkerrors.ErrKeyNotFound,
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			k, ctx := keepertest.MapofmapKeeper(t)
			srv := keeper.NewMsgServerImpl(*k)
			wctx := sdk.WrapSDKContext(ctx)

			_, err := srv.CreateExecutionlayers(wctx, &types.MsgCreateExecutionlayers{Creator: creator,
				Index: strconv.Itoa(0),
			})
			require.NoError(t, err)
			_, err = srv.DeleteExecutionlayers(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				_, found := k.GetExecutionlayers(ctx,
					tc.request.Index,
				)
				require.False(t, found)
			}
		})
	}
}
