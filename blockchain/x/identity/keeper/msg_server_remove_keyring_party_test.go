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
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
)

func Test_msgServer_RemoveKeyringParty(t *testing.T) {
	var defaultKr = types.Keyring{
		Address:       "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:       "testCreator",
		Description:   "testDescription",
		Admins:        []string{"testCreator"},
		Parties:       []string{"testParty"},
		AdminPolicyId: 0,
		Fees:          &types.KeyringFees{KeyReq: 0, SigReq: 0},
		IsActive:      true,
	}

	var wantKr = types.Keyring{
		Address:       "qredokeyring1ph63us46lyw56vrzgaq",
		Creator:       "testCreator",
		Description:   "testDescription",
		Admins:        []string{"testCreator"},
		Parties:       nil,
		AdminPolicyId: 0,
		Fees:          &types.KeyringFees{KeyReq: 0, SigReq: 0},
		IsActive:      true,
	}

	type args struct {
		keyring *types.Keyring
		msg     *types.MsgRemoveKeyringParty
		msg2    *types.MsgRemoveKeyringParty
	}

	tests := []struct {
		name        string
		args        args
		want        *types.MsgRemoveKeyringPartyResponse
		wantKeyring *types.Keyring
		wantErr     bool
		wantErr2    bool
	}{
		{
			name: "PASS: remove party from a keyring",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want:        &types.MsgRemoveKeyringPartyResponse{},
			wantKeyring: &wantKr,
		},
		{
			name: "FAIL: non-admin removing party from a keyring",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "invalidAdmin", "testParty"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: keyring not found",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "invalidKeyring", "testParty"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: party is already removed from the keyring",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgRemoveKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
				msg2:    types.NewMsgRemoveKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want:        &types.MsgRemoveKeyringPartyResponse{},
			wantKeyring: &wantKr,
			wantErr2:    true,
		},
		{
			name: "FAIL: creator no keyring admin",
			args: args{
				keyring: &defaultKr,
				msg:     types.NewMsgRemoveKeyringParty("notKeyringAdmin", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: inactive keyring",
			args: args{
				keyring: &types.Keyring{
					Address:       "qredokeyring1ph63us46lyw56vrzgaq",
					Creator:       "testCreator",
					Description:   "testDescription",
					Admins:        []string{"testCreator"},
					Parties:       []string{},
					AdminPolicyId: 0,
					Fees:          &types.KeyringFees{KeyReq: 0, SigReq: 0},
					IsActive:      false,
				},
				msg: types.NewMsgRemoveKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want:    &types.MsgRemoveKeyringPartyResponse{},
			wantErr: true,
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

			got, err := msgSer.RemoveKeyringParty(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AddKeyringParty() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr2 {
				_, err = msgSer.RemoveKeyringParty(goCtx, tt.args.msg2)
				if (err != nil) != tt.wantErr2 {
					t.Fatalf("AddKeyringParty() error = %v, wantErr %v", err, tt.wantErr2)
				}
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("AddKeyringParty() got = %v, want %v", got, tt.want)
				}
				gotKeyring := ik.GetKeyring(ctx, tt.args.keyring.Address)

				if !reflect.DeepEqual(gotKeyring, tt.wantKeyring) {
					t.Fatalf("NewKeyring() got = %v, want %v", gotKeyring, tt.wantKeyring)
				}
			}
		})
	}

}
