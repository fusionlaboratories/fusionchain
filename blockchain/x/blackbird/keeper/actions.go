package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/policy"
	"github.com/qredo/fusionchain/x/blackbird/types"
)

// RegisterActionHandler registers a handler for a specific action type.
// The handler function is called when the action is executed.
func RegisterActionHandler[ResT any](k *Keeper, actionType string, handlerFn func(sdk.Context, *types.Action) (ResT, error)) {
	if _, ok := k.actionHandlers[actionType]; ok {
		// To be safe and prevent mistakes we shouldn't allow to register
		// multiple handlers for the same action type.
		// However, in the current implementation of Cosmos SDK, this is called
		// twice so we'll ignore the second call.

		// panic(fmt.Sprintf("action handler already registered for %s", actionType))
		return
	}
	k.actionHandlers[actionType] = func(ctx sdk.Context, a *types.Action) (any, error) {
		return handlerFn(ctx, a)
	}
}

// TryExecuteAction uses the provided policy function to determine what is the
// policy currently being applied for the action's message.
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
	policyFn func(sdk.Context, ReqT) (policy.Policy, error),
	handlerFn func(sdk.Context, ReqT) (*ResT, error),
) (*ResT, error) {
	var m sdk.Msg
	err := cdc.UnpackAny(act.Msg, &m)
	if err != nil {
		return nil, err
	}

	msg, ok := m.(ReqT)
	if !ok {
		return nil, fmt.Errorf("invalid message type, expected %T", new(ReqT))
	}

	pol, err := policyFn(ctx, msg)
	if err != nil {
		return nil, err
	}

	signersSet := policy.BuildApproverSet(act.Approvers)

	// Execute action if policy is satified
	if err := pol.Verify(signersSet); err == nil {
		act.Completed = true
		k.SetAction(ctx, act)
		return handlerFn(ctx, msg)
	}

	return nil, nil
}

// AddAction creates a new action for the provided message with initial approvers.
// Who calls this function should also immediately check if the action can be
// executed with the provided initialApprovers.
func (k Keeper) AddAction(ctx sdk.Context, msg sdk.Msg, initialApprovers ...string) (*types.Action, error) {
	wrappedMsg, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	act := types.Action{
		Approvers: initialApprovers,
		Msg:       wrappedMsg,
	}
	k.AppendAction(ctx, &act)
	return &act, nil
}
