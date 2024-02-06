// Copyright 2023 Qredo Ltd.
// This file is part of the Fusion library.
//
// The Fusion library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Fusion library. If not, see https://github.com/qredo/fusionchain/blob/main/LICENSE
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
