package leetcode

func findDiagonalOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0{
		return result
	}
	row := 0
	col := 0
	maxrow := len(matrix)
	maxcol := len(matrix[0])
	sign := 1
	for {
		result = append(result, matrix[row][col])
		if row - 1 >= 0 && col + 1 < maxcol && sign == 1{
			row--
			col++
		}else if col + 1 < maxcol && sign == 1{
			col++
			sign = 0
		}else if row + 1 < maxrow && col - 1 >= 0 && sign == 0{
			col--
			row++
		}else if row + 1 < maxrow && sign == 0{
			row++
			sign = 1
		}else if row + 1 < maxrow && sign == 1{
			row ++
			sign = 0
		}else if col + 1 < maxcol && sign == 0{
			col ++
			sign = 1
		}else{
			break
		}
	}
	return result
}
