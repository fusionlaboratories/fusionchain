// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/identity module sentinel errors
var (
	ErrEmptyDesc       = sdkerrors.Register(ModuleName, 1100, "description is empty")
	ErrDuplicateOwners = sdkerrors.Register(ModuleName, 1101, "duplicate owners")
)
