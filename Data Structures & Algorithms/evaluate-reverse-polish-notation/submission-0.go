func evalRPN(tokens []string) int {
	stack := make([]int,0)
	for _, token := range tokens {
		stack = parseOperator(token, stack)
	}

	return stack[0]
}

func parseOperator(s string, stack []int) []int {
	var res int
	switch s {
	case "+":
		res = stack[len(stack)-2] + stack[len(stack)-1]
	case "-":
		res = stack[len(stack)-2] - stack[len(stack)-1]
	case "*":
		res = stack[len(stack)-2] * stack[len(stack)-1]
	case "/":
		res = stack[len(stack)-2] / stack[len(stack)-1]
	default:
		res, _ = strconv.Atoi(s)
		stack = append(stack, res)
		return stack
	}

	stack = stack[:len(stack)-2]
	stack = append(stack, res)
	return stack
}