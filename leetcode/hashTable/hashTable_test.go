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

func Test807(t *testing.T){
	test807Logf([]int{1, 2, 2, 1}, []int{2, 2}, t)
	test807Logf([]int{1, 2, 2, 1}, []int{2, 3}, t)
	test807Logf([]int{1, 2, 2, 1}, []int{2, 1}, t)
	test807Logf([]int{1, 2, 2, 1}, []int{}, t)
	test807Logf([]int{}, []int{}, t)
	test807Logf([]int{1, 7, 8, 9}, []int{5, 6, 7, 8}, t)
}

func test807Logf(nums1 []int, nums2 []int, t *testing.T){
	t.Logf("Input: %v %v Output: %v", nums1, nums2, intersection(nums1, nums2))
}

func Test808(t *testing.T){
	for i := 1; i <= 1000; i++{
		test808Logf(i, t)
	}
}

func test808Logf(num int, t *testing.T){
	t.Logf("Input: %d Output: %v", num, isHappy(num))
}

func Test811(t *testing.T){
	test811Logf([]int{2, 7, 11, 15}, 9, t)
	test811Logf([]int{2, 7, 11, 15}, 100, t)
	test811Logf([]int{2}, 100, t)
	test811Logf([]int{1, 2, 3, 99}, 100, t)
	test811Logf([]int{1, 2, 3, 99}, 102, t)
	test811Logf([]int{1, 2, 77, 99, 288}, 79, t)
}

func test811Logf(nums []int, target int, t *testing.T){
	t.Logf("Input: %v %d Output: %v", nums, target, twoSum(nums, target))
}

func Test812(t *testing.T){
	test812Logf("egg", "add", t)
	test812Logf("egg", "abc", t)
	test812Logf("foo", "bar", t)
	test812Logf("paper", "title", t)
	test812Logf("abcabcabcc", "cdecdecdee", t)
	test812Logf("abcabcabcc", "cdecdecdef", t)
}

func test812Logf(s string, t string, test *testing.T){
	test.Logf("Input: %s %s Output: %v", s, t, isIsomorphic(s, t))
}

func Test813(t *testing.T){
	test813Logf([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}, t)
	test813Logf([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}, t)
	test813Logf([]string{"a", "b", "c"}, []string{"a", "b", "c"}, t)
	test813Logf([]string{"a", "b", "c"}, []string{"b", "a", "c"}, t)
	test813Logf([]string{"a", "b", "c"}, []string{}, t)
	test813Logf([]string{}, []string{}, t)
	test813Logf([]string{}, []string{"a", "b", "c"}, t)
	test813Logf([]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}, t)
}

func test813Logf(strs1 []string, strs2 []string, t *testing.T){
	t.Logf("Input: %v %v Output: %v", strs1, strs2, findRestaurant(strs1, strs2))
}

func Test815(t *testing.T){
	test815Logf("leetcode", t)
	test815Logf("loveleetcode", t)
	test815Logf("abcabcabcabc", t)
	test815Logf("", t)
	test815Logf("abcabce", t)
}

func test815Logf(s string, t *testing.T){
	t.Logf("Input: %s Output: %d", s, firstUniqChar(s))
}