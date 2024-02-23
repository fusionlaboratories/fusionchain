// Copyright (c) Fusion Laboratories LTD
// SPDX-License-Identifier: BUSL-1.1
package boolparser

func ShuntingYard(s Stack) Stack {
	postfix := Stack{}
	unary := Stack{}
	operators := Stack{}
	for _, v := range s.Values {
		switch v.Type {
		case OPERATOR:
			for !operators.IsEmpty() {
				val := v.Value
				top := operators.Peek().Value
				if (oprData[val].prec <= oprData[top].prec && !oprData[val].rAsoc) ||
					(oprData[val].prec < oprData[top].prec && oprData[val].rAsoc) {
					postfix.Push(operators.Pop())
					continue
				}
				break
			}
			operators.Push(v)
		case UNARY:
			for !unary.IsEmpty() {
				val := v.Value
				top := unary.Peek().Value
				if (oprData[val].prec <= oprData[top].prec && !oprData[val].rAsoc) ||
					(oprData[val].prec < oprData[top].prec && oprData[val].rAsoc) {
					postfix.Push(unary.Pop())
					continue
				}
				break
			}
			unary.Push(v)

		case LPAREN:
			operators.Push(v)
		case RPAREN:
			for i := operators.Length() - 1; i >= 0; i-- {
				if operators.Values[i].Type != LPAREN {
					postfix.Push(operators.Pop())

					if unary.Length() > 0 {
						postfix.Push(unary.Pop())
					}

					continue
				}
				operators.Pop()
				break
			}
		default:
			postfix.Push(v)
			// if unary.Length() != 0 {
			// 	val := unary.Pop()
			// 	postfix.Push(val)
			// }
		}
	}
	unary.EmptyInto(&postfix)
	operators.EmptyInto(&postfix)
	return postfix
}
