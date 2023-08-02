package keeper

import (
	"context"
	"encoding/binary"
	"mapofmap/x/mapofmap/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Addbatch(goCtx context.Context, msg *types.MsgAddbatch) (*types.MsgAddbatchResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	chainid := msg.Chainid
	// batch := msg.Batch
	var sampleBatchId uint64 = 1

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix("/Batch/"+chainid+"/")) //types.PostKey)) // types.PostKey = "Post/value/"
	appendedValue := k.cdc.MustMarshal(msg)
	store.Set(GetBytesFromUint64(sampleBatchId), appendedValue)

	return &types.MsgAddbatchResponse{}, nil
}

func GetBytesFromUint64(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
