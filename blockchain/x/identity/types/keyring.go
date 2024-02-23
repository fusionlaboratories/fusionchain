// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package types

// nolint:stylecheck,st1003
// revive:disable-next-line var-naming

func (k *Keyring) IsParty(address string) bool {
	for _, party := range k.Parties {
		if party == address {
			return true
		}
	}
	return false
}

func (k *Keyring) IsAdmin(address string) bool {
	for _, admin := range k.Admins {
		if admin == address {
			return true
		}
	}
	return false
}

func (k *Keyring) AddParty(address string) {
	k.Parties = append(k.Parties, address)
}

func (k *Keyring) RemoveParty(address string) {
	for i, party := range k.Parties {
		if party == address {
			k.Parties = append(k.Parties[:i], k.Parties[i+1:]...)
			return
		}
	}
}

func (k *Keyring) SetStatus(status bool) {
	k.IsActive = status
}

func (k *Keyring) SetDescription(description string) {
	k.Description = description
}
