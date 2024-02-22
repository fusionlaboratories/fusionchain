// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
		return nil, err
	}
	return &SuiWallet{key: pubkey}, nil
}

func (w *SuiWallet) Address() string {
	tmp := []byte{signatureSchemeFlagED25519}
	tmp = append(tmp, *w.key...)
	addrBytes := blake2b.Sum256(tmp)
	suiAddress := "0x" + hex.EncodeToString(addrBytes[:])[:addressLength]
	return suiAddress
}
