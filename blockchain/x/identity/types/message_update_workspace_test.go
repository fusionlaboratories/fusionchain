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

func TestMsgUpdateWorkspace_NewMsgUpdateWorkspace(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateWorkspace
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 0,
				SignPolicyId:  0,
				Btl:           100,
			},
			err: nil,
		},
		{
			name: "PASS: update policies",
			msg: &MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 1,
				SignPolicyId:  1,
				Btl:           100,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgUpdateWorkspace(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.AdminPolicyId, tt.msg.SignPolicyId, tt.msg.Btl)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgUpdateWorkspace_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 0,
				SignPolicyId:  0,
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

func TestMsgUpdateWorkspace_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWorkspace
	}{
		{
			name: "PASS: valid address",
			msg: MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 0,
				SignPolicyId:  0,
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgUpdateWorkspace, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgUpdateWorkspace_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: &MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 0,
				SignPolicyId:  0,
				Btl:           100,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgUpdateWorkspace{
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
				msg := NewMsgUpdateWorkspace(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.AdminPolicyId, tt.msg.SignPolicyId, tt.msg.Btl)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgUpdateWorkspace_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgUpdateWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: &MsgUpdateWorkspace{
				Creator:       sample.AccAddress(),
				WorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				AdminPolicyId: 0,
				SignPolicyId:  0,
				Btl:           100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgUpdateWorkspace(tt.msg.Creator, tt.msg.WorkspaceAddr, tt.msg.AdminPolicyId, tt.msg.SignPolicyId, tt.msg.Btl)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgUpdateWorkspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateWorkspace
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateWorkspace{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "PASS: valid address",
			msg: MsgUpdateWorkspace{
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
