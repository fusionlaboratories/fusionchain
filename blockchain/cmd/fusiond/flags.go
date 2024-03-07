// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1

package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/qredo/fusionchain/version"
)

const flagLong = "long"

func init() {
	infoCmd.Flags().Bool(flagLong, false, "Print full information")
}

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Print version info",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println(version.Version())
	},
}
