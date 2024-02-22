// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRemoveWorkspaceOwner = "remove_workspace_owner"

var _ sdk.Msg = &MsgRemoveWorkspaceOwner{}

func NewMsgRemoveWorkspaceOwner(creator string, wsAddr string, owner string) *MsgRemoveWorkspaceOwner {
	return &MsgRemoveWorkspaceOwner{
		Creator:       creator,
		WorkspaceAddr: wsAddr,
		Owner:         owner,
	}
}

func (msg *MsgRemoveWorkspaceOwner) Route() string {
	return RouterKey
}

func (msg *MsgRemoveWorkspaceOwner) Type() string {
	return TypeMsgRemoveWorkspaceOwner
}

func (msg *MsgRemoveWorkspaceOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveWorkspaceOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveWorkspaceOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
