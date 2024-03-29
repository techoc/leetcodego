package easy

//给你两个字符串a和b，请返回 这两个字符串中 最长的特殊序列 的长度。如果不存在，则返回 -1。
//「最长特殊序列」定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。
//字符串s的子序列是在从s中删除任意数量的字符后可以获得的字符串。
//例如，"abc" 是 "aebdc" 的子序列，因为删除 "aebdc" 中斜体加粗的字符可以得到 "abc" 。
//"aebdc" 的子序列还包括 "aebdc" 、 "aeb" 和 "" (空字符串)。
//示例 1：
//输入: a = "aba", b = "cdc"
//输出: 3

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}
