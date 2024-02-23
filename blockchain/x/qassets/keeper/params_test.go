// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"testing"

	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/qassets/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	keepers := keepertest.NewTest(t)
	qk := keepers.QassetsKeeper
	ctx := keepers.Ctx
	params := types.DefaultParams()

	qk.SetParams(ctx, params)

	require.EqualValues(t, params, qk.GetParams(ctx))
}
