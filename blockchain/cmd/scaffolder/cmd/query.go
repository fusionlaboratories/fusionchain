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
package cmd

import (
	"fmt"
	"path"

	"github.com/qredo/fusionchain/cmd/scaffolder/casing"
	"github.com/qredo/fusionchain/cmd/scaffolder/editor"
	"github.com/spf13/cobra"
)

type QueryCmdParams struct {
	ModuleName string
	QueryName  string
}

func queryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "query [module] [name]",
		Short: "Scaffold a new query",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			params := QueryCmdParams{
				ModuleName: args[0],
				QueryName:  args[1],
			}

			if err := queryProto(params); err != nil {
				return err
			}

			if err := cliQuery(params); err != nil {
				return err
			}

			if err := cliQueryCommand(params); err != nil {
				return err
			}

			if err := keeperQuery(params); err != nil {
				return err
			}

			return nil
		},
	}
}

func queryProto(params QueryCmdParams) error {
	n := fmt.Sprintf(`// Queries a list of %[1]s items.
rpc %[1]s (Query%[1]sRequest) returns (Query%[1]sResponse) {
  option (google.api.http).get = "/fusionchain/%[3]s/%[2]s";
}
`, params.QueryName, casing.ToSnakeCase(params.QueryName), params.ModuleName)

	addQueryToService := editor.Replacer{
		Substitute: n,
	}

	msgs := fmt.Sprintf(`
message Query%[1]sRequest {}

message Query%[1]sResponse {}`, params.QueryName)
	appendNewQueryMessages := editor.Appender{
		What: msgs,
	}

	return editor.Pipeline(
		path.Join("./proto/fusionchain", params.ModuleName, "query.proto"),
		[]editor.Modifier{
			addQueryToService,
			appendNewQueryMessages,
		},
	)
}

func cliQuery(params QueryCmdParams) error {
	substitute := fmt.Sprintf(`cmd.AddCommand(Cmd%[1]s())`, params.QueryName)
	addCommand := editor.Replacer{
		Substitute: substitute,
	}

	return editor.Pipeline(
		path.Join("./x", params.ModuleName, "client/cli/query.go"),
		[]editor.Modifier{
			addCommand,
		},
	)
}

func cliQueryCommand(params QueryCmdParams) error {
	tmpl := `package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"github.com/qredo/fusionchain/x/{{ .ModuleName }}/types"
)

var _ = strconv.Itoa(0)

func Cmd{{ .QueryName }}() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "{{ .QueryName | ToKebabCase }}",
		Short: "Query {{ .QueryName | ToKebabCase }}",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.Query{{ .QueryName }}Request{}

			res, err := queryClient.{{ .QueryName }}(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}`
	return editor.NewFile(
		path.Join("x", params.ModuleName, "client/cli", fmt.Sprintf("query_%s.go", casing.ToSnakeCase(params.QueryName))),
		tmpl,
		params,
	)
}

func keeperQuery(params QueryCmdParams) error {
	tmpl := `package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/{{ .ModuleName }}/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) {{ .QueryName }}(goCtx context.Context, req *types.Query{{ .QueryName }}Request) (*types.Query{{ .QueryName }}Response, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.Query{{ .QueryName }}Response{}, nil
}`
	return editor.NewFile(
		path.Join("x", params.ModuleName, "keeper", fmt.Sprintf("query_%s.go", casing.ToSnakeCase(params.QueryName))),
		tmpl,
		params,
	)
}
