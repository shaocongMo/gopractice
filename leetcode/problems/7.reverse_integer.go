package leetcode

func reverse(num int) int {
	plus_max_num := 2147483648 - 1
	not_plus_max_num := 2147483648
	isPlus := true
	if num < 0 {
		isPlus = false
	}
	reverseNum := 0
	checkTailZero := true
	for num != 0 {
		remain := num % 10
		num = num / 10
		if remain != 0 {
			checkTailZero = false
			if remain < 0 {
				remain = remain * -1
			}
		} else if checkTailZero && remain == 0 {
			continue
		}
		reverseNum = reverseNum * 10 + remain
		if isPlus && reverseNum > plus_max_num {
			return 0
		} else if !isPlus && reverseNum > not_plus_max_num {
			return 0
		}
	}
	if isPlus {
		return reverseNum
	}
	return reverseNum * -1
}