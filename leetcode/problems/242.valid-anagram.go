package leetcode

func isAnagram(s string, t string) bool {
	runeMap := map[rune]int{}
	for _, sr := range s {
		runeMap[sr] = runeMap[sr] + 1
	}
	for _, tr := range t {
		runeMap[tr] = runeMap[tr] - 1
		if runeMap[tr] == 0 {
			delete(runeMap, tr)
		}
	}
	if len(runeMap) == 0 {
		return true
	}
	return false
}
