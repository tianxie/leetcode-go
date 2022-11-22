// https://leetcode.cn/problems/two-sum/
package p1_two_sum

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	var tests = []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	s := solution1{}
	for _, test := range tests {
		if got := s.twoSum(test.nums, test.target); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}
