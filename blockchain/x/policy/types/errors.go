// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/policy module sentinel errors
var (
	ErrPolicyValidation = sdkerrors.Register(ModuleName, 1200, "policy validation required")
)
