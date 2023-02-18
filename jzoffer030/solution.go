// https://leetcode-cn.com/problems/insert-delete-getrandom-o1/
package jzoffer030

import "math/rand"

type RandomizedSet struct {
	numToLocation map[int]int
	nums          []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	var s RandomizedSet
	s.numToLocation = make(map[int]int)
	s.nums = make([]int, 0)
	return s
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numToLocation[val]; ok {
		return false
	}

	this.numToLocation[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numToLocation[val]; !ok {
		return false
	}

	location := this.numToLocation[val]
	this.numToLocation[this.nums[len(this.nums)-1]] = location
	delete(this.numToLocation, val)
	this.nums[location] = this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	r := rand.Intn(len(this.nums))
	return this.nums[r]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
