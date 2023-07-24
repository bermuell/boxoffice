package simulation

import (
	"math/rand"

	"github.com/bermuell/boxoffice/x/boxoffice/keeper"
	"github.com/bermuell/boxoffice/x/boxoffice/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateConcert(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateConcert{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateConcert simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateConcert simulation not implemented"), nil, nil
	}
}
