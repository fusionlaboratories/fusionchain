// Copyright 2023 Qredo Ltd.
// This file is part of the Fusion library.
//
// The Fusion library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Fusion library. If not, see https://github.com/qredo/fusionchain/blob/main/LICENSE
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/qredo/fusionchain/testutil/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgNewSignatureRequest_NewMsgNewSignatureRequest(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewSignatureRequest
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignatureRequest{
				Creator:        sample.AccAddress(),
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Btl:            1000,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &MsgNewSignatureRequest{tt.msg.Creator, tt.msg.KeyId, tt.msg.DataForSigning, tt.msg.Btl}
			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgNewSignatureRequest_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewSignatureRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewSignatureRequest{
				Creator:        sample.AccAddress(),
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Btl:            1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgNewSignatureRequest_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewSignatureRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewSignatureRequest{
				Creator:        sample.AccAddress(),
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Btl:            1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgNewSignatureRequest, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgNewSignatureRequest_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewSignatureRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignatureRequest{
				Creator:        sample.AccAddress(),
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Btl:            1000,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgNewSignatureRequest{
				Creator: "invalidCreator",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgNewSignatureRequest(tt.msg.Creator, tt.msg.KeyId, tt.msg.DataForSigning, tt.msg.Btl)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgNewSignatureRequest_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewSignatureRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignatureRequest{
				Creator:        sample.AccAddress(),
				KeyId:          1,
				DataForSigning: []byte("778f572f33afab831365d52e563a0ddd"),
				Btl:            1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgNewSignatureRequest(tt.msg.Creator, tt.msg.KeyId, tt.msg.DataForSigning, tt.msg.Btl)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgNewSignatureRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewSignatureRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNewSignatureRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgNewSignatureRequest{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
