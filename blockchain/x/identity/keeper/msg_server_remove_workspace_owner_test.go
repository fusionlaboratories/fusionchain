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

func Test_msgServer_RemoveWorkspaceOwner(t *testing.T) {

	var defaultWsWithOwners = types.Workspace{
		Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
		Creator:       "testOwner",
		Owners:        []string{"testOwner", "testOwner2", "testOwner3"},
		AdminPolicyId: 0,
		SignPolicyId:  0,
	}

	type args struct {
		workspace *types.Workspace
		msg       *types.MsgRemoveWorkspaceOwner
	}
	tests := []struct {
		name          string
		args          args
		want          *types.MsgRemoveWorkspaceOwnerResponse
		wantWorkspace *types.Workspace
		wantErr       bool
	}{
		{
			name: "PASS: remove workspace owner",
			args: args{
				workspace: &defaultWsWithOwners,
				msg:       types.NewMsgRemoveWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner2"),
			},
			want: &types.MsgRemoveWorkspaceOwnerResponse{},
			wantWorkspace: &types.Workspace{
				Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:       "testOwner",
				Owners:        []string{"testOwner", "testOwner3"},
				AdminPolicyId: 0,
				SignPolicyId:  0,
			},
		},
		{
			name: "PASS: remove workspace creator",
			args: args{
				workspace: &defaultWsWithOwners,
				msg:       types.NewMsgRemoveWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner"),
			},
			want: &types.MsgRemoveWorkspaceOwnerResponse{},
			wantWorkspace: &types.Workspace{
				Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:       "testOwner",
				Owners:        []string{"testOwner2", "testOwner3"},
				AdminPolicyId: 0,
				SignPolicyId:  0,
			},
		},
		{
			name: "PASS: remove single owner",
			args: args{
				workspace: &types.Workspace{
					Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
					Creator:       "testOwner",
					Owners:        []string{"testOwner"},
					AdminPolicyId: 0,
					SignPolicyId:  0,
				},
				msg: types.NewMsgRemoveWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner"),
			},
			want: &types.MsgRemoveWorkspaceOwnerResponse{},
			wantWorkspace: &types.Workspace{
				Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:       "testOwner",
				AdminPolicyId: 0,
				SignPolicyId:  0,
			},
		},
		{
			name: "FAIL: workspace is nil or not found",
			args: args{
				workspace: &defaultWsWithOwners,
				msg:       types.NewMsgRemoveWorkspaceOwner("testOwner", "notAWorkspace", "testOwner2"),
			},
			want:    &types.MsgRemoveWorkspaceOwnerResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: removed owner is no owner",
			args: args{
				workspace: &defaultWsWithOwners,
				msg:       types.NewMsgRemoveWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "noOwner"),
			},
			want:    &types.MsgRemoveWorkspaceOwnerResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator is no admin owner",
			args: args{
				workspace: &defaultWsWithOwners,
				msg:       types.NewMsgRemoveWorkspaceOwner("noOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner"),
			},
			want:    &types.MsgRemoveWorkspaceOwnerResponse{},
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

			got, err := msgSer.RemoveWorkspaceOwner(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("RemoveWorkspaceOwner() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("RemoveWorkspaceOwner() got = %v, want %v", got, tt.want)
				}

				gotWorkspace := ik.GetWorkspace(ctx, tt.args.workspace.Address)

				if !reflect.DeepEqual(gotWorkspace, tt.wantWorkspace) {
					t.Errorf("RemoveWorkspaceOwner() got = %v, want %v", gotWorkspace, tt.wantWorkspace)
				}
			}
		})
	}
}
