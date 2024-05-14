package t344

// 344. 反转字符串
// https://leetcode.cn/problems/reverse-string
// 2024-05-14
func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {	// 交换左右两个元素 直到左右指针相遇
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
