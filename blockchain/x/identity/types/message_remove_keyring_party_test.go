// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgRemoveKeyringParty_NewMsgRemoveKeyringParty(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRemoveKeyringParty
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveKeyringParty{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgRemoveKeyringParty(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Party)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgRemoveKeyringParty_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveKeyringParty
	}{
		{
			name: "PASS: valid address",
			msg: MsgRemoveKeyringParty{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgRemoveKeyringParty_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveKeyringParty
	}{
		{
			name: "PASS: valid address",
			msg: MsgRemoveKeyringParty{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgRemoveKeyringParty, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgRemoveKeyringParty_GetSigners(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRemoveKeyringParty
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveKeyringParty{
				Creator:     "qredo1n7x7nv2urvdtc36tvhvc4dg6wfnnwh3cmt9j9w",
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5"},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgRemoveKeyringParty{
				Creator:     "invalid",
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acc, err := sdk.AccAddressFromBech32(tt.msg.Creator)
			if err != nil {
				assert.Panics(t, func() { tt.msg.GetSigners() })
			} else {
				msg := NewMsgRemoveKeyringParty(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Party)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgRemoveKeyringParty_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgRemoveKeyringParty
	}{
		{
			name: "PASS: happy path",
			msg: &MsgRemoveKeyringParty{
				Creator:     "qredo1nexzt4fcc84mgnqwjdhxg6veu97eyy9rgzkczs",
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgRemoveKeyringParty(tt.msg.Creator, tt.msg.KeyringAddr, tt.msg.Party)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgRemoveKeyringParty_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgRemoveKeyringParty
		err  error
	}{
		{
			name: "PASS: valid address",
			msg: MsgRemoveKeyringParty{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
		},
		{
			name: "FAIL: invalid address",
			msg: MsgRemoveKeyringParty{
				Creator:     "invalid_address",
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "qredo1s3qj9p0ymugy6chyrwy3ft2s5u24fc320vdvv5",
			},
			err: sdkerrors.ErrInvalidAddress,
		},
		{
			name: "FAIL: invalid party address",
			msg: MsgRemoveKeyringParty{
				Creator:     sample.AccAddress(),
				KeyringAddr: "qredokeyring1ph63us46lyw56vrzgaq",
				Party:       "invalid_address",
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
