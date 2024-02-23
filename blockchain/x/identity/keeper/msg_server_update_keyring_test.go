// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"reflect"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity"
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
)

func Test_msgServer_UpdateKeyring(t *testing.T) {

	var defaultKr = types.Keyring{
		Address:     "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:     "testCreator",
		Description: "testDescription",
		Admins:      []string{"testCreator"},
		IsActive:    true,
	}

	var wantKr = types.Keyring{
		Address:     "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:     "testCreator",
		Description: "testDescription",
		Admins:      []string{"testCreator"},
		IsActive:    true,
	}

	type args struct {
		keyring *types.Keyring
		msg     *types.MsgUpdateKeyring
	}
	tests := []struct {
		name        string
		args        args
		want        *types.MsgUpdateKeyringResponse
		wantKeyring *types.Keyring
		wantErr     bool
	}{
		{
			name: "PASS: change keyring description",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgUpdateKeyring("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "newDescription", true),
			},
			want: &types.MsgUpdateKeyringResponse{},
			wantKeyring: &types.Keyring{
				Address:     "qredokeyring1ph63us46lyw56vrzgaq",
				Creator:     "testCreator",
				Description: "newDescription",
				Admins:      []string{"testCreator"},
				IsActive:    true,
			},
		},
		{
			name: "FAIL: keyring not found",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgUpdateKeyring("testCreator", "invalidKeyring", "newDescription", true),
			},
			want:    &types.MsgUpdateKeyringResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator no keyring admin",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgUpdateKeyring("noAdmin", "qredokeyring1ph63us46lyw56vrzgaq", "newDescription", true),
			},
			want:    &types.MsgUpdateKeyringResponse{},
			wantErr: true,
		},
		{
			name: "PASS: change keyring status to false",
			args: args{
				keyring: &types.Keyring{
					Address:     "qredokeyring1ph63us46lyw56vrzgaq",
					Creator:     "testCreator",
					Description: "testDescription",
					Admins:      []string{"testCreator"},
					IsActive:    true,
				},
				msg: types.NewMsgUpdateKeyring("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testDescription", false),
			},
			want: &types.MsgUpdateKeyringResponse{},
			wantKeyring: &types.Keyring{
				Address:     "qredokeyring1ph63us46lyw56vrzgaq",
				Creator:     "testCreator",
				Description: "testDescription",
				Admins:      []string{"testCreator"},
				IsActive:    false,
			},
		},
		{
			name: "PASS: change keyring status to true",
			args: args{
				keyring: &types.Keyring{
					Address:     "qredokeyring1ph63us46lyw56vrzgaq",
					Creator:     "testCreator",
					Description: "testDescription",
					Admins:      []string{"testCreator"},
					IsActive:    false,
				},
				msg: types.NewMsgUpdateKeyring("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testDescription", true),
			},
			want:        &types.MsgUpdateKeyringResponse{},
			wantKeyring: &wantKr,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*ik)

			genesis := types.GenesisState{
				Keyrings: []types.Keyring{*tt.args.keyring},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			got, err := msgSer.UpdateKeyring(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateKeyring() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("UpdateKeyring() got = %v, want %v", got, tt.want)
				}
				gotKeyring := ik.GetKeyring(ctx, tt.args.keyring.Address)

				if !reflect.DeepEqual(gotKeyring, tt.wantKeyring) {
					t.Fatalf("UpdateKeyring() got = %v, want %v", gotKeyring, tt.wantKeyring)
				}
			}
		})
	}
}
