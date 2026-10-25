func generateParenthesis(n int) []string {
	return recGenParenthesis("", n, 0)
}

func recGenParenthesis(str string, open, close int) []string {
	if open == 0 && close == 0 {
		return []string{str}
	}

	res := []string{}
	if open > 0 {
		strOpen := str + "("
		res = append(res, recGenParenthesis(strOpen, open-1, close+1)...)
	}

	if close > 0 {
		strClose := str + ")"
		res = append(res, recGenParenthesis(strClose, open, close-1)...)
	}

	return res
}