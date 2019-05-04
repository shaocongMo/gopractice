package leetcode

func Solu(arr []int, S int) bool {
	// return rec_opt(arr, len(arr)-1, S)
	return db_opt(arr, S)
}

func rec_opt(arr []int, i int, S int) bool {
	if S == 0 {
		return true
	} else if i == 0 {
		return arr[i] == S
	} else if arr[i] > S {
		return rec_opt(arr, i-1, S)
	}
	return rec_opt(arr, i-1, S-arr[i]) || rec_opt(arr, i-1, S)
}

func db_opt(arr []int, S int) bool {
	var subset = make([][]bool, len(arr))
	for i := range subset {
		subset[i] = make([]bool, S+1)
	}
	// subset[0][arr[0]] = true
	for i := range subset {
		for j := 0; j < S+1; j++ {
			if i == 0 {
				subset[i][j] = arr[i] == j
			} else if j == 0 {
				subset[i][j] = true
			} else if arr[i] > j {
				subset[i][j] = subset[i-1][j]
			} else {
				subset[i][j] = subset[i-1][j] || subset[i-1][j-arr[i]]
			}
		}
	}
	return subset[len(arr)-1][S]
}
