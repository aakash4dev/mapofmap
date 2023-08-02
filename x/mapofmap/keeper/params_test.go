package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "mapofmap/testutil/keeper"
	"mapofmap/x/mapofmap/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MapofmapKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
