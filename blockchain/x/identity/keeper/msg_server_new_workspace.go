package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/identity/types"
)

func (k msgServer) NewWorkspace(goCtx context.Context, msg *types.MsgNewWorkspace) (*types.MsgNewWorkspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	workspace := &types.Workspace{
		Creator: msg.Creator,
		Owners:  []string{msg.Creator},
	}
	addr := k.CreateWorkspace(ctx, workspace)

	return &types.MsgNewWorkspaceResponse{
		Address: addr,
	}, nil
}
