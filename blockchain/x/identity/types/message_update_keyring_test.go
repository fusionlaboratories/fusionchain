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

func TestMsgUpdateKeyring_NewMsgUpdateKeyring(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateKeyring
		err  error
	}{
		{
			name: "PASS: new description",
			msg: &MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "new_description",
				IsActive:    true,
			},
			err: nil,
		},
		{
			name: "PASS: new status",
			msg: &MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "testKeyring",
				IsActive:    false,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgUpdateKeyring(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Description, tt.msg.IsActive)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgUpdateKeyring_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKeyring
	}{
		{
			name: "PASS: happy path",
			msg: MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "new_description",
				IsActive:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgUpdateKeyring_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKeyring
	}{
		{
			name: "PASS: valid address",
			msg: MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "new_description",
				IsActive:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgUpdateKeyring, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgUpdateKeyring_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateKeyring
	}{
		{
			name: "PASS: happy path",
			msg: &MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "new_description",
				IsActive:    true,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgUpdateKeyring{
				Creator: "invalid",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgUpdateKeyring(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Description, tt.msg.IsActive)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgUpdateKeyring_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateKeyring
	}{
		{
			name: "PASS: happy path",
			msg: &MsgUpdateKeyring{
				Creator:     "qredo1nexzt4fcc84mgnqwjdhxg6veu97eyy9rgzkczs",
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Description: "new_description",
				IsActive:    true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgUpdateKeyring(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Description, tt.msg.IsActive)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgUpdateKeyring_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateKeyring
		err  error
	}{
		{
			name: "PASS: valid address",
			msg: MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				Description: "testDescription",
			},
		},
		{
			name: "FAIL: invalid address",
			msg: MsgUpdateKeyring{
				Creator:     "invalid_address",
				Description: "testDescription",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "FAIL: emtpy description",
			msg: MsgUpdateKeyring{
				Creator:     sample.AccAddress(),
				Description: "",
			},
			err: ErrEmptyDesc,
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
