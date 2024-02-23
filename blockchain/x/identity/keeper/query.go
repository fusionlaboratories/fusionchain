// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper

import (
	"github.com/qredo/fusionchain/x/identity/types"
)

var _ types.QueryServer = Keeper{}
