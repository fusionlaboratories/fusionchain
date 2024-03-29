// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package mpcrelayer

import (
	"time"

	"github.com/cosmos/go-bip39"
)

const (
	mpcRequestKeyLength = 64
	mnemonicKey         = "mnemonic"
	rateLimitPerSecond  = 5
)

// GenerateMnemonic creates a fresh BIP39 mnemonic with 256-bit entropy. TODO - create a shared crypto package
func GenerateMnemonic() (string, error) {
	e, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(e)
}

func requeueKeyItemWithTimeout(c chan *keyRequestQueueItem, item *keyRequestQueueItem, timeout time.Duration) {
	item.retries++
	go func() {
		time.Sleep(timeout)
		c <- item
	}()
}

func requeueSigItemWithTimeout(c chan *signatureRequestQueueItem, item *signatureRequestQueueItem, timeout time.Duration) {
	item.retries++
	go func() {
		time.Sleep(timeout)
		c <- item
	}()
}

func makeThreads(n int) chan struct{} {
	t := make(chan struct{}, n)
	for i := 0; i < n; i++ {
		t <- struct{}{}
	}
	return t
}
