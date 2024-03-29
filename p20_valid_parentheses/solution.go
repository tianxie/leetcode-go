// https://leetcode.cn/problems/valid-parentheses/
package p20_valid_parentheses

type solution interface {
	isValid(s string) bool
}

type solution1 struct{}

func (s1 solution1) isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			// 左括号入栈
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			n := len(stack) - 1
			top := stack[n]
			if s1.isPair(top, c) {
				stack = stack[:n]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func (s1 solution1) isPair(a, b rune) bool {
	return (a == '(' && b == ')') || (a == '[' && b == ']') || (a == '{' && b == '}')
}
