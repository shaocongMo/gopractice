package leetcode

import "testing"

func Test770(t *testing.T) {
	test770Logf([]int{1, 7, 3, 6, 5, 6}, t)
	test770Logf([]int{1, 2, 3}, t)
	test770Logf([]int{1, 1}, t)
	test770Logf([]int{1, 1, 1}, t)
	test770Logf([]int{1, 1, 2, 1, 3, 1, 1, 4, 5}, t)
	test770Logf([]int{1}, t)
	test770Logf([]int{}, t)
	test770Logf([]int{-1, -1, -1, -1, -1, 0}, t)
	test770Logf([]int{-1, -1, -1, -1, 0, 1}, t)
	test770Logf([]int{-1, -1, -1, 0, 1, 1}, t)
	test770Logf([]int{0}, t)
	test770Logf([]int{-1, -1, 0, 1, 1, 0}, t)
}

func test770Logf(nums []int, t *testing.T) {
	t.Logf("Nums: %v Output: %d", nums, pivotIndex(nums))
}

func Test771(t *testing.T) {
	test771Logf([]int{3, 6, 1, 0}, t)
	test771Logf([]int{1, 2, 3, 4}, t)
	test771Logf([]int{1}, t)
	test771Logf([]int{0}, t)
	test771Logf([]int{1, 1, 1, 1, 1, 6, 1, 1, 1, 6}, t)
	test771Logf([]int{9, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}, t)
	test771Logf([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 10}, t)
}

func test771Logf(nums []int, t *testing.T) {
	t.Logf("nums: %v Output: %d", nums, dominantIndex(nums))
}

func Test772(t *testing.T) {
	test772Logf([]int{1, 2, 3, 4}, t)
	test772Logf([]int{9}, t)
	test772Logf([]int{8, 9}, t)
	test772Logf([]int{9, 9}, t)
	test772Logf([]int{0}, t)
}

func test772Logf(nums []int, t *testing.T) {
	t.Logf("nums: %v", nums)
	t.Logf("Output: %d", plusOne(nums))
}

func Test774(t *testing.T) {
	test774Logf([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, t)
	test774Logf([][]int{}, t)
	test774Logf([][]int{{1}}, t)
	test774Logf([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, t)
	test774Logf([][]int{{1, 2, 3, 4}}, t)
	test774Logf([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, t)
}
func test774Logf(matrix [][]int, t *testing.T) {
	t.Logf("matrix: %v Output: %d", matrix, findDiagonalOrder(matrix))
}

func Test775(t *testing.T) {
	test775Logf([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, t)
	test775Logf([][]int{}, t)
	test775Logf([][]int{{1}}, t)
	test775Logf([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}, t)
	test775Logf([][]int{{1, 2, 3, 4}}, t)
	test775Logf([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, t)
	test775Logf([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {10, 11, 12}}, t)
}
func test775Logf(matrix [][]int, t *testing.T) {
	t.Logf("matrix: %v Output: %d", matrix, spiralOrder(matrix))
}

func Test776(t *testing.T) {
	t.Logf("Num: %d Output: %v", 5, generate(5))
	t.Logf("Num: %d Output: %v", 0, generate(0))
	t.Logf("Num: %d Output: %v", 10, generate(10))
	t.Logf("Num: %d Output: %v", 11, generate(11))
	t.Logf("Num: %d Output: %v", 51, generate(51))
}

func Test779(t *testing.T) {
	// test779Logf("11", "1", t)
	// test779Logf("1010", "1011", t)
	// test779Logf("111", "1", t)
	test779Logf("111", "100000", t)
	test779Logf("10000", "1", t)
	test779Logf("10000", "10000", t)
}

func test779Logf(a string, b string, t *testing.T) {
	t.Logf("A: %s B: %s Output: %s", a, b, addBinary(a, b))
}

func Test780(t *testing.T) {
	test780Logf("hello", "ll", t)
	test780Logf("aaaaa", "bba", t)
	test780Logf("aabaaba", "abaab", t)
	test780Logf("aabaaba", "", t)
	test780Logf("", "", t)
	test780Logf("", "ababa", t)
	test780Logf("aaaabc", "abc", t)
	test780Logf("aaabcabc", "abc", t)
	test780Logf("abcaabcabc", "cabc", t)
}

func test780Logf(haystack string, needle string, t *testing.T) {
	t.Logf("Input: %s %s Output: %d", haystack, needle, strStr(haystack, needle))
}

func Test781(t *testing.T) {
	test781Logf([]string{"flower", "flow", "flight"}, t)
	test781Logf([]string{"dog", "racecar", "car"}, t)
	test781Logf([]string{"aacb", "aac", "aac"}, t)
	test781Logf([]string{"aacb", "", "aac"}, t)
	test781Logf([]string{}, t)
}

func test781Logf(strs []string, t *testing.T) {
	t.Logf("Input: %v Output: %s", strs, longestCommonPrefix(strs))
}

func Test783(t *testing.T) {
	test783Logf([]byte{'h', 'e', 'l', 'l', 'o'}, t)
	test783Logf([]byte{'a'}, t)
	test783Logf([]byte{}, t)
	test783Logf([]byte{'a', 'b'}, t)
	test783Logf([]byte{'a', 'b', 'c'}, t)
	test783Logf([]byte{'a', 'b', 'c', 'd'}, t)
}

func test783Logf(s []byte, t *testing.T) {
	t.Logf("Input: %v", s)
	reverseString(s)
	t.Logf("Output: %v", s)
	t.Logf("--------------------")
}

func Test784(t *testing.T) {
	test784Logf([]int{1, 4, 3, 2}, t)
	test784Logf([]int{1, 2, 2, 1}, t)
	test784Logf([]int{1, 1}, t)
	test784Logf([]int{1, 1, 1, 1, 1, 1}, t)
	test784Logf([]int{-1, -1, -1, 1, 2, 1}, t)
	test784Logf([]int{-1, -1, -1, -2}, t)
	test784Logf([]int{-1, -1, 1, 1}, t)
	test784Logf([]int{-1, -1}, t)
	test784Logf([]int{-1, -1, 0, 1}, t)
	test784Logf([]int{-1, -1, 1, 0}, t)
	test784Logf([]int{-1, 0, 0, 0}, t)
}

func test784Logf(nums []int, t *testing.T) {
	t.Logf("Input: %v Output: %d", nums, arrayPairSum(nums))
}
