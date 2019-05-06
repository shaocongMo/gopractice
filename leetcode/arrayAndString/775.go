package leetcode

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}
	row := 0
	col := 0
	min_row := 1
	min_col := 0
	max_row := len(matrix)
	max_col := len(matrix[0])
	// 1 left -> right
	// 2 top -> bottom
	// 3 right -> left
	// 4 bottom -> top
	direction := 1
	old_row := -1
	old_col := -1
	for {
		if old_col != col || old_row != row {
			result = append(result, matrix[row][col])
			old_col = col
			old_row = row
		}
		if col+1 < max_col && direction == 1 {
			col++
		} else if row+1 < max_row && direction == 2 {
			row++
		} else if col-1 >= min_col && direction == 3 {
			col--
		} else if row-1 >= min_row && direction == 4 {
			row--
		} else {
			if direction < 4 {
				direction++
				if direction == 2 && row+1 >= max_row {
					break
				} else if direction == 3 && col-1 < min_col {
					break
				} else if direction == 4 && row-1 < min_row {
					break
				}
			} else {
				min_row++
				min_col++
				max_row--
				max_col--
				if col+1 < max_col && max_row >= min_row && max_col >= min_col {
					direction = 1
				} else {
					break
				}
			}
		}
	}
	return result
}
