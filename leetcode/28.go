package leetcode

//实现strStr()函数。
//给你两个字符串haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回 -1 。
//说明：
//当needle是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。
//对于本题而言，当needle是空字符串时我们应当返回 0 。这与 C 语言的strstr()以及 Java 的indexOf()定义相符。

func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	if len(haystack) == 0 {
		return -1
	}
	for i := 0; i <= len(haystack); i++ {
		if i+len(needle) > len(haystack) {
			return -1
		}
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return 0
}
