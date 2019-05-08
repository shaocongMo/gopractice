package leetcode

import "testing"

func TestBubbleSort(t *testing.T) {
	testBubbleSortLogf([]int{5, 4, 2, 3, 8}, t)
	testBubbleSortLogf([]int{1, 2, 3}, t)
	testBubbleSortLogf([]int{3, 2, 1}, t)
	testBubbleSortLogf([]int{3, 1}, t)
	testBubbleSortLogf([]int{1, 3}, t)
	testBubbleSortLogf([]int{}, t)
	testBubbleSortLogf([]int{-1, 3}, t)
	testBubbleSortLogf([]int{3, -1, 1, 2, -9}, t)
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
	testSelectSortLogf([]int{-1, 3}, t)
	testSelectSortLogf([]int{3, -1, 1, 2, -9}, t)
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
	testInsertSortLogf([]int{-1, 3}, t)
	testInsertSortLogf([]int{3, -1, 1, 2, -9}, t)
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
	testShellSortLogf([]int{-1, 3}, t)
	testShellSortLogf([]int{3, -1, 1, 2, -9}, t)
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
	testMergeSortLogf([]int{-1, 3}, t)
	testMergeSortLogf([]int{3, -1, 1, 2, -9}, t)
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
	testQuickSortLogf([]int{-1, 3}, t)
	testQuickSortLogf([]int{3, -1, 1, 2, -9}, t)
}

func testQuickSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	QuickSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestHeapSort(t *testing.T) {
	testHeapSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testHeapSortLogf([]int{5, 4, 2, 3, 8}, t)
	testHeapSortLogf([]int{1, 2, 3}, t)
	testHeapSortLogf([]int{3, 2, 1}, t)
	testHeapSortLogf([]int{3, 1}, t)
	testHeapSortLogf([]int{1, 3}, t)
	testHeapSortLogf([]int{}, t)
	testHeapSortLogf([]int{-1, 3}, t)
	testHeapSortLogf([]int{3, -1, 1, 2, -9}, t)
}

func testHeapSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	HeapSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestCountingSort(t *testing.T) {
	testCountingSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testCountingSortLogf([]int{5, 4, 2, 3, 8}, t)
	testCountingSortLogf([]int{1, 2, 3}, t)
	testCountingSortLogf([]int{3, 2, 1}, t)
	testCountingSortLogf([]int{3, 1}, t)
	testCountingSortLogf([]int{1, 3}, t)
	testCountingSortLogf([]int{}, t)
	testCountingSortLogf([]int{-1, 3}, t)
	testCountingSortLogf([]int{3, -1, 1, 2, -9}, t)
}

func testCountingSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	CountingSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestBucketSort(t *testing.T) {
	testBucketSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testBucketSortLogf([]int{8, 9, 1, 7, 2, 33, 5, 4, 6, 100}, t)
	testBucketSortLogf([]int{5, 4, 2, 3, 8}, t)
	testBucketSortLogf([]int{1, 2, 3}, t)
	testBucketSortLogf([]int{3, 2, 1}, t)
	testBucketSortLogf([]int{3, 1}, t)
	testBucketSortLogf([]int{1, 3}, t)
	testBucketSortLogf([]int{}, t)
	testBucketSortLogf([]int{-1, 3}, t)
	testBucketSortLogf([]int{3, -1, 1, 2, -9}, t)
}

func testBucketSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	BucketSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}

func TestRedixSort(t *testing.T) {
	testRedixSortLogf([]int{8, 9, 1, 7, 2, 3, 5, 4, 6, 0}, t)
	testRedixSortLogf([]int{8, 9, 1, 7, 2, 33, 5, 4, 6, 100}, t)
	testRedixSortLogf([]int{5, 4, 2, 3, 8}, t)
	testRedixSortLogf([]int{1, 2, 3}, t)
	testRedixSortLogf([]int{3, 2, 1}, t)
	testRedixSortLogf([]int{3, 1}, t)
	testRedixSortLogf([]int{1, 3}, t)
	testRedixSortLogf([]int{1, 0, 3}, t)
	testRedixSortLogf([]int{}, t)
	testRedixSortLogf([]int{-1, 3}, t)
	testRedixSortLogf([]int{3, -1, 1, 2, -9}, t)
}

func testRedixSortLogf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	RedixSort(nums)
	t.Logf("Output: %v", nums)
	t.Logf("----------------")
}
