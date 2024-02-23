// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
		msg  *MsgNewChildWorkspace
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewChildWorkspace{
				Creator:             sample.AccAddress(),
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgNewChildWorkspace(tt.msg.Creator, tt.msg.ParentWorkspaceAddr, 100)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgNewChildWorkspace_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewChildWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: MsgNewChildWorkspace{
				Creator:             sample.AccAddress(),
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgNewChildWorkspace_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewChildWorkspace
	}{
		{
			name: "PASS: valid address",
			msg: MsgNewChildWorkspace{
				Creator:             sample.AccAddress(),
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgNewChildWorkspace, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgNewChildWorkspace_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewChildWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewChildWorkspace{
				Creator:             "qredo1n7x7nv2urvdtc36tvhvc4dg6wfnnwh3cmt9j9w",
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgNewChildWorkspace{
				Creator:             "invalid",
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgNewChildWorkspace(tt.msg.Creator, tt.msg.ParentWorkspaceAddr, tt.msg.Btl)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgNewChildWorkspace_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgNewChildWorkspace
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewChildWorkspace{
				Creator:             "qredo1nexzt4fcc84mgnqwjdhxg6veu97eyy9rgzkczs",
				ParentWorkspaceAddr: "qredoworkspace14a2hpadpsy9h5m6us54",
				Btl:                 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgNewChildWorkspace(tt.msg.Creator, tt.msg.ParentWorkspaceAddr, tt.msg.Btl)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgNewChildWorkspace_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewChildWorkspace
		err  error
	}{
		{
			name: "FAIL: invalid address",
			msg: MsgNewChildWorkspace{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "PASS: valid address",
			msg: MsgNewChildWorkspace{
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
