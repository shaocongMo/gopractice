package leetcode

func addBinary(a string, b string) string {
	maxLen := 0
	if len(a) >= len(b) {
		maxLen = len(a)
	} else {
		maxLen = len(b)
	}
	result := make([]rune, maxLen+1)
	aIndex := len(a) - 1
	bIndex := len(b) - 1
	remain := '0'
	for aIndex >= 0 && bIndex >= 0 {
		if a[aIndex] == '1' && b[bIndex] == '1' {
			result[maxLen] = remain
			remain = '1'
		} else if a[aIndex] == '0' && b[bIndex] == '0' {
			result[maxLen] = remain
			remain = '0'
		} else if remain == '1' {
			result[maxLen] = '0'
		} else {
			result[maxLen] = '1'
		}
		maxLen--
		aIndex--
		bIndex--
	}
	for aIndex >= 0 {
		if a[aIndex] == '1' && remain == '1' {
			result[maxLen] = '0'
		} else if a[aIndex] == '0' && remain == '0' {
			result[maxLen] = '0'
		} else {
			result[maxLen] = '1'
			remain = '0'
		}
		aIndex--
		maxLen--
	}
	for bIndex >= 0 {
		if b[bIndex] == '1' && remain == '1' {
			result[maxLen] = '0'
		} else if b[bIndex] == '0' && remain == '0' {
			result[maxLen] = '0'
		} else {
			result[maxLen] = '1'
			remain = '0'
		}
		bIndex--
		maxLen--
	}
	result[maxLen] = remain
	if remain == '0' {
		return string(result[1:])
	}
	return string(result)
}
