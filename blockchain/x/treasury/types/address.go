// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
	"golang.org/x/crypto/blake2b"
)

func FusionChainAddress(key *Key) (string, error) {
	k, err := key.ToECDSASecp256k1()
	if err != nil {
		return "", err
	}
	var pubkey secp256k1.PubKey
	pubkey.Key = crypto.CompressPubkey(k)
	bech32Address := sdk.AccAddress(pubkey.Address().Bytes()).String()
	return bech32Address, nil
}

func EthereumAddress(key *Key) (string, error) {
	k, err := key.ToECDSASecp256k1()
	if err != nil {
		return "", err
	}
	addr := crypto.PubkeyToAddress(*k)
	return addr.Hex(), nil
}

func CelestiaAddress(key *Key) (string, error) {
	k, err := key.ToECDSASecp256k1()
	if err != nil {
		return "", err
	}
	var pubkey secp256k1.PubKey
	pubkey.Key = crypto.CompressPubkey(k)
	bech32Address := sdk.MustBech32ifyAddressBytes("celestia", pubkey.Address())
	return bech32Address, nil
}

func SuiAddress(key *Key) (string, error) {
	k, err := key.ToEdDSAEd25519()
	if err != nil {
		return "", err
	}
	tmp := []byte{signatureSchemeFlagED25519}
	tmp = append(tmp, *k...)
	addrBytes := blake2b.Sum256(tmp)
	suiAddress := "0x" + hex.EncodeToString(addrBytes[:])[:addressLength]
	return suiAddress, nil
}
