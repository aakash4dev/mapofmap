package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"mapofmap/x/mapofmap/types"
)

func (k msgServer) CreateExecutionlayers(goCtx context.Context, msg *types.MsgCreateExecutionlayers) (*types.MsgCreateExecutionlayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetExecutionlayers(
		ctx,
		msg.Index,
	)
	if isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	var executionlayers = types.Executionlayers{
		Creator: msg.Creator,
		Index:   msg.Index,
		Name:    msg.Name,
		Info:    msg.Info,
		Chainid: msg.Chainid,
	}

	k.SetExecutionlayers(
		ctx,
		executionlayers,
	)
	return &types.MsgCreateExecutionlayersResponse{}, nil
}

func (k msgServer) UpdateExecutionlayers(goCtx context.Context, msg *types.MsgUpdateExecutionlayers) (*types.MsgUpdateExecutionlayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExecutionlayers(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	var executionlayers = types.Executionlayers{
		Creator: msg.Creator,
		Index:   msg.Index,
		Name:    msg.Name,
		Info:    msg.Info,
		Chainid: msg.Chainid,
	}

	k.SetExecutionlayers(ctx, executionlayers)

	return &types.MsgUpdateExecutionlayersResponse{}, nil
}

func (k msgServer) DeleteExecutionlayers(goCtx context.Context, msg *types.MsgDeleteExecutionlayers) (*types.MsgDeleteExecutionlayersResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetExecutionlayers(
		ctx,
		msg.Index,
	)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "index not set")
	}

	// Checks if the the msg creator is the same as the current owner
	if msg.Creator != valFound.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveExecutionlayers(
		ctx,
		msg.Index,
	)

	return &types.MsgDeleteExecutionlayersResponse{}, nil
}
