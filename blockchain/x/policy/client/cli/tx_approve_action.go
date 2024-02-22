// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/qredo/fusionchain/x/policy/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdApproveAction() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "approve-action [action_type] [action_id]",
		Short: "Broadcast message ApproveAction",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			actionID, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgApproveAction(
				clientCtx.GetFromAddress().String(),
				args[0],
				actionID,
				nil,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
