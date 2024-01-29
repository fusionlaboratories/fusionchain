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

var childWs = types.Workspace{
	Address:       "childWs",
	Creator:       "testOwner",
	Owners:        []string{"testOwner"},
	AdminPolicyId: 0,
	SignPolicyId:  0,
}

var wsWithChild = types.Workspace{
	Address:         "qredoworkspace14a2hpadpsy9h5m6us54",
	Creator:         "testOwner",
	Owners:          []string{"testOwner"},
	ChildWorkspaces: []string{"childWs"},
	AdminPolicyId:   0,
	SignPolicyId:    0,
}

func Test_msgServer_AppendChildWorkspace(t *testing.T) {
	type args struct {
		workspace *types.Workspace
		childWs   *types.Workspace
		msg       *types.MsgAppendChildWorkspace
	}
	tests := []struct {
		name          string
		args          args
		want          *types.MsgAppendChildWorkspaceResponse
		wantWorkspace *types.Workspace
		wantErr       bool
	}{
		{
			name: "add child workspace",
			args: args{
				workspace: &defaultWs,
				childWs:   &childWs,
				msg:       types.NewMsgAppendChildWorkspace("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", childWs.Address, 100),
			},
			want: &types.MsgAppendChildWorkspaceResponse{},
			wantWorkspace: &types.Workspace{
				Address:         "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:         "testOwner",
				Owners:          []string{"testOwner"},
				ChildWorkspaces: []string{"childWs"},
				AdminPolicyId:   0,
				SignPolicyId:    0,
			},
			wantErr: false,
		},
		{
			name: "workspace is nil or not found",
			args: args{
				workspace: &defaultWs,
				childWs:   &childWs,
				msg:       types.NewMsgAppendChildWorkspace("testOwner", "notAWorkspace", childWs.Address, 100),
			},
			want:    &types.MsgAppendChildWorkspaceResponse{},
			wantErr: true,
		},
		{
			name: "creator is not an owner",
			args: args{
				workspace: &defaultWs,
				childWs:   &childWs,
				msg:       types.NewMsgAppendChildWorkspace("notAnOwner", "qredoworkspace14a2hpadpsy9h5m6us54", childWs.Address, 100),
			},
			want:    &types.MsgAppendChildWorkspaceResponse{},
			wantErr: true,
		},
		{
			name: "new child is already a child",
			args: args{
				workspace: &wsWithChild,
				childWs:   &childWs,
				msg:       types.NewMsgAppendChildWorkspace("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "childWs", 100),
			},
			want:    &types.MsgAppendChildWorkspaceResponse{},
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
				Workspaces: []types.Workspace{*tt.args.workspace, *tt.args.childWs},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			got, err := msgSer.AppendChildWorkspace(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AppendChildWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AppendChildWorkspace() got = %v, want %v", got, tt.want)
				}

				gotWorkspace := ik.GetWorkspace(ctx, tt.args.workspace.Address)

				if !reflect.DeepEqual(gotWorkspace, tt.wantWorkspace) {
					t.Errorf("AppendChildWorkspace() got = %v, want %v", gotWorkspace, tt.wantWorkspace)
				}
			}
		})
	}
}
