package leetcode

func twoSum(numbers []int, target int) []int {
	size := len(numbers)
	i := 0
	j := size - 1
	for i < j {
		total := numbers[i] + numbers[j]
		if total == target {
			return []int{i + 1, j + 1}
		} else if total > target {
			j--
		} else {
			i++
			// j = size - 1 可处理存在负数情况
			// j = size - 1
		}
	}
	return []int{}
}
