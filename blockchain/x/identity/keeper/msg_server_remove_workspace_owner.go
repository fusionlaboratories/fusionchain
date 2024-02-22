// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
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

func (k msgServer) RemoveWorkspaceOwner(goCtx context.Context, msg *types.MsgRemoveWorkspaceOwner) (*types.MsgRemoveWorkspaceOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
	if ws == nil {
		return nil, fmt.Errorf("workspace not found")
	}

	if !ws.IsOwner(msg.Creator) {
		return nil, fmt.Errorf("creator is not an owner of the workspace")
	}

	if !ws.IsOwner(msg.Owner) {
		return nil, fmt.Errorf("new owner is not an owner of the workspace")
	}

	act, err := k.policyKeeper.AddAction(ctx, msg.Creator, msg, ws.AdminPolicyId, msg.Btl, nil)
	if err != nil {
		return nil, err
	}
	return k.RemoveOwnerActionHandler(ctx, act, &cdctypes.Any{})
}

func (k msgServer) RemoveOwnerPolicyGenerator(ctx sdk.Context, msg *types.MsgRemoveWorkspaceOwner) (policy.Policy, error) {
	ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
	if ws == nil {
		return nil, fmt.Errorf("workspace not found")
	}

	pol := ws.PolicyRemoveOwner()
	return pol, nil
}

func (k msgServer) RemoveOwnerActionHandler(ctx sdk.Context, act *bbirdtypes.Action, payload *cdctypes.Any) (*types.MsgRemoveWorkspaceOwnerResponse, error) {
	return bbird.TryExecuteAction(
		k.policyKeeper,
		k.cdc,
		ctx,
		act,
		payload,
		func(ctx sdk.Context, msg *types.MsgRemoveWorkspaceOwner) (*types.MsgRemoveWorkspaceOwnerResponse, error) {
			ws := k.GetWorkspace(ctx, msg.WorkspaceAddr)
			if ws == nil {
				return nil, fmt.Errorf("workspace not found")
			}

			if !ws.IsOwner(msg.Owner) {
				return nil, fmt.Errorf("owner does not exist")
			}

			ws.RemoveOwner(msg.Owner)

			k.SetWorkspace(ctx, ws)

			return &types.MsgRemoveWorkspaceOwnerResponse{}, nil
		},
	)
}
