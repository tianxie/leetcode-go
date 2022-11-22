// https://leetcode.cn/problems/two-sum/
package p1_two_sum

type solution interface {
	twoSum(nums []int, target int) []int
}

type solution1 struct{}

func (s solution1) twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

type solution2 struct{}

func (s solution2) twoSum(nums []int, target int) []int {
	return nil
}
