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

func TestKeeper_SignRequestById(t *testing.T) {

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
		req      *types.QuerySignatureRequestByIdRequest
	}
	tests := []struct {
		name    string
		args    args
		want    *types.QuerySignatureRequestByIdResponse
		wantErr bool
	}{
		{
			name: "PASS: get a signature request by id",
			args: args{
				signReqs: []types.SignRequest{defaultSignReq},
				req: &types.QuerySignatureRequestByIdRequest{
					Id: 1,
				},
			},
			want: &types.QuerySignatureRequestByIdResponse{
				SignRequest: &defaultSignReq,
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
			name: "FAIL: invalid request",
			args: args{
				signReqs: []types.SignRequest{defaultSignReq},
				req: &types.QuerySignatureRequestByIdRequest{
					Id: 199999,
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
				SignRequests: []types.SignRequest{defaultSignReq},
			}
			treasury.InitGenesis(ctx, *tk, genesis)

			got, err := tk.SignatureRequestById(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignatureRequestById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SignatureRequestById() got = %v, want %v", got, tt.want)
			}
		})
	}
}
