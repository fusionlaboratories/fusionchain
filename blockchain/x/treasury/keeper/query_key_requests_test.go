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
	"github.com/cosmos/cosmos-sdk/types/query"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func TestKeeper_KeyRequests(t *testing.T) {

	type args struct {
		keyReq *types.KeyRequest
		req    *types.QueryKeyRequestsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QueryKeyRequestsResponse
		wantErr bool
	}{
		{
			name: "PASS: return key requests for a workspace and a keyring",
			args: args{
				keyReq: &defaultKeyRequest,
				req: &types.QueryKeyRequestsRequest{
					KeyringAddr: defaultKeyRequest.KeyringAddr,
				},
			},
			want: &types.QueryKeyRequestsResponse{
				Pagination:  &query.PageResponse{Total: 1},
				KeyRequests: []*types.KeyRequest{&defaultKeyRequest},
			},
		},
		// {
		// 	name: "Pass: keyringAddr with no key requests",
		// 	args: args{
		// 		keyReq: &defaultKeyRequest,
		// 		req: &types.QueryKeyRequestsRequest{
		// 			KeyringAddr: "notAKeyringAddr",
		// 		},
		// 	},
		// 	want: &types.QueryKeyRequestsResponse{
		// 		Pagination: &query.PageResponse{},
		// 	},
		// 	wantErr: false,
		// },
		{
			name: "FAIL: request is empty",
			args: args{
				keyReq: &defaultKeyRequest,
				req:    &types.QueryKeyRequestsRequest{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: invalid request status",
			args: args{
				keyReq: &defaultKeyRequest,
				req: &types.QueryKeyRequestsRequest{
					Status: types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED,
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

			got, err := tk.KeyRequests(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("KeyRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyRequests() got = %v, want %v", got, tt.want)
			}
		})
	}
}
