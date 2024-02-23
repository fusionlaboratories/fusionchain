// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package main

import (
	"github.com/qredo/fusionchain/cmd/scaffolder/cmd"
)

// Command scaffolder helps in generating code needed to add new messages and
// queries to existing Cosmos SDK modules.
//
// We are developing this to be a simpler version of Ignite compatible with
// Fusion and focused on our needs.
func main() {
	cmd.Execute()
}
