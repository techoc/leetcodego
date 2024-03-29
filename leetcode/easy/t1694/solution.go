package _1694

import "strings"

func reformatNumber(number string) string {
	//去除空格 和 -
	s := strings.ReplaceAll(number, " ", "")
	s = strings.ReplaceAll(s, "-", "")

	var ans []string
	i := 0
	for ; i+4 < len(s); i += 3 {
		ans = append(ans, s[i:i+3])
	}
	s = s[i:]
	if len(s) < 4 {
		ans = append(ans, s)
	} else {
		ans = append(ans, s[:2], s[2:])
	}
	return strings.Join(ans, "-")
}
