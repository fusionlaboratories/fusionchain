// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"reflect"
	"testing"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity"
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
	policyKeeper "github.com/qredo/fusionchain/x/policy/keeper"
	policyTypes "github.com/qredo/fusionchain/x/policy/types"
)

func Test_msgServer_UpdateWorkspace(t *testing.T) {

	participant := []*policyTypes.PolicyParticipant{
		{
			Abbreviation: "testOwner",
			Address:      "testOwner",
		},
	}

	policy := &policyTypes.BoolparserPolicy{
		Definition:   "testOwner > 0",
		Participants: participant,
	}

	policyPayload, _ := codectypes.NewAnyWithValue(policy)

	type args struct {
		workspace *types.Workspace
		policy    *policyTypes.MsgNewPolicy
		msg       *types.MsgUpdateWorkspace
	}
	tests := []struct {
		name          string
		args          args
		want          *types.MsgUpdateWorkspaceResponse
		wantWorkspace *types.Workspace
		wantErr       bool
	}{
		{
			name: "PASS: add workspace owner",
			args: args{
				workspace: &defaultWs,
				policy:    policyTypes.NewMsgNewPolicy("testOwner", "newPolicy", policyPayload),
				msg:       types.NewMsgUpdateWorkspace("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", 1, 1, 100),
			},
			want: &types.MsgUpdateWorkspaceResponse{},
			wantWorkspace: &types.Workspace{
				Address:       "qredoworkspace14a2hpadpsy9h5m6us54",
				Creator:       "testOwner",
				Owners:        []string{"testOwner"},
				AdminPolicyId: 1,
				SignPolicyId:  1,
			},
			wantErr: false,
		},
		{
			name: "FAIL: workspace is nil or not found",
			args: args{
				workspace: &defaultWs,
				policy:    policyTypes.NewMsgNewPolicy("testOwner", "newPolicy", policyPayload),
				msg:       types.NewMsgUpdateWorkspace("testOwner", "notAWorkspace", 1, 1, 100),
			},
			want:    &types.MsgUpdateWorkspaceResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: no new values",
			args: args{
				workspace: &defaultWs,
				policy:    policyTypes.NewMsgNewPolicy("testOwner", "newPolicy", policyPayload),
				msg:       types.NewMsgUpdateWorkspace("testOwner", "qredoworkspace14a2hpadpsy9h5m6us54", 0, 0, 100),
			},
			want:    &types.MsgUpdateWorkspaceResponse{},
			wantErr: true,
		},
		{
			name: "FAIL: creator is no admin owner",
			args: args{
				workspace: &defaultWs,
				policy:    policyTypes.NewMsgNewPolicy("testOwner", "newPolicy", policyPayload),
				msg:       types.NewMsgUpdateWorkspace("noOwner", "qredoworkspace14a2hpadpsy9h5m6us54", 1, 1, 100),
			},
			want:    &types.MsgUpdateWorkspaceResponse{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			pk := keepers.PolicyKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)

			genesis := types.GenesisState{
				Workspaces: []types.Workspace{*tt.args.workspace},
			}
			identity.InitGenesis(ctx, *ik, genesis)

			msgPolSer := policyKeeper.NewMsgServerImpl(*pk)
			_, err := msgPolSer.NewPolicy(goCtx, tt.args.policy)
			if err != nil {
				t.Fatalf("NewPolicy() error = %v", err)
			}

			msgSer := keeper.NewMsgServerImpl(*ik)
			got, err := msgSer.UpdateWorkspace(goCtx, tt.args.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("UpdateWorkspace() error = %v, wantErr %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("UpdateWorkspace() got = %v, want %v", got, tt.want)
				}

				gotWorkspace := ik.GetWorkspace(ctx, tt.args.workspace.Address)

				if !reflect.DeepEqual(gotWorkspace, tt.wantWorkspace) {
					t.Errorf("UpdateWorkspace() got = %v, want %v", gotWorkspace, tt.wantWorkspace)
				}
			}
		})
	}
}
