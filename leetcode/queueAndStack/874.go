package leetcode

func numSquares(n int) int {
	var squareNums = map[int]bool{}
	for i := 1; i*i <= n; i++ {
		squareNums[i*i] = true
	}
	if squareNums[n] {
		return 1
	}
	return numSquaresBFS(n, squareNums)
}

func numSquaresBFS(n int, squareNums map[int]bool) int {
	var addTimes int = 0
	var cur_nums = []int{}
	var next_nums []int
	var found bool = false
	cur_nums = append(cur_nums, n)
	for len(cur_nums) > 0 && !found {
		addTimes = addTimes + 1
		next_nums = []int{}
		for _, cur_num := range cur_nums {
			for num, _ := range squareNums {
				less := cur_num - num
				if less == 0 {
					found = true
					break
				} else if less > 0 {
					if squareNums[less] {
						found = true
						addTimes = addTimes + 1
						break
					}
					next_nums = append(next_nums, less)
				}
			}
			if found {
				break
			}
		}
		cur_nums = next_nums
	}
	return addTimes
}
