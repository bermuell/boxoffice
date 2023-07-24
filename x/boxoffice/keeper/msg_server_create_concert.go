package keeper

import (
	"context"

	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateConcert(goCtx context.Context, msg *types.MsgCreateConcert) (*types.MsgCreateConcertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateConcertResponse{}, nil
}
