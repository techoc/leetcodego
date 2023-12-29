package t13

// 13.罗马数字转整数
// https://leetcode.cn/problems/roman-to-integer
func romanToInt(s string) int {
	//创建一个map，用于存储罗马数字和对应的数字
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	//创建一个变量，用于存储结果
	result := 0
	//创建一个变量，用于存储字符的长度
	n := len(s)

	// 遍历字符串
	for i := range s {
		//数组切片
		value := m[s[i:i+1]]
		//如果前一个字符小于后一个字符，则减去前一个字符的值
		if i < n-1 && value < m[s[i+1:i+2]] {
			result -= value
		} else {
			result += value
		}
	}
	return result
}
