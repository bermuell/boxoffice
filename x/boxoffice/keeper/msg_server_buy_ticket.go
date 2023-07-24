package keeper

import (
	"context"
	"strconv"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) BuyTicket(goCtx context.Context, msg *types.MsgBuyTicket) (*types.MsgBuyTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.Keeper.GetSystemInfo(ctx)
	if !found {
		panic("SystemInfo not found")
	}

	concertIdx, err := strconv.ParseUint(msg.ConcertIndex, 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(types.ErrInvalidConcertIndex, "concert index not parsable: %s", err)
	}

	if concertIdx < types.DefaultIndex {
		return nil, sdkerrors.Wrapf(types.ErrInvalidConcertIndex, "number needs to be bigger than %d", types.DefaultIndex)
	}

	concert, found := k.Keeper.GetStoredConcert(ctx, msg.ConcertIndex)
	if !found {
		return nil, sdkerrors.Wrapf(types.ErrConcertNotFound, "%s", msg.ConcertIndex)
	}

	if concert.Volume <= 0 {
		return nil, sdkerrors.Wrapf(types.ErrConcertSoldOut, "")
	}

	concert.Volume--
	k.Keeper.SetStoredConcert(ctx, concert)

	_ = ctx

	return &types.MsgBuyTicketResponse{}, nil
}
