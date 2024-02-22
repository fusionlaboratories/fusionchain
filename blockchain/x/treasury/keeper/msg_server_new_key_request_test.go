// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
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

	var defaultKr = idTypes.Keyring{
		Address:       "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:       "testCreator",
		Description:   "testDescription",
		Admins:        []string{"testCreator"},
		Parties:       []string{},
		AdminPolicyId: 0,
		Fees:          &idTypes.KeyringFees{KeyReq: 0, SigReq: 0},
		IsActive:      true,
	}

	var defaultWs = idTypes.Workspace{
		Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
		Creator:       "testCreator",
		Owners:        []string{"testCreator"},
		AdminPolicyId: 0,
		SignPolicyId:  0,
	}

	type args struct {
		keyring   *idTypes.Keyring
		workspace *idTypes.Workspace
		msg       *types.MsgNewKeyRequest
	}
	tests := []struct {
		name           string
		args           args
		wantKeyRequest *types.KeyRequest
		wantErr        bool
	}{
		{
			name: "request a new ecdsa key",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testCreator", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_ECDSA_SECP256K1, 1000),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_ECDSA_SECP256K1,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
			},
		},
		{
			name: "request a new eddsa key",
			args: args{
				keyring:   &defaultKr,
				workspace: &defaultWs,
				msg:       types.NewMsgNewKeyRequest("testCreator", "qredoworkspace14a2hpadpsy9h5m6us54", "qredokeyring1ph63us46lyw56vrzgaq", types.KeyType_KEY_TYPE_EDDSA_ED25519, 1000),
			},
			wantKeyRequest: &types.KeyRequest{
				Id:            1,
				Creator:       "testCreator",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       types.KeyType_KEY_TYPE_EDDSA_ED25519,
				Status:        types.KeyRequestStatus_KEY_REQUEST_STATUS_PENDING,
			},
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

			gotResponse, bool := tk.KeyRequestsRepo().Get(ctx, got.Id)
			if !bool {
				t.Fatalf("KeyRequestsRepo().Get failed, error = %v", bool)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(gotResponse, tt.wantKeyRequest) {
					t.Fatalf("NewKeyRequest() got = %v, want %v", gotResponse, tt.wantKeyRequest)
				}
			}
		})
	}
}
