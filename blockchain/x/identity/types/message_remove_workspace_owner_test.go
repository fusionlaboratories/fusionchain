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

func TestMsgRemoveWorkspaceOwner_NewMsgRemoveWorkspaceOwner(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgRemoveWorkspaceOwner
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgRemoveWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.Owner)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgRemoveWorkspaceOwner_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: MsgRemoveWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgRemoveWorkspaceOwner_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveWorkspaceOwner
	}{
		{
			name: "PASS: valid address",
			msg: MsgRemoveWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgRemoveWorkspaceOwner, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgRemoveWorkspaceOwner_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgRemoveWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveWorkspaceOwner{
				Creator:       "qredo1n7x7nv2urvdtc36tvhvc4dg6wfnnwh3cmt9j9w",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgRemoveWorkspaceOwner{
				Creator:       "invalid",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgRemoveWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.Owner)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgRemoveWorkspaceOwner_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgRemoveWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveWorkspaceOwner{
				Creator:       "qredo1nexzt4fcc84mgnqwjdhxg6veu97eyy9rgzkczs",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Owner:         "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgRemoveWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.Owner)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgRemoveWorkspaceOwner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveWorkspaceOwner
		err  error
	}{
		{
			name: "FAIL: invalid address",
			msg: MsgRemoveWorkspaceOwner{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "PASS: valid address",
			msg: MsgRemoveWorkspaceOwner{
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
