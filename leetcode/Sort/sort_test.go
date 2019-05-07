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

func TestShellSort(t *testing.T) {
	testShellSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testShellSortLogf([]int{5, 4, 2, 3, 8}, t)
	testShellSortLogf([]int{1, 2, 3}, t)
	testShellSortLogf([]int{3, 2, 1}, t)
	testShellSortLogf([]int{3, 1}, t)
	testShellSortLogf([]int{1, 3}, t)
	testShellSortLogf([]int{}, t)
}

func testShellSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	ShellSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestMergeSort(t *testing.T) {
	testMergeSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testMergeSortLogf([]int{5, 4, 2, 3, 8}, t)
	testMergeSortLogf([]int{1, 2, 3}, t)
	testMergeSortLogf([]int{3, 2, 1}, t)
	testMergeSortLogf([]int{3, 1}, t)
	testMergeSortLogf([]int{1, 3}, t)
	testMergeSortLogf([]int{}, t)
}

func testMergeSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	t.Logf("Output: %v", MergeSort(nums))
	t.Logf("----------------")
}

func TestQuickSort(t *testing.T) {
	testQuickSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testQuickSortLogf([]int{5, 4, 2, 3, 8}, t)
	testQuickSortLogf([]int{1, 2, 3}, t)
	testQuickSortLogf([]int{3, 2, 1}, t)
	testQuickSortLogf([]int{3, 1}, t)
	testQuickSortLogf([]int{1, 3}, t)
	testQuickSortLogf([]int{}, t)
}

func testQuickSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	QuickSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}
