// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddWorkspaceOwner = "add_workspace_owner"

var _ sdk.Msg = &MsgAddWorkspaceOwner{}

func NewMsgAddWorkspaceOwner(creator string, wsAddr string, newOwner string, btl uint64) *MsgAddWorkspaceOwner {
	return &MsgAddWorkspaceOwner{
		Creator:       creator,
		WorkspaceAddr: wsAddr,
		NewOwner:      newOwner,
		Btl:           btl,
	}
}

func (msg *MsgAddWorkspaceOwner) Route() string {
	return RouterKey
}

func (msg *MsgAddWorkspaceOwner) Type() string {
	return TypeMsgAddWorkspaceOwner
}

func (msg *MsgAddWorkspaceOwner) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddWorkspaceOwner) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddWorkspaceOwner) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.NewOwner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid new owner address (%s)", err)
	}
	return nil
}
