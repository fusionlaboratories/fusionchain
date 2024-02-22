// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/qassets/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	keepers := keepertest.NewTest(t)
	qk := keepers.QassetsKeeper
	ctx := keepers.Ctx
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	qk.SetParams(ctx, params)

	response, err := qk.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
