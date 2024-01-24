package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	keepers := keepertest.NewTest(t)
	tk := keepers.TreasuryKeeper
	ctx := keepers.Ctx
	return keeper.NewMsgServerImpl(*tk), sdk.WrapSDKContext(ctx)
}
