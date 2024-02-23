// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package boolparser

import (
	"strconv"
	"strings"
)

var oprData = map[string]struct {
	prec  int
	rAsoc bool // true = right // false = left
	fx    func(x, y float64) float64
}{
	// "^": {4, true, func(x, y float64) float64 { return math.Pow(x, y) }},
	"*": {3, false, func(x, y float64) float64 { return x * y }},
	//	"/": {3, false, func(x, y float64) float64 { return x / y }},
	"+": {2, false, func(x, y float64) float64 { return x + y }},
	"-": {2, false, func(x, y float64) float64 { return x - y }},
	">": {2, false, func(x, y float64) float64 { return b2f(x > y) }},
	"<": {2, false, func(x, y float64) float64 { return b2f(x < y) }},
	"&": {2, false, func(x, y float64) float64 { return b2f(f2b(x) && f2b(y)) }},
	"|": {2, false, func(x, y float64) float64 { return b2f(f2b(x) || f2b(y)) }},
	// "!": {2, false, func(x, y float64) float64 { return b2f(!f2b(x)) }},
}

var unaryData = map[string]struct {
	fx func(x float64) float64
}{
	"!": {func(x float64) float64 { return b2f(!f2b(x)) }},
}

func f2b(f float64) bool {
	return f != 0
}
func b2f(b bool) float64 {
	if !b {
		return 0
	}
	return 1
}

// SolvePostfix evaluates and returns the answer of the expression converted to postfix
func SolvePostfix(tokens Stack) float64 {
	stack := Stack{}
	for _, v := range tokens.Values {
		switch v.Type {
		case NUMBER:
			stack.Push(v)
		case UNARY:
			// unary invert
			f := unaryData[v.Value].fx
			var x float64
			x, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			result := f(x)
			stack.Push(Token{NUMBER, strconv.FormatFloat(result, 'f', -1, 64)})
		case OPERATOR:
			f := oprData[v.Value].fx
			var x, y float64
			y, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			x, _ = strconv.ParseFloat(stack.Pop().Value, 64)
			result := f(x, y)
			stack.Push(Token{NUMBER, strconv.FormatFloat(result, 'f', -1, 64)})
		}
	}
	if len(stack.Values) == 0 {
		return 0
	}
	out, _ := strconv.ParseFloat(stack.Values[0].Value, 64)
	return out
}

func Solve(s string) float64 {
	p := NewParser(strings.NewReader(s))
	stack, _ := p.Parse()
	stack = ShuntingYard(stack)
	answer := SolvePostfix(stack)
	return answer
}

func BoolSolve(s string) bool {
	return f2b(Solve(s))
}
