package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mapofmap/x/mapofmap/types"
)

func (k Keeper) ExecutionlayersAll(goCtx context.Context, req *types.QueryAllExecutionlayersRequest) (*types.QueryAllExecutionlayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var executionlayerss []types.Executionlayers
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	executionlayersStore := prefix.NewStore(store, types.KeyPrefix(types.ExecutionlayersKeyPrefix))

	pageRes, err := query.Paginate(executionlayersStore, req.Pagination, func(key []byte, value []byte) error {
		var executionlayers types.Executionlayers
		if err := k.cdc.Unmarshal(value, &executionlayers); err != nil {
			return err
		}

		executionlayerss = append(executionlayerss, executionlayers)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllExecutionlayersResponse{Executionlayers: executionlayerss, Pagination: pageRes}, nil
}

func (k Keeper) Executionlayers(goCtx context.Context, req *types.QueryGetExecutionlayersRequest) (*types.QueryGetExecutionlayersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetExecutionlayers(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetExecutionlayersResponse{Executionlayers: val}, nil
}
