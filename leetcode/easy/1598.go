package easy

func minOperations(logs []string) (ans int) {
	if len(logs) == 0 {
		return 0
	}

	for _, val := range logs {
		if val != "../" && val != "./" {
			ans += 1
		}
		if val == "../" {
			if ans != 0 {
				ans -= 1
			}
		}
		if val == "./" {
			continue
		}
	}
	return ans
}
