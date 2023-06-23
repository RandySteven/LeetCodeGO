package main

func longestValidParentheses(parentheses []string) int {
	stack := []int{-1} //initiate a stack
	max := 0           //initiate max value of parentheses

	for i := 0; i < len(parentheses); i++ {
		/**
		validate if index i == '('
		it will push into stack
		*/
		if parentheses[i] == "(" { //string start with (
			stack = append(stack, i)
		} else { //string start with )
			/**
			1. stack pop
			2. if stack is empty then push index i again into stack
			3. else stack is not empty
			4. if index - the latest stack more than max
			5. max = index - latest stack
			*/
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				if i-stack[len(stack)-1] > max {
					max = i - stack[len(stack)-1]
				}
			}
		}
	}
	return max
}
