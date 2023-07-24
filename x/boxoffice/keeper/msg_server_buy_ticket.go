package keeper

import (
	"context"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) BuyTicket(goCtx context.Context, msg *types.MsgBuyTicket) (*types.MsgBuyTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgBuyTicketResponse{}, nil
}
