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
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
)

func Test_msgServer_AddKeyringParty(t *testing.T) {

	type args struct {
		msg        *types.MsgAddKeyringParty
		msg2       *types.MsgAddKeyringParty
		msgKeyring *types.MsgNewKeyring
	}
	tests := []struct {
		name        string
		args        args
		want        *types.MsgAddKeyringPartyResponse
		wantKeyring *types.Keyring
		wantErr     bool
		wantErr2    bool
	}{
		{
			name: "add a party to a keyring",
			args: args{
				msgKeyring: types.NewMsgNewKeyring("testCreator", "testDescription", 0, 0, 0),
				msg:        types.NewMsgAddKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
				msg2:       types.NewMsgAddKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty2"),
			},
			want: &types.MsgAddKeyringPartyResponse{},
			wantKeyring: &types.Keyring{
				Address:     "qredokeyring1ph63us46lyw56vrzgaq",
				Creator:     "testCreator",
				Description: "testDescription",
				Admins:      []string{"testCreator"},
				Parties:     []string{"testParty"},
				Fees:        &types.KeyringFees{KeyReq: 0, SigReq: 0},
				IsActive:    true,
			},
		},
		{
			name: "keyring not found",
			args: args{
				msgKeyring: types.NewMsgNewKeyring("testCreator", "testDescription", 0, 0, 0),
				msg:        types.NewMsgAddKeyringParty("testCreator", "invalidKeyring", "testParty"),
			},
			want:    &types.MsgAddKeyringPartyResponse{},
			wantErr: true,
		},
		{
			name: "party is already in the keyring",
			args: args{
				msgKeyring: types.NewMsgNewKeyring("testCreator", "testDescription", 0, 0, 0),
				msg:        types.NewMsgAddKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
				msg2:       types.NewMsgAddKeyringParty("testCreator", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want: &types.MsgAddKeyringPartyResponse{},
			wantKeyring: &types.Keyring{
				Address:     "qredokeyring1ph63us46lyw56vrzgaq",
				Creator:     "testCreator",
				Description: "testDescription",
				Admins:      []string{"testCreator"},
				Parties:     []string{"testParty"},
				Fees:        &types.KeyringFees{KeyReq: 0, SigReq: 0},
				IsActive:    true,
			},
			wantErr2: true,
		},
		{
			name: "creator no keyring admin",
			args: args{
				msgKeyring: types.NewMsgNewKeyring("testCreator", "testDescription", 0, 0, 0),
				msg:        types.NewMsgAddKeyringParty("notKeyringAdmin", "qredokeyring1ph63us46lyw56vrzgaq", "testParty"),
			},
			want:    &types.MsgAddKeyringPartyResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ik, ctx := keepertest.IdentityKeeper(t)
			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*ik)
			keyringRes, err := msgSer.NewKeyring(goCtx, tt.args.msgKeyring)
			if err != nil {
				t.Fatalf("NewKeyring() error = %v", err)
			}
			got, err := msgSer.AddKeyringParty(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AddKeyringParty() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr2 {
				_, err = msgSer.AddKeyringParty(goCtx, tt.args.msg2)
				if (err != nil) != tt.wantErr2 {
					t.Fatalf("AddKeyringParty() error = %v, wantErr %v", err, tt.wantErr2)
				}
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Fatalf("AddKeyringParty() got = %v, want %v", got, tt.want)
				}
				gotKeyring := ik.GetKeyring(ctx, keyringRes.Address)

				if !reflect.DeepEqual(gotKeyring, tt.wantKeyring) {
					t.Fatalf("NewKeyring() got = %v, want %v", gotKeyring, tt.wantKeyring)
				}
			}
		})
	}
}
