package leetcode

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	numcount := map[int]int{}
	for _, num := range nums1 {
		numcount[num]++
	}
	insertnums := []int{}
	for _, num := range nums2 {
		if numcount[num] > 0 {
			numcount[num]--
			insertnums = append(insertnums, num)
		}
	}
	return insertnums
}
