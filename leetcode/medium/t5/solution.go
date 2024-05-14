package t5

// 5. 最长回文子串
// https://leetcode.cn/problems/longest-palindromic-substring
// 2024-05-14
func longestPalindrome(s string) string {
	res := "" // 初始化结果字符串

	for i := 0; i < len(s); i++ {
		// 尝试以单个字符 s[i] 为中心扩展，寻找最长回文子串
		s1 := palindrome(s, i, i)
		// 尝试以两个连续字符 s[i] 和 s[i+1] 为中心扩展，寻找最长回文子串
		s2 := palindrome(s, i, i+1)

		// 更新结果字符串为当前找到的最长回文子串
		if len(res) < len(s1) {
			res = s1
		}
		if len(res) < len(s2) {
			res = s2
		}
	}
	return res
}

// palindrome 查找以给定位置 l 和 r 为中心的最长回文子串。
// 参数:
// s - 输入的字符串
// l - 左边扩展的起始位置
// r - 右边扩展的起始位置
// 返回值:
// 返回找到的以位置 l 和 r 为中心的最长回文子串
func palindrome(s string, l int, r int) string {
	// 从中心向两边展开，直到无法扩展或字符不匹配
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	// 返回以位置 l 和 r 为中心的最长回文子串
	return s[l+1 : r]
}
