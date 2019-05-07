package leetcode

func reverseString(s []byte) {
	frontIndex := 0
	tailIndex := len(s) - 1
	for frontIndex < tailIndex {
		s[frontIndex], s[tailIndex] = s[tailIndex],s[frontIndex]
		frontIndex++
		tailIndex--
	}
}
