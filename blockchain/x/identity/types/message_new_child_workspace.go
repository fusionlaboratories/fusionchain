// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
const TypeMsgNewChildWorkspace = "new_child_workspace"

var _ sdk.Msg = &MsgNewChildWorkspace{}

func NewMsgNewChildWorkspace(creator, parentWorkspaceAddr string, btl uint64) *MsgNewChildWorkspace {
	return &MsgNewChildWorkspace{
		Creator:             creator,
		ParentWorkspaceAddr: parentWorkspaceAddr,
		Btl:                 btl,
	}
}

func (msg *MsgNewChildWorkspace) Route() string {
	return RouterKey
}

func (msg *MsgNewChildWorkspace) Type() string {
	return TypeMsgNewChildWorkspace
}

func (msg *MsgNewChildWorkspace) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewChildWorkspace) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewChildWorkspace) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
