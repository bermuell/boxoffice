package keeper

import (
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetStoredConcert set a specific storedConcert in the store from its index
func (k Keeper) SetStoredConcert(ctx sdk.Context, storedConcert types.StoredConcert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredConcertKeyPrefix))
	b := k.cdc.MustMarshal(&storedConcert)
	store.Set(types.StoredConcertKey(
		storedConcert.Index,
	), b)
}

// GetStoredConcert returns a storedConcert from its index
func (k Keeper) GetStoredConcert(
	ctx sdk.Context,
	index string,

) (val types.StoredConcert, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredConcertKeyPrefix))

	b := store.Get(types.StoredConcertKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStoredConcert removes a storedConcert from the store
func (k Keeper) RemoveStoredConcert(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredConcertKeyPrefix))
	store.Delete(types.StoredConcertKey(
		index,
	))
}

// GetAllStoredConcert returns all storedConcert
func (k Keeper) GetAllStoredConcert(ctx sdk.Context) (list []types.StoredConcert) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StoredConcertKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.StoredConcert
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
