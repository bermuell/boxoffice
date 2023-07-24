package boxoffice_test

import (
	"testing"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/testutil/nullify"
	"github.com/bermuell/boxoffice/x/boxoffice"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		SystemInfo: types.SystemInfo{
			NextId: 4,
		},
		StoredConcertList: []types.StoredConcert{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BoxofficeKeeper(t)
	boxoffice.InitGenesis(ctx, *k, genesisState)
	got := boxoffice.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.SystemInfo, got.SystemInfo)
	require.ElementsMatch(t, genesisState.StoredConcertList, got.StoredConcertList)
	// this line is used by starport scaffolding # genesis/test/assert
}
