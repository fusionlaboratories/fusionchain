// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package common

import (
	"fmt"
)

var (
	Version     = "v0.1.0"                                       // Semantic version
	FullVersion = fmt.Sprintf("%s-%v", Version, CommitHash[0:8]) // Full version with commit hash
)
