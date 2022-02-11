package leetcode

// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
// 返回一个表示每个字符串片段的长度的列表。
// 示例：
// 输入：S = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca", "defegde", "hijhklij"。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

// 先扫描一遍整个S，建立一个字典，字典的键是字母，值是这个字母最后一次出现的下标，
// 然后再线性扫描S，设置变量right 用于表示分隔开的子字符串的末尾在S里的下标，变量left表示子字符串的开头在S里的下标
// 如果当前扫描的元素的最后一次出现的下标比right大，就说明子字符串还需要延长，就刷新right，
// 如果当前扫描的元素的下标已经达到了right，就说明这个子字符串完全已经找到了，就可以把答案添加到res中。

func PartitionLabels(s string) []int {
	dic := make(map[byte]int)
	for k, v := range s {
		dic[byte(v)] = k
	}
	res := make([]int, 0)
	//开始字母下标
	start := dic[s[0]]
	end := 0
	for k, v := range s {
		start = max(start, dic[byte(v)])
		if k >= start {
			res = append(res, start-end+1)
			end = start + 1
		}
	}
	return res
}
