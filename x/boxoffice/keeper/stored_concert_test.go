package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/testutil/nullify"
	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNStoredConcert(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.StoredConcert {
	items := make([]types.StoredConcert, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetStoredConcert(ctx, items[i])
	}
	return items
}

func TestStoredConcertGet(t *testing.T) {
	keeper, ctx := keepertest.BoxofficeKeeper(t)
	items := createNStoredConcert(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetStoredConcert(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestStoredConcertRemove(t *testing.T) {
	keeper, ctx := keepertest.BoxofficeKeeper(t)
	items := createNStoredConcert(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveStoredConcert(ctx,
			item.Index,
		)
		_, found := keeper.GetStoredConcert(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestStoredConcertGetAll(t *testing.T) {
	keeper, ctx := keepertest.BoxofficeKeeper(t)
	items := createNStoredConcert(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllStoredConcert(ctx)),
	)
}
