// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
		Status:    status,
		Result:    result,
	}
}

func (msg *MsgFulfilSignatureRequest) Route() string {
	return RouterKey
}

func (msg *MsgFulfilSignatureRequest) Type() string {
	return TypeMsgFulfilSignatureRequest
}

func (msg *MsgFulfilSignatureRequest) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFulfilSignatureRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFulfilSignatureRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
