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

func Test_msgServer_NewChildWorkspace(t *testing.T) {
	type args struct {
		workspace *types.Workspace
		msg       *types.MsgNewChildWorkspace
	}
	tests := []struct {
		name          string
		args          args
		want          *types.MsgNewChildWorkspaceResponse
		wantWorkspace *types.Workspace
		wantErr       bool
	}{
		{
			name: "create new child workspace",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgNewChildWorkspace("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", 100),
			},
			want: &types.MsgNewChildWorkspaceResponse{},
			wantWorkspace: &types.Workspace{
				Address:         "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:         "testOwner",
				Owners:          []string{"testOwner"},
				ChildWorkspaces: []string{"qredoworkspace10j06zdk5gyl6vrss5d5"},
				AdminPolicyId:   0,
				SignPolicyId:    0,
			},
			wantErr: false,
		},
		{
			name: "workspace is nil or not found",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgNewChildWorkspace("testOwner", "notAWorkspace", 100),
			},
			want:    &types.MsgNewChildWorkspaceResponse{},
			wantErr: true,
		},
		{
			name: "creator is not an owner",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgNewChildWorkspace("notAnOwner", "qredoworkspace14a2hpadpsy9h5m6us54", 100),
			},
			want:    &types.MsgNewChildWorkspaceResponse{},
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
				Workspaces: []types.Workspace{*tt.args.workspace},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			got, err := msgSer.NewChildWorkspace(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("NewChildWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewChildWorkspace() got = %v, want %v", got, tt.want)
				}

				gotWorkspace := ik.GetWorkspace(ctx, tt.args.workspace.Address)

				if !reflect.DeepEqual(gotWorkspace, tt.wantWorkspace) {
					t.Errorf("NewChildWorkspace() got = %v, want %v", gotWorkspace, tt.wantWorkspace)
				}
			}
		})
	}
}
