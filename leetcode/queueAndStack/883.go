package leetcode

//可以不用建立新used数组，可用grid代替  把grid 1 的替换为0

func numIslands(grid [][]byte) int {
	var used = make([][]int, len(grid))
	for i, _ := range grid {
		used[i] = make([]int, len(grid[i]))
	}
	var num int = 0
	for i, row := range grid {
		for j, col := range row {
			if col == '1' && used[i][j] == 0 {
				dfs(grid, i, j, used)
				num++
			}
		}
	}
	return num
}

func dfs(grid [][]byte, row int, col int, used [][]int) {
	used[row][col] = 1
	if grid[row][col] == '1' {
		// up
		if row-1 >= 0 && used[row-1][col] == 0 {
			dfs(grid, row-1, col, used)
		}
		//down
		if row+1 < len(grid) && used[row+1][col] == 0 {
			dfs(grid, row+1, col, used)
		}
		//left
		if col-1 >= 0 && used[row][col-1] == 0 {
			dfs(grid, row, col-1, used)
		}
		//right
		if col+1 < len(grid[row]) && used[row][col+1] == 0 {
			dfs(grid, row, col+1, used)
		}
	}
}
