// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/qredo/fusionchain/testutil/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgNewKeyRequest_NewMsgNewKeyRequest(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       1,
				Btl:           1000,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &MsgNewKeyRequest{tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.KeyringAddr, tt.msg.KeyType, tt.msg.Btl}
			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgNewKeyRequest_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       1,
				Btl:           1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgNewKeyRequest_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       1,
				Btl:           1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgNewKeyRequest, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgNewKeyRequest_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       1,
				Btl:           1000,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgNewKeyRequest{
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
				msg := NewMsgNewKeyRequest(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.KeyringAddr, tt.msg.KeyType, tt.msg.Btl)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgNewKeyRequest_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewKeyRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewKeyRequest{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				KeyringAddr:   "qredokeyring1ph63us46lyw56vrzgaq",
				KeyType:       1,
				Btl:           1000,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgNewKeyRequest(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.KeyringAddr, tt.msg.KeyType, tt.msg.Btl)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgNewKeyRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewKeyRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNewKeyRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgNewKeyRequest{
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
