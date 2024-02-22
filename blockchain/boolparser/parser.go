// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package boolparser

import (
	"fmt"
	"io"
)

type buf struct {
	tok Token
	n   int
}

type Parser struct {
	s   *Scanner
	buf buf
}

func NewParser(r io.Reader) *Parser {
	return &Parser{s: NewScanner(r)}
}

func (p *Parser) Scan() (tok Token) {
	if p.buf.n != 0 {
		p.buf.n = 0
		return p.buf.tok
	}

	tok = p.s.Scan()

	p.buf.tok = tok

	return
}

func (p *Parser) ScanIgnoreWhitespace() (tok Token) {
	tok = p.Scan()
	if tok.Type == WHITESPACE {
		tok = p.Scan()
	}
	return
}

func (p *Parser) UnScan() {
	p.buf.n = 1
}

func (p *Parser) Parse() (Stack, error) {
	stack := Stack{}
	for {
		tok := p.ScanIgnoreWhitespace()
		if tok.Type == ERROR {
			return Stack{}, fmt.Errorf("ERROR: %q", tok.Value)
		} else if tok.Type == EOF {
			break
		} else if tok.Type == OPERATOR && tok.Value == "-" {
			lastTok := stack.Peek()
			nextTok := p.ScanIgnoreWhitespace()
			if (lastTok.Type == OPERATOR || lastTok.Value == "" || lastTok.Type == LPAREN) && nextTok.Type == NUMBER {
				stack.Push(Token{NUMBER, "-" + nextTok.Value})
			} else {
				stack.Push(tok)
				p.UnScan()
			}
		} else {
			stack.Push(tok)
		}
	}
	return stack, nil
}
