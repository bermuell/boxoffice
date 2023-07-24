package keeper

import (
	"context"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) StoredConcertAll(goCtx context.Context, req *types.QueryAllStoredConcertRequest) (*types.QueryAllStoredConcertResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var storedConcerts []types.StoredConcert
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	storedConcertStore := prefix.NewStore(store, types.KeyPrefix(types.StoredConcertKeyPrefix))

	pageRes, err := query.Paginate(storedConcertStore, req.Pagination, func(key []byte, value []byte) error {
		var storedConcert types.StoredConcert
		if err := k.cdc.Unmarshal(value, &storedConcert); err != nil {
			return err
		}

		storedConcerts = append(storedConcerts, storedConcert)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllStoredConcertResponse{StoredConcert: storedConcerts, Pagination: pageRes}, nil
}

func (k Keeper) StoredConcert(goCtx context.Context, req *types.QueryGetStoredConcertRequest) (*types.QueryGetStoredConcertResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetStoredConcert(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetStoredConcertResponse{StoredConcert: val}, nil
}
