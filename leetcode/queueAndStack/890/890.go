package leetcode

func decodeString(s string) string {
	stack := []rune{}
	for _, c := range s {
		if c == ']' {
			tmp := []rune{}
			for len(stack) > 0 {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if pop == '[' {
					break
				} else {
					tmp = append(tmp, pop)
				}
			}
			times := 0
			MulNum := 1
			for len(stack) > 0 {
				pop := stack[len(stack)-1]
				if pop >= '0' && pop <= '9' {
					stack = stack[:len(stack)-1]
					times = times + MulNum*int(pop-'0')
					MulNum = MulNum * 10
				} else {
					break
				}
			}
			for i := 0; i < times; i++ {
				tmp_size := len(tmp)
				for j := 0; j < tmp_size; j++ {
					x := tmp[tmp_size-j-1]
					stack = append(stack, x)
				}
			}
		} else {
			stack = append(stack, c)
		}
	}
	return string(stack)
}
