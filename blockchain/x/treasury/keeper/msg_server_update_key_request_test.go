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
	"github.com/qredo/fusionchain/x/treasury"
	"github.com/qredo/fusionchain/x/treasury/keeper"
	"github.com/qredo/fusionchain/x/treasury/types"
)

var defaultKeyRequest = types.KeyRequest{
	Id:            1,
	Creator:       "testCreator",
	WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
	KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
	KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
	Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
}

func Test_msgServer_UpdateKeyRequest(t *testing.T) {

	// too long
	invalidECDSAPubKey := []byte{154, 135, 176, 26, 117, 104, 94, 9, 73, 68, 162, 139, 9, 231, 47, 249, 137, 156, 60, 87, 66, 163}

	// too long
	invalidEdDSAPubkey := []byte{1, 243, 178, 23, 221, 136, 81, 23, 248, 229, 31, 154, 135, 176, 26, 117, 104, 94, 9, 73, 68, 162, 139, 9, 231, 47, 249, 137, 156, 60, 87, 66, 163}

	type args struct {
		keyring    *idTypes.Keyring
		workspace  *idTypes.Workspace
		keyRequest *types.KeyRequest
		msg        *types.MsgUpdateKeyRequest
	}
	tests := []struct {
		name           string
		args           args
		wantKeyRequest *types.KeyRequest
		want           *types.MsgUpdateKeyRequestResponse
		wantErr        bool
	}{
		{
			name: "PASS: return a new ecdsa key",
			args: args{
				keyring:    &defaultKr,
				workspace:  &defaultWs,
				keyRequest: &defaultKeyRequest,
				msg:        types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultECDSAKey.PublicKey}}),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED,
			},
			want: &types.MsgUpdateKeyRequestResponse{},
		},
		{
			name: "PASS: reject the request",
			args: args{
				keyring:    &defaultKr,
				workspace:  &defaultWs,
				keyRequest: &defaultKeyRequest,
				msg:        types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED, types.NewMsgUpdateKeyRequestReject("test")),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED,
				RejectReason:  "test",
			},
			want: &types.MsgUpdateKeyRequestResponse{},
		},
		{
			name: "PASS: return a new eddsa key",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED,
			},
			want: &types.MsgUpdateKeyRequestResponse{},
		},
		{
			name: "PASS: reject eddsa key request",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED, &types.MsgUpdateKeyRequest_RejectReason{RejectReason: "test"}),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_REJECTED,
				RejectReason:  "test",
			},
			want: &types.MsgUpdateKeyRequestResponse{},
		},
		{
			name: "FAIL: request not found",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 999, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyring not found",
			args: args{
				keyring: &idTypes.Keyring{
					Address:       "notAKeyring",
					Creator:       "testCreator",
					Description:   "testDescription",
					Admins:        []string{"testCreator"},
					Parties:       []string{"testCreator"},
					AdminPolicyId: 0,
					Fees:          &idTypes.KeyringFees{KeyReq: 0, SigReq: 0},
					IsActive:      true,
				},
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyring inactive",
			args: args{
				keyring: &idTypes.Keyring{
					Address:       "qredokeyring1ph63us46lyw56vrzgaq",
					Creator:       "testCreator",
					Description:   "testDescription",
					Admins:        []string{"testCreator"},
					Parties:       []string{"testCreator"},
					AdminPolicyId: 0,
					Fees:          &idTypes.KeyringFees{KeyReq: 0, SigReq: 0},
					IsActive:      false,
				},
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator is no keyring party",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("noKeyringParty", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyRequest status is not pending",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyRequest status is unspecified",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_UNSPECIFIED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: defaultEdDSAKey.PublicKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: ecdsa pubkey is invalid",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: invalidECDSAPubKey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: eddsa pubkey is too long",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				keyRequest: &types.KeyRequest{
					Id:            1,
					Creator:       "testCreator",
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
				},
				msg: types.NewMsgUpdateKeyRequest("testCreator", 1, types.KeyRequestStatus_KEY_REQUEST_STATUS_FULFILLED, &types.MsgUpdateKeyRequest_Key{Key: &types.MsgNewKey{PublicKey: invalidEdDSAPubkey}}),
			},
			want:    &types.MsgUpdateKeyRequestResponse{},
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
			tGenesis := types.GenesisState{
				KeyRequests: []types.KeyRequest{*tt.args.keyRequest},
			}
			treasury.InitGenesis(ctx, *tk, tGenesis)

			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*tk)

			got, err := msgSer.UpdateKeyRequest(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateKeyRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {

				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("UpdateKeyRequest() got = %v, want %v", got, tt.want)
				}

				gotKeyRequest, bool := tk.KeyRequestsRepo().Get(ctx, tt.args.keyRequest.Id)
				if !bool {
					t.Fatalf("KeyRequestsRepo() failed, error = %v", bool)
				}

				if !reflect.DeepEqual(gotKeyRequest, tt.wantKeyRequest) {
					t.Fatalf("UpdateKeyRequest() got = %v, want %v", gotKeyRequest, tt.wantKeyRequest)
				}
			}
		})
	}
}
