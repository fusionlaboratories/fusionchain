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

func TestKeeper_KeyRequestById(t *testing.T) {

	type args struct {
		keyReq *types.KeyRequest
		req    *types.QueryKeyRequestByIdRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryKeyRequestByIdResponse
		wantErr bool
	}{
		{
			name: "PASS: get a key request by id",
			args: args{
				keyReq: &defaultKeyRequest,
				req: &types.QueryKeyRequestByIdRequest{
					Id: 1,
				},
			},
			want: &types.QueryKeyRequestByIdResponse{
				KeyRequest: &defaultKeyRequest,
			},
		},
		{
			name: "FAIL: key request not found",
			args: args{
				keyReq: &defaultKeyRequest,
				req: &types.QueryKeyRequestByIdRequest{
					Id: 10000,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: invalid request",
			args: args{
				keyReq: &defaultKeyRequest,
				req:    nil,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			tk := keepers.TreasuryKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)

			genesis := types.GenesisState{
				KeyRequests: []types.KeyRequest{*tt.args.keyReq},
			}
			treasury.InitGenesis(ctx, *tk, genesis)

			got, err := tk.KeyRequestById(goCtx, tt.args.req)
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
