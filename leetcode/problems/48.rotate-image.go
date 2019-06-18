package leetcode

// 对角 镜面 翻转
func rotate(matrix [][]int) {
	n := len(matrix)
	// 对角翻转
	for i := 0; i < len(matrix); i ++ {
		for j := 0; j < len(matrix[i]) - i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][n-i-1]
			matrix[n-j-1][n-i-1] = tmp
		}
	}

	// 镜面翻转
	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < len(matrix[i]); j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-i-1][j]
			matrix[n-i-1][j] = tmp
		}
	}
}

// 每四个数独立旋转
func rotate2(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < len(matrix) / 2; i++ {
		for j := 0; j < len(matrix) / 2 + n % 2; j ++ {
			matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i], matrix[i][j] =
				matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i] 
		}
	}
}

// 外围旋转完，再到内层旋转 时间复杂度 O(n^2)
func rotate3(matrix [][]int) {
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
