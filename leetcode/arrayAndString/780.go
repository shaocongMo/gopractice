package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	fIndex := -1
	for i := 0; i < len(haystack) && len(haystack)-i >= len(needle); i++ {
		if haystack[i] == needle[0] {
			fIndex = i
			nextStartI := -1
			for j := 1; j < len(needle); j++ {
				i++
				if haystack[i] == needle[0] && nextStartI == -1 {
					nextStartI = i
				}
				if haystack[i] != needle[j] {
					fIndex = -1
					break
				}
			}
			if fIndex != -1 {
				return fIndex
			}
			if nextStartI != -1 {
				i = nextStartI - 1
			}
		}
	}
	return fIndex
}
