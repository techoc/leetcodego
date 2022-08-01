package easy

import "strings"

func generateTheString(n int) string {
	//奇数就全部a
	//偶数就是n-1个a  一个b
	if n%2 == 1 {
		//奇数
		return strings.Repeat("a", n)
	}
	//偶数
	return strings.Repeat("a", n-1) + "b"
}
