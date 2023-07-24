package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/testutil/nullify"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestStoredConcertQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.BoxofficeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStoredConcert(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetStoredConcertRequest
		response *types.QueryGetStoredConcertResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetStoredConcertRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetStoredConcertResponse{StoredConcert: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetStoredConcertRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetStoredConcertResponse{StoredConcert: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetStoredConcertRequest{
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
			response, err := keeper.StoredConcert(wctx, tc.request)
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

func TestStoredConcertQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.BoxofficeKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNStoredConcert(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllStoredConcertRequest {
		return &types.QueryAllStoredConcertRequest{
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
			resp, err := keeper.StoredConcertAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StoredConcert), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StoredConcert),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.StoredConcertAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.StoredConcert), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.StoredConcert),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.StoredConcertAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.StoredConcert),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.StoredConcertAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
