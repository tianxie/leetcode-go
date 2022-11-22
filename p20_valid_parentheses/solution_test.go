// https://leetcode.cn/problems/valid-parentheses/
package p20_valid_parentheses

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		s    string
		want bool
	}{
		{"]", false},
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	s := solution1{}
	for _, test := range tests {
		if got := s.isValid(test.s); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %v, want %t", got, test.want)
		}
	}
}
