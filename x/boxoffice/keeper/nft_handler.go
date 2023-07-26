package keeper

import (
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/nft"
)

func (k *Keeper) SetNftPool(ctx sdk.Context, storedConcert *types.StoredConcert) (err error) {
	nftClassId := storedConcert.Name + storedConcert.Index

	_, found := k.nft.GetClass(ctx, nftClassId)
	if !found {
		nftClass := nft.Class{
			Id:   nftClassId,
			Name: storedConcert.Name,
		}
		err = k.nft.SaveClass(ctx, nftClass)
	}
	return
}
