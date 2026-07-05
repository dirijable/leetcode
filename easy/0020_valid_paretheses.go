package easy

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range len(s) {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
        if(len(stack) == 0) {
            return false
        }
		stackIdx := len(stack) - 1
		if stack[stackIdx] == '(' && s[i] != ')' ||
			stack[stackIdx] == '[' && s[i] != ']' ||
			stack[stackIdx] == '{' && s[i] != '}' {
			return false
		}
		stack = stack[:stackIdx]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
