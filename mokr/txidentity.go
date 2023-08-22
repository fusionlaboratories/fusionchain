package main

import (
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/go-bip39"
	"gitlab.qredo.com/qrdochain/fusionchain/crypto/ethsecp256k1"
)

// TxIdentity represents an account on the Fusion Chain.
// It will be used to sign transactions.
type TxIdentity struct {
	Address sdktypes.AccAddress
	PrivKey *ethsecp256k1.PrivKey
}

// NewTxIdentityFromSeed returns a TxIdentity from a seed phrase.
// This is useful in a mock environment or during testing but should not be used in production.
func NewTxIdentityFromSeed(seedPhrase string) (TxIdentity, error) {
	config := sdktypes.GetConfig()
	config.SetBech32PrefixForAccount(addrPrefix, addrPrefix+"pub")

	// Convert the seed phrase to a seed
	seedBytes, err := bip39.NewSeedWithErrorChecking(seedPhrase, "")
	if err != nil {
		return TxIdentity{}, fmt.Errorf("failed to convert seed phrase to seed: %w", err)
	}

	// Create a master key and derive the desired key
	masterKey, ch := hd.ComputeMastersFromSeed(seedBytes)
	derivedKey, err := hd.DerivePrivateKeyForPath(masterKey, ch, derivationPath)
	if err != nil {
		return TxIdentity{}, fmt.Errorf("failed to derive private key: %w", err)
	}

	// Generate a private key object from the bytes
	privKey, pubKey := btcec.PrivKeyFromBytes(derivedKey)
	if err != nil {
		return TxIdentity{}, fmt.Errorf("failed to generate private key: %w", err)
	}

	// Convert the public key to a Cosmos secp256k1.PublicKey
	cosmosPubKey := &ethsecp256k1.PubKey{Key: pubKey.SerializeCompressed()}
	ethermintPrivKey := &ethsecp256k1.PrivKey{
		Key: privKey.Serialize(),
	}

	// Get the address of the public key
	addr := sdktypes.AccAddress(cosmosPubKey.Address().Bytes())

	return TxIdentity{
		Address: addr,
		PrivKey: ethermintPrivKey,
	}, nil
}
