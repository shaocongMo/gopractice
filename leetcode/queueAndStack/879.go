package leetcode

type Node struct {
	index       int
	temperature int
}

func dailyTemperatures(T []int) []int {
	var days = make([]int, len(T))
	var top = -1
	var stack = []Node{}
	for index, temperature := range T {
		if top >= 0 {
			for {
				if top >= 0 && stack[top].temperature < temperature {
					days[stack[top].index] = index - stack[top].index
					stack = stack[:top]
					top--
				} else {
					break
				}
			}
		}
		top++
		stack = append(stack, Node{index: index, temperature: temperature})
	}
	return days
}
