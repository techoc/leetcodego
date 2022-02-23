package leetcode

import "strconv"

//给你两个二进制字符串，返回它们的和（用二进制表示）。

//输入为 非空 字符串且只包含数字1和0。
//示例1:
//输入: a = "11", b = "1"
//输出: "100"
//示例2:
//输入: a = "1010", b = "1011"
//输出: "10101"

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
