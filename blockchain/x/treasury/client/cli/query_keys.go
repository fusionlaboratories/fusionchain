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
package cli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/qredo/fusionchain/x/treasury/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdKeys() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keys [wallet-type] [workspace-addr] [id]",
		Short: "Query Keys, optionally by wallet type, workspace address and key id",
		Args:  cobra.MaximumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			var walletType types.WalletType
			if len(args) > 0 {
				switch strings.ToLower(args[0]) {
				case "ethereum":
					walletType = types.WalletType_WALLET_TYPE_ETH
				case "sepolia":
					walletType = types.WalletType_WALLET_TYPE_ETH_SEPOLIA
				case "all":
					walletType = types.WalletType_WALLET_TYPE_UNSPECIFIED
				default:
					fmt.Printf("invalid wallet type '%s', defaulting to 'all'", args[0])
					walletType = types.WalletType_WALLET_TYPE_UNSPECIFIED
				}
			}

			var id uint64
			if len(args) > 2 {
				id, err = strconv.ParseUint(args[2], 10, 64)
				if err != nil {
					return err
				}
			}

			params := &types.QueryKeysRequest{
				Pagination:    pageReq,
				WorkspaceAddr: "",
				Type:          walletType,
				KeyId:         id,
			}

			if len(args) > 1 {
				params.WorkspaceAddr = args[1]
			}

			res, err := queryClient.Keys(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
