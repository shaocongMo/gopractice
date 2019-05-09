package leetcode

func findMaxConsecutiveOnes(nums []int) int {
	t1, t2 := 0, 0
	for _, num := range nums {
		if num == 0 {
			if t2 > t1 {
				t1 = t2
			}
			t2 = 0
		} else {
			t2 += 1
		}
	}
	if t2 > t1 {
		return t2
	}
	return t1
}
