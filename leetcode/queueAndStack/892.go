package leetcode

type Elem struct {
	x int
	y int
}

func updateMatrix(matrix [][]int) [][]int {

	for i, row := range matrix {
		for j, col := range row {
			if col == 1 {
				updateMatrixBFS(matrix, i, j, []Elem{{x: i, y: j}}, 0)
			}
		}
	}
	return matrix
}

func updateMatrixBFS(matrix [][]int, x int, y int, elems []Elem, target int) {
	var used = make([][]int, len(matrix))
	for i, row := range matrix {
		used[i] = make([]int, len(row))
	}
	used[x][y] = 1
loop:
	for len(elems) > 0 {
		var tmp_elems = []Elem{}
		target++
		for len(elems) > 0 {
			i := elems[0].x
			j := elems[0].y

			elems = elems[1:]
			//up
			if i-1 >= 0 && matrix[i-1][j] != 0 && used[i-1][j] == 0 {
				tmp_elems = append(tmp_elems, Elem{x: i - 1, y: j})
				used[i-1][j] = 1
			} else if i-1 >= 0 && matrix[i-1][j] == 0 {
				matrix[x][y] = target
				break loop
			}
			//down
			if i+1 < len(matrix) && matrix[i+1][j] != 0 && used[i+1][j] == 0 {
				tmp_elems = append(tmp_elems, Elem{x: i + 1, y: j})
				used[i+1][j] = 1
			} else if i+1 < len(matrix) && matrix[i+1][j] == 0 {
				matrix[x][y] = target
				break loop
			}
			//down
			if j-1 >= 0 && matrix[i][j-1] != 0 && used[i][j-1] == 0 {
				tmp_elems = append(tmp_elems, Elem{x: i, y: j - 1})
				used[i][j-1] = 1
			} else if j-1 >= 0 && matrix[i][j-1] == 0 {
				matrix[x][y] = target
				break loop
			}
			//right
			if j+1 < len(matrix[i]) && matrix[i][j+1] != 0 && used[i][j+1] == 0 {
				tmp_elems = append(tmp_elems, Elem{x: i, y: j + 1})
				used[i][j+1] = 1
			} else if j+1 < len(matrix[i]) && matrix[i][j+1] == 0 {
				matrix[x][y] = target
				break loop
			}
		}
		elems = tmp_elems
	}
}
