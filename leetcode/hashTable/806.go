package leetcode

func singleNumber(nums []int) int {
	// min, max := nums[0], nums[0]
	// for _, num := range nums{
	// 	if num < min{
	// 		min = num
	// 	}else if num > max{
	// 		max = num
	// 	}
	// }
	// set := make([]int, max - min + 1)
	// for _, num := range nums{
	// 	set[num-min] += 1
	// }
	// fmt.Println(set)
	// for i, time := range set{
	// 	if time == 1{
	// 		return i+min
	// 	}
	// }
	// return 0

	/* 异或思路 */
	// 异或：相同为0，不同为1. 异或同一个数两次，原数不变。
	result := 0
	for _, num := range nums{
		result = result ^ num
	}
	return result
}