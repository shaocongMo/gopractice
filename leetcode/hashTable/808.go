package leetcode

func isHappy(n int) bool {
	set := make(map[int]bool)
	for set[n] == false && n != 1{
		set[n] = true
		newn := 0
		for n != 0{
			m := n % 10
			newn += m * m
			n = n / 10
		}
		n = newn
	}
	if n == 1{
		return true
	}
	return false
}