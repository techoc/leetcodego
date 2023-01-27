package t2309

import "unicode"

func greatestLetter(s string) string {
	set := map[rune]bool{}
	for _, c := range s {
		set[c] = true
	}
	for i := 'Z'; i >= 'A'; i-- {
		if set[i] && set[unicode.ToLower(i)] {
			return string(i)
		}
	}
	return ""
}
