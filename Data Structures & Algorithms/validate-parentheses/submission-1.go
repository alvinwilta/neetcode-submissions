func isValid(s string) bool {
    stack := make([]byte,0)

	for i:=0;i<len(s);i++ {
		isPair, close := getPair(s[i])
		if isPair {
			stack = append(stack, close)
			continue
		}

		if len(stack)>0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}

// return corresponding closing pair for the bracket
// false if not a open bracket
func getPair(s byte) (bool, byte) {
	switch s {
	case '{':
		return true, '}'
	case '(':
		return true, ')'
	case '[':
		return true, ']'
	default:
		return false, 0
	}
}