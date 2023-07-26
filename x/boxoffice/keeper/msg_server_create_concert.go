package keeper

import (
	"context"
	"strconv"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConcert(goCtx context.Context, msg *types.MsgCreateConcert) (*types.MsgCreateConcertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	systemInfo, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("No SystemInfo found")
	}
	volume, err := strconv.ParseUint(msg.Volume, 10, 64)
	if err != nil {
		return nil, err
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)
	storedConcert := types.StoredConcert{
		Index:  newIndex,
		Name:   msg.Name,
		Volume: volume,
	}

	//TODO: add some validation for concert: storedConcert.Validate()
	k.Keeper.SetNftPool(ctx, &storedConcert)
	k.Keeper.SetStoredConcert(ctx, storedConcert)
	systemInfo.NextId++
	k.Keeper.SetSystemInfo(ctx, systemInfo)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(types.ConcertCreatedEventType,
			sdk.NewAttribute(types.ConcertCreatedEventCreator, msg.Creator),
			sdk.NewAttribute(types.ConcertCreatedEventConcertIndex, newIndex),
		),
	)

	return &types.MsgCreateConcertResponse{}, nil
}
