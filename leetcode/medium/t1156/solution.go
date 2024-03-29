package t1156

// 1156. 单字符重复子串的最大长度
func maxRepOpt1(text string) int {
	count := make(map[rune]int)
	for _, c := range text {
		count[c]++
	}

	res := 0
	i := 0
	for i < len(text) {
		// step1: 找出当前连续的一段 [i, j)
		j := i
		for j < len(text) && text[j] == text[i] {
			j++
		}
		curCnt := j - i

		// step2: 如果这一段长度小于该字符出现的总数，并且前面或后面有空位，则使用 cur_cnt + 1 更新答案
		if curCnt < count[rune(text[i])] && (j < len(text) || i > 0) {
			res = max(res, curCnt+1)
		}

		// step3: 找到这一段后面与之相隔一个不同字符的另一段 [j + 1, k)，如果不存在则 k = j + 1
		k := j + 1
		for k < len(text) && text[k] == text[i] {
			k++
		}
		res = max(res, min(k-i, count[rune(text[i])]))
		i = j
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
