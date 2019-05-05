package leetcode

type MatrixNode struct {
	x int
	y int
}

func updateMatrix(matrix [][]int) [][]int {
	used := make([][]bool, len(matrix))
	queue := []MatrixNode{}
	for i := 0; i < len(matrix); i++ {
		used[i] = make([]bool, len(matrix[i]))
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				queue = append(queue, MatrixNode{x: i, y: j})
				used[i][j] = true
			}
		}
	}
	deep := 0
	for len(queue) > 0 {
		queue_size := len(queue)
		for i := 0; i < queue_size; i++ {
			node := queue[0]
			queue = queue[1:]
			matrix[node.x][node.y] = deep
			// up
			if node.x-1 >= 0 && used[node.x-1][node.y] == false {
				queue = append(queue, MatrixNode{x: node.x - 1, y: node.y})
				used[node.x-1][node.y] = true
			}
			// down
			if node.x+1 < len(matrix) && used[node.x+1][node.y] == false {
				queue = append(queue, MatrixNode{x: node.x + 1, y: node.y})
				used[node.x+1][node.y] = true
			}
			// left
			if node.y-1 >= 0 && used[node.x][node.y-1] == false {
				queue = append(queue, MatrixNode{x: node.x, y: node.y - 1})
				used[node.x][node.y-1] = true
			}
			// right
			if node.y+1 < len(matrix[node.x]) && used[node.x][node.y+1] == false {
				queue = append(queue, MatrixNode{x: node.x, y: node.y + 1})
				used[node.x][node.y+1] = true
			}
		}
		deep++
	}
	return matrix
}
