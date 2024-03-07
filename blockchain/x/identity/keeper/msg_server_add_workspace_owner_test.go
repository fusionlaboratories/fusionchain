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

var defaultWs = types.Workspace{
	Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
	Creator:       "testOwner",
	Owners:        []string{"testOwner"},
	AdminPolicyId: 0,
	SignPolicyId:  0,
}

func Test_msgServer_AddWorkspaceOwner(t *testing.T) {
	type args struct {
		workspace *types.Workspace
		msg       *types.MsgAddWorkspaceOwner
	}
	tests := []struct {
		name          string
		args          args
		want          *types.MsgAddWorkspaceOwnerResponse
		wantWorkspace *types.Workspace
		wantErr       bool
	}{
		{
			name: "PASS: add workspace owner",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgAddWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner2", 100),
			},
			want: &types.MsgAddWorkspaceOwnerResponse{},
			wantWorkspace: &types.Workspace{
				Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:       "testOwner",
				Owners:        []string{"testOwner", "testOwner2"},
				AdminPolicyId: 0,
				SignPolicyId:  0,
			},
			wantErr: false,
		},
		{
			name: "FAIL: workspace is nil or not found",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgAddWorkspaceOwner("testOwner", "notAWorkspace", "testOwner2", 100),
			},
			want:    &types.MsgAddWorkspaceOwnerResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: owner is already owner",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgAddWorkspaceOwner("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner", 100),
			},
			want:    &types.MsgAddWorkspaceOwnerResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator is no admin owner",
			args: args{
				workspace: &defaultWs,
				msg:       types.NewMsgAddWorkspaceOwner("noOwner", "qredoworkspace14a2hpadpsy9h5m6us54", "testOwner", 100),
			},
			want:    &types.MsgAddWorkspaceOwnerResponse{},
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

			got, err := msgSer.AddWorkspaceOwner(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("AddWorkspaceOwner() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("AddWorkspaceOwner() got = %v, want %v", got, tt.want)
				}

				gotWorkspace := ik.GetWorkspace(ctx, tt.args.workspace.Address)

				if !reflect.DeepEqual(gotWorkspace, tt.wantWorkspace) {
					t.Errorf("AddWorkspaceOwner() got = %v, want %v", gotWorkspace, tt.wantWorkspace)
				}
			}
		})
	}
}
