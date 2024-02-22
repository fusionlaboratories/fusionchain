// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package boolparser

type TokenType int

type Token struct {
	Type  TokenType
	Value string
}

var eof = rune(0)

const (
	NUMBER TokenType = iota
	LPAREN
	RPAREN
	CONSTANT
	FUNCTION
	OPERATOR
	UNARY
	WHITESPACE
	ERROR
	EOF
)
