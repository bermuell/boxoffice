package boxoffice

import (
	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set if defined
	if genState.SystemInfo != nil {
		k.SetSystemInfo(ctx, *genState.SystemInfo)
	}
	// Set all the storedConcert
	for _, elem := range genState.StoredConcertList {
		k.SetStoredConcert(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// Get all systemInfo
	systemInfo, found := k.GetSystemInfo(ctx)
	if found {
		genesis.SystemInfo = &systemInfo
	}
	genesis.StoredConcertList = k.GetAllStoredConcert(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
