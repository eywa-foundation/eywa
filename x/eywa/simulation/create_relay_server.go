package simulation

import (
	"math/rand"

	"eywa/x/eywa/keeper"
	"eywa/x/eywa/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgCreateRelayServer(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgCreateRelayServer{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the CreateRelayServer simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "CreateRelayServer simulation not implemented"), nil, nil
	}
}
