package easy

import "unicode"

//	统计 licensePlate 中每个字母的出现次数（忽略大小写）
//	然后遍历 words 中的每个单词
//	若 26 个字母在该单词中的出现次数均不小于在 licensePlate 中的出现次数，则该单词是一个补全词
//	返回最短且最靠前的补全词
func ShortestCompletingWord(licensePlate string, words []string) (ans string) {
	cnt := [26]int{}
	for _, ch := range licensePlate {
		if unicode.IsLetter(ch) {
			cnt[unicode.ToLower(ch)-'a']++
		}
	}

next:
	for _, word := range words {
		c := [26]int{}
		for _, ch := range word {
			c[ch-'a']++
		}
		for i := 0; i < 26; i++ {
			if c[i] < cnt[i] {
				continue next
			}
		}
		if ans == "" || len(word) < len(ans) {
			ans = word
		}
	}
	return
}
