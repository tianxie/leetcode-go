// https://leetcode-cn.com/problems/add-binary/
package jzoffer002

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		a    string
		b    string
		want string
	}{
		{},
		{},
	}

	s := solution1{}
	for _, test := range tests {
		if got := s.addBinary(test.a, test.b); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}
