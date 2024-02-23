// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
const TypeMsgRemoveKeyringParty = "remove_keyring_party"

func NewMsgRemoveKeyringParty(creator, keyringAddr, party string) *MsgRemoveKeyringParty {
	return &MsgRemoveKeyringParty{
		Creator:     creator,
		KeyringAddr: keyringAddr,
		Party:       party,
	}
}

func (msg *MsgRemoveKeyringParty) Route() string { return RouterKey }

func (msg *MsgRemoveKeyringParty) Type() string {
	return TypeMsgRemoveKeyringParty
}

func (msg *MsgRemoveKeyringParty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRemoveKeyringParty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveKeyringParty) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, err = sdk.AccAddressFromBech32(msg.Party)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid party address (%s)", err)
	}
	return nil
}
