package easy

func convertToTitle(columnNumber int) string {
	var ans []byte
	for columnNumber > 0 {
		//在取数之前减一
		columnNumber--
		ans = append(ans, 'A'+byte(columnNumber%26))
		columnNumber /= 26
	}
	//反转字符串
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-1-i] = ans[n-1-i], ans[i]
	}
	return string(ans)
}
