// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper

import (
	"testing"

	cbftdb "github.com/cometbft/cometbft-db"
	"github.com/cometbft/cometbft/libs/log"
	cbftproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	typesparams "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/qredo/fusionchain/x/policy/keeper"
	"github.com/qredo/fusionchain/x/policy/types"
	"github.com/stretchr/testify/require"
)

func PolicyKeeper(t testing.TB, db cbftdb.DB, stateStore storetypes.CommitMultiStore) (*keeper.Keeper, sdk.Context) {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	memStoreKey := storetypes.NewMemoryStoreKey(types.MemStoreKey)

	stateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, db)
	stateStore.MountStoreWithDB(memStoreKey, storetypes.StoreTypeMemory, nil)
	require.NoError(t, stateStore.LoadLatestVersion())

	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)

	ctx := sdk.NewContext(stateStore, cbftproto.Header{}, false, log.NewNopLogger())

	paramsSubspace := typesparams.NewSubspace(cdc,
		types.Amino,
		storeKey,
		memStoreKey,
		"PolicyParams",
	)
	k := keeper.NewKeeper(
		cdc,
		storeKey,
		memStoreKey,
		paramsSubspace,
	)

	// Initialize params
	k.SetParams(ctx, types.DefaultParams())

	return k, ctx
}
