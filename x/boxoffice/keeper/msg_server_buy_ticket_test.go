package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice"
	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/testutil"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func setupMockedMsgServer(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context, *gomock.Controller, *testutil.MockNftKeeper) {
	ctrl := gomock.NewController(t)
	nftMock := testutil.NewMockNftKeeper(ctrl)
	k, ctx := keepertest.BoxofficeKeeperWithMocks(t, nftMock)

	boxoffice.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	return server, *k, context, ctrl, nftMock
}

func TestBuyTicket(t *testing.T) {

	msgServer, _, context, ctrl, nft := setupMockedMsgServer(t)
	defer ctrl.Finish()
	nft.ExpectAny(context)
	msgServer.CreateConcert(context, &types.MsgCreateConcert{
		Creator: alice,
		Name:    "Montreux 2023",
		Volume:  "200",
	})
	resp, err := msgServer.BuyTicket(context, &types.MsgBuyTicket{
		Creator:      alice,
		ConcertIndex: "1",
	})

	require.Nil(t, err)
	require.EqualValues(t, types.MsgBuyTicketResponse{}, *resp)
}
