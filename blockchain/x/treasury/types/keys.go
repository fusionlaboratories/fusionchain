// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

const (
	// ModuleName defines the module name
	ModuleName = "treasury"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_treasury"

	KeyRequestKey = "key_request/value/"

	KeyRequestCountKey = "key_request/count"

	KeyKey = "key/value/"

	KeyCountKey = "key/count"

	SignRequestKey = "sign_request/value/"

	SignRequestCountKey = "sign_request/count"

	WalletKey = "wallet/value/"

	WalletCountKey = "wallet/count"

	SignTransactionRequestKey = "sign_transaction_request/value/"

	SignTransactionRequestCountKey = "sign_transaction_request/count"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
