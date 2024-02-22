// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package casing

import "strings"

func ToKebabCase(s string) string {
	return toCase(s, '-')
}

func ToSnakeCase(s string) string {
	return toCase(s, '_')
}

func toCase(s string, sep rune) string {
	var b strings.Builder
	for i, r := range s {
		if i > 0 && isUpper(r) {
			if _, err := b.WriteRune(sep); err != nil {
				panic(err)
			}
		}
		if isUpper(r) {
			r += ('a' - 'A')
		}
		if _, err := b.WriteRune(r); err != nil {
			panic(err)
		}
	}
	return b.String()
}

func isUpper(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func ToLower(s string) string {
	return strings.ToLower(s)
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
