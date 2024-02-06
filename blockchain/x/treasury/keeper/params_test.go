package keeper_test

import (
	"testing"

	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	keepers := keepertest.NewTest(t)
	tk := keepers.TreasuryKeeper
	ctx := keepers.Ctx
	params := types.DefaultParams()

	tk.SetParams(ctx, params)

	require.EqualValues(t, params, tk.GetParams(ctx))
}
