package keeper

import (
	"context"
	"fmt"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/x/identity/types"
	bbird "github.com/qredo/fusionchain/x/policy/keeper"
	bbirdtypes "github.com/qredo/fusionchain/x/policy/types"
)

func (k msgServer) UpdateWorkspace(goCtx context.Context, msg *types.MsgUpdateWorkspace) (*types.MsgUpdateWorkspaceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
	if ws == nil {
		return nil, fmt.Errorf("workspace not found")
	}

	act, err := k.policyKeeper.AddAction(ctx, msg, ws.AdminPolicyId, msg.Creator)
	if err != nil {
		return nil, err
	}

	return k.UpdateWorkspaceActionHandler(ctx, act, &cdctypes.Any{})
}

func (k msgServer) UpdateWorkspacePolicyGenerator(ctx sdk.Context, msg *types.MsgUpdateWorkspace) (policy.Policy, error) {
	ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
	if ws == nil {
		return nil, fmt.Errorf("workspace not found")
	}

	pol := ws.PolicyUpdateWorkspace()
	return pol, nil
}

func (k msgServer) UpdateWorkspaceActionHandler(ctx sdk.Context, act *bbirdtypes.Action, payload *cdctypes.Any) (*types.MsgUpdateWorkspaceResponse, error) {
	return bbird.TryExecuteAction(
		k.policyKeeper,
		k.cdc,
		ctx,
		act,
		payload,
		func(ctx sdk.Context, msg *types.MsgUpdateWorkspace) (*types.MsgUpdateWorkspaceResponse, error) {
			ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
			if ws == nil {
				return nil, fmt.Errorf("workspace not found")
			}

			if msg.AdminPolicyId != ws.AdminPolicyId {
				_, found := k.policyKeeper.PolicyRepo().Get(ctx, msg.AdminPolicyId)
				if !found {
					return nil, fmt.Errorf("admin policy not found")
				}
				ws.AdminPolicyId = msg.AdminPolicyId
			}

			if msg.SignPolicyId != ws.SignPolicyId {
				_, found := k.policyKeeper.PolicyRepo().Get(ctx, msg.SignPolicyId)
				if !found {
					return nil, fmt.Errorf("sign policy not found")
				}
				ws.SignPolicyId = msg.SignPolicyId
			}

			k.SetWorkspace(ctx, ws)

			return &types.MsgUpdateWorkspaceResponse{}, nil
		},
	)
}
