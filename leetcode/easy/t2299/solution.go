package t2299

import (
	"strings"
	"unicode"
)

// 它有至少 8个字符。
// 至少包含一个小写英文字母。
// 至少包含一个大写英文字母。
// 至少包含一个数字
// 至少包含一个特殊字符。特殊字符为："!@#$%^&*()-+"中的一个。
// 它不包含2个连续相同的字符（比方说"aab"不符合该条件，但是"aba"符合该条件）。
func strongPasswordCheckerII(password string) bool {
	// 至少8个字符
	if len(password) < 8 {
		return false
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool

	for i, ch := range password {
		//不包含两个连续相同的字符
		if i != len(password)-1 && password[i] == password[i+1] {
			return false
		}
		if unicode.IsLower(ch) { //至少包含一个小写字符
			hasLower = true
		} else if unicode.IsUpper(ch) { //至少包含一个大写字符
			hasUpper = true
		} else if unicode.IsDigit(ch) { //至少包含一个数字
			hasDigit = true
		} else if strings.ContainsRune("!@#$%^&*()-+", ch) { //至少包含一个特殊字符
			hasSpecial = true
		}
	}

	return hasLower && hasUpper && hasDigit && hasSpecial
}
