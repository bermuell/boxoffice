package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/bermuell/boxoffice/testutil/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BoxofficeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
