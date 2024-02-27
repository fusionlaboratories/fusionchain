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

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/qredo/fusionchain/testutil/sample"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var metadataTest = MetadataEthereum{
	ChainId: 11155111,
}

var metadataAny *cdctypes.Any

var unsignedTx1 = []byte{2, 65, 62, 38, 30, 38, 35, 30, 34, 61, 38, 31, 37, 63, 38, 30, 30, 38, 32, 35, 32, 30, 38, 39, 34, 39, 39, 33, 66, 34, 35, 36, 36, 36, 62, 32, 61, 37, 38, 34, 33, 34, 37, 31, 31, 64, 31, 61, 32, 30, 64, 32, 61, 39, 37, 33, 33, 63, 30, 37, 61, 35, 33, 31, 38, 38, 37, 30, 65, 33, 35, 66, 61, 39, 33, 31, 61, 30, 30, 30, 30, 38, 30, 38, 30, 38, 30, 38, 30}

func TestMsgNewSignTransactionRequest_NewMsgNewSignTransactionRequest(t *testing.T) {

	metadataAny, _ = cdctypes.NewAnyWithValue(&metadataTest)

	tests := []struct {
		name string
		msg  *MsgNewSignTransactionRequest
		err  error
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignTransactionRequest{
				Creator:             sample.AccAddress(),
				KeyId:               1,
				WalletType:          WalletType_WALLET_TYPE_FUSION,
				UnsignedTransaction: unsignedTx1,
				Btl:                 1000,
				Metadata:            metadataAny,
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := &MsgNewSignTransactionRequest{tt.msg.Creator, tt.msg.KeyId, tt.msg.WalletType, tt.msg.UnsignedTransaction, tt.msg.Btl, tt.msg.Metadata}
			assert.Equalf(t, tt.msg, got, "want", tt.msg)
		})
	}
}

func TestMsgNewSignTransactionRequest_Route(t *testing.T) {

	metadataAny, _ = cdctypes.NewAnyWithValue(&metadataTest)

	tests := []struct {
		name string
		msg  *MsgNewSignTransactionRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewSignTransactionRequest{
				Creator:             sample.AccAddress(),
				KeyId:               1,
				WalletType:          WalletType_WALLET_TYPE_FUSION,
				UnsignedTransaction: unsignedTx1,
				Btl:                 1000,
				Metadata:            metadataAny,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, ModuleName, tt.msg.Route(), "Route()")
		})
	}
}

func TestMsgNewSignTransactionRequest_Type(t *testing.T) {

	metadataAny, _ = cdctypes.NewAnyWithValue(&metadataTest)

	tests := []struct {
		name string
		msg  *MsgNewSignTransactionRequest
	}{
		{
			name: "PASS: valid address",
			msg: &MsgNewSignTransactionRequest{
				Creator:             sample.AccAddress(),
				KeyId:               1,
				WalletType:          WalletType_WALLET_TYPE_FUSION,
				UnsignedTransaction: unsignedTx1,
				Btl:                 1000,
				Metadata:            metadataAny,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, TypeMsgNewSignTransactionRequest, tt.msg.Type(), "Type()")
		})
	}
}

func TestMsgNewSignTransactionRequest_GetSigners(t *testing.T) {

	metadataAny, _ = cdctypes.NewAnyWithValue(&metadataTest)

	tests := []struct {
		name string
		msg  *MsgNewSignTransactionRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignTransactionRequest{
				Creator:             sample.AccAddress(),
				KeyId:               1,
				WalletType:          WalletType_WALLET_TYPE_FUSION,
				UnsignedTransaction: unsignedTx1,
				Btl:                 1000,
				Metadata:            metadataAny,
			},
		},
		{
			name: "FAIL: invalid signer",
			msg: &MsgNewSignTransactionRequest{
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
				msg := NewMsgNewSignTransactionRequest(tt.msg.Creator, tt.msg.KeyId, tt.msg.WalletType, tt.msg.UnsignedTransaction, tt.msg.Btl, tt.msg.Metadata)
				got := msg.GetSigners()

				assert.Equal(t, []sdk.AccAddress{acc}, got)
			}
		})
	}
}

func TestMsgNewSignTransactionRequest_GetSignBytes(t *testing.T) {

	metadataAny, _ = cdctypes.NewAnyWithValue(&metadataTest)

	tests := []struct {
		name string
		msg  *MsgNewSignTransactionRequest
	}{
		{
			name: "PASS: happy path",
			msg: &MsgNewSignTransactionRequest{
				Creator:             sample.AccAddress(),
				KeyId:               1,
				WalletType:          WalletType_WALLET_TYPE_FUSION,
				UnsignedTransaction: unsignedTx1,
				Btl:                 1000,
				Metadata:            nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := NewMsgNewSignTransactionRequest(tt.msg.Creator, tt.msg.KeyId, tt.msg.WalletType, tt.msg.UnsignedTransaction, tt.msg.Btl, tt.msg.Metadata)
			got := msg.GetSignBytes()

			bz := ModuleCdc.MustMarshalJSON(msg)
			sortedBz := sdk.MustSortJSON(bz)

			require.Equal(t, sortedBz, got, "GetSignBytes() result doesn't match sorted JSON bytes")

		})
	}
}

func TestMsgNewSignTransactionRequest_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgNewSignTransactionRequest
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgNewSignTransactionRequest{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgNewSignTransactionRequest{
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
