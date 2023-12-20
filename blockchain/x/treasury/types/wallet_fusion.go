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
	"crypto/ecdsa"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type FusionWallet struct {
	key *ecdsa.PublicKey
}

var _ Wallet = &FusionWallet{}

func NewFusionWallet(k *Key) (*FusionWallet, error) {
	pubkey, err := k.ToECDSASecp256k1()
	if err != nil {
		return nil, err
	}
	return &FusionWallet{key: pubkey}, nil
}

func (w *FusionWallet) Address() string {
	var pubkey secp256k1.PubKey
	pubkey.Key = crypto.CompressPubkey(w.key)
	bech32Address := sdk.AccAddress(pubkey.Address().Bytes()).String()
	return bech32Address
}
