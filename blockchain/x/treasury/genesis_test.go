// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package treasury_test

import (
	"testing"

	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/testutil/nullify"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	keepers := keepertest.NewTest(t)
	tk := keepers.TreasuryKeeper
	ctx := keepers.Ctx
	treasury.InitGenesis(ctx, *tk, genesisState)
	got := treasury.ExportGenesis(ctx, *tk)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
