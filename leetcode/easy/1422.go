package easy

import "strings"

func maxScore(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		//依次按照字符下标分割
		//对比每次求出来的得分，取最大值
		ans = max(ans, strings.Count(s[:i], "0")+strings.Count(s[i:], "1"))
	}
	return
}
