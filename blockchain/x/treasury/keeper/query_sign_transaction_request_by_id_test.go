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
package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func TestKeeper_SignTransactionRequestById(t *testing.T) {
	t.SkipNow()

	var defaultSignTxReq = types.SignTransactionRequest{
		Id:            1,
		SignRequestId: 22,
	}

	type args struct {
		// keys       []*types.Key
		signTxReqs []types.SignTransactionRequest
		req        *types.QuerySignTransactionRequestByIdRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QuerySignTransactionRequestByIdResponse
		wantErr bool
	}{
		{
			name: "PASS: get a sign transaction request by id",
			args: args{
				signTxReqs: []types.SignTransactionRequest{defaultSignTxReq},
				req: &types.QuerySignTransactionRequestByIdRequest{
					Id: 1,
				},
			},
			want: &types.QuerySignTransactionRequestByIdResponse{
				SignTransactionRequest: &defaultSignTxReq,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			tk := keepers.TreasuryKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)

			// need to regenerate genesis.proto to allow for signTxRequests
			genesis := types.GenesisState{
				SignRequests: []types.SignRequest{},
			}
			treasury.InitGenesis(ctx, *tk, genesis)

			got, err := tk.SignTransactionRequestById(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyRequestById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyRequestById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
