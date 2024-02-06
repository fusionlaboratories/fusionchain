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

func TestMsgAddWorkspaceOwner_NewMsgAddWorkspaceOwner(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgAddWorkspaceOwner
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgAddWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgAddWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.NewOwner, 100)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgAddWorkspaceOwner_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: MsgAddWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgAddWorkspaceOwner_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddWorkspaceOwner
	}{
		{
			name: "PASS: valid address",
			msg: MsgAddWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgAddWorkspaceOwner, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgAddWorkspaceOwner_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgAddWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: &MsgAddWorkspaceOwner{
				Creator:       "qredo1n7x7nv2urvdtc36tvhvc4dg6wfnnwh3cmt9j9w",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgAddWorkspaceOwner{
				Creator:       "invalid",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgAddWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.NewOwner, tt.msg.Btl)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgAddWorkspaceOwner_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgAddWorkspaceOwner
	}{
		{
			name: "PASS: happy path",
			msg: &MsgAddWorkspaceOwner{
				Creator:       "qredo1nexzt4fcc84mgnqwjdhxg6veu97eyy9rgzkczs",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgAddWorkspaceOwner(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.NewOwner, tt.msg.Btl)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgAddWorkspaceOwner_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgAddWorkspaceOwner
		err  error
	}{
		{
			name: "PASS: valid address",
			msg: MsgAddWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
		},
		{
			name: "FAIL: invalid address",
			msg: MsgAddWorkspaceOwner{
				Creator:       "invalidAddress",
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
				Btl:           100,
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "FAIL: invalid party address",
			msg: MsgAddWorkspaceOwner{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				NewOwner:      "invalidOwner",
				Btl:           100,
			},
			err: sdkerrors.ErrInvalidAddress,
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
