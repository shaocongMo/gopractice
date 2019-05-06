package leetcode

func longestCommonPrefix(strs []string) string {
	commonPrefix := []byte{}
	index := 0
	var checkC byte
	for len(strs) > 0 {
		for i, str := range strs {
			if index >= len(str) {
				index = -1
				break
			}
			if i == 0 {
				checkC = str[index]
			} else if checkC != str[index] {
				index = -1
				break
			}
		}
		if index == -1 {
			break
		}
		commonPrefix = append(commonPrefix, checkC)
		index++
	}
	return string(commonPrefix)
}
