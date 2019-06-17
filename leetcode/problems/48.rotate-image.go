package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)
	mid := n / 2
	x := 0
	y := 0
	m := 0
	offset := 2
	// 旋转多少组（每4个数字为一组)
	rotateTime := n * n / 4
	for i := 0; i < rotateTime; i++ {
		front := matrix[x][y]
		for j := 1; j <= 4; j++ {
			nextX := y
			if n%2 == 1 && x == mid {
				m = 0
			} else if n%2 == 0 && (x == mid || x == mid-1) {
				m = 1
			} else if x < mid {
				m = n - x*2 - 1
			} else {
				if n%2 == 0 {
					m = (x-mid)*2 + 1
				} else {
					m = (x - mid) * 2
				}
			}
			nextY := x
			if x <= mid-1 {
				nextY += m
			} else {
				nextY -= m
			}
			tmp := matrix[nextX][nextY]
			matrix[nextX][nextY] = front
			front = tmp
			x = nextX
			y = nextY
		}
		if y >= n-offset {
			offset++
			x++
			y = x
		} else {
			y++
		}
	}
}
