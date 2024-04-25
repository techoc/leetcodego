package utils

func StringListIsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		// 在 b 中是否存在 a[i]
		if !Contains(b, a[i]) {
			return false
		}
	}

	return true
}

func Contains(list []string, s string) bool {
	for i := 0; i < len(list); i++ {
		if list[i] == s {
			return true
		}
	}

	return false
}
