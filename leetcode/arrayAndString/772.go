package leetcode

func plusOne(digits []int) []int {
	remain := 1
	plusOneNums := make([]int, len(digits)+1)
	for i := len(digits) - 1; i >= 0; i-- {
		num := digits[i]
		plusOneNums[i+1] = (num + remain) % 10
		remain = (num + remain) / 10
	}
	if remain > 0 {
		plusOneNums[0] = remain
		return plusOneNums
	}
	return plusOneNums[1:]
}
