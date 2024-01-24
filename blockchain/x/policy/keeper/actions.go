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
package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/x/policy/types"
)

// RegisterActionHandler registers a handler for a specific action type.
// The handler function is called when the action is executed.
func RegisterActionHandler[ResT any](k *Keeper, actionType string, handlerFn func(sdk.Context, *types.Action, *codectypes.Any) (ResT, error)) {
	if _, ok := k.actionHandlers[actionType]; ok {
		// To be safe and prevent mistakes we shouldn't allow to register
		// multiple handlers for the same action type.
		// However, in the current implementation of Cosmos SDK, this is called
		// twice so we'll ignore the second call.

		// panic(fmt.Sprintf("action handler already registered for %s", actionType))
		return
	}
	k.actionHandlers[actionType] = func(ctx sdk.Context, a *types.Action, payload *codectypes.Any) (any, error) {
		return handlerFn(ctx, a, payload)
	}
}

func RegisterPolicyGeneratorHandler[ReqT any](k *Keeper, reqType string, handlerFn func(sdk.Context, ReqT) (policy.Policy, error)) {
	if _, ok := k.policyGeneratorHandlers[reqType]; ok {
		// To be safe and prevent mistakes we shouldn't allow to register
		// multiple handlers for the same action type.
		// However, in the current implementation of Cosmos SDK, this is called
		// twice so we'll ignore the second call.

		// panic(fmt.Sprintf("action handler already registered for %s", actionType))
		return
	}

	k.policyGeneratorHandlers[reqType] = func(ctx sdk.Context, a *codectypes.Any) (policy.Policy, error) {
		var m sdk.Msg
		if err := k.cdc.UnpackAny(a, &m); err != nil {
			return nil, err
		}
		return handlerFn(ctx, m.(ReqT))
	}
}

// TryExecuteAction checks if the policy attached to the action is satisfied
// and executes it.
//
// If the policy is satisfied, the provided handler function is executed and
// its response returned. If the policy is still not satisfied, nil is returned.
//
// This function should be called:
// - after an action is created
// - every time there is a change in the approvers set
func TryExecuteAction[ReqT sdk.Msg, ResT any](
	k *Keeper,
	cdc codec.BinaryCodec,
	ctx sdk.Context,
	act *types.Action,
	payload *codectypes.Any,
	handlerFn func(sdk.Context, ReqT) (*ResT, error),
) (*ResT, error) {
	var m sdk.Msg
	err := k.cdc.UnpackAny(act.Msg, &m)
	if err != nil {
		return nil, err
	}

	msg, ok := m.(ReqT)
	if !ok {
		return nil, fmt.Errorf("invalid message type, expected %T", new(ReqT))
	}

	pol, err := PolicyForAction(ctx, k, act)
	if err != nil {
		return nil, err
	}

	signersSet := policy.BuildApproverSet(act.Approvers)

	if err := pol.Verify(signersSet, policy.NewPolicyPayload(cdc, payload), act.GetPolicyDataMap()); err == nil {
		act.Status = types.ActionStatus_ACTION_STATUS_COMPLETED
		k.SetAction(ctx, act)
		return handlerFn(ctx, msg)
	}

	return nil, nil
}

func PolicyForAction(ctx sdk.Context, k *Keeper, act *types.Action) (policy.Policy, error) {
	var (
		pol policy.Policy
		err error
	)

	if act.PolicyId == 0 {
		// if no explicit policy ID specified, try to generate one
		polGen, found := k.policyGeneratorHandlers[act.Msg.TypeUrl]
		if !found {
			return nil, fmt.Errorf("no policy ID specied for action and no policy generator registered for %s", act.Msg.TypeUrl)
		}

		pol, err = polGen(ctx, act.Msg)
		if err != nil {
			return nil, err
		}
	} else {
		p, ok := k.PolicyRepo().Get(ctx, act.PolicyId)
		if !ok {
			return nil, fmt.Errorf("policy not found: %d", act.PolicyId)
		}

		pol, err = types.UnpackPolicy(k.cdc, p)
		if err != nil {
			return nil, err
		}
	}

	return pol, nil
}

// AddAction creates a new action for the provided message with initial approvers.
// Who calls this function should also immediately check if the action can be
// executed with the provided initialApprovers, by calling TryExecuteAction.
func (k Keeper) AddAction(ctx sdk.Context, creator string, msg sdk.Msg, policyID, btl uint64, policyData map[string][]byte) (*types.Action, error) {
	wrappedMsg, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	policyDataKv := make([]*types.KeyValue, 0, len(policyData))
	for k, v := range policyData {
		policyDataKv = append(policyDataKv, &types.KeyValue{Key: k, Value: v})
	}

	// create action object
	act := types.Action{
		Status:     types.ActionStatus_ACTION_STATUS_PENDING,
		Approvers:  []string{},
		PolicyId:   policyID,
		Msg:        wrappedMsg,
		Creator:    creator,
		Btl:        btl,
		PolicyData: policyDataKv,
	}

	// add initial approver
	pol, err := PolicyForAction(ctx, &k, &act)
	if err != nil {
		return nil, err
	}

	creatorAbbr, err := pol.AddressToParticipant(creator)
	if err != nil {
		return nil, err
	}

	if err := act.AddApprover(creatorAbbr); err != nil {
		return nil, err
	}

	// store and return generated action
	k.AppendAction(ctx, &act)
	return &act, nil
}
