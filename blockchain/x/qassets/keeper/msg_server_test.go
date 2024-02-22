// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/qassets/keeper"
	"github.com/qredo/fusionchain/x/qassets/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keepers := keepertest.NewTest(t)
	qk := keepers.QassetsKeeper
	ctx := keepers.Ctx
	return keeper.NewMsgServerImpl(*qk), sdk.WrapSDKContext(ctx)
}
