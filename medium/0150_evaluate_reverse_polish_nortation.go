package medium

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]int, 0, len(tokens))

	for _, token := range tokens {
		if num, err := strconv.Atoi(token); err == nil {
			stack = append(stack, num)
			continue
		}
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		res := ResultByOperation(token, a, b) 

		stack = append(stack, res)
	}

	return stack[0]
}


func ResultByOperation(op string, first, second int) int {
	switch op {
	case "/":
		return first / second
	case "*":
		return  first * second
	case "-":
		return first - second
	case "+":
		return first +second
	}
	return 0
}
