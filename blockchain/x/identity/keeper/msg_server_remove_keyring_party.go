package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/identity/types"
)

func (k msgServer) RemoveKeyringParty(goCtx context.Context, msg *types.MsgRemoveKeyringParty) (*types.MsgRemoveKeyringPartyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kr := k.GetKeyring(ctx, msg.KeyringAddr)
	if kr == nil {
		return nil, fmt.Errorf("keyring not found")
	}

	if !kr.IsActive {
		return nil, fmt.Errorf("keyring is inactive")
	}

	if !kr.IsParty(msg.Party) {
		return nil, fmt.Errorf("party is not a party of the keyring")
	}

	if !kr.IsAdmin(msg.Creator) {
		return nil, fmt.Errorf("tx creator is no keyring admin")
	}

	kr.RemoveParty(msg.Party)
	k.SetKeyring(ctx, kr)

	return &types.MsgRemoveKeyringPartyResponse{}, nil
}
