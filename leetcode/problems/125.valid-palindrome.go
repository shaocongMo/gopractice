package leetcode

func isPalindrome(s string) bool {
	slen := len(s)
	i, j := 0, slen - 1
	var left byte = '0'
	var right byte = '0'
	for ; j > i; i, j = i+1, j-1 {
		for ; i < slen; i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				left = s[i] - 'a'
				break
			}else if s[i] >= 'A' && s[i] <= 'Z' {
				left = s[i] - 'A'
				break
			}else if s[i] >= '0' && s[i] <= '9' {
				left = s[i]
				break
			}
		}
		for ; j >= 0; j-- {
			if s[j] >= 'a' && s[j] <= 'z' {
				right = s[j] - 'a'
				break
			}else if s[j] >= 'A' && s[j] <= 'Z' {
				right = s[j] - 'A'
				break
			}else if s[j] >= '0' && s[j] <= '9' {
				right = s[j]
				break
			}
		}
		if left != right {
			return false
		}
	}
	return true
}
