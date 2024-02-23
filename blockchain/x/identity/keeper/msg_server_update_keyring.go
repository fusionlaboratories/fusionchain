// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/qredo/fusionchain/x/identity/types"
)

func (k msgServer) UpdateKeyring(goCtx context.Context, msg *types.MsgUpdateKeyring) (*types.MsgUpdateKeyringResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	kr := k.GetKeyring(ctx, msg.KeyringAddr)
	if kr == nil {
		return nil, fmt.Errorf("keyring not found")
	}

	// Check if the requester is an admin
	if !kr.IsAdmin(msg.Creator) {
		return nil, fmt.Errorf("keyring updates should be requested by admins")
	}

	kr.SetStatus(msg.IsActive)
	if msg.Description != "" {
		kr.SetDescription(msg.Description)
	}
	k.SetKeyring(ctx, kr)
	return &types.MsgUpdateKeyringResponse{}, nil
}
