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

var defaultKr = idTypes.Keyring{
	Address:       "qredokeyring1ph63us46lyw56vrzgaq",
	Creator:       "testCreator",
	Description:   "testDescription",
	Admins:        []string{"testCreator"},
	Parties:       []string{"testCreator"},
	AdminPolicyId: 0,
	Fees:          &idTypes.KeyringFees{KeyReq: 0, SigReq: 0},
	IsActive:      true,
}

var defaultWs = idTypes.Workspace{
	Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
	Creator:       "testOwner",
	Owners:        []string{"testOwner"},
	AdminPolicyId: 0,
	SignPolicyId:  0,
}

var defaultECDSAKey = types.Key{
	Id:            1,
	WorkspaceAddr: "testWorkspace",
	KeyringAddr:   "testKeyring",
	Type:          types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
	PublicKey:     []byte{3, 49, 216, 1, 30, 80, 59, 231, 205, 4, 237, 36, 101, 79, 9, 119, 30, 114, 19, 70, 242, 44, 91, 95, 238, 20, 245, 197, 193, 86, 22, 120, 35},
}

var defaultEdDSAKey = types.Key{
	Id:            1,
	WorkspaceAddr: "testWorkspace",
	KeyringAddr:   "testKeyring",
	Type:          types.KeyType_KEY_TYPE_EDDSA_ED25519,
	PublicKey:     []byte{243, 178, 23, 221, 136, 81, 23, 248, 229, 31, 154, 135, 176, 26, 117, 104, 94, 9, 73, 68, 162, 139, 9, 231, 47, 249, 137, 156, 60, 87, 66, 163},
}

func Test_msgServer_FulfilSignatureRequest(t *testing.T) {

	var sigRequestPayload = []byte{37, 37, 38, 66, 35, 37, 32, 66, 33, 33, 61, 63, 66, 61, 62, 38, 33, 31, 33, 36, 35, 64, 35, 32, 65, 35, 36, 33, 61, 30, 64, 64, 64, 32, 38, 32, 39, 64, 64, 64, 37, 30, 36, 30, 62, 65, 63, 36, 39, 37, 31, 39, 62, 37, 65, 34, 31, 66, 36, 65, 66, 39, 31, 63}

	var sigPayloadECDSA = []byte{173, 224, 103, 159, 55, 251, 255, 212, 50, 145, 235, 181, 19, 14, 120, 168, 208, 151, 33, 204, 161, 79, 118, 229, 75, 22, 185, 234, 115, 125, 170, 101, 55, 197, 218, 94, 172, 32, 139, 21, 141, 104, 163, 109, 45, 47, 80, 110, 39, 5, 156, 88, 31, 82, 123, 246, 67, 21, 199, 126, 75, 222, 65, 115, 1}
	// removed {173} to match eddsa signature format
	var sigPayloadEdDSA = []byte{224, 103, 159, 55, 251, 255, 212, 50, 145, 235, 181, 19, 14, 120, 168, 208, 151, 33, 204, 161, 79, 118, 229, 75, 22, 185, 234, 115, 125, 170, 101, 55, 197, 218, 94, 172, 32, 139, 21, 141, 104, 163, 109, 45, 47, 80, 110, 39, 5, 156, 88, 31, 82, 123, 246, 67, 21, 199, 126, 75, 222, 65, 115, 1}

	var defaultSigRequest = types.SignRequest{
		Id:             1,
		Creator:        "testCreator",
		KeyId:          1,
		DataForSigning: sigRequestPayload,
		Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
		KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
		Result:         nil,
	}

	var defaultResponseECDSA = types.MsgFulfilSignatureRequest{
		Creator:   "testRequestCreator",
		RequestId: 1,
		Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
		Result:    types.NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
	}

	type args struct {
		key *types.Key
		req *types.SignRequest
		msg *types.MsgFulfilSignatureRequest
	}
	tests := []struct {
		name       string
		args       args
		want       *types.MsgFulfilSignatureRequestResponse
		wantSigReq *types.SignRequest
		wantErr    bool
	}{
		{
			name: "PASS: return signature request - ECDSA",
			args: args{
				key: &defaultECDSAKey,
				req: &defaultSigRequest,
				msg: &defaultResponseECDSA,
			},
			want: &types.MsgFulfilSignatureRequestResponse{},
			wantSigReq: &types.SignRequest{
				Id:             1,
				Creator:        "testCreator",
				KeyId:          1,
				DataForSigning: sigRequestPayload,
				Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Result: &types.SignRequest_SignedData{
					SignedData: sigPayloadECDSA,
				},
			},
			wantErr: false,
		},
		{
			name: "PASS: return signature request - EDDSA",
			args: args{
				key: &defaultEdDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: []byte("test"),
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testRequestCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					Result:    types.NewMsgFulfilSignatureRequestPayload(sigPayloadEdDSA),
				},
			},
			want: &types.MsgFulfilSignatureRequestResponse{},
			wantSigReq: &types.SignRequest{
				Id:             1,
				Creator:        "testCreator",
				KeyId:          1,
				DataForSigning: []byte("test"),
				Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				KeyType:        types.KeyType_KEY_TYPE_EDDSA_ED25519,
				Result: &types.SignRequest_SignedData{
					SignedData: sigPayloadEdDSA,
				},
			},
			wantErr: false,
		},
		{
			name: "PASS: Reject ECDSA Signature Request",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED,
					Result:    types.NewMsgFulfilSignatureRequestReject("test"),
				},
			},
			want: &types.MsgFulfilSignatureRequestResponse{},
			wantSigReq: &types.SignRequest{
				Id:             1,
				Creator:        "testCreator",
				KeyId:          1,
				DataForSigning: sigRequestPayload,
				Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_REJECTED,
				KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Result: &types.SignRequest_RejectReason{
					RejectReason: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "FAIL: Empty Status field in rejection",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_UNSPECIFIED,
					Result:    types.NewMsgFulfilSignatureRequestReject("test"),
				},
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: invalid signature, want eddsa got ecdsa",
			args: args{
				key: &defaultEdDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_EDDSA_ED25519,
					Result:         nil,
				},
				msg: &defaultResponseECDSA,
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: signature request status is already fulfilled",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					Result:    types.NewMsgFulfilSignatureRequestPayload(sigPayloadEdDSA),
				},
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: signature request not found",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 9999,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					Result:    types.NewMsgFulfilSignatureRequestPayload(sigPayloadEdDSA),
				},
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: invalid ecdsa signature length",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					Result: types.NewMsgFulfilSignatureRequestPayload(
						// added 11 on top to max out on length
						[]byte{11, 173, 224, 103, 159, 55, 251, 255, 212, 50, 145, 235, 181, 19, 14, 120, 168, 208, 151, 33, 204, 161, 79, 118, 229, 75, 22, 185, 234, 115, 125, 170, 101, 55, 197, 218, 94, 172, 32, 139, 21, 141, 104, 163, 109, 45, 47, 80, 110, 39, 5, 156, 88, 31, 82, 123, 246, 67, 21, 199, 126, 75, 222, 65, 115, 1},
					),
				},
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: invalid key type",
			args: args{
				key: &defaultECDSAKey,
				req: &types.SignRequest{
					Id:             1,
					Creator:        "testCreator",
					KeyId:          1,
					DataForSigning: sigRequestPayload,
					Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
					KeyType:        types.KeyType_KEY_TYPE_UNSPECIFIED,
					Result:         nil,
				},
				msg: &types.MsgFulfilSignatureRequest{
					Creator:   "testCreator",
					RequestId: 1,
					Status:    types.SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
					Result: types.NewMsgFulfilSignatureRequestPayload(
						[]byte{173, 224, 103, 159, 55, 251, 255, 212, 50, 145, 235, 181, 19, 14, 120, 168, 208, 151, 33, 204, 161, 79, 118, 229, 75, 22, 185, 234, 115, 125, 170, 101, 55, 197, 218, 94, 172, 32, 139, 21, 141, 104, 163, 109, 45, 47, 80, 110, 39, 5, 156, 88, 31, 82, 123, 246, 67, 21, 199, 126, 75, 222, 65, 115, 1},
					),
				},
			},
			want:    &types.MsgFulfilSignatureRequestResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			tk := keepers.TreasuryKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*tk)

			idGenesis := idTypes.GenesisState{
				Keyrings:   []idTypes.Keyring{defaultKr},
				Workspaces: []idTypes.Workspace{defaultWs},
			}
			identity.InitGenesis(ctx, *ik, idGenesis)

			trGenesis := types.GenesisState{
				Keys:         []types.Key{*tt.args.key},
				SignRequests: []types.SignRequest{*tt.args.req},
			}
			treasury.InitGenesis(ctx, *tk, trGenesis)

			got, err := msgSer.FulfilSignatureRequest(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("FulfilSignatureRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("FulfilSignatureRequest() got = %v, want %v", got, tt.want)
				}
				gotSigReq, bool := tk.SignatureRequestsRepo().Get(ctx, tt.args.msg.RequestId)
				if !bool {
					t.Fatalf("SignatureRequestsRepo() got = %v", err)
				}

				if !reflect.DeepEqual(gotSigReq, tt.wantSigReq) {
					t.Fatalf("FulfilSignatureRequest() got = %v, want %v", gotSigReq, tt.wantSigReq)
				}
			}
		})
	}
}
