package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "mapofmap/testutil/keeper"
	"mapofmap/testutil/nullify"
	"mapofmap/x/mapofmap/keeper"
	"mapofmap/x/mapofmap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNExecutionlayers(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Executionlayers {
	items := make([]types.Executionlayers, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetExecutionlayers(ctx, items[i])
	}
	return items
}

func TestExecutionlayersGet(t *testing.T) {
	keeper, ctx := keepertest.MapofmapKeeper(t)
	items := createNExecutionlayers(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetExecutionlayers(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestExecutionlayersRemove(t *testing.T) {
	keeper, ctx := keepertest.MapofmapKeeper(t)
	items := createNExecutionlayers(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveExecutionlayers(ctx,
			item.Index,
		)
		_, found := keeper.GetExecutionlayers(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestExecutionlayersGetAll(t *testing.T) {
	keeper, ctx := keepertest.MapofmapKeeper(t)
	items := createNExecutionlayers(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllExecutionlayers(ctx)),
	)
}
