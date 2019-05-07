package leetcode

import "testing"

func TestBubbleSort(t *testing.T) {
	testBubbleSortLogf([]int{5, 4, 2, 3, 8}, t)
	testBubbleSortLogf([]int{1, 2, 3}, t)
	testBubbleSortLogf([]int{3, 2, 1}, t)
	testBubbleSortLogf([]int{3, 1}, t)
	testBubbleSortLogf([]int{1, 3}, t)
	testBubbleSortLogf([]int{}, t)
}

func testBubbleSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	BubbleSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestSelectSort(t *testing.T) {
	testSelectSortLogf([]int{5, 4, 2, 3, 8}, t)
	testSelectSortLogf([]int{1, 2, 3}, t)
	testSelectSortLogf([]int{3, 2, 1}, t)
	testSelectSortLogf([]int{3, 1}, t)
	testSelectSortLogf([]int{1, 3}, t)
	testSelectSortLogf([]int{}, t)
}

func testSelectSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	SelectSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestInsertSort(t *testing.T) {
	testInsertSortLogf([]int{5, 4, 2, 3, 8}, t)
	testInsertSortLogf([]int{1, 2, 3}, t)
	testInsertSortLogf([]int{3, 2, 1}, t)
	testInsertSortLogf([]int{3, 1}, t)
	testInsertSortLogf([]int{1, 3}, t)
	testInsertSortLogf([]int{}, t)
}

func testInsertSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	InsertSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}