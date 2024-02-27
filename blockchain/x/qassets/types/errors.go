// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/qassets module sentinel errors
var (
	ErrSample = sdkerrors.Register(ModuleName, 1300, "sample error")
)
