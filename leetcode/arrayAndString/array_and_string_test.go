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

func Test785(t *testing.T) {
	test785Logf([]int{2, 7, 11, 15}, 9, t)
	test785Logf([]int{2, 7, 11, 15}, 17, t)
	test785Logf([]int{2}, 17, t)
	test785Logf([]int{}, 17, t)
	test785Logf([]int{2, 7}, 17, t)
	test785Logf([]int{2, 15}, 17, t)
	test785Logf([]int{1, 2, 3, 4, 5, 6}, 4, t)
	test785Logf([]int{-4, -3, -2, -1, 5}, 1, t)
	test785Logf([]int{-4, -3, -2, -1, 5, 10}, 15, t)
	test785Logf([]int{-4, -3, -2, -1, 5, 10}, 1, t)
	test785Logf([]int{-4, -3, -2, -1, 5, 10}, 9, t)
	test785Logf([]int{-4, -3, -2, -1, 5, 10}, 3, t)
}

func test785Logf(nums []int, target int, t *testing.T) {
	t.Logf("Input: %v %d Output: %v", nums, target, twoSum(nums, target))
}

func Test787(t *testing.T) {
	test787Logf([]int{3, 2, 2, 3}, 3, t)
	test787Logf([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, t)
	test787Logf([]int{1}, 2, t)
	test787Logf([]int{}, 2, t)
	test787Logf([]int{2, 2, 2}, 2, t)
	test787Logf([]int{2, 2, 2, 3}, 2, t)
	test787Logf([]int{2, 3, 3}, 2, t)
	test787Logf([]int{3, 3, 2}, 2, t)
}

func test787Logf(nums []int, target int, t *testing.T) {
	t.Logf("Input: %v %d", nums, target)
	end := removeElement(nums, target)
	t.Logf("Output: %v %d", nums[:end], end)
	t.Logf("--------------------")
}

func Test788(t *testing.T) {
	test788Logf([]int{1, 1, 0, 1, 1, 1}, t)
	test788Logf([]int{1, 1, 0, 1, 0, 1}, t)
	test788Logf([]int{}, t)
	test788Logf([]int{0, 0}, t)
	test788Logf([]int{0, 0, 1}, t)
	test788Logf([]int{1, 0, 0}, t)
	test788Logf([]int{1, 1, 0, 0}, t)
	test788Logf([]int{0, 0, 1, 1}, t)
	test788Logf([]int{0, 1, 1, 0, 1, 0}, t)
}

func test788Logf(nums []int, t *testing.T) {
	t.Logf("Input: %v Output: %d", nums, findMaxConsecutiveOnes(nums))
}

func Test789(t *testing.T) {
	test789Logf(7, []int{2, 3, 1, 2, 4, 3}, t)
	test789Logf(7, []int{2, 3, 1, 2, 4, 3, 7}, t)
	test789Logf(7, []int{7, 2, 3, 1, 2, 4, 3}, t)
	test789Logf(7, []int{}, t)
}

func test789Logf(s int, nums []int, t *testing.T) {
	t.Logf("Input: %d %v Output: %v", s, nums, minSubArrayLen(s, nums))
}

func Test791(t *testing.T) {
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 3, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 6, t)
	test791Logf([]int{-1, -100, 3, 99}, 2, t)
	test791Logf([]int{-1, -100, 3, 99}, 4, t)
	test791Logf([]int{-1, -100, 3, 99}, 9, t)
	test791Logf([]int{-1, -100, 3, 99}, 3, t)
	test791Logf([]int{-1, -100, 3, 99}, 0, t)
	test791Logf([]int{}, 10, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 0, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 1, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 2, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 3, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 4, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 5, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 6, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 7, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6, 7}, 8, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6}, 0, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6}, 1, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6}, 2, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6}, 3, t)
	test791Logf([]int{1, 2, 3, 4, 5, 6}, 4, t)
}

func test791Logf(nums []int, k int, t *testing.T) {
	t.Logf("Input: %v %d ", nums, k)
	rotate(nums, k)
	t.Logf("Output: %v", nums)
	t.Logf("-------------------------")
}

func Test792(t *testing.T) {
	test792Logf(3, t)
	test792Logf(4, t)
	test792Logf(5, t)
	test792Logf(6, t)
	test792Logf(7, t)
	test792Logf(8, t)
}

func test792Logf(k int, t *testing.T) {
	t.Logf("Input: %d Output: %v", k, getRow(k))
}

func Test793(t *testing.T) {
	test793Logf("the sky is blue", t)
	test793Logf("  hello world!  ", t)
	test793Logf("a good   example", t)
	test793Logf("a b   c  ", t)
}

func test793Logf(s string, t *testing.T) {
	t.Logf("Input: %s Output: [%s]", s, reverseWords(s))
}

func Test794(t *testing.T) {
	test794Logf("Let's take LeetCode contest", t)
	test794Logf("abc bca aa bc abcd a", t)
}

func test794Logf(s string, t *testing.T) {
	t.Logf("Input: %s Output: [%s]", s, reverseWords2(s))
}

func Test795(t *testing.T) {
	test795Logf([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, t)
	test795Logf([]int{}, t)
	test795Logf([]int{1}, t)
	test795Logf([]int{1, 2, 3, 4}, t)
	test795Logf([]int{1, 2, 3, 4, 4, 4, 4, 4}, t)
	test795Logf([]int{1, 1, 1, 1, 1, 2, 3, 4, 4, 4, 4, 4}, t)
	test795Logf([]int{1, 2, 2, 2, 2, 2, 3}, t)
	test795Logf([]int{1, 1, 1, 1, 1, 2, 3}, t)
}

func test795Logf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	j := removeDuplicates(nums)
	t.Logf("Output: %v", nums[:j])
	t.Logf("------------------------")
}

func Test796(t *testing.T) {
	test796Logf([]int{0, 1, 0, 3, 12}, t)
	test796Logf([]int{0, 1, 0, 0, 0}, t)
	test796Logf([]int{1, 0, 0, 0}, t)
	test796Logf([]int{0, 0, 0, 0, 1}, t)
	test796Logf([]int{0, 0, 0, 0, 0, 2}, t)
	test796Logf([]int{0}, t)
	test796Logf([]int{2}, t)
	test796Logf([]int{}, t)
}

func test796Logf(nums []int, t *testing.T) {
	t.Logf("Input: %v", nums)
	moveZeroes(nums)
	t.Logf("Output: %v", nums)
	t.Logf("------------------------")
}
