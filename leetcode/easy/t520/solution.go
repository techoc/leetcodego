package t520

// 520. 检测大写字母
// https://leetcode.cn/problems/detect-capital
func detectCapitalUse(word string) bool {
	count := 0
	for _, c := range word {
		if c >= 'A' && c <= 'Z' {
			count++
		}
	}
	if count == 0 || count == len(word) || (count == 1 && word[0] >= 'A' && word[0] <= 'Z') {
		return true
	}
	return false
}
