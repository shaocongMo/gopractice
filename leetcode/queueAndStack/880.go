package leetcode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack = []int{}
	for _, token := range tokens {
		stack_top := len(stack) - 1
		switch token {
		case "+":
			stack[stack_top-1] = stack[stack_top-1] + stack[stack_top]
			stack = stack[:stack_top]
		case "-":
			stack[stack_top-1] = stack[stack_top-1] - stack[stack_top]
			stack = stack[:stack_top]
		case "*":
			stack[stack_top-1] = stack[stack_top-1] * stack[stack_top]
			stack = stack[:stack_top]
		case "/":
			stack[stack_top-1] = stack[stack_top-1] / stack[stack_top]
			stack = stack[:stack_top]
		default:
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}
	return stack[0]
}
