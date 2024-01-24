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
package policy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_NewPolicyAnyInGroup(t *testing.T) {
	tests := []struct {
		name          string
		group         []string
		signaturesMap map[string]bool
		expectedErr   bool
	}{
		{
			name:          "empty",
			group:         []string{},
			signaturesMap: map[string]bool{},
			expectedErr:   true,
		},
		{
			name:          "single",
			group:         []string{"foo"},
			signaturesMap: map[string]bool{"foo": true},
		},
		{
			name:          "foo signed",
			group:         []string{"foo", "bar"},
			signaturesMap: map[string]bool{"foo": true},
		},
		{
			name:          "bar signed",
			group:         []string{"foo", "bar"},
			signaturesMap: map[string]bool{"bar": true},
		},
		{
			name:          "all signed",
			group:         []string{"foo", "bar"},
			signaturesMap: map[string]bool{"foo": true, "bar": true},
		},
		{
			name:          "unknown signed",
			group:         []string{"foo", "bar"},
			signaturesMap: map[string]bool{"baz": true},
			expectedErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewAnyInGroupPolicy(tt.group)
			err := p.Verify(tt.signaturesMap, EmptyPolicyPayload(), nil)
			if tt.expectedErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
