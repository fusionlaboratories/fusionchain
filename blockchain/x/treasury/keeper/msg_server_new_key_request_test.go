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
	"github.com/qredo/fusionchain/x/identity"
	idTypes "github.com/qredo/fusionchain/x/identity/types"
	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

func Test_msgServer_NewKeyRequest(t *testing.T) {

	type args struct {
		keyring   *idTypes.Keyring
		workspace *idTypes.Workspace
		msg       *types.MsgNewKeyRequest
	}
	tests := []struct {
		name           string
		args           args
		wantKeyRequest *types.KeyRequest
		want           *types.MsgNewKeyRequestResponse
		wantErr        bool
	}{
		{
			name: "PASS: request a new ecdsa key",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_ECDSA_SECP256K1, 1000),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testOwner",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
			},
			want: &types.MsgNewKeyRequestResponse{Id: 1},
		},
		{
			name: "PASS: request a new eddsa key",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_EDDSA_ED25519, 1000),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testOwner",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
			},
			want: &types.MsgNewKeyRequestResponse{Id: 1},
		},
		{
			name: "FAIL: workspace not found",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testOwner", "notAnOwner", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_ECDSA_SECP256K1, 1000),
			},
			want:    &types.MsgNewKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyring is inactive",
			args: args{
				keyring: &idTypes.Keyring{
					Address:       "qredokeyring1ph63us46lyw56vrzgaq",
					Creator:       "testCreator",
					Description:   "testDescription",
					Admins:        []string{"testCreator"},
					Parties:       []string{},
					AdminPolicyId: 0,
					Fees:          &idTypes.KeyringFees{KeyReq: 0, SigReq: 0},
					IsActive:      false,
				},
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_ECDSA_SECP256K1, 1000),
			},
			want:    &types.MsgNewKeyRequestResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx

			genesis := idTypes.GenesisState{
				Keyrings:   []idTypes.Keyring{*tt.args.keyring},
				Workspaces: []idTypes.Workspace{*tt.args.workspace},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			tk := keepers.TreasuryKeeper
			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*tk)

			got, err := msgSer.NewKeyRequest(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewKeyRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {

				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("NewKeyRequest() got = %v, want %v", got, tt.want)
				}

				gotKeyReq, bool := tk.KeyRequestsRepo().Get(ctx, got.Id)
				if !bool {
					t.Fatalf("KeyRequestsRepo().Get failed, error = %v", bool)
				}

				if !reflect.DeepEqual(gotKeyReq, tt.wantKeyRequest) {
					t.Fatalf("NewKeyRequest() got = %v, want %v", gotKeyReq, tt.wantKeyRequest)
				}
			}
		})
	}
}
