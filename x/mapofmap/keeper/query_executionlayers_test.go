package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "mapofmap/testutil/keeper"
	"mapofmap/testutil/nullify"
	"mapofmap/x/mapofmap/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestExecutionlayersQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MapofmapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExecutionlayers(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetExecutionlayersRequest
		response *types.QueryGetExecutionlayersResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetExecutionlayersRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetExecutionlayersResponse{Executionlayers: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetExecutionlayersRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetExecutionlayersResponse{Executionlayers: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetExecutionlayersRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Executionlayers(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestExecutionlayersQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MapofmapKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNExecutionlayers(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllExecutionlayersRequest {
		return &types.QueryAllExecutionlayersRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExecutionlayersAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Executionlayers), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Executionlayers),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.ExecutionlayersAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Executionlayers), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Executionlayers),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.ExecutionlayersAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Executionlayers),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.ExecutionlayersAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
