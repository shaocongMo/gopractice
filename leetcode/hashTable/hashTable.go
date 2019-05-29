package leetcode

// 811
func twoSum(nums []int, target int) []int{
	numIndex := make(map[int]int)
	for i, num := range nums{
		if numIndex[target - num] > 0{
			return []int{numIndex[target-num] - 1, i}
		}
		numIndex[num] = i + 1
	}
	return []int{}
}

// 812
func isIsomorphic(s string, t string) bool {
	sSet := make(map[byte]int)
	tSet := make(map[byte]int)
	for i, j := 0, 0; i < len(s);{
		if sSet[s[i]] == 0{
			sSet[s[i]] = i + 1
		}
		if tSet[t[j]] == 0{
			tSet[t[j]] = j + 1
		}
		if sSet[s[i]] != tSet[t[j]]{
			return false
		}
		i++
		j++
	}
	return true
}

// 813
func findRestaurant(list1 []string, list2 []string) []string {
	set := make(map[string]int)
	result := []string{}
	for i, str := range list1{
		set[str] = i + 1
	}
	totalLen := len(list1) + len(list2)
	min := -1
	for i, str := range list2{
		if set[str] != 0{
			set[str] += i + totalLen
			if min > set[str] || min == -1{
				min = set[str]
			}
		}
	}
	for str, v := range set{
		if v == min{
			result = append(result, str)
		}
	}
	return result
}

// 815
func firstUniqChar(s string) int {
	set := make(map[rune]int)
	for _, r := range s{
		set[r] ++
	}
	for i, r := range s{
		if set[r] == 1{
			return i
		}
	}
	return -1
}