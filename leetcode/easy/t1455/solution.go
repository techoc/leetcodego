package t1455

import "strings"

// 1455. 检查单词是否为句中其他单词的前缀
// https://leetcode.cn/problems/check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence/
func isPrefixOfWord(sentence, searchWord string) int {
	for i, index, n := 0, 1, len(sentence); i < n; i++ {
		start := i
		for i < n && sentence[i] != ' ' { // 分割单词
			i++
		}
		end := i
		// 判断是否为前缀
		if strings.HasPrefix(sentence[start:end], searchWord) {
			return index
		}
		index++
	}
	return -1
}
