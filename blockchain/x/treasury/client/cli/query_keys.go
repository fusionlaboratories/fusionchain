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
		Use:   "keys [workspace-addr]",
		Short: "Query Keys, optionally by workspace address",
		Args:  cobra.MaximumNArgs(1),
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

			walletTypeArg, err := cmd.Flags().GetString("wallet-type")
			if err != nil {
				return err
			}
			var walletType types.WalletRequestType
			if len(walletTypeArg) > 0 {
				switch strings.ToLower(walletTypeArg) {
				case "ethereum":
					walletType = types.WalletRequestType_WALLET_REQUEST_TYPE_ETH
				case "sepolia":
					walletType = types.WalletRequestType_WALLET_REQUEST_TYPE_ETH_SEPOLIA
				case "all":
					walletType = types.WalletRequestType_WALLET_REQUEST_TYPE_ALL
				default:
					return fmt.Errorf("invalid wallet type %s", walletTypeArg)
				}
			}
			params := &types.QueryKeysRequest{
				Pagination:    pageReq,
				WorkspaceAddr: "",
				Type:          walletType,
			}
			if len(args) > 0 {
				params.WorkspaceAddr = args[0]
			}

			res, err := queryClient.Keys(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cmd.Flags().String("wallet-type", "", "derive an address for type")

	return cmd
}
