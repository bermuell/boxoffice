package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice"
	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerWithOneConcert(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.BoxofficeKeeper(t)

	boxoffice.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateConcert(context, &types.MsgCreateConcert{
		Creator: alice,
		Name:    "Montreux 2023",
		Volume:  "200",
	})
	return server, *k, context
}

func TestBuyTicket(t *testing.T) {
	msgServer, _, context := setupMsgServerWithOneConcert(t)
	resp, err := msgServer.BuyTicket(context, &types.MsgBuyTicket{
		Creator:      alice,
		ConcertIndex: "1",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgBuyTicketResponse{}, *resp)
}
