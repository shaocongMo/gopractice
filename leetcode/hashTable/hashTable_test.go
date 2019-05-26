package leetcode

import "testing"

func Test805(t *testing.T){
	test805Logf([]int{1,2,3,1}, t)
	test805Logf([]int{}, t)
	test805Logf([]int{1}, t)
	test805Logf([]int{1,2}, t)
	test805Logf([]int{1,1}, t)
	test805Logf([]int{1,1,3,4}, t)
	test805Logf([]int{1,2,3,4}, t)
	test805Logf([]int{1,1,1,3,3,4,3,2,4,2}, t)
}

func test805Logf(nums []int, t *testing.T){
	t.Logf("Input: %v Output: %v", nums, containsDuplicate(nums))
}

func Test806(t *testing.T){
	test806Logf([]int{2,2,1}, t)
	test806Logf([]int{4,1,2,1,2}, t)
	test806Logf([]int{1,2,3,1,2,3,7}, t)
	test806Logf([]int{1}, t)
	test806Logf([]int{3,3,2,2,8,1,1,4,4}, t)
}

func test806Logf(nums []int, t *testing.T){
	t.Logf("Input: %v Output: %d", nums, singleNumber(nums))
}
