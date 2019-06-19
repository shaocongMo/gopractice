package leetcode

import "testing"

func Test494(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums, 3, findTargetSumWays(nums, 3))
	nums2 := []int{0, 0, 0, 0, 0, 0, 0, 0, 1}
	t.Logf("Nums: %v S: %d Output: %d", nums2, 1, findTargetSumWays(nums2, 1))
}

func Test112(t *testing.T) {
	test112Log([]int{}, 0, t)
	test112Log([]int{1}, 0, t)
	test112Log([]int{1, 2}, 1, t)
	test112Log([]int{1, 2, 3}, 2, t)
	test112Log([]int{2, 1}, 0, t)
	test112Log([]int{3, 2, 1}, 0, t)
	test112Log([]int{7, 1, 5, 3, 6, 4}, 7, t)
	test112Log([]int{1, 2, 3, 4, 5}, 4, t)
	test112Log([]int{1, 7, 0, 0, 6, 4}, 12, t)
	test112Log([]int{1, 7, 4, 9, 0, 6, 4}, 17, t)
}

func test112Log(prices []int, tar int, t *testing.T) {
	t.Logf("Input: %v Target: %d Output: %v", prices, tar, maxProfit(prices))
}

func Test350(t *testing.T) {
	test350Logf([]int{}, []int{}, t)
	test350Logf([]int{1, 2, 3}, []int{}, t)
	test350Logf([]int{1, 2, 3}, []int{4, 5, 6}, t)
	test350Logf([]int{1, 2, 3}, []int{1}, t)
	test350Logf([]int{1, 2, 3}, []int{1, 4, 5, 6}, t)
	test350Logf([]int{1, 2, 3}, []int{1, 4, 5, 6, 3, 3, 2, 2}, t)
	test350Logf([]int{1, 2, 2, 3}, []int{1, 4, 5, 6, 3, 3, 2, 2}, t)
	test350Logf([]int{-2147483648, 1, 2, 3}, []int{1, -2147483648, -2147483648}, t)
	test350Logf([]int{-1, -2, 3}, []int{-1, 1, 2, 3}, t)
}

func test350Logf(nums1 []int, nums2 []int, t *testing.T) {
	t.Logf("Input: %v %v Output: %v", nums1, nums2, intersect(nums1, nums2))
}

func Test36(t *testing.T) {
	test36Logf([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, t)
	test36Logf([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}, t)
	test36Logf([][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '1'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'}}, t)
	test36Logf([][]byte{
		{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
		{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}, t)
}

func test36Logf(sudu [][]byte, t *testing.T) {
	t.Logf("Input: %v Output: %v", sudu, isValidSudoku(sudu))
}

func Test48(t *testing.T) {
	start := 1
	end := 100
	for i := start; i <= end; i++ {
		nums := make([][]int, i)
		num := 1
		for x := 0; x < i; x++ {
			nums[x] = make([]int, i)
			for y := 0; y < i; y++ {
				nums[x][y] = num
				num++
			}
		}
		test48Logf(nums, t)
	}
}

func test48Logf(matrix [][]int, t *testing.T) {
	t.Logf("Input: ")
	printMatrix(matrix, t)
	rotate(matrix)
	t.Logf("Output: ")
	printMatrix(matrix, t)
}

func printMatrix(matrix [][]int, t *testing.T) {
	for _, row := range matrix {
		t.Logf("%v", row)
	}
	t.Logf("=====================================")
}

func Test8(t *testing.T) {
	test8Logf(10, t)
	test8Logf(1000001, t)
	test8Logf(100000, t)
	test8Logf(120, t)
	test8Logf(102, t)
	test8Logf(-10, t)
	test8Logf(2147483647, t)
	test8Logf(7463847412, t)
	test8Logf(8463847412, t)
	test8Logf(-8463847412, t)
	test8Logf(-9463847412, t)
}

func test8Logf(num int, t *testing.T) {
	t.Logf("Input: %d Output: %d", num, reverse(num))
}
