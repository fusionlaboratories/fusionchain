// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
)

func SimulateMsgAppendChildWorkspace(
	types.AccountKeeper,
	types.BankKeeper,
	keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAppendChildWorkspace{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AppendChildWorkspace simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "AppendChildWorkspace simulation not implemented"), nil, nil
	}
}
