package t2586

// https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range
func vowelStrings(words []string, left int, right int) int {
	// 讲元音字母存入哈希表
	vowels := map[byte]struct{}{'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {}}
	ans := 0
	//遍历列表
	for _, word := range words[left : right+1] {
		// 检测首尾字母是否是元音字母
		if _, ok1 := vowels[word[0]]; ok1 {
			if _, ok2 := vowels[word[len(word)-1]]; ok2 {
				ans++
			}
		}
	}
	return ans
}
