// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/identity/types"
)

func (k msgServer) AddKeyringParty(goCtx context.Context, msg *types.MsgAddKeyringParty) (*types.MsgAddKeyringPartyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kr := k.GetKeyring(ctx, msg.KeyringAddr)
	if kr == nil {
		return nil, fmt.Errorf("keyring not found")
	}

	if !kr.IsActive {
		return nil, fmt.Errorf("keyring is inactive")
	}

	if kr.IsParty(msg.Party) {
		return nil, fmt.Errorf("party is already a party of the keyring")
	}

	if !kr.IsAdmin(msg.Creator) {
		return nil, fmt.Errorf("tx creator is no keyring admin")
	}

	kr.AddParty(msg.Party)
	k.SetKeyring(ctx, kr)

	return &types.MsgAddKeyringPartyResponse{}, nil
}
