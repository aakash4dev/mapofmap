package mapofmap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "mapofmap/testutil/keeper"
	"mapofmap/testutil/nullify"
	"mapofmap/x/mapofmap"
	"mapofmap/x/mapofmap/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		ExecutionlayersList: []types.Executionlayers{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MapofmapKeeper(t)
	mapofmap.InitGenesis(ctx, *k, genesisState)
	got := mapofmap.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.ExecutionlayersList, got.ExecutionlayersList)
	// this line is used by starport scaffolding # genesis/test/assert
}
