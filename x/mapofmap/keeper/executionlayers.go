package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"mapofmap/x/mapofmap/types"
)

// SetExecutionlayers set a specific executionlayers in the store from its index
func (k Keeper) SetExecutionlayers(ctx sdk.Context, executionlayers types.Executionlayers) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionlayersKeyPrefix))
	b := k.cdc.MustMarshal(&executionlayers)
	store.Set(types.ExecutionlayersKey(
		executionlayers.Index,
	), b)
}

// GetExecutionlayers returns a executionlayers from its index
func (k Keeper) GetExecutionlayers(
	ctx sdk.Context,
	index string,

) (val types.Executionlayers, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionlayersKeyPrefix))

	b := store.Get(types.ExecutionlayersKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveExecutionlayers removes a executionlayers from the store
func (k Keeper) RemoveExecutionlayers(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionlayersKeyPrefix))
	store.Delete(types.ExecutionlayersKey(
		index,
	))
}

// GetAllExecutionlayers returns all executionlayers
func (k Keeper) GetAllExecutionlayers(ctx sdk.Context) (list []types.Executionlayers) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionlayersKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Executionlayers
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
