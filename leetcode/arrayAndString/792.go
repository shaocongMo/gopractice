package leetcode

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	for i := 0; i <= rowIndex; i++ {
		for j := i; j >= 0; j-- {
			if i == 0 || j == 0 || j == i {
				result[j] = 1
			} else {
				result[j] = result[j] + result[j-1]
			}
		}
	}
	return result
}
