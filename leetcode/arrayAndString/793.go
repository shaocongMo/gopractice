package leetcode

func reverseWords(s string) string {
	size := len(s)
	word := []byte{}
	result := []byte{}
	for i := 0; i < size; i++ {
		if s[i] != ' ' {
			word = append(word, s[i])
		} else {
			result = appendToResult(result, word)
			word = []byte{}
		}
	}
	result = appendToResult(result, word)
	return string(result)
}

func appendToResult(result []byte, word []byte) []byte {
	if len(word) > 0 {
		if len(result) > 0 {
			word = append(word, ' ')
		}
		result = append(word, result...)
	}
	return result
}
