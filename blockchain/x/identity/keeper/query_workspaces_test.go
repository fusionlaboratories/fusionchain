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

func TestKeeper_Workspaces(t *testing.T) {

	type args struct {
		req            *types.QueryWorkspacesRequest
		msgWorkspace   *types.MsgNewWorkspace
		workspaceCount int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "PASS: create 100 workspaces",
			args: args{
				req: &types.QueryWorkspacesRequest{
					Pagination: nil,
				},
				msgWorkspace:   types.NewMsgNewWorkspace("testOwner", 0, 0),
				workspaceCount: 100,
			},
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			keepers := keepertest.NewTest(t)
			ik := keepers.IdentityKeeper
			ctx := keepers.Ctx
			goCtx := sdk.WrapSDKContext(ctx)
			for i := 0; i < tt.args.workspaceCount; i++ {
				msgSer := keeper.NewMsgServerImpl(*ik)
				_, err := msgSer.NewWorkspace(goCtx, tt.args.msgWorkspace)
				if err != nil {
					t.Fatal(err)
				}
			}
			got, err := ik.Workspaces(goCtx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Workspaces() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got.Workspaces) != tt.want {
				t.Errorf("Workspaces() got = %v, want %v", got, tt.want)
				return
			}
		})
	}
}
