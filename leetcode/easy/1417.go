package easy

func reformat(s string) string {
	charNum, numNum := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			//获取字母的个数
			charNum++
		} else if s[i] >= '0' && s[i] <= '9' {
			//获取数字的个数
			numNum++
		}
	}

	//如果没办法满足 ，则返回空串
	if charNum-numNum > 1 || numNum-charNum > 1 {
		return ""
	}

	res := make([]byte, len(s))
	charIndex, numIndex := 0, 0
	if charNum > numNum {
		numIndex++
	} else {
		charIndex++
	}

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			res[charIndex] = s[i]
			charIndex += 2
		} else if s[i] >= '0' && s[i] <= '9' {
			res[numIndex] = s[i]
			numIndex += 2
		}
	}

	return string(res)
}
