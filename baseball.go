package main

import "strconv"

func calPoints(ops []string) int {
	var stack []int

	for i := range ops {
		num, err := strconv.Atoi(ops[i])

		if err == nil {
			stack = append(stack, num)
		} else {
			if ops[i] == "+" {
				ans := stack[len(stack)-1] + stack[len(stack)-2]
				stack = append(stack, ans)
			} else if ops[i] == "C" {
				stack = stack[:len(stack)-1]
			} else {
				ans := stack[len(stack)-1] * 2
				stack = append(stack, ans)
			}
		}
	}

	ans := 0

	for i := range stack {
		ans += stack[i]
	}

	return ans
}
