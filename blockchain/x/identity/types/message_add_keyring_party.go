// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
const TypeMsgAddKeyringParty = "add_keyring_party"

var _ sdk.Msg = &MsgAddKeyringParty{}

func NewMsgAddKeyringParty(creator, keyringAddr, party string) *MsgAddKeyringParty {
	return &MsgAddKeyringParty{
		Creator:     creator,
		KeyringAddr: keyringAddr,
		Party:       party,
	}
}

func (msg *MsgAddKeyringParty) Route() string {
	return RouterKey
}

func (msg *MsgAddKeyringParty) Type() string {
	return TypeMsgAddKeyringParty
}

func (msg *MsgAddKeyringParty) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddKeyringParty) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddKeyringParty) ValidateBasic() error {
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
