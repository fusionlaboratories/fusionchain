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

func Test_msgServer_NewSignatureRequest(t *testing.T) {
	t.SkipNow()

	type args struct {
		keyring   *idTypes.Keyring
		workspace *idTypes.Workspace
		key       *types.Key
		msg       *types.MsgNewSignatureRequest
	}
	tests := []struct {
		name            string
		args            args
		wantSignRequest *types.SignRequest
		want            *types.MsgNewSignatureRequest
		wantErr         bool
	}{
		{
			name: "PASS: valid signature request",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				key: &types.Key{
					Id:            1,
					WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
					KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
					Type:          types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
					PublicKey:     []byte{3, 49, 216, 1, 30, 80, 59, 231, 205, 4, 237, 36, 101, 79, 9, 119, 30, 114, 19, 70, 242, 44, 91, 95, 238, 20, 245, 197, 193, 86, 22, 120, 35},
				},
				msg: types.NewMsgNewSignatureRequest("testCreator", 1, []byte("778f572f33afab831365d52e563a0ddd"), 1000),
			},
			wantSignRequest: &types.SignRequest{
				Id:             1,
				Creator:        "testCreator",
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Status:         types.SignRequestStatus_SIGN_REQUEST_STATUS_PENDING,
				KeyType:        types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Result:         nil,
			},
			want: &types.MsgNewSignatureRequest{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx

			idGenesis := idTypes.GenesisState{
				Keyrings:   []idTypes.Keyring{*tt.args.keyring},
				Workspaces: []idTypes.Workspace{*tt.args.workspace},
			}
			identity.InitGenesis(ctx, *ik, idGenesis)

			// we need to add default policies to the genesis here as well.
			//polGenesis := polTypes.GenesisState{}

			tk := keepers.TreasuryKeeper
			tGenesis := types.GenesisState{
				Keys: []types.Key{*tt.args.key},
			}
			treasury.InitGenesis(ctx, *tk, tGenesis)

			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*tk)

			got, err := msgSer.NewSignatureRequest(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewSignatureRequest() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {

				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("NewSignatureRequest() got = %v, want %v", got, tt.want)
				}

				gotKeyReq, bool := tk.KeyRequestsRepo().Get(ctx, got.Id)
				if !bool {
					t.Fatalf("NewSignatureRequestRepo().Get failed, error = %v", bool)
				}

				if !reflect.DeepEqual(gotKeyReq, tt.wantSignRequest) {
					t.Fatalf("NewSignatureRequest() got = %v, want %v", gotKeyReq, tt.wantSignRequest)
				}
			}
		})
	}
}
