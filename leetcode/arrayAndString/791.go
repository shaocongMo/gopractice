package leetcode

func rotate(nums []int, k int) {
	size := len(nums)
	if size == 0 || k%size == 0 {
		return
	}
	n := size
	k = k % n
	for i := 0; i < size-1 && k != 0; {
		for j := 0; j < k; j++ {
			nums[i+j], nums[size-k+j] = nums[size-k+j], nums[i+j]
		}
		i += k
		n -= k
		k %= n
	}
}
