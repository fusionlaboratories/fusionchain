// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/qredo/fusionchain/testutil/keeper"
	"github.com/qredo/fusionchain/x/identity/keeper"
	"github.com/qredo/fusionchain/x/identity/types"
)

func TestKeeper_WorkspacesByOwner(t *testing.T) {

	type args struct {
		req          *types.QueryWorkspacesByOwnerRequest
		msgWorkspace *types.MsgNewWorkspace
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "PASS: find by owner",
			args: args{
				req: &types.QueryWorkspacesByOwnerRequest{
					Owner: "testOwner",
				},
				msgWorkspace: types.NewMsgNewWorkspace("testOwner", 0, 0),
			},
			want: 1,
		},
		{
			name: "PASS: wrong owner",
			args: args{
				req: &types.QueryWorkspacesByOwnerRequest{
					Owner: "wrongOwner",
				},
				msgWorkspace: types.NewMsgNewWorkspace("testOwner", 0, 0),
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)
			msgSer := keeper.NewMsgServerImpl(*ik)
			_, err := msgSer.NewWorkspace(goCtx, tt.args.msgWorkspace)
			if err != nil {
				t.Fatal(err)
			}
			got, err := ik.WorkspacesByOwner(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkspacesByOwner() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.want != len(got.Workspaces) {
				t.Errorf("WorkspacesByOwner() got = %v, want %v", got, tt.want)
			}
		})
	}
}
