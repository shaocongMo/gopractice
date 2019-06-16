package leetcode

func isValidSudoku(board [][]byte) bool {
	colCounts := make([]map[byte]int, 9)
	smallCounts := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		colCounts[i] = map[byte]int{}
		smallCounts[i] = map[byte]int{}
	}
	for i := 0; i < len(board); i++ {
		rowCount := map[byte]int{}
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			if rowCount[board[i][j]] > 0 || colCounts[j][board[i][j]] > 0 || smallCounts[i/3*3+j/3][board[i][j]] > 0 {
				return false
			}
			rowCount[board[i][j]]++
			colCounts[j][board[i][j]]++
			smallCounts[i/3*3+j/3][board[i][j]]++
		}
	}
	return true
}
