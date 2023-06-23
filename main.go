package main

import "log"

func main() {
	var inputBaseball []string
	inputBaseball = []string{
		"5", "2", "C", "D", "+",
	}
	inputParenthesses := []string{
		")", "(", "(", ")", ")", ")",
	}
	baseball := calPoints(inputBaseball)
	countAndSayAns := countAndSay(4)
	fizzBuzzAns := fizzBuzz(5)
	longestValidParenthesesAns := longestValidParentheses(inputParenthesses)

	log.Println("Baseball : ", baseball)
	log.Println("Count and say : ", countAndSayAns)
	log.Println("Fizz n Buzz : ", fizzBuzzAns)
	log.Println("Valid Parenthesses : ", longestValidParenthesesAns)
}
