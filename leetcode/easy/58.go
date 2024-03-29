package easy

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	//倒序遍历
	flag := false
	//记录倒序时的空格数
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if flag {
				return len(s) - i - count - 1
			}
			count++
			continue
		}
		if s[i] != ' ' {
			flag = true
		}
	}
	return len(s) - count
}
