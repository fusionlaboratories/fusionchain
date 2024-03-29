// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package treasury

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)

	for i := range genState.KeyRequests {
		k.SetKeyRequest(ctx, &genState.KeyRequests[i])
	}
	k.SetKeyRequestCount(ctx, uint64(len(genState.KeyRequests)))

	for i := range genState.Keys {
		k.SetKey(ctx, &genState.Keys[i])
	}
	k.SetKeyCount(ctx, uint64(len(genState.Keys)))

	for i := range genState.SignRequests {
		k.SetSignRequest(ctx, &genState.SignRequests[i])
	}
	k.SetSignRequestCount(ctx, uint64(len(genState.SignRequests)))
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
