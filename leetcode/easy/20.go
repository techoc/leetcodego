package easy

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
//有效字符串需满足：
//1. 左括号必须用相同类型的右括号闭合。
//2. 左括号必须以正确的顺序闭合。

// IsValid 20.有效的括号
//搞个栈，压栈判断就可以了
func IsValid(s string) bool {
	//特殊情况
	if len(s) == 0 {
		return true
	}
	//特殊情况
	if len(s)%2 != 0 {
		return false
	}
	var stack []rune
	for _, v := range s {
		//如果是左括号，直接入栈
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
		} else {
			//如果栈空，直接返回false
			if len(stack) == 0 {
				return false
			}
			//如果栈顶元素不匹配，直接返回false
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if v == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if v == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}
