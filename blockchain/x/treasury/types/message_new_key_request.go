// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
		KeyringAddr:   keyringAddr,
		KeyType:       keyType,
		Btl:           btl,
	}
}

func (msg *MsgNewKeyRequest) Route() string {
	return RouterKey
}

func (msg *MsgNewKeyRequest) Type() string {
	return TypeMsgNewKeyRequest
}

func (msg *MsgNewKeyRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewKeyRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewKeyRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
