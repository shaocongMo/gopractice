package leetcode

func moveZeroes(nums []int) {
	// zeroNum := 0
	// for i := len(nums) - 2; i >= 0; i-- {
	// 	if nums[i] == 0 {
	// 		j := i
	// 		for ; j < len(nums)-1-zeroNum; j++ {
	// 			nums[j] = nums[j+1]
	// 		}
	// 		zeroNum++
	// 		nums[j] = 0
	// 	}
	// }

	// 把非0往前移动，后面补0
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else {
			nums[n] = nums[i]
			n++
		}
	}
	for i := n; i < len(nums); i++ {
		nums[i] = 0
	}
}
