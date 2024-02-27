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

var sigPayloadECDSA = []byte{173, 224, 103, 159, 55, 251, 255, 212, 50, 145, 235, 181, 19, 14, 120, 168, 208, 151, 33, 204, 161, 79, 118, 229, 75, 22, 185, 234, 115, 125, 170, 101, 55, 197, 218, 94, 172, 32, 139, 21, 141, 104, 163, 109, 45, 47, 80, 110, 39, 5, 156, 88, 31, 82, 123, 246, 67, 21, 199, 126, 75, 222, 65, 115, 1}

func TestMsgFulfillSignatureRequest_NewMsgFulfillSignatureRequest(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgFulfilSignatureRequest
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgFulfilSignatureRequest{
				Creator:   sample.AccAddress(),
				RequestId: 1,
				Status:    SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result:    NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMsgFulfilSignatureRequest(tt.msg.Creator, tt.msg.RequestId, tt.msg.Status, tt.msg.Result)

			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgFulfillSignatureRequest_Route(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgFulfilSignatureRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgFulfilSignatureRequest{
				Creator:   sample.AccAddress(),
				RequestId: 1,
				Status:    SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result:    NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgFulfillSignatureRequest_Type(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgFulfilSignatureRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgFulfilSignatureRequest{
				Creator:   sample.AccAddress(),
				RequestId: 1,
				Status:    SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result:    NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgFulfilSignatureRequest, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgFulfillSignatureRequest_GetSigners(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgFulfilSignatureRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgFulfilSignatureRequest{
				Creator:   sample.AccAddress(),
				RequestId: 1,
				Status:    SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result:    NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgFulfilSignatureRequest{
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
				msg := NewMsgFulfilSignatureRequest(tt.msg.Creator, tt.msg.RequestId, tt.msg.Status, tt.msg.Result)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgFulfillSignatureRequest_GetSignBytes(t *testing.T) {

	tests := []struct {
		name string
		msg  *MsgFulfilSignatureRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgFulfilSignatureRequest{
				Creator:   sample.AccAddress(),
				RequestId: 1,
				Status:    SignRequestStatus_SIGN_REQUEST_STATUS_FULFILLED,
				Result:    NewMsgFulfilSignatureRequestPayload(sigPayloadECDSA),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgFulfilSignatureRequest(tt.msg.Creator, tt.msg.RequestId, tt.msg.Status, tt.msg.Result)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgFulfillSignatureRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgFulfilSignatureRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgFulfilSignatureRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgFulfilSignatureRequest{
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
