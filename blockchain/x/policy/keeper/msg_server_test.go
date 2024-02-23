// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/policy/keeper"
	"github.com/qredo/fusionchain/x/policy/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keepers := keepertest.NewTest(t)
	pk := keepers.PolicyKeeper
	ctx := keepers.Ctx
	return keeper.NewMsgServerImpl(*pk), sdk.WrapSDKContext(ctx)
}
