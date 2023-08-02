package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mapofmap/x/mapofmap/types"
)

func (k Keeper) Getbatch(goCtx context.Context, req *types.QueryGetbatchRequest) (*types.QueryGetbatchResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainid := req.Chainid

	// batchnumber := req.Batchnumber
	var batchnumber uint64 = 1; // take from user. currently taking string

	proof, found := k.GetbatchById(ctx,chainid,batchnumber)
	if found == false {
		return &types.QueryGetbatchResponse{
			Proof: "",
		}, nil
	}else{
		return &types.QueryGetbatchResponse{
			Proof: proof.Batch,
		}, nil
	}
}