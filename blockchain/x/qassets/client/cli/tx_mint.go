// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/qredo/fusionchain/x/qassets/types"
	treasurytypes "github.com/qredo/fusionchain/x/treasury/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint [workspace-addr] [wallet-type] [is-token] [token-name] [token-contract-addr] [amount]",
		Short: "Broadcast message mint",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var walletType treasurytypes.WalletType
			switch strings.ToLower(args[1]) {
			case "ethereum":
				walletType = treasurytypes.WalletType_WALLET_TYPE_ETH
			default:
				return fmt.Errorf("invalid wallet type '%s'", args[1])
			}

			isToken, err := strconv.ParseBool(args[2])
			if err != nil {
				return err
			}
			amount, err := strconv.ParseUint(args[5], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
				clientCtx.GetFromAddress().String(),
				args[0],
				walletType,
				isToken,
				args[3],
				args[4],
				amount,
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
