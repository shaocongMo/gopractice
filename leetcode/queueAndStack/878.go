package leetcode

func isValid(s string) bool {
	var match = map[rune]rune{')': '(', ']': '[', '}': '{'}
	var top = -1
	var stack = []rune{}
	for _, c := range s {
		target := match[c]
		if target == 0 {
			stack = append(stack, c)
			top++
		} else {
			if top >= 0 && stack[top] == target {
				stack = stack[:top]
				top--
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
