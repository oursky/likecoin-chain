package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/likecoin/likechain/x/likenft/keeper"
	"github.com/likecoin/likechain/x/likenft/types"
)

func SimulateMsgSellNFT(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgSellNFT{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the SellNFT simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "SellNFT simulation not implemented"), nil, nil
	}
}
