package leetcode

func reverseWords2(s string) string {
	result := []byte{}
	word := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			word = append([]byte{s[i]}, word...)
		} else {
			if len(word) > 0 {
				result = append(result, word...)
				word = []byte{}
			}
			result = append(result, s[i])
		}
	}
	if len(word) > 0 {
		result = append(result, word...)
	}
	return string(result)
}
