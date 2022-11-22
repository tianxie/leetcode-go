package jzoffer001

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		a    int
		b    int
		want int
	}{
		{15, 2, 7},
		{7, -3, -2},
		{0, 1, 0},
		{1, 1, 1},
	}

	s := solution1{}
	for _, test := range tests {
		if got := s.divide(test.a, test.b); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %d, want %d", got, test.want)
		}
	}
}
