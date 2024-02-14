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

func TestKeeper_SignRequest(t *testing.T) {

	var defaultSignReq = types.SignRequest{
		Id:             1,
		Creator:        "testCreator",
		KeyId:          1,
		DataForSigning: []byte("test"),
		Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
		KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
	}

	type args struct {
		signReqs []types.SignRequest
		req      *types.QuerySignatureRequestsRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QuerySignatureRequestsResponse
		wantErr bool
	}{
		{
			name: "PASS: get all signature requests",
			args: args{
				signReqs: []types.SignRequest{defaultSignReq},
				req: &types.QuerySignatureRequestsRequest{
					KeyringAddr: defaultECDSAKey.KeyringAddr,
				},
			},
			want: &types.QuerySignatureRequestsResponse{
				Pagination:   &query.PageResponse{Total: 1},
				SignRequests: []*types.SignRequest{&defaultSignReq},
			},
		},
		{
			name: "FAIL: invalid request",
			args: args{
				signReqs: []types.SignRequest{defaultSignReq},
				req:      nil,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: keyring address does not match",
			args: args{
				signReqs: []types.SignRequest{defaultSignReq},
				req: &types.QuerySignatureRequestsRequest{
					KeyringAddr: "notAKeyring",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: key not found",
			args: args{
				signReqs: []types.SignRequest{{
					Id:      1,
					Creator: "testCreator",
					KeyId:   50756,
				}},
				req: &types.QuerySignatureRequestsRequest{
					KeyringAddr: defaultKr.Address,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "FAIL: requested key has wrong status",
			args: args{
				signReqs: []types.SignRequest{{
					Id:      1,
					Creator: "testCreator",
					KeyId:   1,
					Status:  types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				}},
				req: &types.QuerySignatureRequestsRequest{
					KeyringAddr: "testKeyring",
					Status:      types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
				},
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
				Keys:         []types.Key{defaultECDSAKey},
				SignRequests: []types.SignRequest{tt.args.signReqs[0]},
			}
			treasury.InitGenesis(ctx, *tk, genesis)

			got, err := tk.SignatureRequests(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignatureRequests() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignatureRequests() got = %v, want %v", got, tt.want)
			}
		})
	}
}
